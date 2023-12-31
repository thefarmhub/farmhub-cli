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

// Project represents a single project with minimal fields.
type Project struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Type string `json:"mode"`
}

// ProjectsResponse represents the response structure for the projects query.
type GetProjectsResponse struct {
	Viewer struct {
		Projects struct {
			Nodes []Project
		}
	}
}

// GetProjects fetches a list of projects.
func (c *Client) GetProjects(ctx context.Context) ([]Project, error) {
	req := graphql.NewRequest(`
		query GetAllProjects {
			viewer {
				projects {
					nodes {
						id
						name
						mode
					}
				}
			}
		}
	`)

	// Set the Authorization header
	req.Header.Set("Authorization", "Bearer "+c.token)

	var respData GetProjectsResponse
	err := graphql.NewClient(c.endpoint).Run(ctx, req, &respData)
	if err != nil {
		return nil, err
	}

	return respData.Viewer.Projects.Nodes, nil
}

// Sensor represents detailed information about a sensor.
type Sensor struct {
	ID                          string `json:"id"`
	Name                        string `json:"name"`
	Description                 string `json:"description"`
	Active                      bool   `json:"active"`
	Endpoint                    string `json:"endpoint"`
	IoTThingName                string `json:"iotThingName"`
	IoTCertificatePem           string `json:"iotCertificatePem"`
	IoTCertificatePrivateKey    string `json:"iotCertificatePrivateKey"`
	IoTRootCertificateAuthority string `json:"iotRootCertificateAuthority"`
	Logs []struct {
		ID      string `json:"id"`
		Name    string `json:"name"`
		IoTTopic string `json:"iotTopic"`
	} `json:"logs"`
}

// SensorsResponse represents the response structure for the sensors query.
type SensorsResponse struct {
	Viewer struct {
		Projects struct {
			Nodes []struct {
				Sensors struct {
					Nodes []Sensor
				} `json:"sensors"`
			} `json:"nodes"`
		} `json:"projects"`
	}
}

// GetSensorsByProjectID fetches sensors for a given project ID.
func (c *Client) GetSensorsByProjectID(ctx context.Context, projectId string) ([]Sensor, error) {
	if projectId == "" {
		return nil, errors.New("project ID is required")
	}

	req := graphql.NewRequest(`
		query GetSensors($projectId: ID!) {
			viewer {
				projects(id: $projectId) {
					nodes {
						sensors {
							nodes {
								id
								name
								description
								active
								endpoint
								iotThingName
								iotCertificatePem
								iotCertificatePrivateKey
								iotRootCertificateAuthority
								logs {
									id
									name
									iotTopic
								}
							}
						}
					}
				}
			}
		}
	`)

	req.Var("projectId", projectId)

	// Set the Authorization header
	req.Header.Set("Authorization", "Bearer "+c.token)

	var respData SensorsResponse
	err := graphql.NewClient(c.endpoint).Run(ctx, req, &respData)
	if err != nil {
		return nil, err
	}

	// Assuming that projects will have at least one node.
	if len(respData.Viewer.Projects.Nodes) == 0 {
		return nil, errors.New("no project found with the given ID")
	}

	return respData.Viewer.Projects.Nodes[0].Sensors.Nodes, nil
}
