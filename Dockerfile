FROM debian:stable-slim

LABEL MAINTAINER="Eduardo Tavares <edutav@gmail.com>"
LABEL APP="devbook-api"
LABEL VERSION="0.0.1"

COPY . /app

EXPOSE 3000

WORKDIR /app

CMD [ "./devbook-api" ]
