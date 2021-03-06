package main

import "github.com/spf13/pflag"

const (
	// configuration defaults support local development (i.e. "go run ...")

	// Server
	defaultServerAddress = "0.0.0.0"
	defaultServerPort    = "9090"

	// Gateway
	defaultGatewayEnable      = true
	defaultGatewayAddress     = "0.0.0.0"
	defaultGatewayPort        = "8080"
	defaultGatewayURL         = "/myapp/v1/"
	defaultGatewaySwaggerFile = "pkg/pb/service.swagger.json"

	// Database
	defaultDatabaseEnable = true
	// DSN example: "postgres://postgres:postgres@postgres:5432/atlas_db?sslmode=disable"
	defaultDatabaseDSN      = ""
	defaultDatabaseType     = "postgres"
	defaultDatabaseAddress  = "0.0.0.0"
	defaultDatabasePort     = "5432"
	defaultDatabaseName     = "myapp"
	defaultDatabaseUser     = "postgres"
	defaultDatabasePassword = "postgres"
	defaultDatabaseSSL      = "disable"
	defaultDatabaseOption   = ""

	// Health
	defaultInternalEnable    = true
	defaultInternalAddress   = "0.0.0.0"
	defaultInternalPort      = "8081"
	defaultInternalHealth    = "/healthz"
	defaultInternalReadiness = "/ready"

	defaultConfigDirectory = "deploy/"
	defaultConfigFile      = ""
	defaultSecretFile      = ""
	defaultApplicationID   = "myapp"

	// Heartbeat
	defaultKeepaliveTime    = 10
	defaultKeepaliveTimeout = 20

	// Logging
	defaultLoggingLevel = "debug"
)

var (
	// define flag overrides
	flagServerAddress = pflag.String("server.address", defaultServerAddress, "adress of gRPC server")
	flagServerPort    = pflag.String("server.port", defaultServerPort, "port of gRPC server")

	flagGatewayEnable      = pflag.Bool("gateway.enable", defaultGatewayEnable, "enable gatway")
	flagGatewayAddress     = pflag.String("gateway.address", defaultGatewayAddress, "address of gateway server")
	flagGatewayPort        = pflag.String("gateway.port", defaultGatewayPort, "port of gateway server")
	flagGatewayURL         = pflag.String("gateway.endpoint", defaultGatewayURL, "endpoint of gateway server")
	flagGatewaySwaggerFile = pflag.String("gateway.swaggerFile", defaultGatewaySwaggerFile, "directory of swagger.json file")

	flagDatabaseEnable   = pflag.Bool("database.enable", defaultDatabaseEnable, "enable database")
	flagDatabaseDSN      = pflag.String("database.dsn", defaultDatabaseDSN, "DSN of the database")
	flagDatabaseType     = pflag.String("database.type", defaultDatabaseType, "type of the database")
	flagDatabaseAddress  = pflag.String("database.address", defaultDatabaseAddress, "address of the database")
	flagDatabasePort     = pflag.String("database.port", defaultDatabasePort, "port of the database")
	flagDatabaseName     = pflag.String("database.name", defaultDatabaseName, "name of the database")
	flagDatabaseUser     = pflag.String("database.user", defaultDatabaseUser, "database username")
	flagDatabasePassword = pflag.String("database.password", defaultDatabasePassword, "database password")
	flagDatabaseSSL      = pflag.String("database.ssl", defaultDatabaseSSL, "database ssl mode")
	flagDatabaseOption   = pflag.String("database.option", defaultDatabaseOption, "define custom option to db driver")

	flagInternalEnable    = pflag.Bool("internal.enable", defaultInternalEnable, "enable internal http server")
	flagInternalAddress   = pflag.String("internal.address", defaultInternalAddress, "address of internal http server")
	flagInternalPort      = pflag.String("internal.port", defaultInternalPort, "port of internal http server")
	flagInternalHealth    = pflag.String("internal.health", defaultInternalHealth, "endpoint for health checks")
	flagInternalReadiness = pflag.String("internal.readiness", defaultInternalReadiness, "endpoint for readiness checks")

	flagConfigDirectory = pflag.String("config.source", defaultConfigDirectory, "directory of the configuration file")
	flagConfigFile      = pflag.String("config.file", defaultConfigFile, "directory of the configuration file")
	flagSecretFile      = pflag.String("config.secret.file", defaultSecretFile, "directory of the secrets configuration file")
	flagApplicationID   = pflag.String("app.id", defaultApplicationID, "identifier for the application")

	flagKeepaliveTime    = pflag.Int("config.keepalive.time", defaultKeepaliveTime, "default value, in seconds, of the keepalive time")
	flagKeepaliveTimeout = pflag.Int("config.keepalive.timeout", defaultKeepaliveTimeout, "default value, in seconds, of the keepalive timeout")

	flagLoggingLevel = pflag.String("logging.level", defaultLoggingLevel, "log level of application")
)
