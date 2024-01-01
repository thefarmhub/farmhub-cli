package model

type Log struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	IoTTopic string `json:"iotTopic"`
	Metric   string `json:"metric"`
}
