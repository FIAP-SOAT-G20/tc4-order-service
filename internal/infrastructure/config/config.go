package config

import (
	"log"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

type Config struct {
	// Database settings
	DBDSN          string
	DBMaxOpenConns int
	DBMaxIdleConns int
	DBMaxLifetime  time.Duration

	// Server settings
	ServerPort                    string
	ServerReadTimeout             time.Duration
	ServerWriteTimeout            time.Duration
	ServerIdleTimeout             time.Duration
	ServerGracefulShutdownTimeout time.Duration

	// Environment
	Environment string

	// Mercado Pago
	FakeMercadoPagoURL             string
	FakeMercadoPagoNotificationURL string
	MercadoPagoToken               string
	MercadoPagoURL                 string
	MercadoPagoTimeout             time.Duration
	MercadoPagoRetryCount          int
	MercadoPagoNotificationURL     string

	// JWT Settings
	JWTSecret     string
	JWTExpiration time.Duration
}

func LoadConfig() *Config {
	if err := godotenv.Load(); err != nil && !os.IsNotExist(err) {
		log.Printf("Warning: .env file not found or error loading it: %v", err)
	}

	dbMaxOpenConns, _ := strconv.Atoi(getEnv("DB_MAX_OPEN_CONNS", "25"))
	dbMaxIdleConns, _ := strconv.Atoi(getEnv("DB_MAX_IDLE_CONNS", "25"))
	dbMaxLifetime, _ := time.ParseDuration(getEnv("DB_CONN_MAX_LIFETIME", "5m"))

	serverReadTimeout, _ := time.ParseDuration(getEnv("SERVER_READ_TIMEOUT", "10s"))
	serverWriteTimeout, _ := time.ParseDuration(getEnv("SERVER_WRITE_TIMEOUT", "10s"))
	serverIdleTimeout, _ := time.ParseDuration(getEnv("SERVER_IDLE_TIMEOUT", "60s"))
	serverGracefulShutdownTimeout, _ := time.ParseDuration(getEnv("SERVER_GRACEFUL_SHUTDOWN_SEC_TIMEOUT", "5s"))

	mercadoPagoTimeout, _ := time.ParseDuration(getEnv("MERCADO_PAGO_TIMEOUT", "10s"))
	mercadoPagoRetryCount, _ := strconv.Atoi(getEnv("MERCADO_PAGO_RETRY_COUNT", "2"))

	jwtExpirationStr := getEnv("JWT_EXPIRATION", "24h")
	jwtExpiration, err := time.ParseDuration(jwtExpirationStr)
	if err != nil {
		log.Printf("Warning: invalid JWT_EXPIRATION value %q: %v. Using default value 24h.", jwtExpirationStr, err)
		jwtExpiration = 24 * time.Hour
	}

	return &Config{
		// Database settings
		DBDSN:          getEnv("DB_DSN", "host=localhost port=5432 user=postgres password=postgres dbname=fastfood_10soat_g22_tc3 sslmode=disable"),
		DBMaxOpenConns: dbMaxOpenConns,
		DBMaxIdleConns: dbMaxIdleConns,
		DBMaxLifetime:  dbMaxLifetime,

		// Server settings
		ServerPort:                    getEnv("SERVER_PORT", "8080"),
		ServerReadTimeout:             serverReadTimeout,
		ServerWriteTimeout:            serverWriteTimeout,
		ServerIdleTimeout:             serverIdleTimeout,
		ServerGracefulShutdownTimeout: serverGracefulShutdownTimeout,

		// Environment
		Environment: getEnv("ENVIRONMENT", "development"),

		// Mercado Pago
		FakeMercadoPagoURL:             getEnv("FAKE_MERCADO_PAGO_URL", "url"),
		FakeMercadoPagoNotificationURL: getEnv("FAKE_MERCADO_PAGO_NOTIFICATION_URL", "url"),
		MercadoPagoToken:               getEnv("MERCADO_PAGO_TOKEN", "token"),
		MercadoPagoURL:                 getEnv("MERCADO_PAGO_URL", "url"),
		MercadoPagoTimeout:             mercadoPagoTimeout,
		MercadoPagoRetryCount:          mercadoPagoRetryCount,
		MercadoPagoNotificationURL:     getEnv("MERCADO_PAGO_NOTIFICATION_URL", "url"),

		// JWT Settings
		JWTSecret:     getEnv("JWT_SECRET", "SUPER_SECRET_KEY_DONT_TELL_ANYONE"),
		JWTExpiration: jwtExpiration,
	}
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
