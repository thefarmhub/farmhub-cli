package fhclient

import (
	"context"
	"errors"

	"github.com/machinebox/graphql"
)

const defaultEndpoint = "https://api.farmhub.ag/graphql"

// Client represents an authentication client.
type Client struct {
	endpoint string
	token    string
}

// NewClient creates a new authentication client with the given token.
func NewClient() *Client {
	return &Client{
		endpoint: defaultEndpoint,
	}
}

// SetToken sets the token for the client.
func (c *Client) SetToken(token string) {
	c.token = token
}

// LoginResponse represents the response structure for the Login mutation.
type LoginResponse struct {
	AuthLogin struct {
		Token string
	}
}

// Login performs a login request with the provided email and password.
func (c *Client) Login(ctx context.Context, email, password string) (string, error) {
	if email == "" || password == "" {
		return "", errors.New("email and password are required")
	}

	req := graphql.NewRequest(`
		mutation Login($email: String!, $password: String!, $source: String) {
			authLogin(email: $email, password: $password, source: $source) {
				token
			}
		}
	`)

	req.Var("email", email)
	req.Var("password", password)
	req.Var("source", "farmhub-cli")

	var respData LoginResponse
	err := graphql.NewClient(c.endpoint).Run(ctx, req, &respData)
	if err != nil {
		return "", err
	}

	return respData.AuthLogin.Token, nil
}
