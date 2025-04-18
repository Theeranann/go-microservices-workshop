package configs

type Configs struct {
	PostgreSQL PostgreSQL
	App        Fiber
	Redis    Redis
}

type Fiber struct {
	Host string
	Port string
	Cors string
}

// Database
type PostgreSQL struct {
	Host     string
	Port     string
	Protocol string
	Username string
	Password string
	Database string
	SSLMode  string
}

type Redis struct {
	Host     string
	Port     string
	Password string
}