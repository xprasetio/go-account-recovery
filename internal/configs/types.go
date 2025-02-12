package configs

type (
	Config struct {
		Service       Service
		Database      DatabaseConfig
		EmailConfig    EmailConfig
	}

	Service struct {
		Port      string
		SecretKey string
	}

	DatabaseConfig struct {
		DataSourceName string
	}
	
	EmailConfig struct {
		SMTPHost     string
		SMTPPort     int
		SMTPUsername string
		SMTPPassword string
		FromEmail    string
		FromName     string
	}
)
