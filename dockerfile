FROM golang:1.21

WORKDIR /app

COPY . .

ENV TODO_PORT=7450 TODO_DBFILE=. TODO_PASSWORD=12345678 CGO_ENABLED=0 GOOS=linux GOARCH=amd64

RUN go mod download
RUN go mod tidy

RUN go build -o /mytracker main.go

CMD ["/mytracker"]