package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/cmd/carapace/cmd/action"
	"github.com/rsteube/carapace-bridge/pkg/actions/bridge"
	shlex "github.com/rsteube/carapace-shlex"
	"github.com/rsteube/carapace/pkg/style"
	"github.com/spf13/cobra"
)

var diffCmd = &cobra.Command{
	Use:   "--diff",
	Short: "",
	Args:  cobra.MinimumNArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		var stdout, stderr bytes.Buffer
		rootCmd := cmd.Root()
		rootCmd.SetArgs(append([]string{"_carapace", "export", ""}, args...))
		rootCmd.SetOut(&stdout)
		rootCmd.SetErr(&stderr)
		rootCmd.Execute()

		var export struct {
			Values []struct {
				Value       string `json:"value"`
				Description string `json:"description,omitempty"`
				Style       string `json:"style,omitempty"`
			}
		}
		if err := json.Unmarshal(stdout.Bytes(), &export); err != nil {
			return err
		}

		command, _bridge, _ := strings.Cut(args[0], "/")
		macro := "bridge.Zsh"
		switch _bridge {
		case "argcomplete":
			macro = "bridge.Argcomplete"
		case "bash":
			macro = "bridge.Bash"
		case "carapace":
			macro = "bridge.Carapace"
		case "clap":
			macro = "bridge.CarapaceBin"
		case "click":
			macro = "bridge.Click"
		case "cobra":
			macro = "bridge.Cobra"
		case "complete":
			macro = "bridge.Complete"
		case "fish":
			macro = "bridge.Fish"
		case "inshellisense":
			macro = "bridge.Inshellisense"
		case "kingpin":
			macro = "bridge.Kingpin"
		case "powershell":
			macro = "bridge.Powershell"
		case "urfavecli":
			macro = "bridge.Urfavecli"
		case "yargs":
			macro = "bridge.Yargs"
		default:
			_bridge = "zsh"

		}
		joined := shlex.Join(args)
		fmt.Printf("diff --git a/%v b/carapace\n", _bridge)
		fmt.Printf("--- a/%v/carapace --macro %v %v\n", _bridge, macro, joined) // TODO macro by _bridge
		fmt.Printf("+++ b/carapace/carapace %v export %v\n", command, joined)
		for _, v := range export.Values {
			s := v.Value
			if v.Description != "" {
				s += fmt.Sprintf(" (%v)", v.Description)
			}

			switch v.Style {
			case style.Red:
				s = "- " + s
			case style.Green:
				s = "+ " + s
			default:
				s = "  " + s
			}

			fmt.Println(s)
		}
		return nil
	},
}

func init() {
	carapace.Gen(diffCmd).Standalone()
	diffCmd.Flags().SetInterspersed(false)

	carapace.Gen(diffCmd).PositionalCompletion(
		carapace.ActionMultiPartsN("/", 2, func(c carapace.Context) carapace.Action {
			switch len(c.Parts) {
			case 0:
				return action.ActionCompleters(action.CompleterOpts{Internal: true, Spec: true})
			default:
				return carapace.ActionStyledValues(
					"argcomplete", style.Default,
					"bash", "#d35673",
					"carapace", style.Default,
					"clap", style.Default,
					"click", style.Default,
					"cobra", style.Default,
					"complete", style.Default,
					"fish", "#7ea8fc",
					"inshellisense", style.Default,
					"kingpin", style.Default,
					"macro", style.Default,
					"powershell", "#e8a16f",
					"urfavecli", style.Default,
					"yargs", style.Default,
					"zsh", "#efda53",
				)
			}
		}),
	)

	carapace.Gen(diffCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			command, _bridge, _ := strings.Cut(c.Args[0], "/")
			return carapace.Diff(
				carapace.ActionCallback(func(c carapace.Context) carapace.Action {
					switch _bridge {
					case "argcomplete":
						return bridge.ActionArgcomplete(command)
					case "bash":
						return bridge.ActionBash(command)
					case "carapace":
						return bridge.ActionCarapace(command)
					case "clap":
						return bridge.ActionClap(command)
					case "click":
						return bridge.ActionClick(command)
					case "cobra":
						return bridge.ActionCobra(command)
					case "complete":
						return bridge.ActionComplete(command)
					case "fish":
						return bridge.ActionFish(command)
					case "inshellisense":
						return bridge.ActionInshellisense(command)
					case "kingpin":
						return bridge.ActionKingpin(command)
					case "powershell":
						return bridge.ActionPowershell(command)
					case "urfavecli":
						return bridge.ActionUrfavecli(command)
					case "yargs":
						return bridge.ActionYargs(command)
					default:
						return bridge.ActionZsh(command)
					}
				}),
				bridge.ActionCarapaceBin(command),
			).Shift(1)
		}),
	)
}
