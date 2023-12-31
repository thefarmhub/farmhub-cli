package generate

import (
	"context"

	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/thefarmhub/farmhub-cli/internal/fhclient"
)

var project string

// NewGenerateCommand creates a new Cobra command for the generate process.
func NewGenerateCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "generate",
		Short: "Generate sensor data",
		RunE: func(cmd *cobra.Command, args []string) error {
			var err error
			projectId := project

			client := fhclient.NewClient()
			client.SetToken(viper.GetString("auth.token"))

			if projectId == "" {
				projectId, err = selectProject(client)
				if err != nil {
					return err
				}
			}

			selectedSensorId, err := selectSensor(client, projectId)
			if err != nil {
				return err
			}

			pterm.Success.Println("Selected sensor ID:", selectedSensorId)
			return nil
		},
	}

	cmd.Flags().StringVarP(&project, "project", "p", "", "Select the project by ID")

	return cmd
}

// selectProject allows the user to select a project interactively.
func selectProject(client *fhclient.Client) (string, error) {
    projects, err := client.GetProjects(context.Background())
    if err != nil {
        pterm.Fatal.Println("Error fetching projects:", err)
        return "", err
    }

    projectNames := make([]string, len(projects))
    projectMap := make(map[string]string)
    for i, project := range projects {
        projectName := project.Name + " (" + project.ID + ")"
        projectNames[i] = projectName
        projectMap[projectName] = project.ID
    }

    selectedProject, err := pterm.DefaultInteractiveSelect.
        WithDefaultText("Select the project: ").
        WithOptions(projectNames).
        Show()
    if err != nil {
        pterm.Fatal.Println("Error selecting project:", err)
        return "", err
    }

    return projectMap[selectedProject], nil
}


// selectSensor allows the user to select a sensor from a list.
func selectSensor(client *fhclient.Client, projectId string) (string, error) {
    sensors, err := client.GetSensorsByProjectID(context.Background(), projectId)
    if err != nil {
        pterm.Fatal.Println("Error fetching sensors:", err)
        return "", err
    }

    sensorNames := make([]string, len(sensors))
    sensorMap := make(map[string]string)
    for i, sensor := range sensors {
        sensorName := sensor.Name + " (" + sensor.ID + ")"
        sensorNames[i] = sensorName
        sensorMap[sensorName] = sensor.ID
    }

    selectedSensor, err := pterm.DefaultInteractiveSelect.
        WithDefaultText("Select the sensor: ").
        WithOptions(sensorNames).
        Show()
    if err != nil {
        pterm.Fatal.Println("Error selecting sensor:", err)
        return "", err
    }

    return sensorMap[selectedSensor], nil
}

