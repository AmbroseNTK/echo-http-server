# build multistage

FROM golang:1.20-alpine AS builder

WORKDIR /

COPY . .

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /

COPY --from=builder /main .

CMD ["/main"]

EXPOSE 8080