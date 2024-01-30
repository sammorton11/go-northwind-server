FROM golang:latest

WORKDIR /app

COPY . .

RUN go mod  download

RUN go build -o main .

# Copy the SQLite database file into the container
COPY northwind.db /usr/src/app/northwind.db


EXPOSE 8080

CMD ["./main"]
