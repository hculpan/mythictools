# syntax=docker/dockerfile:1

FROM golang:1.20

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . ./

COPY .env ./

RUN go build -o /mythictools

EXPOSE 3000

CMD [ "/mythictools" ]