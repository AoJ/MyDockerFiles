apt-get install -y software-properties-common
add-apt-repository -y ppa:chris-lea/node.js
apt-get update
apt-get install -y nodejs && apt-get clean

echo '\n# Node.js\nexport PATH="node_modules/.bin:$PATH"' >> /root/.bash_profile
