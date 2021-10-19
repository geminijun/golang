FROM ubuntu
ENV MY_SERVICE_PORT=80
ENV VERSION=1.0
LABEL multi.label1="value1" multi.label2="value2" other="value3"
COPY bin/amd64/httpserver /httpserver
EXPOSE 80
ENTRYPOINT /httpserver
