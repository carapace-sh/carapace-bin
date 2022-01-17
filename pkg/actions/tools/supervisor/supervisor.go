package supervisor

import (
	"errors"
	"os"

	"github.com/rsteube/carapace"
)

func configPath(c carapace.Context) (string, error) {
	// TODO ../etc/supervisord.conf (Relative to the executable)
	// TODO ../supervisord.conf (Relative to the executable)
	paths := []string{
		c.Dir + "/supervisord.conf",
		c.Dir + "/etc/supervisord.conf",
		"/etc/supervisord.conf",
		"/etc/supervisor/supervisord.conf",
	}

	for _, path := range paths {
		if _, err := os.Stat(path); err == nil {
			return path, nil
		}
	}
	return "", errors.New("could not find config file")
}
