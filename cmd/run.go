package cmd

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Run the arbitrage bot",
	Long:  "Start the arbitrage bot. Make sure you have configured everything correctly.",
	Run: func(cmd *cobra.Command, args []string) {
		logrus.Info("Hello World!")
		logrus.Info("You are running this program in " + viper.GetString("config") + " mode.")
	},
}

func init() {
	rootCmd.AddCommand(runCmd)
}
