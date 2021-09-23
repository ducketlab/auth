package cmd

import (
	"errors"
	"github.com/ducketlab/auth/config"
	"github.com/spf13/cobra"
)

var (
	confType string
	confFile string
)

var serviceCmd = &cobra.Command{
	Use:   "start",
	Short: "authorization center service",
	Long:  "authorization center service",
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := loadGlobalConfig(confType); err != nil {
			return err
		}

		return nil
	},
}

func loadGlobalConfig(configType string) error {
	switch configType {
	case "file":
		err := config.LoadConfigFromToml(confFile)
		if err != nil {
			return err
		}
	case "env":
		err := config.LoadConfigFromEnv()
		if err != nil {
			return err
		}
		return nil
	default:
		return errors.New("unknown config type")
	}
	return nil
}
