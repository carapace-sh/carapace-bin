package cmd

import (
	"github.com/rsteube/carapace-bin/completers/http_completer/cmd"
)

func Execute() error {
	return cmd.ExecuteHttps()
}
