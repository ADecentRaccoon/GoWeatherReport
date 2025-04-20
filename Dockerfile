FROM golang:1.24-bookworm

WORKDIR /app
COPY . .

RUN go mod tidy && go build -o myapp /app/cmd/main.go

CMD [ "./myapp",  "echo", "app-started"]