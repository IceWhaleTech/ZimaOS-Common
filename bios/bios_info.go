package bios

import (
	"os"
	"strings"
)

const (
	ZIMABLADE = "ZimaBlade"

	ZIMABOARD  = "ZimaBoard"
	ZIMABOARD2 = "ZimaBoard2"

	ZIMACUBE    = "ZimaCube"
	ZIMACUBEPRO = "ZimaCube Pro"
)

func GetModel() string {
	data, err := os.ReadFile("/sys/class/dmi/id/board_name")
	if err != nil {
		return ""
	}
	boardName := strings.ToLower(strings.TrimSpace(string(data)))
	boardName = strings.ReplaceAll(boardName, " ", "")

	data, err = os.ReadFile("/sys/class/dmi/id/board_version")
	if err != nil {
		return ""
	}
	boardVersion := strings.ToLower(strings.TrimSpace(string(data)))
	boardVersion = strings.ReplaceAll(boardVersion, " ", "")

	info := boardName + " " + boardVersion

	switch {
	case strings.Contains(info, "zimacubepro"):
		return ZIMACUBEPRO
	case strings.Contains(info, "zimacube"):
		return ZIMACUBE
	case strings.Contains(info, "zimaboard2"):
		return ZIMABOARD2
	case strings.Contains(info, "zimaboard"):
		return ZIMABOARD
	case strings.Contains(info, "zbb001"):
		return ZIMABLADE
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
