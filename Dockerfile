FROM golang:1.24-bookworm

WORKDIR /app
COPY . /app
RUN go mod tidy

EXPOSE 8080
CMD [ "go", "run", "/app/cmd/main.go"]
