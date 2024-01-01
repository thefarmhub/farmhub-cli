package model

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
	Logs                        []Log  `json:"logs"`
}
