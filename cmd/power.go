package cmd

import (
	"errors"
	"fmt"
	"github.com/gilbsgilbs/YamahaWXA50RemoteControl/lib"
	"github.com/spf13/viper"

	"github.com/spf13/cobra"
)

// powerCmd represents the power command
var powerCmd = &cobra.Command{
	Use:   "power on/off",
	Short: "Powers the WXA on or off.",
	Args: cobra.RangeArgs(1, 1),
	RunE: func(cmd *cobra.Command, args []string) error {
		endpoint := viper.GetString("endpoint")
		action := args[0]

		var err error = nil
		if action == "on" {
			_, err = lib.PowerOn(endpoint)
		} else if action == "off" {
			_, err = lib.PowerOff(endpoint)
		} else {
			err = errors.New(
				fmt.Sprintf(`unknown action "%s".`, action))
		}

		return err
	},
}

func init() {
	rootCmd.AddCommand(powerCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// powerCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// powerCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
