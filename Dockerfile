FROM golang:1.24.1-bullseye

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o godocker ./cmd/app/

EXPOSE 8081

CMD ["./godocker"]