/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"text/tabwriter"
	"time"

	"github.com/DWethmar/bookmarks/bookmark"
	"github.com/spf13/cobra"
)

const (
	padding = 1
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "ls",
	Short: "List all bookmarks",
	Long:  `List all bookmarks that are saved in the bookmark manager`,
	RunE:  runList,
}

// runList represents the command to run when the list command is specified
func runList(cmd *cobra.Command, _ []string) error {
	lib, err := setupBookmarks(loadLibraryOptions{
		Verbose: cmd.Flag("verbose").Changed,
		DBName:  appName,
	})
	if err != nil {
		return err
	}
	bookmarks, err := lib.List()
	if err != nil {
		return fmt.Errorf("failed to list bookmarks: %w", err)
	}
	if len(bookmarks) == 0 {
		return nil
	}
	table(bookmarks)
	return nil
}

// table prints a table of bookmarks to the console
func table(bookmarks []*bookmark.Bookmark) {
	tw := tabwriter.NewWriter(os.Stdout, 0, 0, padding, ' ', 0)
	fmt.Fprintln(tw, "Title\tContent\tCreated At")
	for _, b := range bookmarks {
		fmt.Fprintf(tw, "%s\t%s\t%s\n", b.Title, b.Content, b.CreatedAt.Format(time.DateTime))
	}
	tw.Flush()
}

func init() {
	rootCmd.AddCommand(listCmd)
}
