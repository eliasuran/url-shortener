FROM golang:1.21.6

WORKDIR /app

COPY go.mod .
COPY api.go .

RUN go get
RUN go build -o bin .

ENTRYPOINT [ "/app/bin" ]
