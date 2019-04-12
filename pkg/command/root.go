package command

import (
	"cmdb/pkg/utils"
	"fmt"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
)

var (
	cfgFile string
)

var RootCmd = &cobra.Command{
	Use: "CMDB",
}

func init() {
	cobra.OnInitialize(initConfig)
	RootCmd.PersistentFlags().StringVar(&cfgFile, "config", "./conf/config.yaml", "config file")
}

func Execute() {
	RootCmd.AddCommand(versionCmd)
	RootCmd.AddCommand(xormCmd)
	RootCmd.Execute()
}

func initConfig() {
	// Don't forget to read kubeconfig either from cfgFile or from home directory!
	if cfgFile != "" {
		// Use kubeconfig file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search kubeconfig in home directory with name ".cobra" (without extension).
		viper.AddConfigPath(home + "./conf")
		viper.SetConfigName("config")
	}

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("Can't read config:", err)
		os.Exit(1)
	}
	utils.LoggerInit()


}


