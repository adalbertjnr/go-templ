FROM golang:bullseye as builder
WORKDIR /app
RUN go install github.com/a-h/templ/cmd/templ@latest
RUN templ generate
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o app

FROM alpine:3.18.2
WORKDIR /app
RUN addgroup --system alpinegroup && adduser --system alpineuser -g alpinegroup
RUN chown -R alpineuser:alpinegroup /app
USER alpineuser
COPY --from=builder /app .
EXPOSE 5000
CMD ["./app"]