FROM alpine3.6
# ENV GO_ENV production

ADD ./migrate .
ADD ./seed .

CMD ./migrate
CMD ./seed