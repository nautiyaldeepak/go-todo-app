FROM golang:1.22 AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o goapp .

FROM alpine:3.19.1
COPY --from=builder /app ./app
WORKDIR app
EXPOSE 3000
CMD ["./goapp"]