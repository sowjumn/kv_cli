FROM golang:1.17

RUN mkdir -p /github.com/sowjumn/interview/devoted

WORKDIR /github.com/sowjumn/interview/devoted

COPY . .

RUN go get -d -v ./...

RUN go install -v ./...

RUN go build -o kv_cli .

CMD ["./kv_cli"]