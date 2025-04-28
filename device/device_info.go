package device

import (
	"os"

	"github.com/tidwall/gjson"
)

func GetDeviceType() (string, error) {
	data, err := os.ReadFile("/run/zimaos/device-info.json")
	if err != nil {
		return "", err
	}

	return gjson.GetBytes(data, "device.type").String(), nil
}
