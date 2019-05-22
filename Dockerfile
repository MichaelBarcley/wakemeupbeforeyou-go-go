FROM golang:1.12.5-alpine3.9
RUN mkdir /app
ADD . /app
WORKDIR /app
RUN apk update && apk add git
RUN go get golang.org/x/time/rate
RUN go get -u github.com/tidwall/gjson
RUN go build -o main .
CMD ["/app/main"]