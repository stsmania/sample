FROM golang:1.21
ENV GIN_MODE release

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . ./

EXPOSE 8080

CMD ["go", "run", "main.go"]
