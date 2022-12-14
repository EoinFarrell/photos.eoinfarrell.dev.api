FROM golang:1.18

WORKDIR /app

COPY go.mod go.mod
COPY go.sum go.sum
RUN go mod download

ADD . /app

RUN go build -o photos-eoinfarrelll-dev-api ./cmd/main.go

EXPOSE 8080

ENTRYPOINT [ "/app/photos-eoinfarrelll-dev-api" ]