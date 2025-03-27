FROM  golang:1.23.4

WORKDIR /app

COPY go.mod go.sum ./

COPY . .

RUN ls 

RUN go mod download


RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o app

EXPOSE 8080

CMD ["./app"]