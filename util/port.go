package util

import (
	"fmt"
)

func HttpPort() string {
	port := "5000"

	if GetEnvVar("PORT") != "" {
		port = GetEnvVar("PORT")
	}
	return fmt.Sprintf(":%s", port)
}
