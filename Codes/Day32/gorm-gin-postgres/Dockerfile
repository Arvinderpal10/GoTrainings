FROM golang:1.25-alpine

WORKDIR /app

# Copy everything at once (no caching, but simpler)
COPY . .

# Download dependencies and build
RUN go mod download && go build -o main .

EXPOSE 8080

CMD ["./main"]