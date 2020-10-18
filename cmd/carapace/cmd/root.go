package cmd

//go:generate go run ../../generate/gen.go

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "carapace [flags] [COMPLETER] [bash|elvish|fish|oil|powershell|xonsh|zsh]",
	Short: `multi-shell multi-command argument completer`,
	Example: `
  bash:       source <(carapace _carapace)
  elvish:     eval (carapace _carapace|slurp)
  fish:       carapace _carapace | source
  oil:        source <(carapace _carapace)
  powershell: carapace _carapace | Out-String | Invoke-Expression
  xonsh:      exec($(carapace _carapace))
`,
	Args:      cobra.MinimumNArgs(1),
	ValidArgs: completers,
	Run: func(cmd *cobra.Command, args []string) {
		// since flag parsing is disabled do this manually
		switch args[0] {
		case "-h":
			cmd.Usage()
		case "--help":
			cmd.Usage()
		case "--list":
			fmt.Println(strings.Join(completers, "\n"))
		case "_carapace":
			switch determineShell() {
			case "bash":
				fmt.Println(bash(completers))
			case "elvish":
				fmt.Println(elvish(completers))
			case "fish":
				fmt.Println(fish(completers))
			case "oil":
				fmt.Println(bash(completers))
			case "powershell":
				fmt.Println(powershell(completers))
			case "xonsh":
				fmt.Println(xonsh(completers))
			case "zsh":
				fmt.Println(zsh(completers))
			}
			// TODO lazy completion script for all completers
		default:
			invokeCompleter(args[0])
		}

	},
	FParseErrWhitelist: cobra.FParseErrWhitelist{
		UnknownFlags: true,
	},
	DisableFlagParsing: true,
}

func invokeCompleter(completer string) {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	outC := make(chan string)
	// copy the output in a separate goroutine so printing can't block indefinitely
	go func() {
		var buf bytes.Buffer
		io.Copy(&buf, r)
		outC <- buf.String()
	}()

	os.Args[1] = "_carapace"
	executeCompleter(completer)

	w.Close()
	out := <-outC
	os.Stdout = old
	patched := strings.Replace(string(out), "carapace _carapace", "carapace "+completer, -1)         // general callback
	patched = strings.Replace(patched, "'carapace', '_carapace'", "'carapace', '"+completer+"'", -1) // xonsh callback
	fmt.Print(patched)

}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.Flags().Bool("list", false, "list completers")
}
