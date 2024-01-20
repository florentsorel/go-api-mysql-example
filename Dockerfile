FROM golang:1.22 as builder

WORKDIR /app
COPY go.* ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /tmp/build/api ./cmd/api

FROM gcr.io/distroless/static-debian12:nonroot

ENV MYSQL_DSN "test-user:test-password@tcp(database:3306)/test?parseTime=true"

COPY --from=builder /tmp/build/api /
CMD ["/api"]
