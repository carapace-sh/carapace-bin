package bridges

import "os"

func Bridges() map[string]string {
	return map[string]string{
		"gh":       "cobra",
		"kubectl":  "cobra",
		"minikube": "cobra",
	}

}
