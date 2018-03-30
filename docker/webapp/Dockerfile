FROM node:8.11
MAINTAINER Kevin <work@didelotkev.ovh>

# Install git in order to clone the project later
RUN apt-get update && \
    apt-get install -y git

# Small trick for https problem
RUN git config --global http.sslverify false

RUN mkdir /webapp

WORKDIR /webapp
RUN git clone https://github.com/DidelotK/heat_loadbalanced_autoscaling_group_poc

WORKDIR /webapp/heat_loadbalanced_autoscaling_group_poc/webapp/front
RUN npm install
RUN npm run build

