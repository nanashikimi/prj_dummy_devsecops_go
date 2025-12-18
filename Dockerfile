FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY go.mod ./
RUN go mod download || true
#--Trying to download dependencies, but don't fail, if such don't exist--

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o server .

# --------------------
FROM alpine:3.23

WORKDIR /app

COPY --from=builder /app/server .

RUN adduser -D -u 10001 appuser
USER appuser

EXPOSE 8080

CMD ["./server"]
