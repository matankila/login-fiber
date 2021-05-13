ARG  BUILDER_IMAGE=golang:alpine
############################
# STEP 1 build executable binary
############################
FROM ${BUILDER_IMAGE} as builder

RUN apk add --no-cache git make bash

# Install git + SSL ca certificates.
# Git is required for fetching the dependencies.
# Ca-certificates is required to call HTTPS endpoints.
RUN apk update && apk add --no-cache git ca-certificates tzdata && update-ca-certificates

# Create appuser
ENV USER=appuser
ENV UID=10001

# See https://stackoverflow.com/a/55757473/12429735
RUN adduser \
    --disabled-password \
    --gecos "" \
    --home "/nonexistent" \
    --shell "/sbin/nologin" \
    --no-create-home \
    --uid "${UID}" \
    "${USER}"

ENV GO111MODULE=on
ENV CGO_ENABLED 0
WORKDIR $GOPATH/src
COPY . .
RUN go mod download
RUN go mod verify
RUN go get -u github.com/swaggo/swag/cmd/swag
RUN swag init -g cmd/main.go

# Build the binary
RUN make build LOCAL=false

# Run Test & return coverage %
RUN make test

############################
# STEP 2 build a small image
############################
FROM scratch

# Import from builder.
COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /etc/group /etc/group

# Copy our static executable
COPY --from=builder /go/bin/main /go/bin/main

# Use an unprivileged user.
USER appuser:appuser

# Run the hello binary.
ENTRYPOINT ["/go/bin/main"]
EXPOSE 8080