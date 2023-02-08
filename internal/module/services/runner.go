package services

import (
	"encoding/json"
	"fmt"

	"github.com/dlazz/windows-management-rest/internal/executor"
)

type Runner struct {
}

func NewRunner() *Runner {
	return &Runner{}
}

func (r *Runner) GetServices() ([]*WindowsService, error) {
	res := struct {
		Ok      bool              `json:"Ok"`
		Message []*WindowsService `json:"Message"`
		Error   string            `json:"Error"`
	}{}

	if err := executor.Wrap(getServiceList, &res); err != nil {
		return nil, err
	}

	if !res.Ok {
		return nil, fmt.Errorf(fmt.Sprint(res.Error))
	}

	return res.Message, nil
}

func (r *Runner) GetService(name string) (*WindowsService, error) {
	res := struct {
		Ok      bool            `json:"Ok"`
		Message *WindowsService `json:"Message"`
		Error   string          `json:"Error"`
	}{}

	if err := executor.Wrap(fmt.Sprintf(getService, name), &res); err != nil {
		return nil, err
	}

	if !res.Ok {
		return nil, fmt.Errorf(fmt.Sprint(res.Error))
	}

	return res.Message, nil
}

func (r *Runner) StopService(name string) (err error) {
	res := struct {
		Ok      bool   `json:"Ok"`
		Message string `json:"Message"`
		Error   string `json:"Error"`
	}{}

	out, err := executor.RunPowershellCommand(fmt.Sprintf(stopService, name))
	if err != nil {
		return err
	}
	if err := json.Unmarshal(out, &res); err != nil {
		return err
	}
	if !res.Ok {
		return fmt.Errorf(fmt.Sprint(res.Error))
	}

	return nil
}

func (r *Runner) StartService(name string) (err error) {
	res := struct {
		Ok      bool   `json:"Ok"`
		Message string `json:"Message"`
		Error   string `json:"Error"`
	}{}

	out, err := executor.RunPowershellCommand(fmt.Sprintf(startService, name))
	if err != nil {
		return err
	}
	if err := json.Unmarshal(out, &res); err != nil {
		return err
	}
	if !res.Ok {
		return fmt.Errorf(fmt.Sprint(res.Error))
	}

	return nil
}

func (r *Runner) RestartService(name string) (err error) {
	res := struct {
		Ok      bool   `json:"Ok"`
		Message string `json:"Message"`
		Error   string `json:"Error"`
	}{}

	out, err := executor.RunPowershellCommand(fmt.Sprintf(restartService, name))
	if err != nil {
		return err
	}
	if err := json.Unmarshal(out, &res); err != nil {
		return err
	}
	if !res.Ok {
		return fmt.Errorf(fmt.Sprint(res.Error))
	}

	return nil
}
