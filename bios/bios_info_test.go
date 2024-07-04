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
}
