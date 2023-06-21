FROM ubuntu/bind9

RUN apt update \
  && apt install -y \
  bind9-doc \
  dnsutils \
  geoip-bin \
  mariadb-server \
  net-tools

# Copy configuration files
COPY named.conf.options /etc/bind/
COPY named.conf.local /etc/bind/
COPY db.angelita.co /etc/bind/

# Expose Ports
EXPOSE 53/udp

# Start the Name Service
CMD ["/usr/sbin/named", "-g", "-c", "/etc/bind/named.conf", "-u", "bind"]

# docker build -t dns-projetoredes:latest ./bind9
# docker run -d --name bind9 --net my_network --ip 172.29.0.2 dns-projetoredes  
