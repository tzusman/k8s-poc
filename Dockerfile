FROM golang:1.14

WORKDIR /app

COPY main.go .
RUN go build main.go

ENTRYPOINT [ "./main" ]