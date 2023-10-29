FROM golang:latest

WORKDIR /app

COPY . .

RUN go build -o myapp

EXPOSE 80

CMD ["./myapp"]
