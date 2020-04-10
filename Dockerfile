FROM golang:1.13-alpine

# Copy
RUN mkdir /app
ADD main.go /app/
ADD go.mod /app/
WORKDIR /app

# Build
RUN go mod vendor
RUN CGO_ENABLED=0 go build -o main .

# Start
EXPOSE 8080
CMD ["/app/main"]
