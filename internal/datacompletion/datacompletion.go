package datacompletion

import (
	"fmt"
	"reflect"

	"github.com/pterm/pterm"
	"github.com/thefarmhub/farmhub-cli/internal/model"
)

// Complete fills the empty fields of ConfigVariables based on the datatype tag.
func Complete[S any](config *S, sensor *model.Sensor) {
	val := reflect.ValueOf(config).Elem()
	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		datatype := val.Type().Field(i).Tag.Get("datatype")
		name := val.Type().Field(i).Tag.Get("name")
		metric := val.Type().Field(i).Tag.Get("metric")

		switch datatype {
		case "ssid":
			if field.String() == "" {
				ssid, _ := pterm.DefaultInteractiveTextInput.Show("Enter WiFi SSID")
				field.SetString(ssid)
			}
		case "password":
			if field.String() == "" {
				password, _ := pterm.DefaultInteractiveTextInput.Show("Enter WiFi Password")
				field.SetString(password)
			}
		case "topic":
			if field.String() == "" {
				field.SetString(selectLog(name, metric, sensor))
			}
		}
	}
}

// selectLog displays logs and allows the user to select one.
func selectLog(name, metric string, sensor *model.Sensor) string {
	// First, try to find a log that matches the metric directly.
	for _, log := range sensor.Logs {
		if log.Metric == metric {
			return log.IoTTopic
		}
	}

	// If no direct match is found, allow the user to choose.
	logOptions := make([]string, len(sensor.Logs))
	logMap := make(map[string]string)

	for i, log := range sensor.Logs {
		logOptions[i] = log.Name
		logMap[log.Name] = log.IoTTopic
	}

	selectedLogName, _ := pterm.DefaultInteractiveSelect.
		WithOptions(logOptions).
		Show(fmt.Sprintf("Select a notebook for %s:", name))

	// Get the IoTTopic from the map using the selected log name
	if iotTopic, exists := logMap[selectedLogName]; exists {
		return iotTopic
	}

	return ""
}
