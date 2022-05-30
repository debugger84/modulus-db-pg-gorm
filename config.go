package db

import (
	application "github.com/debugger84/modulus-application"
	"time"
)

type ModuleConfig struct {
	host                 string
	port                 int
	user                 string
	pass                 string
	name                 string
	sslMode              string
	maxIdleConns         int
	maxOpenConns         int
	connMaxLifetime      time.Duration
	preferSimpleProtocol *bool
	loggingEnabled       *bool
	slowQueryLimit       int
}

func NewModuleConfig() *ModuleConfig {
	return &ModuleConfig{}
}
func (s *ModuleConfig) ProvidedServices() []interface{} {
	return []interface{}{
		NewGormLogger,
		NewDb,
		func() *ModuleConfig {
			return s
		},
	}
}

func (s *ModuleConfig) SetHost(host string) {
	s.host = host
}

func (s *ModuleConfig) SetPort(port int) {
	s.port = port
}

func (s *ModuleConfig) SetUser(user string) {
	s.user = user
}

func (s *ModuleConfig) SetPass(pass string) {
	s.pass = pass
}

func (s *ModuleConfig) SetName(name string) {
	s.name = name
}

func (s *ModuleConfig) SetSslMode(sslMode string) {
	s.sslMode = sslMode
}

func (s *ModuleConfig) SetMaxIdleConns(maxIdleConns int) {
	s.maxIdleConns = maxIdleConns
}

func (s *ModuleConfig) SetMaxOpenConns(maxOpenConns int) {
	s.maxOpenConns = maxOpenConns
}

func (s *ModuleConfig) SetConnMaxLifetime(connMaxLifetime time.Duration) {
	s.connMaxLifetime = connMaxLifetime
}

func (s *ModuleConfig) SetPreferSimpleProtocol(preferSimpleProtocol bool) {
	s.preferSimpleProtocol = &preferSimpleProtocol
}

func (s *ModuleConfig) SetLoggingEnabled(loggingEnabled bool) {
	s.loggingEnabled = &loggingEnabled
}

func (s *ModuleConfig) SetSlowQueryLimit(slowQueryLimit int) {
	s.slowQueryLimit = slowQueryLimit
}

func (s *ModuleConfig) InitConfig(config application.Config) error {
	if s.host == "" {
		s.host = config.GetEnv("DB_HOST")
	}
	if s.port == 0 {
		s.port = config.GetEnvAsInt("DB_PORT")
	}
	if s.user == "" {
		s.user = config.GetEnv("DB_USER")
	}
	if s.pass == "" {
		s.pass = config.GetEnv("DB_PASSWORD")
	}
	if s.name == "" {
		s.name = config.GetEnv("DB_NAME")
	}
	if s.sslMode == "" {
		s.sslMode = config.GetEnv("DB_SSL_MODE")
	}
	if s.maxIdleConns == 0 {
		s.maxIdleConns = config.GetEnvAsInt("DB_MAX_IDLE_CONNS")
	}
	if s.maxOpenConns == 0 {
		s.maxOpenConns = config.GetEnvAsInt("DB_MAX_OPEN_CONNS")
	}
	if s.connMaxLifetime == 0 {
		s.connMaxLifetime = time.Second * time.Duration(config.GetEnvAsInt("DB_CONN_MAX_LIFETIME"))
	}
	if s.maxOpenConns == 0 {
		s.maxOpenConns = config.GetEnvAsInt("DB_MAX_OPEN_CONNS")
	}
	if s.preferSimpleProtocol == nil {
		val := config.GetEnvAsBool("DB_PREFER_SIMPLE_PROTOCOL")
		s.preferSimpleProtocol = &val
	}
	if s.loggingEnabled == nil {
		val := config.GetEnvAsBool("DB_LOGGING_ENABLED")
		s.loggingEnabled = &val
	}
	if s.slowQueryLimit == 0 {
		s.slowQueryLimit = config.GetEnvAsInt("DB_SLOW_QUERY_LOGGING_LIMIT")
	}

	return nil
}
