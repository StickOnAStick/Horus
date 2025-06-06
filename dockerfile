FROM golang:1.24.3-bookworm AS build-stage
# Set dest for copy
WORKDIR /app
COPY go.mod ./
# Download go modules
RUN go mod download
COPY *.go ./
#Build
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o /horus

# Second stage : run
FROM gcr.io/distroless/static:nonroot AS build-release-stage
# Copy the build files over to a new dir in this second image
COPY --from=build-stage /horus /horus
# Copy static files
COPY  static /static
EXPOSE 8080
USER nonroot:nonroot
ENTRYPOINT [ "/horus" ]