package cmd

import (
	"fmt"

	"github.com/awproksel/assume-role/internal"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(cmdBecome)
}

var cmdBecome = &cobra.Command{
	Use:   "become [profile]",
	Short: "Returns commands to set related environmental variables in parent shell",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		creds, err := internal.AssumeRoleViaProfile(args[0])

		if err != nil {
			fmt.Println("Error while executing assume-role", err)
		}

		internal.SourceableBashEnv(creds)
	},
}
