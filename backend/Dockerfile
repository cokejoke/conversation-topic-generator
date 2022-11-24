FROM golang:1.18-alpine

RUN apk add --no-cache gcc musl-dev

WORKDIR /app

VOLUME /app/data

COPY . .

RUN go install github.com/mattn/go-sqlite3
RUN go mod download

CMD ["go", "run", "."]