package config

// Config is the application config
type Config struct {
	Server   Server
	Database Database
	Log      Log
}

// Log settings
type Log struct {
	Level string `default:"info"`
	Path  string
}

// Database settings
type Database struct {
	Type       string `default:"sqlite3"`
	Connection string `default:"data/ghz.db"`
}

// Server settings
type Server struct {
	Port uint `default:"80"`
}

// Read the config file
func Read(path string) (*Config, error) { _ = "STUB: not implemented"; return nil, nil }
