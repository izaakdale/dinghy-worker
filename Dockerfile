FROM golang:1.20-alpine as builder
WORKDIR /

COPY . .
RUN go mod download

RUN go build -o dinghy-worker .

FROM scratch
WORKDIR /bin
COPY --from=builder /dinghy-worker /bin

CMD [ "/bin/dinghy-worker" ]