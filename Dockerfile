# First stage: Get Golang image from DockerHub.
FROM golang:1.23.1@sha256:2fe82a3f3e006b4f2a316c6a21f62b66e1330ae211d039bb8d1128e12ed57bf1 AS backend-builder

# Set our working directory for this stage.
WORKDIR /backendcompile

# Copy all of our files.
COPY . .

# Get and install all dependencies.
RUN CGO_ENABLED=0 go build -o abdimasa ./cmd/app/main.go

# Last stage: discard everything except our executables.
FROM alpine:latest AS prod

# Set our next working directory.
WORKDIR /build

# Copy our executable and our built React application.
COPY --from=backend-builder /backendcompile/abdimasa .

# Declare entrypoints and activation commands.
EXPOSE 3003
ENTRYPOINT ["./abdimasa"]