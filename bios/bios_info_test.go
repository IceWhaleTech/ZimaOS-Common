package bios

import "testing"

func TestGetModel(t *testing.T) {
	t.Run("ZimaCube", func(t *testing.T) {
		model := GetModel()
		t.Log("Model:", model)
		if model != ZIMACUBE {
			t.Errorf("Expected %s, got %s", ZIMACUBE, model)
		}
	})
	t.Run("Zimacube Pro", func(t *testing.T) {
		model := GetModel()
		t.Log("Model:", model)
		if model != ZIMACUBEPRO {
			t.Errorf("Expected %s, got %s", ZIMACUBEPRO, model)
		}
	})

	t.Run("ZMB 1.0", func(t *testing.T) {
		model := GetModel()
		t.Log("Model:", model)
		if model != ZIMABOARD {
			t.Errorf("Expected %s, got %s", ZIMABOARD, model)
		}
	})
}

func TestGetDeviceImageByModel(t *testing.T) {
	t.Run("ZimaCube", func(t *testing.T) {
		img, err := GetDeviceImageByModel()
		t.Log("Image:", img)
		if err != nil {
			t.Errorf("Expected no error, got %s", err)
		}
	})
}
