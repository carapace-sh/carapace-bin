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
		"3des-cbc", "aes128-cbc", "aes128-ctr", "aes192-cbc", "aes192-ctr", "aes256-cbc", "aes256-ctr", "arcfour", "arcfour128", "arcfour256", "blowfish-cbc", "cast128-cbc",
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
			"AddKeysToAgent":                   carapace.ActionValues("ask", "confirm", "no", "yes"),
			"AddressFamily":                    carapace.ActionValues("any", "inet", "inet6"),
			"BatchMode":                        carapace.ActionValues("no", "yes").StyleF(style.ForKeyword),
			"BindAddress":                      carapace.ActionValues("no", "yes").StyleF(style.ForKeyword),
			"CanonicalDomains":                 carapace.ActionValues(),
			"CanonicalizeFallbackLocal":        carapace.ActionValues("no", "yes").StyleF(style.ForKeyword),
			"CanonicalizeHostname":             carapace.ActionValues("always", "no", "yes").StyleF(style.ForKeyword),
			"CanonicalizeMaxDots":              carapace.ActionValues(),
			"CanonicalizePermittedCNAMEs":      carapace.ActionValues(),
			"CASignatureAlgorithms":            carapace.ActionValues(),
			"CertificateFile":                  carapace.ActionFiles(),
			"ChallengeResponseAuthentication":  carapace.ActionValues("no", "yes").StyleF(style.ForKeyword),
			"CheckHostIP":                      carapace.ActionValues("no", "yes").StyleF(style.ForKeyword),
			"Ciphers":                          ActionCiphers(),
			"ClearAllForwardings":              carapace.ActionValues("no", "yes").StyleF(style.ForKeyword),
			"Compression":                      carapace.ActionValues("no", "yes").StyleF(style.ForKeyword),
			"ConnectionAttempts":               carapace.ActionValues(),
			"ConnectTimeout":                   carapace.ActionValues(),
			"ControlMaster":                    carapace.ActionValues("ask", "auto", "autoask", "no", "yes"),
			"ControlPath":                      carapace.ActionValues(),
			"ControlPersist":                   carapace.ActionValues("no", "yes").StyleF(style.ForKeyword),
			"DynamicForward":                   carapace.ActionValues(),
			"EscapeChar":                       carapace.ActionValues(),
			"ExitOnForwardFailure":             carapace.ActionValues("no", "yes").StyleF(style.ForKeyword),
			"FingerprintHash":                  carapace.ActionValues("md5", "sha256"),
			"ForwardAgent":                     carapace.ActionValues("no", "yes").StyleF(style.ForKeyword),
			"ForwardX11":                       carapace.ActionValues("no", "yes").StyleF(style.ForKeyword),
			"ForwardX11Timeout":                carapace.ActionValues("no", "yes").StyleF(style.ForKeyword),
			"ForwardX11Trusted":                carapace.ActionValues("no", "yes").StyleF(style.ForKeyword),
			"GatewayPorts":                     carapace.ActionValues("no", "yes").StyleF(style.ForKeyword),
			"GlobalKnownHostsFile":             carapace.ActionFiles(),
			"GSSAPIAuthentication":             carapace.ActionValues("no", "yes").StyleF(style.ForKeyword),
			"HashKnownHosts":                   carapace.ActionValues("no", "yes").StyleF(style.ForKeyword),
			"Host":                             net.ActionHosts(),
			"HostbasedAuthentication":          carapace.ActionValues("no", "yes").StyleF(style.ForKeyword),
			"HostbasedKeyTypes":                carapace.ActionValues(),
			"HostKeyAlgorithms":                ActionHostKeyAlgorithms(),
			"HostKeyAlias":                     carapace.ActionValues(),
			"Hostname":                         carapace.ActionValues(),
			"IdentitiesOnly":                   carapace.ActionValues("no", "yes").StyleF(style.ForKeyword),
			"IdentityAgent":                    carapace.ActionValues(),
			"IdentityFile":                     carapace.ActionFiles(),
			"IPQoS":                            carapace.ActionValues("af11", "af12", "af13", "af21", "af22", "af23", "af31", "af32", "af33", "af41", "af42", "af43", "cs0", "cs1", "cs2", "cs3", "cs4", "cs5", "cs6", "cs7", "ef", "le", "lowdelay", "none", "reliability", "throughput"),
			"KbdInteractiveAuthentication":     carapace.ActionValues("no", "yes").StyleF(style.ForKeyword),
			"KbdInteractiveDevices":            carapace.ActionValues(),
			"KexAlgorithms":                    carapace.ActionValues(),
			"LocalCommand":                     carapace.ActionValues(),
			"LocalForward":                     carapace.ActionValues(),
			"LogLevel":                         carapace.ActionValues("DEBUG", "DEBUG1", "DEBUG2", "DEBUG3", "ERROR", "FATAL", "INFO", "QUIET", "VERBOSE"),
			"MACs":                             carapace.ActionValues(),
			"Match":                            carapace.ActionValues(),
			"NoHostAuthenticationForLocalhost": carapace.ActionValues("no", "yes").StyleF(style.ForKeyword),
			"NumberOfPasswordPrompts":          carapace.ActionValues(),
			"PasswordAuthentication":           carapace.ActionValues("no", "yes").StyleF(style.ForKeyword),
			"PermitLocalCommand":               carapace.ActionValues("no", "yes").StyleF(style.ForKeyword),
			"PKCS11Provider":                   carapace.ActionValues(),
			"Port":                             net.ActionPorts(),
			"PreferredAuthentications":         carapace.ActionValues(),
			"ProxyCommand":                     carapace.ActionValues(),
			"ProxyJump":                        carapace.ActionValues(),
			"ProxyUseFdpass":                   carapace.ActionValues(),
			"PubkeyAcceptedKeyTypes":           carapace.ActionValues(),
			"PubkeyAuthentication":             carapace.ActionValues("no", "yes").StyleF(style.ForKeyword),
			"RekeyLimit":                       carapace.ActionValues(),
			"RemoteCommand":                    carapace.ActionValues(),
			"RemoteForward":                    carapace.ActionValues(),
			"RequestTTY":                       carapace.ActionValues("auto", "force", "no", "yes"),
			"SendEnv":                          carapace.ActionValues(),
			"ServerAliveInterval":              carapace.ActionValues(),
			"ServerAliveCountMax":              carapace.ActionValues(),
			"SetEnv":                           carapace.ActionValues(),
			"StreamLocalBindMask":              carapace.ActionValues(),
			"StreamLocalBindUnlink":            carapace.ActionValues("no", "yes").StyleF(style.ForKeyword),
			"StrictHostKeyChecking":            carapace.ActionValues("ask", "no", "yes"),
			"TCPKeepAlive":                     carapace.ActionValues("no", "yes").StyleF(style.ForKeyword),
			"Tunnel":                           carapace.ActionValues("ethernet", "no", "point-to-point", "yes"),
			"TunnelDevice":                     carapace.ActionValues(),
			"UpdateHostKeys":                   carapace.ActionValues("ask", "no", "yes"),
			"User":                             os.ActionUsers(),
			"UserKnownHostsFile":               carapace.ActionFiles(),
			"VerifyHostKeyDNS":                 carapace.ActionValues("ask", "no", "yes"),
			"VisualHostKey":                    carapace.ActionValues("no", "yes").StyleF(style.ForKeyword),
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
