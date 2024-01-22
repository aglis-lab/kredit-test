# Build the application from source
FROM golang:1.21.6 AS build-stage

WORKDIR /app

COPY ./ ./
RUN go mod download

RUN go build -o /kredit ./cmd/app/main.go

# Run the tests in the container
FROM build-stage AS run-test-stage
RUN go test -v ./...

# Deploy the application binary into a lean image
FROM ubuntu:22.04 AS build-release-stage

WORKDIR /

COPY --from=build-stage /kredit /kredit

EXPOSE 5005

USER nonroot:nonroot

ENTRYPOINT ["/kredit"]
