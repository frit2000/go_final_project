FROM golang:1.21.5

WORKDIR /app

COPY . .

ENV TODO_PORT=7450 TODO_DBFILE=. TODO_PASSWORD=12345678 
ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64

RUN go mod download

RUN go build -o /mytracker main.go

CMD ["/mytracker"]