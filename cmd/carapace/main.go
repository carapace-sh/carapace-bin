package main

import "github.com/rsteube/carapace-bin/cmd/carapace/cmd"

var version = "develop"

// go:generate go run ../../carapace-generate/gen.go
func main() {
	cmd.Execute(version)
}
