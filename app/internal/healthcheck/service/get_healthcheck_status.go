package healthcheck

import (
	"id-manager/config"
	"net/http"
	"os"
)

type Status struct {
	Status   int    `json:"status"`
	Hostname string `json:"hostname"`
	Service  string `json:"service"`
}

func GetHealthcheckStatus(c config.Config) (int, Status) {
	status := http.StatusOK

	return status, Status{
		Status:   status,
		Hostname: getHostName(),
		Service:  c.ServiceName,
	}
}

func getHostName() string {
	hostname, err := os.Hostname()
	if err != nil {
		return "UNKNOWN_HOST"
	}

	return hostname
}
