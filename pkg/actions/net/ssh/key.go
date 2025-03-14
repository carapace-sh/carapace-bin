package ssh

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace/pkg/style"
)

type key struct {
	Type        string
	Key         string
	Description string
}

func readPublicKey(name string) (*key, error) {
	file, err := os.OpenFile(name, os.O_RDONLY, os.ModePerm)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	if scanner.Scan() {
		if fields := strings.Fields(scanner.Text()); len(fields) > 2 {
			return &key{Type: fields[0], Key: fields[1], Description: fields[2]}, nil
		} else if len(fields) > 1 {
			return &key{Type: fields[0], Key: fields[1]}, nil

		}
	}
	return nil, fmt.Errorf("failed to read public key from %v", name)
}

// ActionPublicKeys completes public keys
//
//	/home/user/.ssh/id_rsa.pub (ssh-rsa user)
//	/home/user/.ssh/another.pub (ssh-rsa user@another)
func ActionPublicKeys() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		abs, err := c.Abs("~/.ssh/")
		if err != nil {
			return carapace.ActionMessage(err.Error())
		}

		entries, err := os.ReadDir(abs)
		if err != nil {
			return carapace.ActionMessage(err.Error())
		}

		vals := make([]string, 0)
		for _, entry := range entries {
			if !entry.IsDir() && strings.HasSuffix(entry.Name(), ".pub") {
				path := abs + entry.Name()
				if k, err := readPublicKey(path); err == nil {
					vals = append(vals, path, k.Type+" "+k.Description)
				}
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
		abs, err := c.Abs("~/.ssh/")
		if err != nil {
			return carapace.ActionMessage(err.Error())
		}

		entries, err := os.ReadDir(abs)
		if err != nil {
			return carapace.ActionMessage(err.Error())
		}

		vals := make([]string, 0)
		for _, entry := range entries {
			if !entry.IsDir() && strings.HasSuffix(entry.Name(), ".pub") {
				path := abs + entry.Name()
				privatePath := strings.TrimSuffix(path, ".pub")
				if _, err := os.Stat(privatePath); !os.IsNotExist(err) {
					if k, err := readPublicKey(path); err == nil {
						vals = append(vals, privatePath, k.Type+" "+k.Description)
					}
				}
			}
		}
		return carapace.ActionValuesDescribed(vals...).StyleF(style.ForPath)
	})
}

// ActionSigningKeys completes the contents of public keys for which a private key exists
func ActionSigningKeys() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		abs, err := c.Abs("~/.ssh/")
		if err != nil {
			return carapace.ActionMessage(err.Error())
		}

		entries, err := os.ReadDir(abs)
		if err != nil {
			return carapace.ActionMessage(err.Error())
		}

		vals := make([]string, 0)
		for _, entry := range entries {
			if entry.IsDir() || !strings.HasSuffix(entry.Name(), ".pub") {
				continue
			}

			path := abs + entry.Name()
			privatePath := strings.TrimSuffix(path, ".pub")

			if _, err := os.Stat(privatePath); err != nil {
				continue
			}

			if contents, err := os.ReadFile(path); err == nil {
				pubKeyLine := strings.Split(string(contents), "\n")[0]
				pubKeyFields := strings.Fields(pubKeyLine)
				if len(pubKeyFields) > 1 {
					vals = append(vals, pubKeyFields[0]+" "+pubKeyFields[1], privatePath)
				}
			}
		}
		return carapace.ActionValuesDescribed(vals...)
	})
}
