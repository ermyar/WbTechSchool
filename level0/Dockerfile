FROM golang:alpine

COPY . .

RUN go mod download
RUN go build -o app ./cmd/app/...

EXPOSE 8080

CMD ["./app"]