FROM golang:1.12 AS build

WORKDIR /workspace

COPY . .

ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64

RUN go get -v .
RUN go build -o /app.out .

# ==============================

FROM scratch

WORKDIR /app

COPY --from=build /app.out .

CMD ["/app/app.out"]
