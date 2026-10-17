package common

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

func AddServerFlags(cmd *cobra.Command) {
	cmd.Flags().Bool("allow-anonymous", false, "admit anonymous callers on plain TCP listeners")
	cmd.Flags().String("block-dir", "", "discover executable block plugins in DIR (repeatable)")
	cmd.Flags().Bool("block-plugin-reload", false, "reload block plugins after source changes")
	cmd.Flags().String("config", "", "read the config file at PATH")
	cmd.Flags().Bool("exit-empty", false, "stop when no sessions or control connections remain")
	cmd.Flags().String("extra-listen", "", "also listen on unix, http, tailnet, tailnet-serve, or login-server ENDPOINT (repeatable)")
	cmd.Flags().String("listen", "", "listen on ENDPOINT")
	cmd.Flags().String("log-format", "text", "write server logs as FORMAT: text or json")
	cmd.Flags().String("max-blocks", "", "refuse new blocks once the server holds N across all sessions")
	cmd.Flags().String("max-panes", "", "refuse new panes once a window holds N")
	cmd.Flags().String("max-sessions", "", "refuse new sessions once the server holds N")
	cmd.Flags().String("max-windows", "", "refuse new windows once a session holds N")
	cmd.Flags().String("socket-mode", "0600", "set Unix socket permissions to MODE")
	cmd.Flags().String("state-dir", "", "store persistent server state below DIR")
	cmd.Flags().Bool("tailnet-allow-identified", false, "admit every identified caller through the host's tailscaled")
	cmd.Flags().Bool("tailscale", false, "also serve the API over the tailnet as an embedded tsnet node")
	cmd.Flags().String("tailscale-advertise-tag", "", "advertise ACL TAG (tag:name) when --tailscale enrolls the node (repeatable)")
	cmd.Flags().Bool("tailscale-allow-tailnet-members", false, "allow any tailnet member permitted by ACLs to reach Rex")
	cmd.Flags().String("tailscale-auth-key-file", "", "enroll --tailscale with the one-time auth key read from FILE")
	cmd.Flags().String("tailscale-hostname", "rex", "tailnet device NAME for --tailscale")
	cmd.Flags().String("tailscaled-socket", "", "use the tailscaled LocalAPI socket at PATH")
	cmd.Flags().String("websocket-origin-pattern", "", "allow browser WebSockets from cross-origin PATTERN (repeatable)")

	carapace.Gen(cmd).FlagCompletion(carapace.ActionMap{
		"block-dir":               carapace.ActionDirectories(),
		"config":                  carapace.ActionFiles(),
		"log-format":              carapace.ActionValues("json", "text"),
		"state-dir":               carapace.ActionDirectories(),
		"tailscale-auth-key-file": carapace.ActionFiles(),
		"tailscaled-socket":       carapace.ActionFiles(),
	})
}
