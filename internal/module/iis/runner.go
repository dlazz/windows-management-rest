package iis

import (
	"fmt"

	"github.com/dlazz/windows-management-rest/internal/executor"
)

type Runner struct {
}

func NewRunner() *Runner {
	return &Runner{}
}

func (r *Runner) GetWebsiteList() ([]*Website, error) {
	res := struct {
		Ok      bool       `json:"Ok"`
		Message []*Website `json:"Message"`
		Error   string     `json:"Error"`
	}{}

	if err := executor.Wrap(getWebSiteList, &res); err != nil {
		return nil, err
	}
	fmt.Println(res)
	if !res.Ok {
		return nil, fmt.Errorf(fmt.Sprint(res.Error))
	}
	return res.Message, nil
}

func (r *Runner) GetWebsite(name string) (*Website, error) {
	res := struct {
		Ok      bool     `json:"Ok"`
		Message *Website `json:"Message"`
		Error   string   `json:"Error"`
	}{}

	if err := executor.Wrap(fmt.Sprintf(getWebSite, name), &res); err != nil {
		return nil, err
	}
	if !res.Ok {
		return nil, fmt.Errorf(fmt.Sprint(res.Error))
	}
	return res.Message, nil
}

func (r *Runner) StopWebsite(name string) (string, error) {
	res := struct {
		Ok      bool   `json:"Ok"`
		Message string `json:"Message"`
		Error   string `json:"Error"`
	}{}

	if err := executor.Wrap(fmt.Sprintf(stopWebSite, name), &res); err != nil {
		return "", err
	}
	if !res.Ok {
		return "", fmt.Errorf(fmt.Sprint(res.Error))
	}
	return res.Message, nil
}

func (r *Runner) StartWebsite(name string) (string, error) {
	res := struct {
		Ok      bool   `json:"Ok"`
		Message string `json:"Message"`
		Error   string `json:"Error"`
	}{}
	if err := executor.Wrap(fmt.Sprintf(startWebSite, name), &res); err != nil {
		return "", err
	}
	if !res.Ok {
		return "", fmt.Errorf(fmt.Sprint(res.Error))
	}
	return res.Message, nil
}

func (r *Runner) GetAppPoolList() ([]*AppPool, error) {
	res := struct {
		Ok      bool       `json:"Ok"`
		Message []*AppPool `json:"Message"`
		Error   string     `json:"Error"`
	}{}

	if err := executor.Wrap(getAppPoolList, &res); err != nil {
		return nil, err
	}
	fmt.Println(res)
	if !res.Ok {
		return nil, fmt.Errorf(fmt.Sprint(res.Error))
	}
	return res.Message, nil
}

func (r *Runner) StartWebAppPool(name string) (string, error) {
	res := struct {
		Ok      bool   `json:"Ok"`
		Message string `json:"Message"`
		Error   string `json:"Error"`
	}{}
	if err := executor.Wrap(fmt.Sprintf(startWebAppPool, name), &res); err != nil {
		return "", err
	}
	if !res.Ok {
		return "", fmt.Errorf(fmt.Sprint(res.Error))
	}
	return res.Message, nil
}

func (r *Runner) StopWebAppPool(name string) (string, error) {
	res := struct {
		Ok      bool   `json:"Ok"`
		Message string `json:"Message"`
		Error   string `json:"Error"`
	}{}
	if err := executor.Wrap(fmt.Sprintf(stopWebAppPool, name), &res); err != nil {
		return "", err
	}
	if !res.Ok {
		return "", fmt.Errorf(fmt.Sprint(res.Error))
	}
	return res.Message, nil
}
