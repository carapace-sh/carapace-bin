package tsh

import (
	"encoding/json"
	"strings"

	"github.com/carapace-sh/carapace"
)

// ActionSessions completes active sessions of the given kinds
//
//	8f7a2e02-40e2-4bd7-9f52-1f2f6f0e8e5c (ssh alice@node1)
//	3a1c9b1e-5a2a-4f0e-8b6c-9d1e2f3a4b5c (k8s alice@kube)
func ActionSessions(kinds string) carapace.Action {
	args := []string{"sessions", "ls", "--format", "json"}
	for kind := range strings.SplitSeq(kinds, ",") {
		if kind = strings.TrimSpace(kind); kind != "" {
			args = append(args, "--kind", kind)
		}
	}

	return carapace.ActionExecCommand("tsh", args...)(func(output []byte) carapace.Action {
		var sessions []struct {
			Spec struct {
				SessionID         string `json:"session_id"`
				Kind              string `json:"kind"`
				Login             string `json:"login"`
				ClusterName       string `json:"cluster_name"`
				TargetHostname    string `json:"target_hostname"`
				KubernetesCluster string `json:"kubernetes_cluster"`
			} `json:"spec"`
		}

		if err := json.Unmarshal(output, &sessions); err != nil {
			return carapace.ActionMessage(err.Error())
		}

		vals := make([]string, 0)
		for _, session := range sessions {
			target := session.Spec.TargetHostname
			if target == "" {
				target = session.Spec.KubernetesCluster
			}

			desc := session.Spec.Kind
			if target != "" {
				if session.Spec.Login != "" {
					desc += " " + session.Spec.Login + "@" + target
				} else {
					desc += " " + target
				}
			}
			if session.Spec.ClusterName != "" {
				desc += " (" + session.Spec.ClusterName + ")"
			}
			vals = append(vals, session.Spec.SessionID, desc)
		}
		return carapace.ActionValuesDescribed(vals...)
	})
}
