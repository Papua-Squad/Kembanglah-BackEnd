FROM golang:alpine

RUN apk update && apk add --no-cache git

WORKDIR /app

COPY . ./

RUN go mod tidy

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o main main.go

EXPOSE 3000

ENTRYPOINT ["/app/main"]