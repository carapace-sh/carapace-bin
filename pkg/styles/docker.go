package styles

import "github.com/rsteube/carapace/pkg/style"

var Docker = struct {
	Container string `desc:"docker containers"`
	Image     string `desc:"docker images"`
	Network   string `desc:"docker networks"`
	Node      string `desc:"docker nodes"`
	Secret    string `desc:"docker secrets"`
	Service   string `desc:"docker services"`
	Volume    string `desc:"docker volumes"`
}{
	Container: style.Blue,
	Image:     style.Yellow,
	Network:   style.Gray,
	Node:      style.Cyan,
	Secret:    style.Underlined,
	Service:   style.Green,
	Volume:    style.Bold,
}

func init() {
	style.Register("docker", &Docker)
}
