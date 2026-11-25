// package ssh contains ssh related actions
package ssh

import (
	"strings"

	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/net"
	"github.com/carapace-sh/carapace-bin/pkg/actions/os"
	sshkey "github.com/carapace-sh/carapace-ssh/key"
	"github.com/carapace-sh/carapace-ssh/option"
	"github.com/carapace-sh/carapace/pkg/style"
)

func actionQuery(query string, fallback ...string) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		return carapace.ActionExecCommand("ssh", "-Q", query)(func(output []byte) carapace.Action {
			lines := strings.Split(strings.TrimSpace(string(output)), "\n")
			if len(lines) == 1 && lines[0] == "" {
				lines = nil
			}
			if len(lines) > 0 {
				return carapace.ActionValues(lines...)
			}
			return carapace.ActionValues(fallback...)
		})
	})
}

func actionYesNo() carapace.Action {
	return carapace.ActionValues("no", "yes").StyleF(style.ForKeyword)
}

func actionYesNoAsk(values ...string) carapace.Action {
	return carapace.ActionValues(values...).StyleF(style.ForKeyword)
}

func actionMatch() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		criteria := option.MatchCriteria()
		vals := make([]string, 0, len(criteria))
		for _, criterion := range criteria {
			vals = append(vals, criterion.Name)
		}
		return carapace.ActionValues(vals...)
	})
}

func actionIdentityAgent() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		vals := []string{"none", "SSH_AUTH_SOCK"}
		seen := map[string]struct{}{"none": {}, "SSH_AUTH_SOCK": {}}
		for _, socket := range sshkey.AgentSocketPaths(func(name string) string {
			if value, ok := c.LookupEnv(name); ok {
				return value
			}
			return ""
		}) {
			if _, ok := seen[socket]; ok {
				continue
			}
			seen[socket] = struct{}{}
			vals = append(vals, socket)
		}
		return carapace.ActionValues(vals...)
	})
}

func actionKnownHostsFiles() carapace.Action {
	return carapace.Batch(
		carapace.ActionValues("none"),
		carapace.ActionFiles(),
	).ToA()
}

// ActionCiphers completes ciphers
//
//	3des-cbc
//	aes128-cbc
func ActionCiphers() carapace.Action {
	return actionQuery("cipher", "3des-cbc", "aes128-cbc", "aes128-ctr", "aes192-cbc", "aes192-ctr", "aes256-cbc", "aes256-ctr", "arcfour", "arcfour128", "arcfour256", "blowfish-cbc", "cast128-cbc")
}

// ActionHostKeyAlgorithms completes host key algorithms
//
//	ssh-ed25519
//	ssh-ed25519-cert-v01@openssh.com
func ActionHostKeyAlgorithms() carapace.Action {
	return actionQuery("HostKeyAlgorithms")
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
			"BatchMode":                        actionYesNo(),
			"BindAddress":                      actionYesNo(),
			"CanonicalDomains":                 carapace.ActionValues(),
			"CanonicalizeFallbackLocal":        actionYesNo(),
			"CanonicalizeHostname":             actionYesNoAsk("always", "no", "yes"),
			"CanonicalizeMaxDots":              carapace.ActionValues(),
			"CanonicalizePermittedCNAMEs":      carapace.ActionValues(),
			"CASignatureAlgorithms":            actionQuery("sig"),
			"CertificateFile":                  carapace.ActionFiles(),
			"ChallengeResponseAuthentication":  actionYesNo(),
			"CheckHostIP":                      actionYesNo(),
			"Ciphers":                          ActionCiphers(),
			"ClearAllForwardings":              actionYesNo(),
			"Compression":                      actionYesNo(),
			"ConnectionAttempts":               carapace.ActionValues(),
			"ConnectTimeout":                   carapace.ActionValues(),
			"ControlMaster":                    carapace.ActionValues("ask", "auto", "autoask", "no", "yes"),
			"ControlPath":                      carapace.ActionFiles(),
			"ControlPersist":                   actionYesNo(),
			"DynamicForward":                   carapace.ActionValues(),
			"EscapeChar":                       carapace.ActionValues(),
			"ExitOnForwardFailure":             actionYesNo(),
			"FingerprintHash":                  carapace.ActionValues("md5", "sha256"),
			"ForwardAgent":                     actionYesNo(),
			"ForwardX11":                       actionYesNo(),
			"ForwardX11Timeout":                carapace.ActionValues(),
			"ForwardX11Trusted":                actionYesNo(),
			"GatewayPorts":                     actionYesNo(),
			"GlobalKnownHostsFile":             actionKnownHostsFiles(),
			"GSSAPIAuthentication":             actionYesNo(),
			"HashKnownHosts":                   actionYesNo(),
			"Host":                             ActionHosts(""),
			"HostbasedAuthentication":          actionYesNo(),
			"HostbasedKeyTypes":                actionQuery("key"),
			"HostKeyAlgorithms":                ActionHostKeyAlgorithms(),
			"HostKeyAlias":                     carapace.ActionValues(),
			"Hostname":                         carapace.ActionValues(),
			"IdentitiesOnly":                   actionYesNo(),
			"IdentityAgent":                    actionIdentityAgent(),
			"IdentityFile":                     carapace.ActionFiles(),
			"IPQoS":                            carapace.ActionValues("af11", "af12", "af13", "af21", "af22", "af23", "af31", "af32", "af33", "af41", "af42", "af43", "cs0", "cs1", "cs2", "cs3", "cs4", "cs5", "cs6", "cs7", "ef", "le", "lowdelay", "none", "reliability", "throughput"),
			"KbdInteractiveAuthentication":     actionYesNo(),
			"KbdInteractiveDevices":            carapace.ActionValues(),
			"KexAlgorithms":                    actionQuery("kex"),
			"LocalCommand":                     carapace.ActionValues(),
			"LocalForward":                     carapace.ActionValues(),
			"LogLevel":                         carapace.ActionValues("DEBUG", "DEBUG1", "DEBUG2", "DEBUG3", "ERROR", "FATAL", "INFO", "QUIET", "VERBOSE"),
			"MACs":                             actionQuery("mac"),
			"Match":                            actionMatch(),
			"NoHostAuthenticationForLocalhost": actionYesNo(),
			"NumberOfPasswordPrompts":          carapace.ActionValues(),
			"PasswordAuthentication":           actionYesNo(),
			"PermitLocalCommand":               actionYesNo(),
			"PKCS11Provider":                   carapace.ActionFiles(),
			"Port":                             net.ActionPorts(),
			"PreferredAuthentications":         carapace.ActionValues("gssapi-with-mic", "hostbased", "keyboard-interactive", "password", "publickey"),
			"ProxyCommand":                     carapace.ActionValues(),
			"ProxyJump":                        ActionHosts(""),
			"ProxyUseFdpass":                   actionYesNo(),
			"PubkeyAcceptedKeyTypes":           actionQuery("key"),
			"PubkeyAuthentication":             actionYesNo(),
			"RekeyLimit":                       carapace.ActionValues(),
			"RemoteCommand":                    carapace.ActionValues(),
			"RemoteForward":                    carapace.ActionValues(),
			"RequestTTY":                       carapace.ActionValues("auto", "force", "no", "yes"),
			"SendEnv":                          carapace.ActionValues(),
			"ServerAliveInterval":              carapace.ActionValues(),
			"ServerAliveCountMax":              carapace.ActionValues(),
			"SetEnv":                           carapace.ActionValues(),
			"StreamLocalBindMask":              carapace.ActionValues(),
			"StreamLocalBindUnlink":            actionYesNo(),
			"StrictHostKeyChecking":            carapace.ActionValues("accept-new", "ask", "no", "off", "yes").StyleF(style.ForKeyword),
			"TCPKeepAlive":                     actionYesNo(),
			"Tunnel":                           carapace.ActionValues("ethernet", "no", "point-to-point", "yes"),
			"TunnelDevice":                     carapace.ActionValues(),
			"UpdateHostKeys":                   carapace.ActionValues("ask", "no", "yes"),
			"User":                             os.ActionUsers(),
			"UserKnownHostsFile":               actionKnownHostsFiles(),
			"VerifyHostKeyDNS":                 carapace.ActionValues("ask", "no", "yes"),
			"VisualHostKey":                    actionYesNo(),
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
