package device

import "testing"

func TestGetDeviceType(t *testing.T) {
	t.Run("", func(t *testing.T) {
		model := GetDeviceModel()
		t.Log("Model:", model)
	})
}
