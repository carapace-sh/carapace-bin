package tsh

import (
	"encoding/json"

	"github.com/carapace-sh/carapace"
)

// ActionMFADevices completes MFA devices
//
//	phone (totp)
//	yubikey (webauthn)
func ActionMFADevices() carapace.Action {
	return carapace.ActionExecCommand("tsh", "mfa", "ls", "--format", "json")(func(output []byte) carapace.Action {
		var devices []struct {
			Metadata struct {
				Name string `json:"name"`
			} `json:"metadata"`
			ID       string          `json:"id"`
			TOTP     json.RawMessage `json:"totp"`
			U2F      json.RawMessage `json:"u2f"`
			Webauthn json.RawMessage `json:"webauthn"`
			SSO      json.RawMessage `json:"sso"`
			Browser  json.RawMessage `json:"browser"`
		}

		if err := json.Unmarshal(output, &devices); err != nil {
			return carapace.ActionMessage(err.Error())
		}

		vals := make([]string, 0)
		for _, device := range devices {
			deviceType := ""
			switch {
			case device.TOTP != nil:
				deviceType = "totp"
			case device.U2F != nil:
				deviceType = "u2f"
			case device.Webauthn != nil:
				deviceType = "webauthn"
			case device.SSO != nil:
				deviceType = "sso"
			case device.Browser != nil:
				deviceType = "browser"
			}
			vals = append(vals, device.Metadata.Name, deviceType)
			if device.ID != "" && device.ID != device.Metadata.Name {
				vals = append(vals, device.ID, device.Metadata.Name)
			}
		}
		return carapace.ActionValuesDescribed(vals...)
	})
}
