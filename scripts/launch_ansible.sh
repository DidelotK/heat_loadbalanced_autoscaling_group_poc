#!/bin/sh
echo "Clone git project"
cd /opt
git clone https://github.com/DidelotK/heat_loadbalanced_autoscaling_group_poc

echo "Install ansible roles"
cd /opt/heat_loadbalanced_autoscaling_group_poc/ansible
ansible-galaxy install -p ./roles -r requirements.yml

echo "Launch ansible"
ansible-playbook main.yml

