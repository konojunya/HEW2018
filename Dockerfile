FROM alpine:3.6
WORKDIR /root

#EXPOSE 8000

ADD ./cmd/main main

ENTRYPOINT ["./main"]
