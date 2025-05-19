package httpd

// Config contains server setup.
type Config struct {
	Interface string `env:"HTTPD_INTERFACE"`
	Port      string `env:"HTTPD_PORT"`
}
