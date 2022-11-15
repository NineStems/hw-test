FROM golang:1.19 as build

WORKDIR /source

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o ./app/sender_calendar ./cmd/sender

FROM scratch

COPY --from=build source/app/sender_calendar .
COPY --from=build source/configs/sender_config.yaml configs/sender_config.yaml

ENTRYPOINT ["/sender_calendar"]



