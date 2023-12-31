package login

import (
	"context"
	"errors"
	"fmt"
	"os"
	"regexp"

	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/stripe/stripe-cli/pkg/ansi"
	"github.com/thefarmhub/farmhub-cli/internal/fhclient"
)

func NewLoginCommand() *cobra.Command {
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9.!#$%&'*+/=?^_{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$`)

	loginCmd := &cobra.Command{
		Use:   "login",
		Short: "Log in to your account",
		RunE: func(cmd *cobra.Command, args []string) error {
			email, err := promptForInput("Email", validateEmail(emailRegex))
			if err != nil {
				return err
			}

			password, err := promptForInput("Password", validatePassword)
			if err != nil {
				return err
			}

			err = performLogin(email, password)
			if err == nil {
				fmt.Println("You have been successfully logged in.")
			}

			return err
		},
	}

	return loginCmd
}

// validateEmail returns a function that validates an email address.
func validateEmail(regex *regexp.Regexp) func(string) error {
	return func(input string) error {
		if !regex.MatchString(input) || len(input) < 3 || len(input) > 254 {
			return errors.New("invalid email address")
		}
		return nil
	}
}

// validatePassword validates the input password.
func validatePassword(input string) error {
	if len(input) < 3 {
		return errors.New("invalid password")
	}
	return nil
}

// promptForInput prompts the user for input and validates it using the provided validation function.
func promptForInput(label string, validate func(string) error) (string, error) {
	prompt := promptui.Prompt{
		Label:    label,
		Validate: validate,
	}
	if label == "Password" {
		prompt.Mask = '•'
	}
	return prompt.Run()
}

// performLogin performs the login process.
func performLogin(email, password string) error {
	s := ansi.StartNewSpinner("Logging you in...", os.Stdout)
	defer ansi.StopSpinner(s, "", os.Stdout)

	client := fhclient.NewClient()
	token, err := client.Login(context.Background(), email, password)
	if err != nil {
		return err
	}

	viper.Set("auth.token", token)
	return viper.WriteConfig()
}
