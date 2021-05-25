package main

import "github.com/rsteube/carapace-bin/cmd/carapace/cmd"

var version = "develop"

// go:generate go run ../../generate/gen.go
func main() {
	cmd.Execute(version)
}
