package cmd

import (
	"fmt"
	"os"

	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/nguyenvanduocit/system-care/cmd/configcmd"
	"github.com/nguyenvanduocit/system-care/cmd/pushcmd"
)

var cfgFile string

var RootCmd = &cobra.Command{
	Use: "system-care",
	RunE: func(cmd *cobra.Command, args []string) error {
		return cmd.Usage()
	},
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	RootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", fmt.Sprintf("config file (default is ./%s.yaml or $HOME/%s.yaml)", ".system-care", ".system-care"))
	RootCmd.PersistentFlags().StringP("token", "t", "", "token")
	RootCmd.PersistentFlags().StringP("bucket", "b", "", "bucket id")
	RootCmd.PersistentFlags().StringP("org", "o", "", "org")
	RootCmd.PersistentFlags().StringP("server", "s", "", "server url")
	viper.BindPFlags(RootCmd.PersistentFlags())

	RootCmd.AddCommand(configcmd.RootCmd)
	RootCmd.AddCommand(pushcmd.RootCmd)
}

func initConfig() {
	viper.SetEnvPrefix("sc")
	viper.AutomaticEnv()
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		viper.AddConfigPath(home)
		viper.AddConfigPath(".")
		viper.SetConfigName(".system-care")
	}

	if err := viper.ReadInConfig(); err != nil {
		// log.Println(err)
	}
}
