## ----------------------------------------------------------------------------
## Build
## ----------------------------------------------------------------------------

FROM golang:1.19.1-bullseye AS build

WORKDIR /app
COPY . .

RUN go get ./...
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /casinoguest ./cmd/casinoguest

## ----------------------------------------------------------------------------
## Deploy
## ----------------------------------------------------------------------------

FROM debian:11.5

WORKDIR /opt/CasinoGuest
RUN useradd nonroot --user-group --no-create-home
RUN chown -R nonroot:nonroot /opt/CasinoGuest && chmod -R 755 /opt/CasinoGuest

USER nonroot:nonroot

COPY --from=build /casinoguest .

EXPOSE 8080
ENTRYPOINT ["/opt/CasinoGuest/casinoguest"]

