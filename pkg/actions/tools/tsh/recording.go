package tsh

import (
	"encoding/json"

	"github.com/carapace-sh/carapace"
)

// ActionRecordings completes recorded sessions
//
//	8f7a2e02-40e2-4bd7-9f52-1f2f6f0e8e5c (ssh alice@node1)
//	3a1c9b1e-5a2a-4f0e-8b6c-9d1e2f3a4b5c (k8s alice@kube)
func ActionRecordings() carapace.Action {
	return carapace.ActionExecCommand("tsh", "recordings", "ls", "--format", "json")(func(output []byte) carapace.Action {
		var recordings []struct {
			SessionID         string `json:"sid"`
			Protocol          string `json:"proto"`
			User              string `json:"user"`
			ServerHostname    string `json:"server_hostname"`
			KubernetesCluster string `json:"kubernetes_cluster"`
		}

		if err := json.Unmarshal(output, &recordings); err != nil {
			return carapace.ActionMessage(err.Error())
		}

		vals := make([]string, 0)
		for _, recording := range recordings {
			if recording.SessionID == "" {
				continue
			}

			target := recording.ServerHostname
			if target == "" {
				target = recording.KubernetesCluster
			}

			desc := recording.Protocol
			if target != "" {
				if recording.User != "" {
					desc += " " + recording.User + "@" + target
				} else {
					desc += " " + target
				}
			}
			vals = append(vals, recording.SessionID, desc)
		}
		return carapace.ActionValuesDescribed(vals...)
	})
}
