package spacer

import (
	"bytes"
	"os"
)

const SpacerPath = "/var/lib/casaos/spacer"

func InitSpacer() error {
	if _, err := os.Stat(SpacerPath); os.IsNotExist(err) {
		os.WriteFile(SpacerPath, bytes.Repeat([]byte{0}, 20<<20), 0o777)
	}
	return nil
}

func WithSpacer(f func() error) (err error) {
	defer InitSpacer()
	if err = f(); err != nil {
		os.Remove(SpacerPath)
		err = f()
	}
	return err
}
