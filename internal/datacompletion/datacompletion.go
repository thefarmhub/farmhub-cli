package datacompletion

import (
	"context"
	"errors"
	"fmt"
	"reflect"

	"github.com/pterm/pterm"
	"github.com/thefarmhub/farmhub-cli/internal/fhclient"
	"github.com/thefarmhub/farmhub-cli/internal/model"
)

type Completer struct {
	client  *fhclient.Client
	sensor  *model.Sensor
	project *model.Project
}

func NewCompleter(project *model.Project, sensor *model.Sensor) *Completer {
	return &Completer{
		client:  fhclient.NewClient(),
		sensor:  sensor,
		project: project,
	}
}

// Complete fills the empty fields of ConfigVariables based on the datatype tag.
func (c *Completer) Complete(config any) {
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
				topic, err := c.selectLog(name, metric)
				if err != nil {
					pterm.Error.Println(err)
					continue
				}
				field.SetString(topic)
			}
		}
	}
}

// selectLog displays logs and allows the user to select one
func (c *Completer) selectLog(name, metric string) (string, error) {
	for _, log := range c.sensor.Logs {
		if log.Metric == metric {
			return log.IoTTopic, nil
		}
	}

	msg := fmt.Sprintf("Notebook not found for %s. Do you want to create a new log?", name)
	createLog, _ := pterm.DefaultInteractiveConfirm.Show(msg)
	if !createLog {
		return "", errors.New("log creation declined")
	}

	newLog, err := c.client.CreateLog(context.Background(), c.project.ID, name, metric)
	if err != nil {
		return "", err
	}

	c.sensor.Logs = append(c.sensor.Logs, *newLog)

	var logIds []string
	for _, log := range c.sensor.Logs {
		logIds = append(logIds, log.ID)
	}

	_, err = c.client.UpdateSensorLogs(context.Background(), c.sensor.ID, logIds)
	if err != nil {
		return "", err
	}

	return newLog.IoTTopic, nil
}
