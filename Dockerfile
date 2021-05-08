FROM golang:latest AS build_base
WORKDIR /src
COPY . .
RUN make build-server

FROM alpine:latest
COPY --from=build_base /src/bin/pgmeta /pgmeta/server
EXPOSE 8080
CMD ["/pgmeta/server"]
