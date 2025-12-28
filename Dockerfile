FROM golang:1.21

WORKDIR /app

COPY . .
RUN go mod download



RUN pwd
RUN ls

RUN CGO_ENABLED=0 GOOS=linux go build -o /task_app ./cmd/server/main.go

EXPOSE 8080
CMD ["/task_app"]