package cmd_test

import (
	"path/filepath"
	"runtime"
	"testing"

	"github.com/DWethmar/bookmarks/cmd"
)

func TestConfigDir(t *testing.T) {
	t.Run("Windows", func(t *testing.T) {
		// skip if not windows
		if runtime.GOOS != "windows" {
			t.Skip("skipping test on non-windows system")
		}

		t.Setenv("APPDATA", "C:\\Users\\<User>\\AppData\\Roaming")
		got := cmd.ConfigDir("windows", "bookmark")

		// Normalize the expected path to avoid mismatched separators
		want := filepath.Clean("C:\\Users\\<User>\\AppData\\Roaming\\bookmark")

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

	t.Run("Darwin", func(t *testing.T) {
		// skip if not darwin
		if runtime.GOOS != "darwin" {
			t.Skip("skipping test on non-darwin system")
		}

		t.Setenv("XDG_CONFIG_HOME", "/home/user/.config")
		got := cmd.ConfigDir("darwin", "bookmark")

		// Normalize the expected path to avoid mismatched separators
		want := filepath.Clean("/home/user/.config/bookmark")

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

	t.Run("Linux", func(t *testing.T) {
		// skip if not linux
		if runtime.GOOS != "linux" {
			t.Skip("skipping test on non-linux system")
		}

		t.Setenv("XDG_CONFIG_HOME", "/home/user/.config")
		got := cmd.ConfigDir("linux", "bookmark")

		// Normalize the expected path to avoid mismatched separators
		want := filepath.Clean("/home/user/.config/bookmark")

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

	t.Run("Fallback", func(t *testing.T) {
		// skip if windows, darwin or linux
		if runtime.GOOS != "linux" {
			t.Skip("skipping test on windows, darwin or linux system")
		}

		t.Setenv("HOME", "/home/user")
		got := cmd.ConfigDir("fallback", "bookmark")

		// Normalize the expected path to avoid mismatched separators
		want := filepath.Clean("/home/user/bookmark")

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
}
