#!/usr/bin/env bash


VERSION="0.0.1"
CONFIG=/opt/deploy/deploy.conf
LOG=/var/log/deploy.log
TEST=0
REF=
ENV=

#
# Output usage information.
#

usage() {
  cat <<-EOF

  Usage: deploy [options] <env> [command]

  Options:

    -c, --config <path>  set config path. defaults to ./deploy.conf
    -T, --no-tests       ignore test hook
    -V, --version        output program version
    -h, --help           output help information

  Commands:

    setup                run remote setup commands
    config [key]         output config file or [key]
    curr[ent]            output current release commit
    exec|run <cmd>       execute the given <cmd>
    [ref]                deploy to [ref], the 'ref' setting, or latest tag

EOF
}

#
# Abort with <msg>
#

abort() {
  echo $(date +"%Y-%m-%d %T.%3N")" ERROR $@" 1>&2
  exit 1
}

#
# Log <msg>.
#

log() {
  echo $(date +"%Y-%m-%d %T.%3N")" INFO $@"
}

#
# Set configuration file <path>.
#

set_config_path() {
  test -f $1 || abort invalid --config path
  CONFIG=$1
}

#
# Check if config <section> exists.
#

config_section() {
  grep "^\[$1" $CONFIG &> /dev/null
}

#
# Get config value by <key>.
#

config_get() {
  local key=$1
  test -n "$key" \
    && grep "^\[$ENV" -A 20 $CONFIG \
    | grep "^$key" \
    | head -n 1 \
    | cut -d ' ' -f 2-999
}

#
# Output version.
#

version() {
  echo $VERSION
}

#
# Run the given remote <cmd>.
#

run() {
  echo "\"$@\"" >> $LOG
  $@
}


#
# Output config or [key].
#

config() {
  if test $# -eq 0; then
    cat $CONFIG
  else
    config_get $1
  fi
}

#
# Execute hook <name> relative to the path configured.
#

hook() {
  test -n "$1" || abort hook name required
  local hook=$1
  local path=`config_get path`
  local cmd=`config_get $hook`
  if test -n "$cmd"; then
    log "executing $hook \`$cmd\`"
    run "php $path/$cmd 2>&1 | tee -a $LOG; \
      exit \${PIPESTATUS[0]}"
    test $? -eq 0
  else
    log hook $hook
  fi
}

#
# Run setup.
#

setup() {
  local path=`config_get path`
  local repo=`config_get repoPath`
  log checking paths
  # run "ls $path > /dev/null"
  # test $? -eq 0 || abort setup paths failed

  # run "ls $repo > /dev/null"
  # test $? -eq 0 || abort repo path failed
}

#
# Deploy [ref].
#

deploy() {
  local ref=$1
  local repo=`config_get repoPath`
  local path=`config_get path`
  local www=`config_get wwwPath`
  local app=`config_get appDir`
  log deploying

  hook pre-deploy || abort pre-deploy hook failed

  # reset HEAD
  log resetting HEAD to $ref
  run "git --git-dir=$repo --work-tree=$path reset --hard $ref"
  test $? -eq 0 || abort git fetch failed

  # link current
  run "ln -sfn $path/$app $www"
  test $? -eq 0 || abort symlink failed


  hook post-deploy || abort post-deploy hook failed

  if test $TEST -eq 1; then
    hook test
    if test $? -ne 0; then
      log tests failed
    fi
  else
    log ignoring tests
  fi

  # done
  log successfully deployed $ref
}

#
# Get current commit.
#

current_commit() {
  local repo=`config_get repoPath`
  local path=`config_get path`

  run "git --git-dir=$repo --work-tree=$path rev-parse --short HEAD"
}


#
# Require environment arg.
#

require_env() {
  config_section $ENV || abort "[$ENV] config section not defined"
  test -z "$ENV" && abort "<env> required"
}

#
# Ensure all changes are committed and pushed before deploying.
#

check_for_local_changes() {
  local repo=`config_get repoPath`
  local path=`config_get path`

  git --git-dir=$repo --work-tree=$path --no-pager diff --exit-code --quiet          || abort "commit or stash your changes before deploying"
  git --git-dir=$repo --work-tree=$path --no-pager diff --exit-code --quiet --cached || abort "commit your staged changes before deploying"
  [ -z "`git rev-list @{upstream}.. -n 1`" ]       || abort "push your changes before deploying"
}

# parse argv

while test $# -ne 0; do
  arg=$1; shift
  case $arg in
    -h|--help) usage; exit ;;
    -V|--version) version; exit ;;
    -c|--config) set_config_path $1; shift ;;
    run|exec) require_env; run "cd `config_get path` && $@"; exit ;;
    curr|current) require_env; current_commit; exit ;;
    setup) require_env; setup $@; exit ;;
    config) config $@; exit ;;
    *)
      if test -z "$ENV"; then
        ENV=$arg;
      else
        REF="$REF $arg";
      fi
      ;;
  esac
done

require_env
# check_for_local_changes

# deploy
deploy "${REF:-`config_get ref`}"

