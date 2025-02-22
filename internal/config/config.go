package config

type Config struct {
	WatchDir    string   // Directory to watch
	IgnoreDirs  []string // Directories to ignore
	EntryPoint  string   // Main file to run
	BuildFlags  []string // Go build flags
	IncludeExt  []string // File extensions to watch
	ExcludeExt  []string // File extensions to ignore
	DelayMS     int      // Delay before restart in milliseconds
	DebugMode   bool     // Enable debug logging
	ClearScreen bool     // Clear screen on reload
}

func NewDefaultConfig() *Config {
	return &Config{
		WatchDir:    ".",
		IgnoreDirs:  []string{".git", "vendor", "node_modules"},
		EntryPoint:  "main.go",
		BuildFlags:  []string{},
		IncludeExt:  []string{".go"},
		ExcludeExt:  []string{".tmp"},
		DelayMS:     100,
		DebugMode:   false,
		ClearScreen: true,
	}
}
