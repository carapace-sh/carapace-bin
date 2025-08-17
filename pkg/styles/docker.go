package styles

import "github.com/carapace-sh/carapace/pkg/style"

var Docker = struct {
	Container string `description:"docker containers"`
	Image     string `description:"docker images"`
	Network   string `description:"docker networks"`
	Node      string `description:"docker nodes"`
	Secret    string `description:"docker secrets"`
	Service   string `description:"docker services"`
	Volume    string `description:"docker volumes"`
}{
	Container: style.Blue, // TODO deprecated (now styled based on state)
	Image:     style.Yellow,
	Network:   style.Magenta,
	Node:      style.Cyan,
	Secret:    style.Underlined,
	Service:   style.Green,
	Volume:    style.Bold,
}

func init() {
	style.Register("docker", &Docker)
}
