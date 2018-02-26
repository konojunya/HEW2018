FROM alpine:3.6
WORKDIR /root

ENV PORT 8000

ADD ./cmd/main .
ADD ./.env .
ADD ./public ./public
ADD ./view ./view

EXPOSE 8000

ENTRYPOINT ["./main"]
