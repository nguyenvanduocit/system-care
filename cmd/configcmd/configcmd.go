package configcmd

import (
	"fmt"

	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var shouldSave bool

var RootCmd = &cobra.Command{
	Use: "config",
	RunE: func(cmd *cobra.Command, args []string) error {
		if shouldSave {
			configOutPath, err := getDefaultConfigPath()
			if err != nil {
				return err
			}

			if err := viper.WriteConfigAs(configOutPath); err != nil {
				return err
			}
			fmt.Printf("Configs was save to: %s\n", configOutPath)
			return nil
		}
		return printConfig()
	},
}

func init() {
	RootCmd.Flags().BoolVar(&shouldSave, "save", false, "save config to default path")
}

func getDefaultConfigPath() (string, error) {
	defaultPath := viper.ConfigFileUsed()
	if defaultPath != "" {
		return defaultPath, nil
	}

	homedirPath, err := homedir.Dir()
	if err != nil {
		return "", err
	}
	return homedirPath + "/" + ".system-care.yml", nil
}

func printConfig() error {
	allSettings := viper.AllSettings()
	for key, value := range allSettings {
		fmt.Printf("%s: %v\n", key, value)
	}
	return nil
}
