package nix

import (
	"encoding/json"
	"strings"

	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace/pkg/style"
)

// ActionLocalChannels completes local channels
//
//	nixos (https://nixos.org/channels/nixos-22.05)
//	nixpkgs (https://nixos.org/channels/nixpkgs-unstable)
func ActionLocalChannels() carapace.Action {
	return carapace.ActionExecCommand("nix-channel", "--list")(func(output []byte) carapace.Action {
		lines := strings.Split(string(output), "\n")

		vals := make([]string, 0)
		for _, line := range lines[:len(lines)-1] {
			fields := strings.Fields(line)
			vals = append(vals, fields[0], strings.Join(fields[1:], " "))
		}
		return carapace.ActionValuesDescribed(vals...)
	})
}

// ActionRemoteChannels completes remote channels
//
//	nixos-21.11-small (unmaintained)
//	nixos-22.05 (stable)
func ActionRemoteChannels() carapace.Action {
	return carapace.ActionExecCommand("curl", "https://monitoring.nixos.org/prometheus/api/v1/query?query=channel_revision")(func(output []byte) carapace.Action {
		var response struct {
			Data struct {
				Result []struct {
					Metric struct {
						Channel string
						Status  string
					}
				}
			}
		}

		if err := json.Unmarshal(output, &response); err != nil {
			return carapace.ActionMessage(err.Error())
		}

		styleForStatus := func(s string) string {
			switch s {
			case "rolling":
				return style.Blue
			case "stable":
				return style.Green
			case "unmaintained":
				return style.Red
			default:
				return style.Default
			}
		}

		vals := make([]string, 0)
		for _, result := range response.Data.Result {
			vals = append(vals, result.Metric.Channel, result.Metric.Status, styleForStatus(result.Metric.Status))
		}
		return carapace.ActionStyledValuesDescribed(vals...)
	})
}
