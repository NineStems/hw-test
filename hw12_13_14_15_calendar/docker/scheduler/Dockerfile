FROM golang:1.19 as build

WORKDIR /source

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o ./app/calendar_scheduler ./cmd/scheduler

FROM scratch

COPY --from=build source/app/calendar_scheduler .
COPY --from=build source/configs/scheduler_config.yaml configs/scheduler_config.yaml

ENTRYPOINT ["/calendar_scheduler"]



