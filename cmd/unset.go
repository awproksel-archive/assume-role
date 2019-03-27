package cmd

import (
	"github.com/awproksel/assume-role/internal"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(cmdUnset)
}

var cmdUnset = &cobra.Command{
	Use:   "unset",
	Short: "Returns commands to unset related environmental variables in parent shell",
	Run: func(cmd *cobra.Command, args []string) {
		internal.SourceableUnsetBashEnv()
	},
}
