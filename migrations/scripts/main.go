package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

const (
	postgresURLEnv = "POSTGRES_URL"
	migrationsPath = "file://migrations"
)

// Config holds application configuration
type Config struct {
	PostgresURL string
}

// LoadConfig loads configuration from environment variables
func LoadConfig() (*Config, error) {
	// Load .env file if present (fails silently if not found)
	_ = godotenv.Load(".env")

	postgresURL, exists := os.LookupEnv(postgresURLEnv)
	if !exists {
		return nil, fmt.Errorf("environment variable %s not set", postgresURLEnv)
	}

	return &Config{
		PostgresURL: postgresURL,
	}, nil
}

// NewMigrator creates a new database migrator
func NewMigrator(dbURL string) (*migrate.Migrate, error) {
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	// Verify database connection
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("database connection test failed: %w", err)
	}

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to create postgres driver: %w", err)
	}

	m, err := migrate.NewWithDatabaseInstance(migrationsPath, "postgres", driver)
	if err != nil {
		return nil, fmt.Errorf("failed to create migrator: %w", err)
	}

	return m, nil
}

// GetVersion retrieves and displays the current migration version
func GetVersion(m *migrate.Migrate) (uint, bool, error) {
	version, dirty, err := m.Version()
	if err != nil && err != migrate.ErrNilVersion {
		return 0, false, fmt.Errorf("failed to get version: %w", err)
	}

	if err == migrate.ErrNilVersion {
		fmt.Println("Database has no migration version (not initialized)")
		return 0, false, nil
	}

	fmt.Printf("Current database version: %d (dirty: %v)\n", version, dirty)
	return version, dirty, nil
}

// MigrateUp migrates the database to the latest version
func MigrateUp(m *migrate.Migrate) error {
	fmt.Println("Migrating database up to latest version...")

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		return fmt.Errorf("migration failed: %w", err)
	} else if err == migrate.ErrNoChange {
		fmt.Println("Database already at latest version")
	} else {
		fmt.Println("Database migration successful")
	}

	return verifyMigration(m)
}

// MigrateDown rolls back the most recent migration
func MigrateDown(m *migrate.Migrate) error {
	fmt.Println("Rolling back the most recent migration...")

	if err := m.Steps(-1); err != nil {
		return fmt.Errorf("rollback failed: %w", err)
	}

	fmt.Println("Database rollback successful")
	return verifyMigration(m)
}

// MigrateToVersion migrates to a specific version
func MigrateToVersion(m *migrate.Migrate, version uint) error {
	fmt.Printf("Migrating database to version %d...\n", version)

	if err := m.Migrate(version); err != nil && err != migrate.ErrNoChange {
		return fmt.Errorf("migration to version %d failed: %w", version, err)
	} else if err == migrate.ErrNoChange {
		fmt.Printf("Database already at version %d\n", version)
	} else {
		fmt.Printf("Database successfully migrated to version %d\n", version)
	}

	return verifyMigration(m)
}

// verifyMigration checks if the database is in a dirty state after migration
func verifyMigration(m *migrate.Migrate) error {
	_, dirty, err := GetVersion(m)
	if err != nil {
		return err
	}

	if dirty {
		return fmt.Errorf("database is in a dirty state; manual intervention required")
	}

	return nil
}

func main() {
	// Define command-line flags
	upFlag := flag.Bool("up", false, "Migrate database to the latest version")
	downFlag := flag.Bool("down", false, "Roll back the most recent migration")
	versionFlag := flag.Uint("to", 0, "Migrate to a specific version")
	showVersionFlag := flag.Bool("version", false, "Show current database version")
	flag.Parse()

	// Load configuration
	config, err := LoadConfig()
	if err != nil {
		log.Fatalf("Configuration error: %v", err)
	}

	// Create migrator
	m, err := NewMigrator(config.PostgresURL)
	if err != nil {
		log.Fatalf("Initialization error: %v", err)
	}
	defer m.Close()

	// Execute the requested operation
	switch {
	case *upFlag:
		if err := MigrateUp(m); err != nil {
			log.Fatalf("Error: %v", err)
		}
	case *downFlag:
		if err := MigrateDown(m); err != nil {
			log.Fatalf("Error: %v", err)
		}
	case *versionFlag > 0:
		if err := MigrateToVersion(m, *versionFlag); err != nil {
			log.Fatalf("Error: %v", err)
		}
	case *showVersionFlag:
		if _, _, err := GetVersion(m); err != nil {
			log.Fatalf("Error: %v", err)
		}
	default:
		fmt.Println("No action specified. Use -up, -down, -to, or -version flags.")
		flag.PrintDefaults()
	}
}
