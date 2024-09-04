package console

import (
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var rootCmd = &cobra.Command{
	Use:   "kodinggo",
	Short: "Kodinggo App",
	Long:  `Kodinggo is a CLI application for managing your story.`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	setupLogger()
	loadEnv()
}

func loadEnv() {
	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.AddConfigPath("config/")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
}

func setupLogger() {
	log := logrus.New()

	// Set JSON formatter
	log.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
		ForceColors:   true,
	})
}
