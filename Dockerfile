FROM golang:alpine as builder

WORKDIR /paygo
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

#create new alpine image
FROM alpine:latest
RUN apk add --no-cache ca-certificates 

WORKDIR /paygo/
COPY --from=builder /paygo/config /paygo/config/
COPY --from=builder /paygo/main .

EXPOSE 8080
CMD ["./main"]

