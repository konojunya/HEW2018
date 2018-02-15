FROM alpine:3.6
WORKDIR /root

ADD ./cmd/main .
ADD ./public ./public
ADD ./view ./view

EXPOSE 8000

ENV MODE production
ENV FIREBASE_TOKEN $FIREBASE_TOKEN

CMD ["./main"]
