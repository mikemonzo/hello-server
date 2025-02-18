FROM golang:1.24.0-alpine

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

ENV GO111MODULE=on

RUN go build -o hello-server .

EXPOSE 9000

CMD [ "./hello-server" ]