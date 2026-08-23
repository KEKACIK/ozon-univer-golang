FROM scratch

WORKDIR /app
COPY ./loms/bin/app .
COPY ./loms/config.yaml .
COPY .env .


ENTRYPOINT ["./app"]
