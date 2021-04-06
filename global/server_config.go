package global

import "strconv"

//Config stores configuration to start an http server
type Config struct {
	Port            port   `json:"port"`
	ShutdownTimeout int    `json:"shutdown_timeout"`
	Mode            string `json:"mode"`
	Name            string `json:"name"`
	MonitoringKey   string `json:"monitoring_key"`
}

type port int

func (p port) ToString() string {
	return ":" + strconv.Itoa(int(p))
}
