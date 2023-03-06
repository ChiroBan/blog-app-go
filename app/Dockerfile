FROM golang:alpine

RUN mkdir -p /var/www

WORKDIR /var/www

COPY . /var/www

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o app

EXPOSE 8080
ENTRYPOINT ["/var/www/app"]