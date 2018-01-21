FROM centos:centos6
MAINTAINER Imagine teng

ENV SSH_PASSWORD=teng

#Install base tool
#RUN yum install -y vim

# Copy files for setting
ADD . /opt/

# Setup default path
WORKDIR /home

# Private expose
EXPOSE 22 80 8080

# Volume for web server install
VOLUME ["/home/websit","/home/config","home/logs"]

# Start run shell
CMD ["bash"]