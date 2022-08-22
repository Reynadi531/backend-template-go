FROM golang:1.19 as builder

WORKDIR /app
COPY . .

RUN go mod download
RUN CGO_ENABLED=0 make build

FROM gcr.io/distroless/static-debian11
COPY --from=builder /app/bin/main /app
CMD ["/app"]
