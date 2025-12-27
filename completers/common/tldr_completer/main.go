package main

import (
	tealdeer "github.com/carapace-sh/carapace-bin/completers/common/tldr_completer/tealdeer/cmd"
	tldr_python_client "github.com/carapace-sh/carapace-bin/completers/common/tldr_completer/tldr-python-client/cmd"
	"github.com/carapace-sh/carapace-bridge/pkg/choices"
)

func main() {
	switch choice, _ := choices.Get("tldr"); {
	case choice.Match("tldr/tealdeer@common"):
		tealdeer.Execute()
	case choice.Match("tldr/tldr-python-client@common"):
		tldr_python_client.Execute()
	default:
		tealdeer.Execute()
	}
}
