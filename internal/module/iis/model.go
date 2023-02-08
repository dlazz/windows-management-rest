package iis

type Website struct {
	Name            string `json:"name"`
	ID              int    `json:"id"`
	State           string `json:"state"`
	PhysicalPath    string `json:"shysicalPath"`
	ApplicationPool string `json:"applicationPool"`
}

type AppPool struct {
	Name                  string `json:"name"`
	State                 string `json:"state"`
	ManagedRunTimeVersion string `json:"managedRunTimeVersion"`
}
