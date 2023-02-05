FROM golang:1.20.0 as build
WORKDIR /app
COPY ./ .
RUN go build


FROM alpine:3.17.1

WORKDIR /app

RUN adduser -D rusty
RUN chown -R rusty /app
RUN chgrp -R rusty /app

COPY --from=build /app/rusty-chain /app

USER rusty
ENTRYPOINT [ "/app/rusty-chain" ]