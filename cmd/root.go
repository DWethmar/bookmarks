package cmd

import (
	"os"

	"github.com/DWethmar/bookmarks/ui"
	"github.com/spf13/cobra"
)

const appName = "bookmarks"

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   appName,
	Short: "a bookmark manager for webresources",
	Long:  `A bookmark manager for webresources, that allows you to save, list and delete bookmarks`,
	RunE:  runRootCmd,
}

// runRootCmd represents the command to run when no subcommands are specified
func runRootCmd(cmd *cobra.Command, _ []string) error {
	lib, err := loadLibrary(loadLibraryOptions{
		Verbose: cmd.Flag("verbose").Changed,
		DBName:  appName,
	})
	if err != nil {
		return err
	}
	return ui.Run(lib)
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().BoolP("verbose", "v", false, "verbose output")
}
