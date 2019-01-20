package cmd

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/url"
	"os"

	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

var endpoint string
var debug bool

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "wxa50",
	Short: "Remote control for Yamaha WXA-50 amplifier.",
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		endpoint := viper.GetString("endpoint")
		if endpoint == "" {
			return errors.New(`missing required flag "endpoint"`)
		}
		url, err := url.Parse(endpoint)
		if err != nil || url.Scheme == "" || url.Host == "" {
			return errors.New("invalid endpoint provided")
		}

		return nil
	},
	// 	Long: `A longer description that spans multiple lines and likely contains
	// examples and usage of using your application. For example:
	//
	// Cobra is a CLI library for Go that empowers applications.
	// This application is a tool to generate the needed files
	// to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	//	Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.config/wxa50/wxa50.yml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	rootCmd.PersistentFlags().StringVar(&endpoint, "endpoint", "", `HTTP endpoint for for the amplifier. e.g. "http://192.168.0.17"`)
	_ = viper.BindPFlag("endpoint", rootCmd.PersistentFlags().Lookup("endpoint"))

	rootCmd.PersistentFlags().BoolVar(&debug, "debug", false, `Enable debug mode`)
	_ = viper.BindPFlag("debug", rootCmd.PersistentFlags().Lookup("debug"))
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".wxa50" (without extension).
		viper.AddConfigPath(home + "/.config/wxa50")
		viper.SetConfigName("config")
	}

	viper.AutomaticEnv() // read in environment variables that match

	_ = viper.ReadInConfig()

	if debug := viper.GetBool("debug"); debug == false {
		log.SetOutput(ioutil.Discard)
	}

	if configFileUsed := viper.ConfigFileUsed(); configFileUsed != "" {
		log.Println("Using config file:", configFileUsed)
	}
}
