FROM alpine:edge


ENV MOSQUITTO_VERSION 1.4.15

RUN echo 'http://dl-4.alpinelinux.org/alpine/edge/main' >> /etc/apk/repositories


RUN apk --update add git build-base libressl-dev libwebsockets-dev c-ares-dev util-linux-dev hiredis-dev curl-dev libxslt docbook-xsl hiredis libwebsockets libuuid c-ares libressl curl ca-certificates && rm -rf /var/cache/apk/*


ADD http://mosquitto.org/files/source/mosquitto-${MOSQUITTO_VERSION}.tar.gz /build/
RUN cd /build && tar xvfz mosquitto-${MOSQUITTO_VERSION}.tar.gz && \
    cd /build/mosquitto-${MOSQUITTO_VERSION} && make


RUN cd /build && git clone https://github.com/jpmens/mosquitto-auth-plug.git && \
    cd /build/mosquitto-auth-plug && cp config.mk.in config.mk && \
    cd
    sed -i.bak 's/BACKEND_REDIS \?= no/BACKEND_REDIS \?= yes/' /build/mosquitto-auth-plug/config.mk && \
    sed -i.bak 's/BACKEND_POSTGRES \?= yes/BACKEND_POSTGRES \?= yes/' /build/mosquitto-auth-plug/config.mk && \
    sed -i.bak 's/MOSQUITTO_SRC =/MOSQUITTO_SRC = \/build\/mosquitto-${MOSQUITTO_VERSION}/' /build/mosquitto-auth-plug/config.mk && \
    cd /build/mosquitto-auth-plug && make

CMD ["/bin/sh"]