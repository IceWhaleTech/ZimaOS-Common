package bios

import (
	"os"
	"strings"
)

const (
	ZIMACUBE    = "ZimaCube"
	ZIMACUBEPRO = "ZimaCube Pro"
)

func GetModel() string {
	src := "/sys/class/dmi/id/board_version"
	_, err := os.Stat(src)
	if os.IsNotExist(err) {
		return ""
	}

	content, err := os.ReadFile(src)
	if err != nil {
		return ""
	}

	model := strings.ToLower(string(content))
	model = strings.ReplaceAll(model, " ", "")
	model = strings.ReplaceAll(model, "\n", "")

	if model == "zimacube" {
		return ZIMACUBE
	}
	if model == "zimacubepro" {
		return ZIMACUBEPRO
	}

	return ""
}

func GetSerialNumber() (string, error) {
	data, err := os.ReadFile("/sys/class/dmi/id/board_serial")
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(data)), nil
}

func IsIceWhaleProduct() bool {
	b, err := os.ReadFile("/sys/class/dmi/id/board_vendor")
	return err == nil && strings.Contains(strings.ToLower(string(b)), "icewhale")
}
