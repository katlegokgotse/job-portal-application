FROM golang:1.20

WORKDIR /backend

COPY . .

RUN go mod tidy
RUN go build -o main .

EXPOSE 8080

CMD ["./main"]