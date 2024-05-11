# Build Stage
FROM golang:1.20 AS build-stage

LABEL app="build-InternalTransfersSystem"
LABEL REPO="https://github.com/wdevarshi/InternalTransfersSystem"

ENV PROJPATH=/go/src/github.com/wdevarshi/InternalTransfersSystem

# Because of https://github.com/docker/docker/issues/14914
ENV PATH=$PATH:$GOROOT/bin:$GOPATH/bin

ADD . /go/src/github.com/wdevarshi/InternalTransfersSystem
WORKDIR /go/src/github.com/wdevarshi/InternalTransfersSystem

RUN make build-alpine

# Final Stage
FROM alpine:latest

ARG GIT_COMMIT
ARG VERSION
LABEL REPO="https://github.com/wdevarshi/InternalTransfersSystem"
LABEL GIT_COMMIT=$GIT_COMMIT
LABEL VERSION=$VERSION

# add tz data
RUN apk add --no-cache tzdata

# Because of https://github.com/docker/docker/issues/14914
ENV PATH=$PATH:/opt/InternalTransfersSystem/bin

WORKDIR /opt/InternalTransfersSystem/bin

COPY --from=build-stage /go/src/github.com/wdevarshi/InternalTransfersSystem/bin/InternalTransfersSystem /opt/InternalTransfersSystem/bin/
RUN chmod +x /opt/InternalTransfersSystem/bin/InternalTransfersSystem

# Create appuser
RUN adduser -D -g '' InternalTransfersSystem
USER InternalTransfersSystem

ENTRYPOINT ["/opt/InternalTransfersSystem/bin/InternalTransfersSystem"]
