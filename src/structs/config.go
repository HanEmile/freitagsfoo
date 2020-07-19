package structs

// Configuration bundles the indivudial configurations
type Configuration struct {
	HTTP HTTPConfig
}

// HTTPConfig defines an HTTP config
type HTTPConfig struct {
	BindIP string
	Port   int
}
