############################
# STEP 1 backend build base
############################
FROM golang:1.17-alpine3.13 as backend-build-base
RUN apk add --update --no-cache git ca-certificates build-base
WORKDIR /build
COPY go.mod .
COPY go.sum .
RUN go mod download -x

############################
# STEP 2 frontend build base
############################
FROM node:16-alpine as frontend-build-base
WORKDIR /app
COPY frontend/package.json .
COPY frontend/yarn.lock .
RUN yarn install --frozen-lockfile

############################
# STEP 3 build backend executable
############################
FROM backend-build-base AS backend-build
COPY . .
RUN CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -o /build/bin/box cmd/service/main.go

############################
# STEP 4 build frontend
############################
FROM frontend-build-base AS frontend-build
COPY frontend/ .
RUN yarn build

############################
# STEP 5 Finalize image
############################
FROM nginx:1.21-alpine
WORKDIR /app

# Copy backend server
COPY --from=backend-build /build/bin/box server

# Copy frontend
RUN rm -rf /usr/share/nginx/html/*
COPY nginx/ /etc/nginx/
COPY --from=frontend-build /app/build/ /usr/share/nginx/html