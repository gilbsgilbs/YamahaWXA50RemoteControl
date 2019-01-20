package cmd

import (
	"errors"
	"fmt"
	"github.com/antchfx/xmlquery"
	"github.com/gilbsgilbs/YamahaWXA50RemoteControl/lib"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"strconv"
)

// volumeCmd represents the source command
var volumeCmd = &cobra.Command{
	Use:   "volume [inc/dec/set value]",
	Short: "Get, set, increase or decrease the volume.",
	Long: `Get, set, increase or decrease the volume.

If an action and a value are specified, increase or decrease volume by ` +
		`or set the volume to the opposite of the specified value in cB (= 0.1dB).` +
		`e.g. "set 450" will set the volume to -45.0dB. "inc 50" will increase the ` +
		`volume by 5.0dB.`,
	Args: cobra.MaximumNArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		endpoint := viper.GetString("endpoint")

		node, err := lib.GetParams(endpoint)
		if err != nil {
			return err
		}
		lvlValue := xmlquery.FindOne(node, "//Volume/Lvl/Val/text()").Data

		if len(args) == 0 {
			fmt.Println(lvlValue)
			return nil
		}

		if len(args) <= 1 {
			return errors.New("expected two arguments")
		}

		action := args[0]
		value := args[1]

		valueAsInt, err := strconv.ParseInt(value, 10, 0)
		if err != nil {
			return errors.New(
				fmt.Sprintf(
					`cannot convert value "%s" to a positive integer.`,
					value))
		}

		if action == "set" {
			_, err = lib.SetVolume(endpoint, -1*int(valueAsInt))
			return err
		}

		if action != "inc" && action != "dec" {
			return errors.New(
				fmt.Sprintf(
					`unknown action "%s".`,
					action))
		}

		var sign int64 = 1
		if action == "dec" {
			sign = -1
		}

		lvlValueAsInt, _ := strconv.ParseInt(lvlValue, 10, 0)
		lvlValueAsInt += sign * valueAsInt

		_, err = lib.SetVolume(endpoint, int(lvlValueAsInt))

		return err
	},
}

func init() {
	rootCmd.AddCommand(volumeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// sourceCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// sourceCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
