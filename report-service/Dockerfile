FROM golang:1.23.4
WORKDIR /app
COPY . .

RUN go mod tidy
RUN go build -o report-service main.go
EXPOSE 8083
CMD ["./report-service"]
