FROM gcr.io/distroless/static-debian12

WORKDIR /app
COPY ./cart/bin/app .
COPY ./cart/config.yaml .
COPY .env .


ENTRYPOINT ["./app"]
