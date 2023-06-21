FROM nginx

RUN apt update \
  && apt install -y \
  dnsutils \
  net-tools \
  apache2-utils \
  && mkdir -p /var/www \
  && mkdir -p /etc/nginx/ssl \
  && adduser --system --group --quiet --home /home/user1 user1 \
  && adduser --system --group --quiet --home /home/user2 user2 

RUN htpasswd -c -b /home/user1/.htpasswd user1 123
RUN htpasswd -c -b /home/user2/.htpasswd user2 123

# Copy configuration files
COPY nginx.conf		/etc/nginx/
COPY ssl		      /etc/nginx/ssl
COPY index.html		/var/www/
COPY user1        /home/user1
COPY user2        /home/user2

# Expose Ports
EXPOSE 80
EXPOSE 443
EXPOSE 6060 

# docker build -t nginx-server01:latest ./nginx
# docker run -d --name nginx --net my_network --ip 172.29.0.3 nginx-server01 && docker network connect intraNet nginx 
