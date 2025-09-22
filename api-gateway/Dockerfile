FROM golang:1.23-alpine

WORKDIR /app

# Install git & certs for module download
RUN apk add --no-cache git ca-certificates

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o server .

EXPOSE 8080
CMD ["./server"]
