package cmd

import (
	"github.com/rsteube/carapace-bin/completers/brotli_completer/cmd"
)

/**
Description for go:generate
	Use: "unbrotli",
	Short: "compress or decompress files",
*/

func Execute() error {
	return cmd.ExecuteUnbrotli()
}
