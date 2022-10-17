package grpc

import (
	"net"

	v1 "github.com/calendar/hw12_13_14_15_calendar/api/v1"
	"github.com/calendar/hw12_13_14_15_calendar/common"
	"github.com/calendar/hw12_13_14_15_calendar/internal/config"
)

type Server struct {
	v1.UnimplementedCalendarServer
	cfg *config.Config
	log common.Logger
	ln  net.Listener
}
