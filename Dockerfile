FROM alpine:3.6
WORKDIR /root

ADD ./cmd/main .
ADD ./public ./public
ADD ./view ./view
ADD config.yml config.yml

EXPOSE 8000

# herokuではENTRYPOINTではなくCMDを使う
CMD ["./main"]