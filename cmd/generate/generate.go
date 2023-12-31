package generate

import (
	"context"
	"fmt"

	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/thefarmhub/farmhub-cli/internal/fhclient"
	"github.com/thefarmhub/farmhub-cli/internal/kit"
	"github.com/thefarmhub/farmhub-cli/internal/model"
)

var (
	kitFlag string
	project string
)

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

			sensor, err := selectSensor(client, projectId)
			if err != nil {
				return err
			}

			hardware := mustSelectKit()

			output, err := hardware.GenerateCode(sensor)
			if err != nil {
				return err
			}

			fmt.Println(output)

			return nil
		},
	}

	cmd.Flags().StringVarP(&kitFlag, "kit", "k", "", "Select the kit to use")
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
func selectSensor(client *fhclient.Client, projectId string) (*model.Sensor, error) {
	sensors, err := client.GetSensorsByProjectID(context.Background(), projectId)
	if err != nil {
		pterm.Fatal.Println("Error fetching sensors:", err)
		return nil, err
	}

	sensorNames := make([]string, len(sensors))
	sensorMap := make(map[string]*model.Sensor)
	for i, sensor := range sensors {
		sensorName := sensor.Name + " (" + sensor.ID + ")"
		sensorNames[i] = sensorName
		sensorMap[sensorName] = &sensor
	}

	selectedSensorName, err := pterm.DefaultInteractiveSelect.
		WithDefaultText("Select the sensor: ").
		WithOptions(sensorNames).
		Show()
	if err != nil {
		pterm.Fatal.Println("Error selecting sensor:", err)
		return nil, err
	}

	return sensorMap[selectedSensorName], nil
}

func mustSelectKit() kit.Kit {
	var err error
	selectedKit := kitFlag

	if selectedKit == "" {
		selectedKit, err = pterm.DefaultInteractiveSelect.
			WithDefaultText("Select the sensor: ").
			WithOptions(kit.GetKitNames()).
			Show()
		if err != nil {
			pterm.Fatal.Println("Error selecting kit:", err)
		}
	}

	return kit.GetKitByName(selectedKit)
}
