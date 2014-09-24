package utils

type Config struct {
	// Fixtures loader switcher
	LoadFixtures bool

	// Mongo URL
	MgoURL string

	// Database name
	MgoDB string
}

func LoadConfig(conf *Config) {
	conf.LoadFixtures = false

	conf.MgoURL = "localhost"
	conf.MgoDB = "myapi"
}
