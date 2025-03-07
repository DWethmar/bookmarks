package cmd

import (
	"fmt"
	"log/slog"
	"os"
	"path"
	"path/filepath"
	"runtime"

	"github.com/DWethmar/bookmarks/bookmark"
	"github.com/DWethmar/bookmarks/bookmark/json"
)

// ConfigDir returns the appropriate configuration directory for the given OS.
func ConfigDir(goos, appName string) string {
	var configDir string

	switch goos {
	case "windows":
		configDir = os.Getenv("APPDATA") // Typically "C:\\Users\\<User>\\AppData\\Roaming"
		if configDir == "" {
			configDir = filepath.Join(os.Getenv("USERPROFILE"), "AppData", "Roaming")
		}
	case "darwin", "linux":
		configDir = os.Getenv("XDG_CONFIG_HOME")
		if configDir == "" {
			configDir = filepath.Join(os.Getenv("HOME"), ".config")
		}
	default:
		// Fallback to home directory
		configDir = os.Getenv("HOME")
	}

	// Ensure fully normalized paths
	return filepath.Clean(filepath.Join(configDir, appName))
}

// Logger returns a new logger instance.
func Logger(verbose bool) *slog.Logger {
	level := slog.LevelInfo
	if verbose {
		level = slog.LevelDebug
	}
	return slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{
		Level: level,
	}))
}

// loadLibraryOptions are the options for loading a library.
type loadLibraryOptions struct {
	Verbose bool
	DBName  string
}

// loadLibrary loads a library.
func loadLibrary(o loadLibraryOptions) (*bookmark.Library, error) {
	logger := Logger(o.Verbose)
	workDir := ConfigDir(runtime.GOOS, appName)
	// make sure the workdir exists
	if err := os.MkdirAll(workDir, 0755); err != nil {
		return nil, err
	}
	logger.Debug(
		"workDir",
		slog.String("workDir", workDir),
		slog.String("appName", appName),
		slog.String("goos", runtime.GOOS),
		slog.String("dbName", o.DBName),
	)
	store := json.NewStore(path.Join(workDir, fmt.Sprintf("%s.json", o.DBName)))
	l := bookmark.NewLibrary(logger, store)
	return l, nil
}
