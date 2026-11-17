package common

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

func AddServerFlags(cmd *cobra.Command) {
	cmd.Flags().Bool("allow-anonymous", false, "admit anonymous callers on plain TCP listeners")
	cmd.Flags().StringArray("block-dir", nil, "discover executable block plugins in DIR (repeatable)")
	cmd.Flags().Bool("block-plugin-reload", false, "reload block plugins after source changes")
	cmd.Flags().String("config", "", "read the config file at PATH")
	cmd.Flags().Bool("exit-empty", false, "stop when no sessions or control connections remain")
	cmd.Flags().StringArray("extra-listen", nil, "also listen on ENDPOINT (repeatable)")
	cmd.Flags().String("listen", "", "listen on ENDPOINT")
	cmd.Flags().String("log-format", "text", "write server logs as FORMAT: text or json")
	cmd.Flags().String("log-level", "", "log at LEVEL and above: debug, info, warn, or error")
	cmd.Flags().String("max-blocks", "1024", "refuse new blocks once the server holds N across all sessions")
	cmd.Flags().String("max-panes", "64", "refuse new panes once a window holds N")
	cmd.Flags().String("max-sessions", "64", "refuse new sessions once the server holds N")
	cmd.Flags().String("max-windows", "64", "refuse new windows once a session holds N")
	cmd.Flags().Bool("restore", true, "restore the saved sessions")
	cmd.Flags().String("socket-mode", "0600", "set Unix socket permissions to MODE")
	cmd.Flags().String("state-dir", "", "store persistent server state below DIR")
	cmd.Flags().Bool("tailnet-allow-identified", false, "admit every identified caller through the host's tailscaled")
	cmd.Flags().Bool("tailscale", false, "also serve the API over the tailnet as an embedded tsnet node")
	cmd.Flags().StringArray("tailscale-advertise-tag", nil, "advertise ACL TAG when --tailscale enrolls the node (repeatable)")
	cmd.Flags().Bool("tailscale-allow-tailnet-members", false, "allow any tailnet member permitted by ACLs to reach Rex")
	cmd.Flags().String("tailscale-auth-key-file", "", "enroll --tailscale with the one-time auth key read from FILE")
	cmd.Flags().String("tailscale-control-url", "", "register the tailnet node with the control server at URL")
	cmd.Flags().String("tailscale-hostname", "rex", "tailnet device NAME for --tailscale")
	cmd.Flags().String("tailscaled-socket", "", "use the tailscaled LocalAPI socket at PATH")
	cmd.Flags().StringArray("websocket-origin-pattern", nil, "allow browser WebSockets from cross-origin PATTERN (repeatable)")

	carapace.Gen(cmd).FlagCompletion(carapace.ActionMap{
		"block-dir":               carapace.ActionDirectories(),
		"config":                  carapace.ActionFiles(),
		"log-format":              carapace.ActionValues("json", "text"),
		"log-level":               carapace.ActionValues("debug", "error", "info", "warn"),
		"state-dir":               carapace.ActionDirectories(),
		"tailscale-auth-key-file": carapace.Batch(carapace.ActionValues("-"), carapace.ActionFiles()).ToA(),
		"tailscaled-socket":       carapace.ActionFiles(),
	})
}
