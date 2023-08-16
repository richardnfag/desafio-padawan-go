FROM golang:1.20 AS build

WORKDIR /usr/src/app

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .

RUN go build -v -o /usr/local/bin/app cmd/main.go
RUN curl -sSf https://atlasgo.sh | bash -s -- -y

EXPOSE 8000

CMD ["app"]
