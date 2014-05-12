# We're being careful here in that we're only cleaning up containers which are no longer running
docker ps -a | grep "Exit" | awk '{print $1}' | while read -r id ; do
  sudo docker rm $id
done

# We only remove images which don't have a name
sudo docker images | grep "^<none>" | head -n 1 | awk 'BEGIN { FS = "[ \t]+" } { print $3 }'  | while read -r id ; do
  sudo docker rmi $id
done

# docker rm $(docker ps -a -q)
# docker rmi $(docker images -q)
