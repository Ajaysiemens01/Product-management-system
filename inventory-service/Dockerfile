FROM golang:1.23.4
WORKDIR /app
COPY . .
RUN go mod tidy
RUN go build -o inventory-service main.go
EXPOSE 8082
CMD ["./inventory-service"]
