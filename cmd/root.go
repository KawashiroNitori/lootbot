package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var isDebug bool

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "lootbot",
	Short: "FFXIV raid equipment loot tracker bot for Kaiheila",
	Long: `FF14 机器人，可追踪零式装备需求及分配情况，并提供开黑啦机器人功能。
`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

var RootCmd = rootCmd

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.lootbot.yaml)")
	rootCmd.PersistentFlags().BoolVarP(&isDebug, "debug", "d", false, "使用调试模式")

}
