package fhclient

import (
	"context"
	"errors"

	"github.com/machinebox/graphql"
	"github.com/spf13/viper"
	"github.com/thefarmhub/farmhub-cli/internal/model"
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
		token: viper.GetString("auth.token"),
	}
}

// SetToken sets the token for the client.
func (c *Client) SetToken(token string) {
	c.token = token
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

	type loginResponse struct {
		AuthLogin struct {
			Token string
		}
	}

	var respData loginResponse
	err := graphql.NewClient(c.endpoint).Run(ctx, req, &respData)
	if err != nil {
		return "", err
	}

	return respData.AuthLogin.Token, nil
}

// GetProjects fetches a list of projects.
func (c *Client) GetProjects(ctx context.Context) ([]model.Project, error) {
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

	type getProjectsResponse struct {
		Viewer struct {
			Projects struct {
				Nodes []model.Project
			}
		}
	}

	var respData getProjectsResponse
	err := graphql.NewClient(c.endpoint).Run(ctx, req, &respData)
	if err != nil {
		return nil, err
	}

	return respData.Viewer.Projects.Nodes, nil
}

// GetSensorsByProjectID fetches sensors for a given project ID.
func (c *Client) GetSensorsByProjectID(ctx context.Context, projectId string) ([]model.Sensor, error) {
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
									metric
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

	type sensorsResponse struct {
		Viewer struct {
			Projects struct {
				Nodes []struct {
					Sensors struct {
						Nodes []model.Sensor
					} `json:"sensors"`
				} `json:"nodes"`
			} `json:"projects"`
		}
	}

	var respData sensorsResponse
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

// CreateSensor creates a new sensor with the provided details.
func (c *Client) CreateSensor(ctx context.Context, projectId, name string) (model.Sensor, error) {
	if projectId == "" || name == "" {
		return model.Sensor{}, errors.New("project ID and name are required")
	}

	req := graphql.NewRequest(`
		mutation CreateSensor($projectId: ID!, $name: String!) {
			sensorCreate(projectId: $projectId, input: { name: $name }) {
				id
				name
				description
				active
				endpoint
				iotThingName
				iotCertificatePem
				iotCertificatePrivateKey
				iotRootCertificateAuthority
			}
		}
	`)

	req.Var("projectId", projectId)
	req.Var("name", name)

	// Set the Authorization header
	req.Header.Set("Authorization", "Bearer "+c.token)

	type createSensorResponse struct {
		SensorCreate model.Sensor
	}

	var respData createSensorResponse
	err := graphql.NewClient(c.endpoint).Run(ctx, req, &respData)
	if err != nil {
		return model.Sensor{}, err
	}

	return respData.SensorCreate, nil
}

// CreateLog creates a new log with the provided details.
func (c *Client) CreateLog(ctx context.Context, projectId, name, metric string) (*model.Log, error) {
	if projectId == "" || name == "" || metric == "" {
		return nil, errors.New("project ID, name, and metric are required")
	}

	req := graphql.NewRequest(`
		mutation CreateLog($projectId: ID!, $name: String!, $metric: MetricEnum!) {
			logCreate(
				projectId: $projectId
				input: {
					name: $name
					metric: $metric
				}
			) {
				id
				name
				metric
				iotTopic
			}
		}
	`)

	req.Var("projectId", projectId)
	req.Var("name", name)
	req.Var("metric", metric)

	// Set the Authorization header
	req.Header.Set("Authorization", "Bearer "+c.token)

	type createLogResponse struct {
		LogCreate model.Log
	}

	var respData createLogResponse
	err := graphql.NewClient(c.endpoint).Run(ctx, req, &respData)
	if err != nil {
		return nil, err
	}

	return &respData.LogCreate, nil
}

// UpdateSensorLogs updates the logs associated with a sensor.
func (c *Client) UpdateSensorLogs(ctx context.Context, sensorId string, logIds []string) ([]model.Log, error) {
	if sensorId == "" || len(logIds) == 0 {
		return nil, errors.New("sensor ID and logs are required")
	}

	req := graphql.NewRequest(`
		mutation UpdateSensorLogs($sensorId: ID!, $logs: [ID!]) {
			sensorUpdateLogs(sensorId: $sensorId, input: { logs: $logs }) {
				id
				logs {
					id
					name
					iotTopic
				}
			}
		}
	`)

	req.Var("sensorId", sensorId)
	req.Var("logs", logIds)

	// Set the Authorization header
	req.Header.Set("Authorization", "Bearer "+c.token)

	type updateSensorLogsResponse struct {
		SensorUpdateLogs model.Sensor
	}

	var respData updateSensorLogsResponse
	err := graphql.NewClient(c.endpoint).Run(ctx, req, &respData)
	if err != nil {
		return nil, err
	}

	return respData.SensorUpdateLogs.Logs, nil
}
