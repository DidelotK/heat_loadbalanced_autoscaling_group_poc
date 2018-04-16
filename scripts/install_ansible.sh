#!/bin/sh
echo "Update yum repo"
yum update -y

echo "Install pip"
yum update -y
yum install -y python-pip  
pip install -U pip

echo "Install git"
yum install -y git

echo "Install ansible"
pip install ansible

echo "Clear yum cache"
yum clean all
rm -rf /var/cache/yum

