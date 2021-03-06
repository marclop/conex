// Package conex provides easy to use Docker Integration with Testing.
package conex

import (
	"testing"
)

// We keep logger here because the filename is show along the logs, this means
// conex.go is put before each log makes the source more clear to the user.
func logf(t *testing.T, f string, args ...interface{}) {
	t.Logf(f, args...)
}

//Same story as above.
func fatalf(t *testing.T, f string, args ...interface{}) {
	t.Fatalf(f, args...)
}

// Manager the conex container manager.
type Manager interface {
	Run(m *testing.M, images ...string) int
	Box(t *testing.T, config *Config) Container
}

// Container is a simple interface to a docker
// container.
type Container interface {
	ID() string
	Name() string
	Image() string
	Address() string

	Drop()

	//TODO: Yo.
	// Ports() []string
}

// Config contains the configuration data about a container.
type Config struct {
	Image      string   // Name of the image as it was passed by the operator (e.g. could be symbolic)
	Env        []string // List of environment variable to set in the container
	Cmd        []string // Command to run when starting the container
	Hostname   string   // Hostname
	Domainname string   // Domainname
	User       string   // User that will run the command(s) inside the container, also support user:group
}
