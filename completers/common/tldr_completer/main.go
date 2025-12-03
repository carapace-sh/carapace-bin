package main

import (
	"os"

	tealdeer "github.com/carapace-sh/carapace-bin/completers/common/tldr_completer/tealdeer/cmd"
	tldr_python_client "github.com/carapace-sh/carapace-bin/completers/common/tldr_completer/tldr-python-client/cmd"
)

func main() {
	switch os.Getenv("CARAPACE_VARIANT") { // TODO implement generic way to choose completer (patch init script as well?)
	case "tealdeer":
		tealdeer.Execute() // TODO option pattern to modify cmd and maybe set an annotation in cobra for the variant?
	default:
		tldr_python_client.Execute()
	}
}
