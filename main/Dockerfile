FROM golang:latest

WORKDIR /main

COPY . .

RUN go mod download

RUN go build -o main main.go

CMD [ "./main" ]