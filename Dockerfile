
# syntax=docker/dockerfile:1
ARG ARCH=
FROM ${ARCH}golang:1.20.1-bullseye
LABEL maintainer="Ken Ellorando (kenellorando.com)"
LABEL source="github.com/kenellorando/maya"
WORKDIR /maya
COPY ./* ./
RUN go mod download
RUN go build -o /maya/maya-server

RUN useradd -s /bin/bash maya
RUN chown maya:maya /maya/ /maya/* /maya/maya-server
RUN chmod u+wrx /maya/ /maya/* 

USER maya
CMD /maya/maya-server -token $DISCORD_TOKEN
