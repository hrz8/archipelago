# Archipelago Chat using WebSocket

This project is a part of Technical Test of Software Engineering in Archipelago.

## Project Structure
```
.
├── Dockerfile
├── README.md
├── chat
│   ├── chat.go
│   └── handler.go
├── cmd
│   ├── root.go
│   ├── serve.go
│   └── version.go
├── config
│   └── config.go
├── docker
│   └── nginx.conf
├── docker-compose.prod.yaml
├── docker-compose.yaml
├── go.mod
├── go.sum
├── main.go
└── web
    ├── README.md
    ├── dist
    │   ├── assets
    │   │   ├── index-Dt1i8YJl.css
    │   │   └── index-vlisaxTd.js
    │   ├── index.html
    │   └── vite.svg
    ├── index.html
    ├── package.json
    ├── public
    │   └── vite.svg
    ├── src
    │   ├── App.vue
    │   ├── assets
    │   │   └── vue.svg
    │   ├── components
    │   │   └── ChatWindow.vue
    │   ├── main.js
    │   ├── scss
    │   │   └── style.scss
    │   └── style.css
    ├── vite.config.js
    └── yarn.lock
```

## Prerequisites

- **Golang** If you want to run in your machine without docker
- **NodeJS** If you want to run in your machine without docker (for build UI)
- **Docker** installed on your local machine

## Docker compose

Use Docker compose to run the server without manual setup (ref: `docker-compose.yaml`). You can start the app by following command:
```bash
docker compose up -d
```

🚀 Open your browser `http://localhost:8080/`. Voila! The dashboard is ready. 🚀 

## Manual Guide

You can also run the app using golang binary if you are already installed golang on your local machine.

```bash
# Setup frontend
cd web
yarn install
# Setup backend
cd ..
go mod tidy
# Run the server
go run . serve
```
