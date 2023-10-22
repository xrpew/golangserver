FROM golang:1.21

WORKDIR /app

COPY . .

RUN go build -o mi_aplicacion

EXPOSE 8081

CMD ["./mi_aplicacion"]
