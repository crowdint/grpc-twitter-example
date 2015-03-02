# -*- mode: ruby -*-
# vi: set ft=ruby :

$setup = <<-SETUP
  apt-get update && upgrade
  sudo apt-get install automake autoconf build-tools build-essential libtool g++-multilib g++

  ## Install latest protobufs
  wget https://github.com/google/protobuf/releases/download/v3.0.0-alpha-2/protobuf-cpp-3.0.0-alpha-2.tar.gz
  tar -xvf protobuf-cpp-3.0.0-alpha-2.tar.gz
  cd protobuf-3.0.0-alpha-2/
  ./autogen.sh
  ./configure
  make
  make check
  sudo make install
  type protoc >/dev/null 2>&1 || { echo >&2 "Failed to install protobufs"; exit 1; }

  ## Install Go
  wget https://storage.googleapis.com/golang/go1.4.2.linux-amd64.tar.gz
  tar -C /usr/local -xzf go1.4.2.linux-amd64.tar.gz

  echo 'export PATH=$PATH:/usr/local/go/bin
  export GOPATH=/home/vagrant/go
  export GOBIN=$GOPATH/bin' >> /home/vagrant/.bash_profile
  source ~/.bash_profile

  mkdir -p $GOPATH/src/github.com/crowdint
  ln -s /vagrant /home/vagrant/go/src/github.com/crowdint/grpc-twitter-example
SETUP

Vagrant.configure(2) do |config|
  config.vm.box = "hashicorp/precise64"
  config.vm.provision "shell", inline: $script
end
