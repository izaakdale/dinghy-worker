FROM golang:1.20-alpine as builder
WORKDIR /

COPY . .
RUN go mod download

RUN go build -o dinghy-worker .

# CMD [ "/dinghy-worker" ]

FROM alpine
WORKDIR /
COPY --from=builder /dinghy-worker /

CMD [ "/dinghy-worker" ]