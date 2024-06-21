package config

import (
	"github.com/spf13/viper"
	"os"
	"testing"
)

func TestNewEnv(t *testing.T) {
	t.Run("should return environment variables", func(t *testing.T) {
		os.Setenv("USER_SERVICE_DB_USERNAME", "test_username")
		os.Setenv("USER_SERVICE_DB_PASSWORD", "test_password")
		os.Setenv("USER_SERVICE_DB_HOST", "test_host")
		os.Setenv("USER_SERVICE_DB_PORT", "test_port")
		os.Setenv("USER_SERVICE_DB_NAME", "test_dbname")
		os.Setenv("USER_SERVICE_DB_MAX_IDLE_CONNS", "10")
		os.Setenv("USER_SERVICE_DB_MAX_OPEN_CONNS", "20")
		os.Setenv("USER_SERVICE_DB_MAX_LIFETIME", "30")

		env := NewEnv()

		if env.DBUsername != "test_username" {
			t.Errorf("Expected DBUsername to be 'test_username', got %s", env.DBUsername)
		}

		if env.DBPassword != "test_password" {
			t.Errorf("Expected DBPassword to be 'test_password', got %s", env.DBPassword)
		}

		if env.DBHost != "test_host" {
			t.Errorf("Expected DBHost to be 'test_host', got %s", env.DBHost)
		}

		if env.DBPort != "test_port" {
			t.Errorf("Expected DBPort to be 'test_port', got %s", env.DBPort)
		}

		if env.DBName != "test_dbname" {
			t.Errorf("Expected DBName to be 'test_dbname', got %s", env.DBName)
		}

		if env.DBMaxIdleConns != 10 {
			t.Errorf("Expected DBMaxIdleConns to be 10, got %d", env.DBMaxIdleConns)
		}

		if env.DBMaxOpenConns != 20 {
			t.Errorf("Expected DBMaxOpenConns to be 20, got %d", env.DBMaxOpenConns)
		}

		if env.DBMaxLifetime != 30 {
			t.Errorf("Expected DBMaxLifetime to be 30, got %d", env.DBMaxLifetime)
		}
	})

	t.Run("should return zero values when environment variables are not set", func(t *testing.T) {
		viper.Reset()
		os.Clearenv()

		env := NewEnv()

		if env.DBUsername != "" {
			t.Errorf("Expected DBUsername to be '', got %s", env.DBUsername)
		}

		if env.DBPassword != "" {
			t.Errorf("Expected DBPassword to be '', got %s", env.DBPassword)
		}

		if env.DBHost != "" {
			t.Errorf("Expected DBHost to be '', got %s", env.DBHost)
		}

		if env.DBPort != "" {
			t.Errorf("Expected DBPort to be '', got %s", env.DBPort)
		}

		if env.DBName != "" {
			t.Errorf("Expected DBName to be '', got %s", env.DBName)
		}

		if env.DBMaxIdleConns != 0 {
			t.Errorf("Expected DBMaxIdleConns to be 0, got %d", env.DBMaxIdleConns)
		}

		if env.DBMaxOpenConns != 0 {
			t.Errorf("Expected DBMaxOpenConns to be 0, got %d", env.DBMaxOpenConns)
		}

		if env.DBMaxLifetime != 0 {
			t.Errorf("Expected DBMaxLifetime to be 0, got %d", env.DBMaxLifetime)
		}
	})
}
