package cmd

import (
	"fmt"
	"os"
	"syscall"

	_ "github.com/lib/pq"
	"github.com/param108/profile/api/store"
	"github.com/urfave/cli/v2"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/term"
)

var (
	createUserUsername string

	createUserCommand = &cli.Command{
		Name:   "create_user",
		Usage:  "create a new email user with username and password",
		Action: createUserCmd,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "username",
				Usage:       "username for the new user",
				Required:    true,
				Destination: &createUserUsername,
			},
		},
	}
)

func createUserCmd(c *cli.Context) error {
	if createUserUsername == "" {
		return fmt.Errorf("username is required")
	}

	// Get password from user input (hidden)
	fmt.Print("Enter password: ")
	passwordBytes, err := term.ReadPassword(int(syscall.Stdin))
	if err != nil {
		return fmt.Errorf("failed to read password: %w", err)
	}
	fmt.Println() // Add newline after password input

	password := string(passwordBytes)
	if len(password) < 6 {
		return fmt.Errorf("password must be at least 6 characters")
	}

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("failed to hash password: %w", err)
	}

	// Create store
	db, err := store.NewStore()
	if err != nil {
		return fmt.Errorf("failed to connect to database: %w", err)
	}

	// Get writer from environment
	writer := os.Getenv("WRITER")
	if writer == "" {
		return fmt.Errorf("WRITER environment variable is required")
	}

	// Create the user
	emailUser, err := db.CreateEmailUser(createUserUsername, string(hashedPassword), writer)
	if err != nil {
		return fmt.Errorf("failed to create user: %w", err)
	}

	fmt.Printf("User created successfully:\n")
	fmt.Printf("ID: %s\n", emailUser.ID)
	fmt.Printf("Username: %s\n", emailUser.UserName)
	fmt.Printf("Writer: %s\n", emailUser.Writer)

	return nil
}

func init() {
	register(createUserCommand)
}