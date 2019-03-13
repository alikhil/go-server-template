package hello

// Context - is struct containing all needed things to handle requests
// it also contains config and may contain connection objects to databases
type Context struct {
	Config *Config
}

// GetContext - returns new context based on config
// in real apps this method will try to connect to databases and configure other dependent services
func GetContext(cfg *Config) (*Context, error) {
	return &Context{Config: cfg}, nil
}
