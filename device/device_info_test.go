package device

import "testing"

func TestGetDeviceModel(t *testing.T) {
	t.Run("", func(t *testing.T) {
		model := GetDeviceModel()
		t.Log("Model:", model)
	})
}
