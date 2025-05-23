FROM  golang:1.23.4

WORKDIR /app

COPY go.mod go.sum ./

COPY . .

RUN ls -a

RUN go mod download

RUN go install github.com/air-verse/air@latest

EXPOSE 8080

CMD ["air", "-c", ".air.toml"]