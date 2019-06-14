FROM harbor.haodai.net/base/alpine:3.7cgo
WORKDIR /app

MAINTAINER wenzhenglin(http://g.haodai.net/wenzhenglin/ali-alert.git)

COPY ali-alert /app

CMD /app/ali-alert
ENTRYPOINT ["./ali-alert"]

EXPOSE 8080