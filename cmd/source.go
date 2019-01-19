package cmd

import (
	"fmt"
	"github.com/antchfx/xmlquery"
	"github.com/gilbsgilbs/YamahaWXA50RemoteControl/lib"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// sourceCmd represents the source command
var sourceCmd = &cobra.Command{
	Use:   "source [new source]",
	Short: "Get or define the source.",
	Long: "Get or define the source. e.g. AUX, Spotify, …",
	Args: cobra.MaximumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		endpoint := viper.GetString("endpoint")
		if len(args) == 0 {
			node, err := lib.GetParams(endpoint)
			if err != nil {
				return err
			}
			sourceValue := xmlquery.FindOne(node, "//Input/Input_Sel/text()").Data
			fmt.Println(sourceValue)
		} else {
			_, err := lib.SetSource(endpoint, args[0])
			if err != nil {
				return err
			}
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(sourceCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// sourceCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// sourceCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
