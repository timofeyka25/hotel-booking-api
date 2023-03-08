package db

type Config struct {
	Username string `env:"DB_USERNAME" envDefault:"timofeyka.com.03"`
	Password string `env:"DB_PASSWORD" envDefault:"5jZF7HxWAXob"`
	DbName   string `env:"DB_NAME" envDefault:"hotelbookingdb"`
	Host     string `env:"DB_HOST" envDefault:"ep-ancient-term-974725-pooler.eu-central-1.aws.neon.tech"`
}
