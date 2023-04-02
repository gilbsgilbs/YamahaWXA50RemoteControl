package cmd

import (
	"fmt"
	"github.com/antchfx/xmlquery"
	"github.com/gilbsgilbs/YamahaWXA50RemoteControl/lib"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// muteCmd represents the source command
var muteCmd = &cobra.Command{
	Use:   "mute [on/off/toggle]",
	Short: "Mute or unmute the volume, or get the current mute status.",
	Args:  cobra.MaximumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		endpoint := viper.GetString("endpoint")

		if len(args) == 0 {
			node, err := lib.GetParams(endpoint)
			if err != nil {
				return err
			}
			muteValue := xmlquery.FindOne(node, "//Volume/Mute/text()").Data
			fmt.Println(muteValue)
			return nil
		}

		action := args[0]

		var err error
		if action == "on" {
			_, err = lib.Mute(endpoint)
		} else if action == "off" {
			_, err = lib.Unmute(endpoint)
		} else if action == "toggle" {
			var node *xmlquery.Node
			node, err = lib.GetParams(endpoint)
			if err != nil {
				return err
			}
			muteValue := xmlquery.FindOne(node, "//Volume/Mute/text()").Data
			if muteValue == "Off" {
				_, err = lib.Mute(endpoint)
			} else if muteValue == "On" {
				_, err = lib.Unmute(endpoint)
			} else {
				err = fmt.Errorf(`unknown mute value "%s".`, muteValue)
			}
		} else {
			err = fmt.Errorf(`unknown mute action "%s".`, action)
		}

		return err
	},
}

func init() {
	rootCmd.AddCommand(muteCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// sourceCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// sourceCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
