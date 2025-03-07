package cmd

import (
	"fmt"
	"os"

	"github.com/DWethmar/bookmarks/ui"
	"github.com/spf13/cobra"
)

const appName = "bookmarks"

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   appName,
	Short: "A bookmark manager for webresources",
	Long:  `A bookmark manager for webresources, that allows you to save, search and delete bookmarks`,
	RunE:  runRootCmd,
}

// runRootCmd represents the command to run when no subcommands are specified
func runRootCmd(cmd *cobra.Command, _ []string) error {
	var err error
	lib, err := setupBookmarks(loadLibraryOptions{
		Verbose: cmd.Flag("verbose").Changed,
		DBName:  appName,
	})
	if err != nil {
		return err
	}
	// if a query is provided, search for bookmarks
	if q := cmd.Flag("search").Value.String(); q != "" {
		bookmarks, sErr := lib.Search(q)
		if sErr != nil {
			return fmt.Errorf("failed to search bookmarks: %w", sErr)
		}
		if len(bookmarks) == 0 {
			return nil
		}
		table(bookmarks)
		return nil
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
	rootCmd.Flags().StringP("search", "s", "", "search for bookmarks")
}
