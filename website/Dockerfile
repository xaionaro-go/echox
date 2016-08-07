FROM reg.lab.st/armor:0.1.0
MAINTAINER Vishal Rana <vr@labstack.com>

COPY armor.json /etc
COPY public /srv/www

EXPOSE 80

CMD ["-c", "/etc/armor.json"]
