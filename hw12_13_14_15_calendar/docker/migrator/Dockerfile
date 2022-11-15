FROM golang:1.19 as build

WORKDIR /source

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o ./app/migrator ./cmd/migrator

FROM scratch

COPY --from=build source/app/migrator .
COPY --from=build source/configs/calendar_config.yaml configs/calendar_config.yaml
COPY --from=build source/migrations migrations

ENTRYPOINT ["/migrator"]



