package bridges

type t string

const (
	cobra    t = "cobra"
	complete   = "complete"
	yargs      = "yargs"
)

var bridges = map[string]t{
	"gh":       cobra,
	"kubectl":  complete,
	"minikube": cobra,
}
