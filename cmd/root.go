package cmd

import (
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var RootCmd = &cobra.Command{
	Use: "auth",
	Short: "auth a service authentication and authentication center",
	Long: "auth ...",
	RunE: func(cmd *cobra.Command, args []string) error {
		return errors.New("no flag find")
	},
}

func Execute()  {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}