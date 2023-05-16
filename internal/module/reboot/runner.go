package reboot

import (
	"fmt"

	"github.com/dlazz/windows-management-rest/internal/executor"
)

type Runner struct {
}

func NewRunner() *Runner {
	return &Runner{}
}

func (r *Runner) reboot() (message string, err error) {
	res := struct {
		Ok      bool   `json:"Ok"`
		Message string `json:"Message"`
		Error   string `json:"Error"`
	}{}

	if err := executor.Wrap(rebootComputer, &res); err != nil {
		return message, err
	}

	if !res.Ok {
		return message, fmt.Errorf(fmt.Sprint(res.Error))
	}

	return res.Message, nil
}
