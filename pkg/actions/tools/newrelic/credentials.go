package newrelic

import (
	"encoding/json"
	"os"

	"github.com/rsteube/carapace"
)

type profile struct {
	Region    string
	AccountID int
}

func getProfiles() (map[string]profile, error) {
	path, err := carapace.Context{}.Abs("~/.newrelic/credentials.json")
	if err != nil {
		return nil, err
	}

	content, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	profiles := make(map[string]profile)
	if err := json.Unmarshal(content, &profiles); err != nil {
		return nil, err
	}

	if defaultProfile, err := getDefaultProfile(); err == nil {
		profiles[""] = profiles[defaultProfile]
	}
	return profiles, nil
}

func getDefaultProfile() (string, error) {
	path, err := carapace.Context{}.Abs("~/.newrelic/default-profile.json")
	if err != nil {
		return "", err
	}

	content, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}

	var name string
	if err := json.Unmarshal(content, &name); err != nil {
		return "", err
	}
	return name, nil
}
