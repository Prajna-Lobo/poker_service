FROM golang:1.19

WORKDIR /app

COPY . ./

RUN go mod download

RUN go build -o /poker_service app/main.go

EXPOSE 8080

CMD [ "/poker_service" ]