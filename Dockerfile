FROM node:22-slim as frontend

WORKDIR /app/web
COPY web/package.json web/yarn.lock ./
RUN yarn install

COPY web .
RUN yarn build

FROM golang:1.23 as backend

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o app .

FROM debian:stable-slim

WORKDIR /app
RUN apt-get update && apt-get install -y ca-certificates && rm -rf /var/lib/apt/lists/*
COPY --from=backend /app/app .
COPY --from=frontend /app/web/dist ./web/dist

EXPOSE 8080

ENTRYPOINT ["./app"]
