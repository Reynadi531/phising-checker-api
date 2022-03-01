FROM golang:1.17-alpine

RUN apk update && \
    apk add make git bash curl openssl alpine-sdk --no-cache

WORKDIR /app

COPY . .

RUN make build

ENV PORT 3000
EXPOSE 3000

CMD ["./bin/main"]
