FROM golang:1.23 as builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download -x
COPY . .
RUN go build -o kubeit ./server/main.go && ls -l /app

FROM builder as runner
CMD ["./kubeit"]
