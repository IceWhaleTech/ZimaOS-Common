package bios

import (
	"io"
	"os"
	"strings"
)

const (
	ZIMACUBE    = "ZimaCube"
	ZIMACUBEPRO = "ZimaCubePro"
)

func GetModel() string {
	src := "/sys/class/dmi/id/board_version"
	_, err := os.Stat(src)
	if os.IsNotExist(err) {
		return ""
	} else {
		file, err := os.Open(src)
		if err != nil {
			return ""
		}
		defer file.Close()
		content, err := io.ReadAll(file)
		if err != nil {
			return ""
		}
		if string(content) == "ZimaCube" {
			return ZIMACUBE
		}
		if string(content) == "ZimaCubePro" {
			return ZIMACUBEPRO
		}
		return ""

	}
}
func GetSerialNumber() string {
	src := "/sys/class/dmi/id/board_version"
	_, err := os.Stat(src)
	//ccc
	if os.IsNotExist(err) {
		return ""
	} else {
		file, err := os.Open(src)
		if err != nil {
			return ""
		}
		defer file.Close()
		content, err := io.ReadAll(file)
		if err != nil {
			return ""
		}
		return string(content)
	}
}

func IsIceWhaleProduct() bool {
	src := "/sys/class/dmi/id/board_vendor"
	_, err := os.Stat(src)
	if os.IsNotExist(err) {
		return false
	} else {
		file, err := os.Open(src)
		if err != nil {
			return false
		}
		defer file.Close()
		content, err := io.ReadAll(file)
		if err != nil {
			return false
		}
		if strings.Contains(strings.ToLower(string(content)), "icewhale") {
			return true
		}
		return false
	}
}
