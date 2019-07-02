FROM golang:1.12 AS build

WORKDIR /workspace

COPY . .

ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64

RUN go get -v .
RUN go build -o /app.out .

# ==============================

FROM alpine:3.8

WORKDIR /app

RUN apk add --no-cache ca-certificates

COPY --from=build /app.out .

RUN echo "0       *       *       *       *       /app/app.out" >> /etc/crontabs/root

CMD [ "crond", "-f" ]
