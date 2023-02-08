package srvc

import (
	"github.com/kardianos/service"
)

var Config = &service.Config{
	Name:        "Wmr",
	DisplayName: "Windows Management Rest",
	Description: "Rest API for managing windows",
}

type srvc struct {
	runner func() error
}

func New(runner func() error) *srvc {
	return &srvc{
		runner: runner,
	}
}
func (s *srvc) Start(service.Service) error {
	// Start should not block. Do the actual work async.
	go s.run()
	return nil
}
func (s *srvc) run() {
	s.runner()
}
func (s *srvc) Stop(service.Service) error {
	// Stop should not block. Return with a few seconds.
	return nil
}
