// Package plugin contains the Traefik plugin for adding headers based on the
// [net/http.Request.RemoteAddr] field.
package plugin

import (
	"context"
	"errors"
	"net/http"
	"strings"
)

var errMissingHeaderConfig = errors.New("missing header config: must set at least one of headers.port, headers.ip, or headers.address")

// Config the plugin configuration.
type Config struct {
	Headers ConfigHeaders `json:"headers,omitempty"`
}

// ConfigHeaders defines the headers to use for the different values.
type ConfigHeaders struct {
	Port    string `json:"port,omitempty"`
	IP      string `json:"ip,omitempty"`
	Address string `json:"address,omitempty"`
}

// CreateConfig creates the default plugin configuration.
func CreateConfig() *Config {
	return &Config{
		Headers: ConfigHeaders{},
	}
}

// RemoteAddrPlugin is the main handler model for this Traefik plugin.
type RemoteAddrPlugin struct {
	next    http.Handler
	headers ConfigHeaders
	name    string
}

// New created a new RemoteAddrPlugin.
func New(_ context.Context, next http.Handler, config *Config, name string) (http.Handler, error) {
	if config.Headers == (ConfigHeaders{}) {
		return nil, errMissingHeaderConfig
	}

	return &RemoteAddrPlugin{
		headers: config.Headers,
		next:    next,
		name:    name,
	}, nil
}

func (a *RemoteAddrPlugin) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	//ip, port, _ := strings.Cut(req.RemoteAddr, ":")
        parts := strings.Split(reqRemoteAddr, ":")

        if len(parts) == 2 {
                ip := parts[0]
                port := parts[1]
                fmt.Printf("IP: %s, Port: %s\n", ip, port)
        } else {
                fmt.Println("Formato inválido")
        }

	if a.headers.IP != "" {
		req.Header.Set(a.headers.IP, ip)
	}
	if a.headers.Port != "" {
		req.Header.Set(a.headers.Port, port)
	}
	if a.headers.Address != "" {
		req.Header.Set(a.headers.Address, req.RemoteAddr)
	}

	a.next.ServeHTTP(rw, req)
}
