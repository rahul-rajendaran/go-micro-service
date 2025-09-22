FROM golang:1.24-alpine


WORKDIR /app

# Install git & certs for module download
RUN apk add --no-cache git ca-certificates

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o server .

EXPOSE 4002
CMD ["./server"]
