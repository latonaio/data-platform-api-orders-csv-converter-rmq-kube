package config

import (
	"fmt"
	"os"
	"strconv"
)

type Conf struct {
	RMQ *RMQ
}

func NewConf() *Conf {
	return &Conf{
		RMQ: newRMQ(),
	}
}

func getEnv(key, fallback string) string {
	val := os.Getenv(key)
	if val == "" {
		val = fallback
	}
	return val
}

func getEnvInt(key string, fallback int) int {
	rawVal := os.Getenv(key)
	val, err := strconv.Atoi(rawVal)
	if err != nil {
		fmt.Fprintf(os.Stderr, "environment %s required number type: %+v", key, err)
		val = fallback
	}
	return val
}
