FROM golang:1.19 as build

WORKDIR /source

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o ./app/calendar ./cmd/calendar

FROM scratch

EXPOSE 8080
EXPOSE 8081

COPY --from=build /usr/share/zoneinfo /usr/share/zoneinfo

COPY --from=build source/app/calendar .
COPY --from=build source/configs/calendar_config.yaml configs/calendar_config.yaml

ENTRYPOINT ["/calendar"]



