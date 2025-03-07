package cmd

import (
	"errors"
	"fmt"
	"time"

	"github.com/DWethmar/bookmarks/bookmark"
	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a bookmark",
	Long:  "Add a bookmark to the bookmark manager",
	RunE:  runAddCmd,
}

// runAddCmd represents the command to run when the add command is specified
func runAddCmd(cmd *cobra.Command, args []string) error {
	title, err := cmd.Flags().GetString("title")
	if err != nil {
		return fmt.Errorf("failed to get name flag: %w", err)
	}
	if len(args) == 0 {
		return errors.New("no content provided")
	}
	content := args[0]
	// create bookmark
	b := &bookmark.Bookmark{
		Title:     title,
		Content:   content,
		CreatedAt: time.Now(),
	}
	lib, err := setupBookmarks(loadLibraryOptions{
		Verbose: cmd.Flag("verbose").Changed,
		DBName:  appName,
	})
	if err != nil {
		return err
	}
	if err = lib.Add(cmd.Context(), b); err != nil {
		return fmt.Errorf("failed to add bookmark: %w", err)
	}
	return nil
}

func init() {
	rootCmd.AddCommand(addCmd)
	addCmd.Flags().StringP("title", "t", "", "title of the bookmark")
}
