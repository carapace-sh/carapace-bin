package tsh

import (
	"encoding/json"

	"github.com/carapace-sh/carapace"
)

// ActionRequests completes access requests
//
//	2f9e5c1e-2c9e-4e1f-9f3a-6b8c0d1e2f3a (alice admin)
//	7c8d9e0f-1a2b-3c4d-5e6f-7a8b9c0d1e2f (bob auditor)
func ActionRequests() carapace.Action {
	return carapace.ActionExecCommand("tsh", "request", "ls", "--format", "json")(func(output []byte) carapace.Action {
		var requests []struct {
			Metadata struct {
				Name string `json:"name"`
			} `json:"metadata"`
			Spec struct {
				User  string   `json:"user"`
				Roles []string `json:"roles"`
			} `json:"spec"`
		}

		if err := json.Unmarshal(output, &requests); err != nil {
			return carapace.ActionMessage(err.Error())
		}

		vals := make([]string, 0)
		for _, request := range requests {
			desc := request.Spec.User
			for _, role := range request.Spec.Roles {
				desc += " " + role
			}
			vals = append(vals, request.Metadata.Name, desc)
		}
		return carapace.ActionValuesDescribed(vals...)
	})
}

// ActionRequestableRoles completes requestable roles
//
//	admin (access admin)
//	auditor (read only access)
func ActionRequestableRoles() carapace.Action {
	return carapace.ActionExecCommand("tsh", "request", "search", "--roles", "--format", "json")(func(output []byte) carapace.Action {
		var roles []struct {
			Role        string
			Description string
		}

		if err := json.Unmarshal(output, &roles); err != nil {
			return carapace.ActionMessage(err.Error())
		}

		vals := make([]string, 0)
		for _, role := range roles {
			vals = append(vals, role.Role, role.Description)
		}
		return carapace.ActionValuesDescribed(vals...)
	})
}
