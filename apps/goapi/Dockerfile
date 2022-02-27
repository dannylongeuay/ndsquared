ARG  BUILDER_IMAGE=golang:buster
ARG  DISTROLESS_IMAGE=gcr.io/distroless/base
############################
# STEP 1 build executable binary
############################
FROM ${BUILDER_IMAGE} as builder

# Ensure ca-certficates are up to date
RUN update-ca-certificates

WORKDIR $GOPATH/ndsquared/goapi/

COPY go.mod .

ENV GO111MODULE=on
RUN go mod download
RUN go mod verify

COPY . .

# Generate swagger docs
RUN go install github.com/swaggo/swag/cmd/swag@latest
RUN swag init -d src/ -ot go

# Build the binary
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -a -installsuffix cgo -o /go/bin/goapi ndsquared/goapi/src

############################
# STEP 2 build a small image
############################
# using base nonroot image
# user:group is nobody:nobody, uid:gid = 65534:65534
FROM ${DISTROLESS_IMAGE}

# Copy our static executable
COPY --from=builder /go/bin/goapi /go/bin/goapi

# Run the hello binary.
ENTRYPOINT ["/go/bin/goapi"]