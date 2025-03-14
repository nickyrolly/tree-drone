package main

import (
	"os"
	"testing"

	"github.com/nickyrolly/tree-drone/handler"
	"github.com/nickyrolly/tree-drone/repository"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

func TestNewServer(t *testing.T) {
	// --- Test Case 1: Database URL from configuration ---
	t.Run("Database URL from Config", func(t *testing.T) {
		// Setup Viper configuration
		cfg := viper.New()
		cfg.Set("database.url", "postgres://localhost:5432/tree_drone?sslmode=disable")
		cfg.Set("database.driver", "postgres")

		// Call the function
		server := newServer(cfg)

		// Assertions
		assert.NotNil(t, server)
		assert.NotNil(t, server.Repository)
	})

	// --- Test Case 2: Database URL from environment variable ---
	t.Run("Database URL from Environment", func(t *testing.T) {
		// Setup environment variable
		os.Setenv("DATABASE_URL", "postgres://localhost:5432/tree_drone?sslmode=disable")
		defer os.Unsetenv("DATABASE_URL") // Clean up after the test

		// Setup Viper configuration (no database.url)
		cfg := viper.New()
		cfg.Set("database.driver", "postgres")

		// Call the function
		server := newServer(cfg)

		// Assertions
		assert.NotNil(t, server)
		assert.NotNil(t, server.Repository)
	})

	// --- Test Case 3: Empty database driver ---
	t.Run("Empty database driver", func(t *testing.T) {
		// Setup environment variable
		os.Setenv("DATABASE_URL", "postgres://localhost:5432/tree_drone?sslmode=disable")
		defer os.Unsetenv("DATABASE_URL") // Clean up after the test

		// Setup Viper configuration (no database.url and database.driver)
		cfg := viper.New()
		cfg.Set("database.url", "")

		// Call the function
		server := newServer(cfg)

		// Assertions
		assert.NotNil(t, server)
		assert.NotNil(t, server.Repository)
	})
	// --- Test Case 4: Empty all ---
	t.Run("Empty All", func(t *testing.T) {
		// Setup environment variable
		defer os.Unsetenv("DATABASE_URL") // Clean up after the test

		// Setup Viper configuration (no database.url and database.driver)
		cfg := viper.New()

		// Call the function
		server := newServer(cfg)

		// Assertions
		assert.NotNil(t, server)
		assert.NotNil(t, server.Repository)
	})
}

// mock
type MockRepository struct {
	Options repository.NewRepositoryOptions
}

func (m *MockRepository) SetEstate(estate *repository.Estate) error {
	return nil
}
func (m *MockRepository) CreateEstate(e *repository.Estate) error {
	return nil
}

func (m *MockRepository) GetOptions() repository.NewRepositoryOptions {
	return m.Options
}

func NewMockRepository(opts repository.NewRepositoryOptions) *MockRepository {
	return &MockRepository{
		Options: opts,
	}
}

// ensure that MockRepository implement RepositoryInterface
var _ repository.RepositoryInterface = (*MockRepository)(nil)

func TestMain(m *testing.M) {
	// Set up anything needed before running the tests.
	// For example, create a test database, etc.

	// Run the tests.
	exitCode := m.Run()

	// Tear down anything needed after running the tests.
	// For example, drop the test database, etc.

	// Exit with the appropriate code.
	os.Exit(exitCode)
}

func TestHandlerServerCreation(t *testing.T) {
	// Create a mock repository.
	mockRepo := NewMockRepository(repository.NewRepositoryOptions{
		Driver: "postgres",
		Url:    "postgres://localhost:5432/tree_drone?sslmode=disable",
	})

	// Create server options with the mock repository.
	opts := handler.NewServerOptions{
		Repository: mockRepo,
	}

	// Create the server.
	server := handler.NewServer(opts)

	// Assertions.
	assert.NotNil(t, server)
	assert.NotNil(t, server.Repository)          //Check if the server.Repository is not nil.
	optsRepo := server.Repository.GetOptions()   // We call getOptions() method.
	assert.Equal(t, "postgres", optsRepo.Driver) //Now we check the values using the getOptions().
	assert.Equal(t, "postgres://localhost:5432/tree_drone?sslmode=disable", optsRepo.Url)
}
