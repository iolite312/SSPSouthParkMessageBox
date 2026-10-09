FROM golang:1.26 AS base
WORKDIR /app/api

FROM base as deps
COPY api/go.mod api/go.sum ./
RUN go mod download

FROM deps as build
COPY api/ ./
RUN CGO_ENABLED=0 GOOS=linux go build -o /ssp-sp-messaging "./cmd"

FROM alpine:3.24 as runner
COPY --from=build /ssp-sp-messaging /ssp-sp-messaging
EXPOSE 8080

# Run
CMD ["/ssp-sp-messaging"]