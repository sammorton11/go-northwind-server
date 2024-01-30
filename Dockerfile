FROM golang:latest

WORKDIR /app

COPY . .

RUN go mod  download

RUN go build -o main .


# Set the path for the SQLite database file
ENV DB_PATH /usr/src/app/northwind.db

# Check if the file exists before copying it
RUN test -e $DB_PATH || (mkdir -p $(dirname $DB_PATH) && cp ./db/northwind.db $DB_PATH)

# Expose port 8080
EXPOSE 8080

# Run the application
CMD ["./main"]

