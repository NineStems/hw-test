FROM golang:1.19 as build

WORKDIR /source

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o ./app/sender_calendar ./cmd/sender

FROM scratch

EXPOSE 8080

COPY --from=build source/app/calendar .
COPY --from=build source/configs/sender_config.yaml configs/sender_config.yaml

ENTRYPOINT ["/calendar"]



