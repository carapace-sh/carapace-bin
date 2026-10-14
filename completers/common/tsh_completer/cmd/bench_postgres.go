package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bench_postgresCmd = &cobra.Command{
	Use:    "postgres",
	Short:  "Run PostgreSQL database benchmark tests.",
	Hidden: true,
	Run:    func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bench_postgresCmd).Standalone()

	bench_postgresCmd.Flags().String("db-name", "", "Database name where benchmark queries will be executed.")
	bench_postgresCmd.Flags().String("db-user", "", "Database user used to connect to the target database. The user must have enough permissions on the database to execute all the benchmark queries.")
	benchCmd.AddCommand(bench_postgresCmd)
}
