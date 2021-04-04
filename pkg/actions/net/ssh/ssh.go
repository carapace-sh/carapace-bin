// package ssh contains ssh related actions
package ssh

import "github.com/rsteube/carapace"

// ActionCiphers completes ciphers
//   3des-cbc
//   aes128-cbc
func ActionCiphers() carapace.Action {
	return carapace.ActionValues(
		"3des-cbc", "aes128-cbc", "aes192-cbc", "aes256-cbc", "aes128-ctr", "aes192-ctr", "aes256-ctr", "arcfour128", "arcfour256", "arcfour", "blowfish-cbc", "cast128-cbc",
	)
}

// ActionOptions completes options and their values
//   AddKeysToAgent
//   AddressFamily=inet
func ActionOptions() carapace.Action {
	return carapace.ActionMultiParts("=", func(c carapace.Context) carapace.Action {
		options := map[string]carapace.Action{
			"AddKeysToAgent":                   carapace.ActionValues(),
			"AddressFamily":                    carapace.ActionValues("any", "inet", "inet6"),
			"BatchMode":                        carapace.ActionValues("yes", "no"),
			"BindAddress":                      carapace.ActionValues("yes", "no"),
			"CanonicalDomains":                 carapace.ActionValues(),
			"CanonicalizeFallbackLocal":        carapace.ActionValues(),
			"CanonicalizeHostname":             carapace.ActionValues(),
			"CanonicalizeMaxDots":              carapace.ActionValues(),
			"CanonicalizePermittedCNAMEs":      carapace.ActionValues(),
			"CASignatureAlgorithms":            carapace.ActionValues(),
			"CertificateFile":                  carapace.ActionValues(),
			"ChallengeResponseAuthentication":  carapace.ActionValues("yes", "no"),
			"CheckHostIP":                      carapace.ActionValues("yes", "no"),
			"Ciphers":                          carapace.ActionValues(),
			"ClearAllForwardings":              carapace.ActionValues(),
			"Compression":                      carapace.ActionValues("yes", "no"),
			"ConnectionAttempts":               carapace.ActionValues(),
			"ConnectTimeout":                   carapace.ActionValues(),
			"ControlMaster":                    carapace.ActionValues(),
			"ControlPath":                      carapace.ActionValues(),
			"ControlPersist":                   carapace.ActionValues(),
			"DynamicForward":                   carapace.ActionValues(),
			"EscapeChar":                       carapace.ActionValues(),
			"ExitOnForwardFailure":             carapace.ActionValues("yes", "no"),
			"FingerprintHash":                  carapace.ActionValues(),
			"ForwardAgent":                     carapace.ActionValues("yes", "no"),
			"ForwardX11":                       carapace.ActionValues("yes", "no"),
			"ForwardX11Timeout":                carapace.ActionValues(),
			"ForwardX11Trusted":                carapace.ActionValues("yes", "no"),
			"GatewayPorts":                     carapace.ActionValues("yes", "no"),
			"GlobalKnownHostsFile":             carapace.ActionFiles(),
			"GSSAPIAuthentication":             carapace.ActionValues("yes", "no"),
			"HashKnownHosts":                   carapace.ActionValues("yes", "no"),
			"Host":                             carapace.ActionValues(),
			"HostbasedAuthentication":          carapace.ActionValues("yes", "no"),
			"HostbasedKeyTypes":                carapace.ActionValues(),
			"HostKeyAlgorithms":                carapace.ActionValues(),
			"HostKeyAlias":                     carapace.ActionValues(),
			"Hostname":                         carapace.ActionValues(),
			"IdentitiesOnly":                   carapace.ActionValues("yes", "no"),
			"IdentityAgent":                    carapace.ActionValues(),
			"IdentityFile":                     carapace.ActionFiles(),
			"IPQoS":                            carapace.ActionValues(),
			"KbdInteractiveAuthentication":     carapace.ActionValues(),
			"KbdInteractiveDevices":            carapace.ActionValues(),
			"KexAlgorithms":                    carapace.ActionValues(),
			"LocalCommand":                     carapace.ActionValues(),
			"LocalForward":                     carapace.ActionValues(),
			"LogLevel":                         carapace.ActionValues("QUIET", "FATAL", "ERROR", "INFO", "VERBOSE", "DEBUG", "DEBUG1", "DEBUG2", "DEBUG3"),
			"MACs":                             carapace.ActionValues(),
			"Match":                            carapace.ActionValues(),
			"NoHostAuthenticationForLocalhost": carapace.ActionValues(),
			"NumberOfPasswordPrompts":          carapace.ActionValues(),
			"PasswordAuthentication":           carapace.ActionValues(),
			"PermitLocalCommand":               carapace.ActionValues(),
			"PKCS11Provider":                   carapace.ActionValues(),
			"Port":                             carapace.ActionValues(),
			"PreferredAuthentications":         carapace.ActionValues(),
			"ProxyCommand":                     carapace.ActionValues(),
			"ProxyJump":                        carapace.ActionValues(),
			"ProxyUseFdpass":                   carapace.ActionValues(),
			"PubkeyAcceptedKeyTypes":           carapace.ActionValues(),
			"PubkeyAuthentication":             carapace.ActionValues(),
			"RekeyLimit":                       carapace.ActionValues(),
			"RemoteCommand":                    carapace.ActionValues(),
			"RemoteForward":                    carapace.ActionValues(),
			"RequestTTY":                       carapace.ActionValues(),
			"SendEnv":                          carapace.ActionValues(),
			"ServerAliveInterval":              carapace.ActionValues(),
			"ServerAliveCountMax":              carapace.ActionValues(),
			"SetEnv":                           carapace.ActionValues(),
			"StreamLocalBindMask":              carapace.ActionValues(),
			"StreamLocalBindUnlink":            carapace.ActionValues(),
			"StrictHostKeyChecking":            carapace.ActionValues("yes", "no", "ask"),
			"TCPKeepAlive":                     carapace.ActionValues("yes", "no"),
			"Tunnel":                           carapace.ActionValues("yes", "point-to-point", "ethernet", "no"),
			"TunnelDevice":                     carapace.ActionValues(),
			"UpdateHostKeys":                   carapace.ActionValues(),
			"User":                             carapace.ActionValues(),
			"UserKnownHostsFile":               carapace.ActionFiles(),
			"VerifyHostKeyDNS":                 carapace.ActionValues("yes", "no", "ask"),
			"VisualHostKey":                    carapace.ActionValues("yes", "no"),
			"XAuthLocation":                    carapace.ActionFiles(),
		}

		switch len(c.Parts) {
		case 0:
			keys := make([]string, 0, len(options))
			for key := range options {
				keys = append(keys, key)
			}
			return carapace.ActionValues(keys...).Invoke(c).Suffix("=").ToA()
		case 1:
			if val, ok := options[c.Parts[0]]; ok {
				return val
			} else {
				return carapace.ActionValues()
			}
		default:
			return carapace.ActionValues()
		}
	})
}
