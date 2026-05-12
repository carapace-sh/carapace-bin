package embed

import (
	"testing"

	"github.com/spf13/pflag"
)

func TestFindSubcommandFlagArgSkipsRootFlagValues(t *testing.T) {
	rootFlags := pflag.NewFlagSet("root", pflag.ContinueOnError)
	rootFlags.String("sudoflags", "", "Pass arguments to sudo")
	rootFlags.BoolP("verbose", "v", false, "be verbose")

	subcommandFlags := pflag.NewFlagSet("subcommands", pflag.ContinueOnError)
	subcommandFlags.BoolP("query", "Q", false, "Query the package database")
	subcommandFlags.BoolP("sync", "S", false, "Synchronize packages")

	tests := []struct {
		name string
		args []string
		want string
	}{
		{
			name: "short subcommand after root string flag",
			args: []string{"--sudoflags", "-S", "-Q", "--sudo"},
			want: "-Q",
		},
		{
			name: "long subcommand after root string flag with equals",
			args: []string{"--sudoflags=-S", "--query", "--sudo"},
			want: "--query",
		},
		{
			name: "combined short subcommand",
			args: []string{"-v", "-Syu", "--needed"},
			want: "-S",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, ok := findSubcommandFlagArg(tt.args, rootFlags, subcommandFlags, false)
			if !ok {
				t.Fatal("expected subcommand flag")
			}
			if got != tt.want {
				t.Fatalf("expected %q, got %q", tt.want, got)
			}
		})
	}
}

func TestFindSubcommandFlagArgStopsAtPositional(t *testing.T) {
	rootFlags := pflag.NewFlagSet("root", pflag.ContinueOnError)
	subcommandFlags := pflag.NewFlagSet("subcommands", pflag.ContinueOnError)
	subcommandFlags.BoolP("query", "Q", false, "Query the package database")

	if got, ok := findSubcommandFlagArg([]string{"package", "-Q"}, rootFlags, subcommandFlags, false); ok {
		t.Fatalf("expected no subcommand flag, got %q", got)
	}
}
