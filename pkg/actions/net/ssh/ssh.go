// package ssh contains ssh related actions
package ssh

import (
	"strings"

	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/net"
	"github.com/carapace-sh/carapace-bin/pkg/actions/os"
	"github.com/carapace-sh/carapace/pkg/style"
)

// ActionCiphers completes ciphers
//
//	3des-cbc
//	aes128-cbc
func ActionCiphers() carapace.Action {
	return carapace.ActionValues(
		"3des-cbc", "aes128-cbc", "aes192-cbc", "aes256-cbc", "aes128-ctr", "aes192-ctr", "aes256-ctr", "arcfour128", "arcfour256", "arcfour", "blowfish-cbc", "cast128-cbc",
	)
}

// ActionHostKeyAlgorithms completes host key algorithms
//
//	ssh-ed25519
//	ssh-ed25519-cert-v01@openssh.com
func ActionHostKeyAlgorithms() carapace.Action {
	return carapace.ActionExecCommand("ssh", "-Q", "hostkeyalgorithms")(func(output []byte) carapace.Action {
		lines := strings.Split(string(output), "\n")
		return carapace.ActionValues(lines[:len(lines)-1]...)
	})
}

// ActionOptions completes options and their values
//
//	AddKeysToAgent
//	AddressFamily=inet
func ActionOptions() carapace.Action {
	return carapace.ActionMultiParts("=", func(c carapace.Context) carapace.Action {
		options := map[string]carapace.Action{
			"AddKeysToAgent":                   carapace.ActionValues("yes", "ask", "confirm", "no"),
			"AddressFamily":                    carapace.ActionValues("any", "inet", "inet6"),
			"BatchMode":                        carapace.ActionValues("yes", "no").StyleF(style.ForKeyword),
			"BindAddress":                      carapace.ActionValues("yes", "no").StyleF(style.ForKeyword),
			"CanonicalDomains":                 carapace.ActionValues(),
			"CanonicalizeFallbackLocal":        carapace.ActionValues("yes", "no").StyleF(style.ForKeyword),
			"CanonicalizeHostname":             carapace.ActionValues("yes", "no", "always").StyleF(style.ForKeyword),
			"CanonicalizeMaxDots":              carapace.ActionValues(),
			"CanonicalizePermittedCNAMEs":      carapace.ActionValues(),
			"CASignatureAlgorithms":            carapace.ActionValues(),
			"CertificateFile":                  carapace.ActionFiles(),
			"ChallengeResponseAuthentication":  carapace.ActionValues("yes", "no").StyleF(style.ForKeyword),
			"CheckHostIP":                      carapace.ActionValues("yes", "no").StyleF(style.ForKeyword),
			"Ciphers":                          ActionCiphers(),
			"ClearAllForwardings":              carapace.ActionValues("yes", "no").StyleF(style.ForKeyword),
			"Compression":                      carapace.ActionValues("yes", "no").StyleF(style.ForKeyword),
			"ConnectionAttempts":               carapace.ActionValues(),
			"ConnectTimeout":                   carapace.ActionValues(),
			"ControlMaster":                    carapace.ActionValues("yes", "no", "ask", "auto", "autoask"),
			"ControlPath":                      carapace.ActionValues(),
			"ControlPersist":                   carapace.ActionValues("yes", "no").StyleF(style.ForKeyword),
			"DynamicForward":                   carapace.ActionValues(),
			"EscapeChar":                       carapace.ActionValues(),
			"ExitOnForwardFailure":             carapace.ActionValues("yes", "no").StyleF(style.ForKeyword),
			"FingerprintHash":                  carapace.ActionValues("md5", "sha256"),
			"ForwardAgent":                     carapace.ActionValues("yes", "no").StyleF(style.ForKeyword),
			"ForwardX11":                       carapace.ActionValues("yes", "no").StyleF(style.ForKeyword),
			"ForwardX11Timeout":                carapace.ActionValues("yes", "no").StyleF(style.ForKeyword),
			"ForwardX11Trusted":                carapace.ActionValues("yes", "no").StyleF(style.ForKeyword),
			"GatewayPorts":                     carapace.ActionValues("yes", "no").StyleF(style.ForKeyword),
			"GlobalKnownHostsFile":             carapace.ActionFiles(),
			"GSSAPIAuthentication":             carapace.ActionValues("yes", "no").StyleF(style.ForKeyword),
			"HashKnownHosts":                   carapace.ActionValues("yes", "no").StyleF(style.ForKeyword),
			"Host":                             net.ActionHosts(),
			"HostbasedAuthentication":          carapace.ActionValues("yes", "no").StyleF(style.ForKeyword),
			"HostbasedKeyTypes":                carapace.ActionValues(),
			"HostKeyAlgorithms":                ActionHostKeyAlgorithms(),
			"HostKeyAlias":                     carapace.ActionValues(),
			"Hostname":                         carapace.ActionValues(),
			"IdentitiesOnly":                   carapace.ActionValues("yes", "no").StyleF(style.ForKeyword),
			"IdentityAgent":                    carapace.ActionValues(),
			"IdentityFile":                     carapace.ActionFiles(),
			"IPQoS":                            carapace.ActionValues("af11", "af12", "af13", "af21", "af22", "af23", "af31", "af32", "af33", "af41", "af42", "af43", "cs0", "cs1", "cs2", "cs3", "cs4", "cs5", "cs6", "cs7", "ef", "le", "lowdelay", "throughput", "reliability", "none"),
			"KbdInteractiveAuthentication":     carapace.ActionValues("yes", "no").StyleF(style.ForKeyword),
			"KbdInteractiveDevices":            carapace.ActionValues(),
			"KexAlgorithms":                    carapace.ActionValues(),
			"LocalCommand":                     carapace.ActionValues(),
			"LocalForward":                     carapace.ActionValues(),
			"LogLevel":                         carapace.ActionValues("QUIET", "FATAL", "ERROR", "INFO", "VERBOSE", "DEBUG", "DEBUG1", "DEBUG2", "DEBUG3"),
			"MACs":                             carapace.ActionValues(),
			"Match":                            carapace.ActionValues(),
			"NoHostAuthenticationForLocalhost": carapace.ActionValues("yes", "no").StyleF(style.ForKeyword),
			"NumberOfPasswordPrompts":          carapace.ActionValues(),
			"PasswordAuthentication":           carapace.ActionValues("yes", "no").StyleF(style.ForKeyword),
			"PermitLocalCommand":               carapace.ActionValues("yes", "no").StyleF(style.ForKeyword),
			"PKCS11Provider":                   carapace.ActionValues(),
			"Port":                             net.ActionPorts(),
			"PreferredAuthentications":         carapace.ActionValues(),
			"ProxyCommand":                     carapace.ActionValues(),
			"ProxyJump":                        carapace.ActionValues(),
			"ProxyUseFdpass":                   carapace.ActionValues(),
			"PubkeyAcceptedKeyTypes":           carapace.ActionValues(),
			"PubkeyAuthentication":             carapace.ActionValues("yes", "no").StyleF(style.ForKeyword),
			"RekeyLimit":                       carapace.ActionValues(),
			"RemoteCommand":                    carapace.ActionValues(),
			"RemoteForward":                    carapace.ActionValues(),
			"RequestTTY":                       carapace.ActionValues("yes", "no", "force", "auto"),
			"SendEnv":                          carapace.ActionValues(),
			"ServerAliveInterval":              carapace.ActionValues(),
			"ServerAliveCountMax":              carapace.ActionValues(),
			"SetEnv":                           carapace.ActionValues(),
			"StreamLocalBindMask":              carapace.ActionValues(),
			"StreamLocalBindUnlink":            carapace.ActionValues("yes", "no").StyleF(style.ForKeyword),
			"StrictHostKeyChecking":            carapace.ActionValues("yes", "no", "ask"),
			"TCPKeepAlive":                     carapace.ActionValues("yes", "no").StyleF(style.ForKeyword),
			"Tunnel":                           carapace.ActionValues("yes", "point-to-point", "ethernet", "no"),
			"TunnelDevice":                     carapace.ActionValues(),
			"UpdateHostKeys":                   carapace.ActionValues("yes", "no", "ask"),
			"User":                             os.ActionUsers(),
			"UserKnownHostsFile":               carapace.ActionFiles(),
			"VerifyHostKeyDNS":                 carapace.ActionValues("yes", "no", "ask"),
			"VisualHostKey":                    carapace.ActionValues("yes", "no").StyleF(style.ForKeyword),
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
