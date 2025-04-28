package device

import (
	"os"

	"github.com/tidwall/gjson"
)

const (
	ZIMACUBE    = "ZimaCube"
	ZIMACUBEPRO = "ZimaCube-Pro"
	ZIMABOARD   = "ZimaBoard"
	ZIMABLADE   = "ZimaBlade"
	ZIMABOARD2  = "ZimaBoard2"
	ZIMAOS      = "ZimaOS"
)

func GetDeviceType() string {
	data, err := os.ReadFile("/run/zimaos/device-info.json")
	if err != nil {
		return ZIMAOS
	}

	return gjson.GetBytes(data, "device.model").String()
}
