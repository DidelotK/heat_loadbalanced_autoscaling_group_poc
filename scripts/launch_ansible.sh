#!/bin/sh
echo "Clone git project"
cd /opt
git clone https://github.com/DidelotK/heat_loadbalanced_autoscaling_group_poc

echo "Launch ansible"
cd /opt/heat_loadbalanced_autoscaling_group_poc
ansible-playbook main.yml

