# syntax=docker/dockerfile:1
FROM golang:latest
WORKDIR /src
COPY . .
EXPOSE 6700
RUN go mod download
RUN go build -o /bin/buffer ./cmd/buffer
ENTRYPOINT [ "/bin/buffer" ]