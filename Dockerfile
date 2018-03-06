FROM alpine:3.6
WORKDIR /root
ENV GO_ENV production

RUN set -x \
  && apk upgrade --no-cache \
  && apk --update add tzdata \
  && cp /usr/share/zoneinfo/Asia/Tokyo /etc/localtime \
  && apk del tzdata \
  && rm -rf /var/cache/apk/*

RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*

ADD ./cmd/main .
ADD ./public ./public
ADD ./view ./view
ADD ./config.yml ./config.yml
ADD ./cmd/migrate .
ADD ./cmd/seed .

EXPOSE 8000

# Run it
ENTRYPOINT ["./main"]
