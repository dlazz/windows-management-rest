package services

type WindowsService struct {
	Name        string `json:"Name"`
	DisplayName string `json:"DisplayName"`
	Description string `json:"Description"`
	ProcessId   int    `json:"ProcessId"`
	Status      string `json:"Status"`
	StartMode   string `json:"StartMode"`
	State       string `json:"State"`
}
