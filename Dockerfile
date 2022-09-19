## ----------------------------------------------------------------------------
## Build
## ----------------------------------------------------------------------------

FROM golang:1.19.1-bullseye AS build

WORKDIR /app
COPY . .

ARG version
ENV VERSION ${version:-'0.0.0-develop'}
RUN make casino

## ----------------------------------------------------------------------------
## Deploy
## ----------------------------------------------------------------------------

FROM debian:11.5

WORKDIR /opt/CasinoGuest
COPY --from=build /app/dist/casinoguest .

RUN useradd nonroot --user-group --no-create-home
RUN chown -R nonroot:nonroot /opt/CasinoGuest && chmod -R 755 /opt/CasinoGuest

USER nonroot:nonroot

EXPOSE 8080
ENTRYPOINT ["/opt/CasinoGuest/casinoguest"]

