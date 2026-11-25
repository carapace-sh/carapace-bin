package ssh

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-ssh/key"
	"github.com/carapace-sh/carapace/pkg/style"
)

type publicKey struct {
	Type        string
	Key         string
	Description string
}

func readPublicKey(name string) (*publicKey, error) {
	file, err := os.Open(name)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	if scanner.Scan() {
		if fields := strings.Fields(scanner.Text()); len(fields) > 2 {
			return &publicKey{Type: fields[0], Key: fields[1], Description: fields[2]}, nil
		} else if len(fields) > 1 {
			return &publicKey{Type: fields[0], Key: fields[1]}, nil
		}
	}
	return nil, fmt.Errorf("failed to read public key from %v", name)
}

func sshDir(c carapace.Context) string {
	if home, ok := c.LookupEnv("HOME"); ok && home != "" {
		return key.DefaultDir(home)
	}
	if home, err := os.UserHomeDir(); err == nil && home != "" {
		return key.DefaultDir(home)
	}
	return ""
}

// ActionPublicKeys completes public keys
//
//	/home/user/.ssh/id_rsa.pub (ssh-rsa user)
//	/home/user/.ssh/another.pub (ssh-rsa user@another)
func ActionPublicKeys() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		dir := sshDir(c)
		if dir == "" {
			return carapace.ActionValues()
		}

		pairs, err := key.DiscoverKeyPairs(dir)
		if err != nil {
			return carapace.ActionMessage(err.Error())
		}

		vals := make([]string, 0)
		for _, pair := range pairs {
			if pair.Public == "" {
				continue
			}
			if k, err := readPublicKey(pair.Public); err == nil {
				description := k.Type
				if k.Description != "" {
					description += " " + k.Description
				}
				vals = append(vals, pair.Public, description)
			}
		}
		return carapace.ActionValuesDescribed(vals...).StyleF(style.ForPath)
	})
}

// ActionPrivateKeys completes private keys
//
//	/home/user/.ssh/id_rsa (ssh-rsa user)
//	/home/user/.ssh/another (ssh-rsa user@another)
func ActionPrivateKeys() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		dir := sshDir(c)
		if dir == "" {
			return carapace.ActionValues()
		}

		pairs, err := key.DiscoverKeyPairs(dir)
		if err != nil {
			return carapace.ActionMessage(err.Error())
		}

		vals := make([]string, 0)
		for _, pair := range pairs {
			if pair.Private == "" || pair.Public == "" {
				continue
			}
			if k, err := readPublicKey(pair.Public); err == nil {
				description := k.Type
				if k.Description != "" {
					description += " " + k.Description
				}
				vals = append(vals, pair.Private, description)
			}
		}
		return carapace.ActionValuesDescribed(vals...).StyleF(style.ForPath)
	})
}

// ActionSigningKeys completes the contents of public keys for which a private key exists
//
//	ssh-ed25519 AAAAC3NzaC1lZDI1NTE5AAAAI... (id_ed25519)
//	ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQ... (id_rsa)
func ActionSigningKeys() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		dir := sshDir(c)
		if dir == "" {
			return carapace.ActionValues()
		}

		pairs, err := key.DiscoverKeyPairs(dir)
		if err != nil {
			return carapace.ActionMessage(err.Error())
		}

		vals := make([]string, 0)
		for _, pair := range pairs {
			if pair.Private == "" || pair.Public == "" {
				continue
			}
			contents, err := os.ReadFile(pair.Public)
			if err != nil {
				continue
			}
			pubKeyLine, _, _ := strings.Cut(string(contents), "\n")
			pubKeyFields := strings.Fields(pubKeyLine)
			if len(pubKeyFields) > 1 {
				vals = append(vals, pubKeyFields[0]+" "+pubKeyFields[1], filepath.Base(pair.Private))
			}
		}
		return carapace.ActionValuesDescribed(vals...)
	})
}
