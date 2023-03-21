FROM golang:1.18.3-alpine3.16 AS builder

WORKDIR /build
COPY /build/tmaicerp /home
COPY /build/commonApp/storage/logs/tmaic.log /home/commonApp/storage/logs/tmaic.log
COPY /build/commonApp/storage/logs/sql.log /home/commonApp/storage/logs/sql.log
COPY /build/commonApp/config/config.yml /home/commonApp/config/config.yml
COPY /build/commonApp/config/platform.yaml /home/commonApp/config/platform.yaml


EXPOSE 6000
EXPOSE 5000
EXPOSE 5001
EXPOSE 5002
EXPOSE 5003






