FROM alpine:3.6
WORKDIR /root

ADD ./cmd/main .
ADD ./public ./public
ADD ./view ./view
ADD ./dbconfig.yml ./dbconfig.yml

EXPOSE 8000

ENTRYPOINT ["./main"]
