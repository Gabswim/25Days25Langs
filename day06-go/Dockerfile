FROM golang:1.23.3-alpine
WORKDIR /usr/src
COPY solution.go .
RUN go build -o solution solution.go
CMD ["./solution"]