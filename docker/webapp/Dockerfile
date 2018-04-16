FROM node:8.11 AS builder
MAINTAINER Kevin <work@didelotkev.ovh>

# Install git in order to clone the project later
RUN apt-get update && \
    apt-get install -y git

RUN mkdir /webapp

WORKDIR /webapp
RUN git clone https://github.com/DidelotK/heat_loadbalanced_autoscaling_group_poc

WORKDIR /webapp/heat_loadbalanced_autoscaling_group_poc/webapp/front
RUN yarn install
RUN yarn run build


FROM golang:1.10
COPY --from=builder /webapp /go/src/app

RUN apt-get update && apt-get install -y --no-install-recommends \
           stress-ng \
    && rm -rf /var/lib/apt/lists/*

WORKDIR /go/src/app
RUN ls
RUN go get -d -v ./...
RUN go install -v ./...

EXPOSE 8080

WORKDIR /go/src/app/heat_loadbalanced_autoscaling_group_poc/webapp/server/src/app
CMD ["app"]

