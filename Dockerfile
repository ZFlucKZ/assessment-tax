FROM golang:1.22.2 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 go build -o assessment-tax

FROM alpine:3.19 AS runner

WORKDIR /app

COPY --from=builder /app/assessment-tax ./

CMD ["/app/assessment-tax"]