package cmd

import (
	"fmt"
	log "github.com/sirupsen/logrus"
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
	// Set config path to 'configs' dir
	viper.AddConfigPath(filepath.Join(".", "configs"))

	// Set config type to 'yaml'
	viper.SetConfigType("yaml")

	// Set config name to
	viper.SetConfigName(viper.GetString("config"))

	// Load environment variables
	viper.AutomaticEnv()

	// Read config
	if err := viper.ReadInConfig(); err != nil {
		//TODO Log error and exit
		fmt.Println("no config")
	}
}

// Initialize the logger
func initLogger() {
	// Set the logging format
	switch viper.GetString("format") {
	case "text":
		log.SetFormatter(&log.TextFormatter{FullTimestamp: true})
	case "json":
		log.SetFormatter(&log.JSONFormatter{PrettyPrint: true})
	}

	// Activate/ Deactivate function specific reporting
	// Adds significant overhead if activated
	log.SetReportCaller(viper.GetBool("functionTracing"))

	// Set the logging level
	switch viper.GetString("logLevel") {
	case "trace":
		log.SetLevel(log.TraceLevel)
	case "debug":
		log.SetLevel(log.DebugLevel)
	case "info":
		log.SetLevel(log.InfoLevel)
	case "warn":
		log.SetLevel(log.WarnLevel)
	case "error":
		log.SetLevel(log.ErrorLevel)
	case "fatal":
		log.SetLevel(log.FatalLevel)
	case "panic":
		log.SetLevel(log.PanicLevel)
	}
}
