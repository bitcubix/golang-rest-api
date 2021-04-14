package health

import "github.com/bitcubix/golang-rest-api/pkg/log"

type Status string

const (
	StatusOK  Status = "ok"
	StatusErr Status = "error"
)

type Service struct {
	logger log.Logger
}

func NewService(logger log.Logger) *Service {
	return &Service{
		logger: logger.WithPrefix("services").WithFields(log.Fields{"service": "health", "part": "service"}),
	}
}

func (s *Service) GetStatus() Status {
	// TODO check system

	status := StatusOK

	s.logger.Infof("server status: %s", status)

	return status
}
