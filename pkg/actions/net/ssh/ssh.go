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

// ActionCiphers completes ciphers
//
//	3des-cbc
//	aes128-cbc
func ActionCiphers() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		return carapace.ActionExecCommand("ssh", "-Q", "cipher")(func(output []byte) carapace.Action {
			lines := strings.Split(strings.TrimSpace(string(output)), "\n")
			if len(lines) == 1 && lines[0] == "" {
				lines = nil
			}
			if len(lines) > 0 {
				return carapace.ActionValues(lines...)
			}
			return carapace.ActionValues("3des-cbc", "aes128-cbc", "aes128-ctr", "aes192-cbc", "aes192-ctr", "aes256-cbc", "aes256-ctr", "arcfour", "arcfour128", "arcfour256", "blowfish-cbc", "cast128-cbc")
		})
	}).Tag("ciphers").Uid("ssh", "cipher")
}

// ActionHostKeyAlgorithms completes host key algorithms
//
//	ssh-ed25519
//	ssh-ed25519-cert-v01@openssh.com
func ActionHostKeyAlgorithms() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		return carapace.ActionExecCommand("ssh", "-Q", "HostKeyAlgorithms")(func(output []byte) carapace.Action {
			lines := strings.Split(strings.TrimSpace(string(output)), "\n")
			if len(lines) == 1 && lines[0] == "" {
				return carapace.ActionValues()
			}
			return carapace.ActionValues(lines...)
		})
	}).Tag("host key algorithms").Uid("ssh", "host-key-algorithm")
}

func actionQuery(query, tag, uid string, fallback ...string) carapace.Action {
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
	}).Tag(tag).Uid("ssh", uid)
}

func actionYesNo() carapace.Action {
	return carapace.ActionValues("no", "yes").StyleF(style.ForKeyword)
}

func actionYesNoAsk(values ...string) carapace.Action {
	return carapace.ActionValues(values...).StyleF(style.ForKeyword)
}

// ActionKexAlgorithms completes key exchange algorithms
func ActionKexAlgorithms() carapace.Action {
	return actionQuery("kex", "kex algorithms", "kex-algorithm")
}

// ActionMACs completes message authentication code algorithms
func ActionMACs() carapace.Action {
	return actionQuery("mac", "mac algorithms", "mac-algorithm")
}

// ActionSignatureAlgorithms completes signature algorithms
func ActionSignatureAlgorithms() carapace.Action {
	return actionQuery("sig", "signature algorithms", "signature-algorithm")
}

// ActionKeyTypes completes key types
func ActionKeyTypes() carapace.Action {
	return actionQuery("key", "key types", "key-type")
}

// ActionMatchCriteria completes Match criteria
func ActionMatchCriteria() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		criteria := option.MatchCriteria()
		vals := make([]string, 0, len(criteria))
		for _, criterion := range criteria {
			vals = append(vals, criterion.Name)
		}
		return carapace.ActionValues(vals...)
	}).Tag("match criteria").Uid("ssh", "match-criterion")
}

// ActionIdentityAgents completes identity agent values
func ActionIdentityAgents() carapace.Action {
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
			"CASignatureAlgorithms":            ActionSignatureAlgorithms(),
			"CertificateFile":                  carapace.ActionFiles(),
			"ChallengeResponseAuthentication":  actionYesNo(),
			"CheckHostIP":                      actionYesNo(),
			"Ciphers":                          ActionCiphers(),
			"ClearAllForwardings":              actionYesNo(),
			"Compression":                      actionYesNo(),
			"ConnectionAttempts":               carapace.ActionValues("1", "2", "3", "4", "5", "10"),
			"ConnectTimeout":                   carapace.ActionValues("0", "5", "10", "30", "60"),
			"ControlMaster":                    carapace.ActionValues("ask", "auto", "autoask", "no", "yes"),
			"ControlPath":                      carapace.ActionFiles(),
			"ControlPersist":                   carapace.ActionValues("0", "1h", "1m", "10m", "30m", "no", "yes").StyleF(style.ForKeyword),
			"DynamicForward":                   carapace.ActionValues(),
			"EscapeChar":                       carapace.ActionValues("none", "~"),
			"ExitOnForwardFailure":             actionYesNo(),
			"FingerprintHash":                  carapace.ActionValues("md5", "sha256"),
			"ForwardAgent":                     actionYesNo(),
			"ForwardX11":                       actionYesNo(),
			"ForwardX11Timeout":                carapace.ActionValues("0", "1h", "5m", "10m", "30m"),
			"ForwardX11Trusted":                actionYesNo(),
			"GatewayPorts":                     actionYesNo(),
			"GlobalKnownHostsFile":             carapace.Batch(carapace.ActionValues("none"), carapace.ActionFiles()).ToA(),
			"GSSAPIAuthentication":             actionYesNo(),
			"HashKnownHosts":                   actionYesNo(),
			"Host":                             ActionHosts(""),
			"HostbasedAuthentication":          actionYesNo(),
			"HostbasedKeyTypes":                ActionKeyTypes(),
			"HostKeyAlgorithms":                ActionHostKeyAlgorithms(),
			"HostKeyAlias":                     carapace.ActionValues(),
			"Hostname":                         carapace.ActionValues(),
			"IdentitiesOnly":                   actionYesNo(),
			"IdentityAgent":                    ActionIdentityAgents(),
			"IdentityFile":                     carapace.ActionFiles(),
			"IPQoS":                            carapace.ActionValues("af11", "af12", "af13", "af21", "af22", "af23", "af31", "af32", "af33", "af41", "af42", "af43", "cs0", "cs1", "cs2", "cs3", "cs4", "cs5", "cs6", "cs7", "ef", "le", "lowdelay", "none", "reliability", "throughput"),
			"KbdInteractiveAuthentication":     actionYesNo(),
			"KbdInteractiveDevices":            carapace.ActionValues("bsdauth", "pam", "skey"),
			"KexAlgorithms":                    ActionKexAlgorithms(),
			"KnownHostsCommand":                carapace.ActionFiles(),
			"LocalCommand":                     carapace.ActionValues(),
			"LocalForward":                     carapace.ActionValues(),
			"LogLevel":                         carapace.ActionValues("DEBUG", "DEBUG1", "DEBUG2", "DEBUG3", "ERROR", "FATAL", "INFO", "QUIET", "VERBOSE"),
			"MACs":                             ActionMACs(),
			"Match":                            ActionMatchCriteria(),
			"NoHostAuthenticationForLocalhost": actionYesNo(),
			"NumberOfPasswordPrompts":          carapace.ActionValues("1", "2", "3"),
			"PasswordAuthentication":           actionYesNo(),
			"PermitLocalCommand":               actionYesNo(),
			"PKCS11Provider":                   carapace.ActionFiles(),
			"Port":                             net.ActionPorts(),
			"PreferredAuthentications":         carapace.ActionValues("gssapi-with-mic", "hostbased", "keyboard-interactive", "password", "publickey"),
			"ProxyCommand":                     carapace.ActionValues(),
			"ProxyJump":                        ActionHosts(""),
			"ProxyUseFdpass":                   actionYesNo(),
			"PubkeyAcceptedKeyTypes":           ActionKeyTypes(),
			"PubkeyAuthentication":             actionYesNo(),
			"RekeyLimit":                       carapace.ActionValues("1G", "1h", "default"),
			"RemoteCommand":                    carapace.ActionValues(),
			"RemoteForward":                    carapace.ActionValues(),
			"RequestTTY":                       carapace.ActionValues("auto", "force", "no", "yes"),
			"RevokedHostKeys":                  carapace.ActionFiles(),
			"SecurityKeyProvider":              carapace.ActionValues("internal", "none"),
			"SendEnv":                          carapace.ActionValues(),
			"ServerAliveCountMax":              carapace.ActionValues("0", "1", "3", "6", "10"),
			"ServerAliveInterval":              carapace.ActionValues("0", "15", "30", "60", "120"),
			"SetEnv":                           carapace.ActionValues(),
			"StreamLocalBindMask":              carapace.ActionValues(),
			"StreamLocalBindUnlink":            actionYesNo(),
			"StrictHostKeyChecking":            carapace.ActionValues("accept-new", "ask", "no", "off", "yes").StyleF(style.ForKeyword),
			"TCPKeepAlive":                     actionYesNo(),
			"Tunnel":                           carapace.ActionValues("ethernet", "no", "point-to-point", "yes"),
			"TunnelDevice":                     carapace.ActionValues(),
			"UpdateHostKeys":                   carapace.ActionValues("ask", "no", "yes"),
			"User":                             os.ActionUsers(),
			"UserKnownHostsFile":               carapace.Batch(carapace.ActionValues("none"), carapace.ActionFiles()).ToA(),
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
