FROM nginx:latest
MAINTAINER Kevin <work@didelotkev.ovh>

# Add all necessary configurations
COPY conf/webapp.conf /etc/nginx/conf.d/webapp.conf
COPY conf/nginx.conf /etc/nginx/nginx.conf

# Expose HTTP port
EXPOSE 80

VOLUME /var/log/nginx

CMD ["nginx", "-g", "daemon off;"]

