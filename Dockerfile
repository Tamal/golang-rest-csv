# Build
FROM golang:1.16-alpine AS build

WORKDIR /app

COPY . .

# RUN go mod download
RUN go build -o emp_service ./cmd/emp


# Run

FROM alpine:3.13

WORKDIR /
COPY --from=build /app/emp_service .

EXPOSE 8082

ENTRYPOINT [ "/emp_service" ]