# Build stage
FROM golang:1.23.4-alpine3.21 AS builder
WORKDIR /app
COPY . .
RUN apk --no-cache add curl
RUN apk --no-cache add make
RUN make dependencies
RUN ls
RUN go mod download
RUN make build

# Run stage
FROM alpine:3.21
WORKDIR /app
COPY --from=builder /app/blog .
COPY --from=builder /app/content ./content
COPY .env .
COPY posts.json .

ENTRYPOINT [ "/app/blog" ]
EXPOSE 80