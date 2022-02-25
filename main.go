package main

import (
	"console/cmd/core"
	"console/config/mysql"
	"github.com/spf13/cobra"
	"log"
)

var versionCmd = &cobra.Command {
	Use: "version",
	Short: "version subcommand show git version info.",
}

func init() {
	versionCmd.AddCommand(core.Cmd)
	mysql.InitGorm()
}

func main () {
	if err := versionCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}