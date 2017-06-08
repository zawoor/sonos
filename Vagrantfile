# -*- mode: ruby -*-
# vi: set ft=ruby :

Vagrant.configure("2") do |config|

    config.vm.box = "switch/centos-6.7"
	config.vm.network "public_network"
    config.ssh.forward_agent = true

    config.vm.provision "file", source: "~/.ssh/config", destination: "/home/vagrant/.ssh/config"

    config.vm.provision "shell", inline: 
<<SCRIPT
		service iptables stop
		curl --silent --location https://rpm.nodesource.com/setup_6.x | bash -
		yum install -y nodejs golang
		cd /vagrant
		go build; 
SCRIPT
end