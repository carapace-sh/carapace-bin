package nix

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"path"
	"strings"

	"github.com/rsteube/carapace"
)

//go:embed attributeComplete.nix
var attributeCompleteScript string

type AttributeOpts struct {
	Source  string
	Include string
	Arg     []string
	ArgStr  []string
}

// ActionAttributes completes attributes
//
//	firefox
//	git
func ActionAttributes(opts AttributeOpts) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		if opts.Source == "" {
			opts.Source = "default.nix"
		}

		attrPath := c.CallbackValue

		if strings.HasPrefix(opts.Source, "<") {
			// expression is a local channel, use as-is
		} else if strings.HasPrefix(opts.Source, "http:") || strings.HasPrefix(opts.Source, "https:") {
			// expression is a url, wrap in fetchTarball
			opts.Source = fmt.Sprintf("(fetchTarball %s)", opts.Source)
		} else if strings.HasPrefix(opts.Source, "channel:") {
			// expression is a channel alias, convert to url
			channelName := opts.Source[8:]
			opts.Source = fmt.Sprintf("(fetchTarball https://nixos.org/channels/%s/nixexprs.tar.xz)", channelName)
		} else {
			// expression is a local file
			filePath := path.Clean(opts.Source)
			if path.Base(filePath) == filePath {
				// paths must have at least one `/`
				filePath = fmt.Sprintf("./%s", filePath)
			}
			opts.Source = filePath
		}
		expression := fmt.Sprintf("import %s", opts.Source)

		// nix itself handles quotes weirdly so we just avoid them altogether
		if strings.Contains(attrPath, "\"") {
			return carapace.ActionMessage("attrPath may not contain double-quotes")
		}

		// strip right-most attr
		lastDotIndex := strings.LastIndex(attrPath, ".")
		if lastDotIndex >= 0 {
			attrPath = attrPath[0:lastDotIndex]
		} else {
			attrPath = ""
		}

		args := []string{
			"--eval", "--json",
			"--expr", attributeCompleteScript,
			"--arg", "__carapaceInput__", expression,
			"--argstr", "__carapaceAttrPath__", attrPath,
		}
		if opts.Include != "" {
			args = append(args, "-I", opts.Include)
		}

		// pass through provided --arg and --argstr
		for index, value := range opts.Arg {
			if index%2 == 0 {
				args = append(args, "--arg")
			}
			args = append(args, value)
		}
		for index, value := range opts.ArgStr {
			if index%2 == 0 {
				args = append(args, "--argstr")
			}
			args = append(args, value)
		}

		return carapace.ActionExecCommand("nix-instantiate", args...)(func(output []byte) carapace.Action {
			var data []string
			json.Unmarshal(output, &data)

			prefix := ""
			if attrPath != "" {
				prefix = fmt.Sprintf("%s.", attrPath)
			}

			return carapace.ActionValues(data...).Prefix(prefix).NoSpace()
		})
	})
}
