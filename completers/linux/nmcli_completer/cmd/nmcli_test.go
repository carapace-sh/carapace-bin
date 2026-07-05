package cmd

import (
	"testing"

	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace/pkg/sandbox"
)

func TestNmcliDevice(t *testing.T) {
	sandbox.Package(t, "github.com/carapace-sh/carapace-bin/completers/linux/nmcli_completer")(func(s *sandbox.Sandbox) {
		s.Run("device", "").
			Expect(carapace.ActionValuesDescribed(
				"checkpoint", "Run a command with a configuration checkpoint",
				"connect", "Connect the device",
				"delete", "Delete the software devices",
				"disconnect", "Disconnect devices",
				"down", "Disconnect device and prevent auto-activation",
				"lldp", "Display information about neighboring devices learned through LLDP",
				"modify", "Modify one or more proties currently active on the device",
				"monitor", "Monitor device activity",
				"reapply", "Attempt to update device with changes",
				"set", "Modify device properties",
				"show", "Show details of device",
				"status", "Show status for all devices",
				"up", "Connect the device",
				"wifi", "Perform operation on Wi-Fi devices",
			).Tag("commands"))

		s.Run("device", "wifi", "").
			Expect(carapace.ActionValuesDescribed(
				"connect", "Connect to a Wi-Fi network",
				"hotspot", "Create a Wi-Fi hotspot",
				"list", "List available Wi-Fi access points",
				"rescan", "Request that NetworkManager immediately re-scan for available access points",
				"show-password", "Show details of active Wi-Fi networks including secrets",
			).Tag("commands"))
	})
}

func TestNmcliGeneral(t *testing.T) {
	sandbox.Package(t, "github.com/carapace-sh/carapace-bin/completers/linux/nmcli_completer")(func(s *sandbox.Sandbox) {
		s.Run("general", "").
			Expect(carapace.ActionValuesDescribed(
				"hostname", "Get or change persistent system hostname",
				"logging", "Get or change NetworkManager logging level and domains",
				"permissions", "Show caller permissions for authenticated operations",
				"reload", "Reload NetworkManager's configuration",
				"status", "Show overall status of NetworkManager",
			).Tag("commands"))
	})
}

func TestNmcliConnection(t *testing.T) {
	sandbox.Package(t, "github.com/carapace-sh/carapace-bin/completers/linux/nmcli_completer")(func(s *sandbox.Sandbox) {
		s.Run("connection", "").
			Expect(carapace.ActionValuesDescribed(
				"add", "add a connection",
				"clone", "Clone an existing connection profile",
				"delete", "Delete an existing connection profile",
				"down", "Deactivate a connection",
				"edit", "Edit an existing connection profile in an interactive editor",
				"export", "Export a connection",
				"import", "Import an external configuration",
				"load", "Load or reload one or more connection files from disk",
				"migrate", "Migrate connection profiles to a different settings plugin",
				"modify", "Modify one or more properties of the connection profile",
				"monitor", "Monitor an existing connection profile",
				"reload", "Reload all connection files from disk",
				"show", "show details for specified connection",
				"up", "Activate a connection",
			).Tag("commands"))
	})
}
