FROM golang:1.21
WORKDIR /go/src

RUN go install github.com/cosmtrek/air@v1.44.0

CMD ["air", "-c", ".air.toml"]