package device

import (
	"os"

	"github.com/tidwall/gjson"
)

func GetDeviceType() string {
	data, err := os.ReadFile("/run/zimaos/device-info.json")
	if err != nil {
		return ""
	}

	return gjson.GetBytes(data, "device.type").String()
}
