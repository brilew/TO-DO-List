package config

import "os"

const (
	DefaultPort   = "7540"
	DefaultWebDir = "web"
)

type Config struct {
	Port   string // адрес для http.Server, например "[::]:7540"
	WebDir string
}

func Load() Config {
	port := os.Getenv("TODO_PORT")
	if port == "" {
		port = DefaultPort
	}

	// Убираем ведущее двоеточие, если оно есть, чтобы не получить "::7540"
	if port[0] == ':' {
		port = port[1:]
	}

	return Config{
		// "[::]" слушает и IPv4, и IPv6 — решает проблему с localhost на Windows,
		// где localhost резолвится в ::1.
		Port:   "[::]:" + port,
		WebDir: DefaultWebDir,
	}
}
