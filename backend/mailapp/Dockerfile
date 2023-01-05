FROM golang:1.18-alpine

WORKDIR /mailapp
COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN go build -v

EXPOSE 3000

CMD ["./mailapp"]