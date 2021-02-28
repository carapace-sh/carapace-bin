package cmd

import (
	"strings"

	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/net"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "ssh",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolS("4", "4", false, "Forces ssh to use IPv4 addresses only")
	rootCmd.Flags().BoolS("6", "6", false, "Forces ssh to use IPv6 addresses only")
	rootCmd.Flags().BoolS("A", "A", false, "Enables forwarding of connections from an authentication agent")
	rootCmd.Flags().BoolS("a", "a", false, "Disables forwarding of the authentication agent connection.")
	rootCmd.Flags().StringS("B", "B", "", "Bind to given address before attempting to connect")
	rootCmd.Flags().StringS("b", "b", "", "Use bind_address on the local machine as the source address of the connection")
	rootCmd.Flags().BoolS("C", "C", false, "Requests compression of all data")
	rootCmd.Flags().StringS("c", "c", "", "Selects the cipher specification for encrypting the session")
	rootCmd.Flags().StringS("D", "D", "", "Specifies a local “dynamic” application-level port forwarding")
	rootCmd.Flags().StringS("E", "E", "", "Append debug logs to log_file instead of standard error")
	rootCmd.Flags().StringS("e", "e", "", "Sets the escape character for sessions with a pty")
	rootCmd.Flags().StringS("F", "F", "", "Specifies an alternative per-user configuration file")
	rootCmd.Flags().BoolS("f", "f", false, "Requests ssh to go to background just before command execution")
	rootCmd.Flags().BoolS("G", "G", false, "Causes ssh to print its configuration after evaluating Host and Match blocks and exit")
	rootCmd.Flags().BoolS("g", "g", false, "Allows remote hosts to connect to local forwarded ports")
	rootCmd.Flags().StringS("I", "I", "", "Specify the PKCS#11 shared library ssh should use")
	rootCmd.Flags().StringS("i", "i", "", "Selects a file from which the identity (private key) is read")
	rootCmd.Flags().StringS("J", "J", "", "Connect to the target host by using given jump host")
	rootCmd.Flags().BoolS("K", "K", false, "Enables GSSAPI-based authentication and forwarding of GSSAPI credentials")
	rootCmd.Flags().BoolS("k", "k", false, "Disables forwarding of GSSAPI credentials to the server")
	rootCmd.Flags().StringS("L", "L", "", "Specifies that connections are to be forwarded")
	rootCmd.Flags().StringS("l", "l", "", "Specifies the user to log in as on the remote machine")
	rootCmd.Flags().BoolS("M", "M", false, "Places the ssh client into “master” mode for connection sharing")
	rootCmd.Flags().StringS("m", "m", "", "A comma-separated list of MAC algorithms")
	rootCmd.Flags().BoolS("N", "N", false, "Do not execute a remote command.  This is useful for just forwarding ports.")
	rootCmd.Flags().BoolS("n", "n", false, "Redirects stdin from /dev/null")
	rootCmd.Flags().StringS("O", "O", "", "Control an active connection multiplexing master process")
	rootCmd.Flags().StringS("o", "o", "", "Can be used to give options in the format used in the configuration file")
	rootCmd.Flags().StringS("p", "p", "", "Port to connect to on the remote host")
	rootCmd.Flags().StringS("Q", "Q", "", "Queries ssh for the algorithms supported")
	rootCmd.Flags().BoolS("q", "q", false, "Quiet mode")
	rootCmd.Flags().StringS("R", "R", "", "Specifies that connections ar to be forwarded")
	rootCmd.Flags().StringS("S", "S", "", "Specifies the location of a control socket")
	rootCmd.Flags().BoolS("s", "s", false, "May be used to request invocation of a subsystem on the remote system")
	rootCmd.Flags().BoolS("T", "T", false, "Disable pseudo-terminal allocation")
	rootCmd.Flags().BoolS("t", "t", false, "Force pseudo-terminal allocation")
	rootCmd.Flags().BoolS("V", "V", false, "Display the version number and exit")
	rootCmd.Flags().BoolS("v", "v", false, "Verbose mode")
	rootCmd.Flags().StringS("W", "W", "", "Requests that standard input and output on the client be forwarded")
	rootCmd.Flags().StringS("w", "w", "", "Requests tunnel device forwarding")
	rootCmd.Flags().BoolS("X", "X", false, "Enables X11 forwarding")
	rootCmd.Flags().BoolS("x", "x", false, "Disables X11 forwarding")
	rootCmd.Flags().BoolS("Y", "Y", false, "Enables trusted X11 forwarding")
	rootCmd.Flags().BoolS("y", "y", false, "Send log information using the syslog system module")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		// TODO complete bindings
		// "B"
		// "b"
		"c": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return ActionCiphers().Invoke(c).Filter(c.Parts).ToA()
		}),
		// "D"
		"E": carapace.ActionFiles(),
		"F": carapace.ActionFiles(),
		"I": carapace.ActionFiles(),
		"i": carapace.ActionFiles(),
		// "J"
		// "L"
		// ";"
		// "m"
		"O": carapace.ActionValuesDescribed(
			"check", "check that the master process is running",
			"forward", "request forwardings without command execution",
			"cancel", "cancel forwardings",
			"exit", "request the master to exit",
			"stop", "request the master to stop accepting further multiplexing requests",
		),
		"o": ActionOptions(),
		// "p"
		"Q": ActionQueryOptions(),
		// "R"
		"S": carapace.ActionFiles(),
		// "W"
		// "w"
	})

	carapace.Gen(rootCmd).PositionalCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if strings.Contains(c.CallbackValue, "@") {
				prefix := strings.SplitN(c.CallbackValue, "@", 2)[0]
				return net.ActionHosts().Invoke(c).Prefix(prefix + "@").ToA()
			} else {
				return net.ActionHosts()
			}
		}),
	)
}

func ActionCiphers() carapace.Action {
	return carapace.ActionValues(
		"3des-cbc", "aes128-cbc", "aes192-cbc", "aes256-cbc", "aes128-ctr", "aes192-ctr", "aes256-ctr", "arcfour128", "arcfour256", "arcfour", "blowfish-cbc", "cast128-cbc",
	)
}

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

func ActionQueryOptions() carapace.Action {
	return carapace.ActionValuesDescribed(
		"cipher", "supported symmetric ciphers",
		"cipher-auth", "supported symmetric ciphers that support authenticated encryption",
		"help", "supported query terms for use with the -Q flag",
		"mac", "supported message integrity codes",
		"kex", "key exchange algorithms",
		"key", "key types",
		"key-cert", "certificate key types",
		"key-plain", "non-certificate key types",
		"key-sig", "all key types and signature algorithms",
		"protocol-version", "supported SSH protocol versions",
		"sig", "supported signature algorithms",
	)
}
