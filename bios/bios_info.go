package bios

import (
	"embed"
	"encoding/base64"
	"os"
	"strings"
)

const (
	ZIMABLADE = "ZimaBlade"

	ZIMABOARD  = "ZimaBoard"
	ZIMABOARD2 = "ZimaBoard V2"

	ZIMACUBE    = "ZimaCube"
	ZIMACUBEPRO = "ZimaCube Pro"
)

//go:embed assets/*
var assets embed.FS

func GetModel() string {
	src := "/sys/class/dmi/id/board_version"
	data, err := os.ReadFile(src)
	if err != nil {
		return ""
	}

	model := strings.ToLower(strings.TrimSpace(string(data)))
	model = strings.ReplaceAll(model, " ", "")
	model = strings.ReplaceAll(model, "\n", "")

	switch model {
	case "zimacube":
		return ZIMACUBE
	case "zimacubepro":
		return ZIMACUBEPRO
	case "zmb1.0":
		return ZIMABOARD
	case "":
		return ZIMABOARD2
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

func GetDeviceImageByModel() (string, error) {
	getImageBase64 := func(name string) (string, error) {
		data, err := assets.ReadFile(name)
		if err != nil {
			return "", err
		}

		imgBase64 := base64.StdEncoding.EncodeToString(data)

		return imgBase64, nil
	}

	model := GetModel()
	switch model {
	case ZIMACUBE:
		return getImageBase64("assets/cube.png")
	case ZIMACUBEPRO:
		return getImageBase64("assets/cube.png")
	case ZIMABOARD:
		return getImageBase64("assets/board.png")
	case ZIMABOARD2:
		return getImageBase64("assets/board2.png")
	case ZIMABLADE:
		return getImageBase64("assets/board.png")
	default:
		return getImageBase64("assets/other.png")
	}
}
