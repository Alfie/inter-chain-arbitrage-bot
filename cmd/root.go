package cmd

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"path/filepath"
)

var configName string
var logLevel string

var rootCmd = &cobra.Command{
	Use:   "inter-chain-arbitrage-bot",
	Short: "The first truly arbitrage-bot in the cosmos",
	Long:  "The inter-chain-arbitrage-bot (ICAB) is the first community-driven open-source project that allows everyone to take advantage of arbitrage opportunities in the cosmos.",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		initLogger()
	},
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	cobra.OnInitialize(initConfig)

	/* PERSISTENT FLAGS */
	rootCmd.PersistentFlags().StringVar(&configName, "config", "", "Name of the config file (default: production)")
	rootCmd.PersistentFlags().StringVar(&logLevel, "logLevel", "", "Level of the logger")

	/* FLAG BINDING */
	viper.BindPFlag("config", rootCmd.PersistentFlags().Lookup("config"))
	viper.BindPFlag("logLevel", rootCmd.PersistentFlags().Lookup("logLevel"))

	/* DEFAULT ENVIRONMENT VARIABLES */
	viper.SetDefault("config", "production")
}

// Load the configuration and set the environment variables
func initConfig() {
	// Set config path
	viper.AddConfigPath(filepath.Join(".", "configs"))

	// Set config type
	viper.SetConfigType("yaml")

	// Set config name
	viper.SetConfigName(viper.GetString("config"))

	// Load environment variables
	viper.AutomaticEnv()

	// Read config
	if err := viper.ReadInConfig(); err != nil {
		logrus.Fatal("No configuration file with the name '", viper.GetString("config"), "' found.")
	}
}

// Initialize the logger
func initLogger() {
	// Set the logging format
	switch viper.GetString("format") {
	case "text":
		logrus.SetFormatter(&logrus.TextFormatter{FullTimestamp: true})
	case "json":
		logrus.SetFormatter(&logrus.JSONFormatter{PrettyPrint: true})
	default: // text
		logrus.SetFormatter(&logrus.TextFormatter{FullTimestamp: true})
	}

	// Activate/ Deactivate function specific reporting
	// Adds significant overhead if activated
	logrus.SetReportCaller(viper.GetBool("functionTracing"))

	// Set the logging level
	switch viper.GetString("logLevel") {
	case "trace":
		logrus.SetLevel(logrus.TraceLevel)
	case "debug":
		logrus.SetLevel(logrus.DebugLevel)
	case "info":
		logrus.SetLevel(logrus.InfoLevel)
	case "warn":
		logrus.SetLevel(logrus.WarnLevel)
	case "error":
		logrus.SetLevel(logrus.ErrorLevel)
	case "fatal":
		logrus.SetLevel(logrus.FatalLevel)
	case "panic":
		logrus.SetLevel(logrus.PanicLevel)
	default: // info
		logrus.SetLevel(logrus.InfoLevel)
	}
}
