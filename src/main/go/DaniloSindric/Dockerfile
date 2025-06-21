# syntax=docker/dockerfile:1
ARG GO_VERSION=1.24.0
FROM --platform=$BUILDPLATFORM golang:${GO_VERSION} AS build
COPY . src/
RUN cd src && go build .

FROM scratch AS final
COPY --from=build /go/src/1brc /
