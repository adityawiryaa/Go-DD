#builder
FROM golang:1.21-alpine AS builder

ENV GOOS=linux
ENV GOARCH=amd64

WORKDIR /home
COPY . .
RUN go build -o go-driven-design cmd/main.go

FROM alpine:latest AS runner

RUN apk add tzdata
COPY --from=builder /home/go-driven-design .
EXPOSE 6001
CMD ./go-driven-design