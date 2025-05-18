package main

import (
	"strings"
	"testing"
)

func TestBiodataValues(t *testing.T) {
	if Biodata.Name != "Efendi Sugiantoro" {
		t.Errorf("Expected Name to be 'Efendi Sugiantoro', got '%s'", Biodata.Name)
	}
	if Biodata.Age != 25 {
		t.Errorf("Expected Age to be 25, got %d", Biodata.Age)
	}
	if Biodata.Address != "Jl. Veteran No. 21, Kediri, Jawa Timur, Indonesia" {
		t.Errorf("Expected Address to be 'Jl. Veteran No. 21, Kediri, Jawa Timur, Indonesia', got '%s'", Biodata.Address)
	}
	if Biodata.Email != "efendi.sugiantoro@example.com" {
		t.Errorf("Expected Email to be 'efendi.sugiantoro@example.com', got '%s'", Biodata.Email)
	}
}

func TestGetBiodataInfo(t *testing.T) {
	info := GetBiodataInfo()
	if !strings.Contains(info, Biodata.Name) {
		t.Errorf("GetBiodataInfo() output does not contain Name")
	}
	if !strings.Contains(info, "25") {
		t.Errorf("GetBiodataInfo() output does not contain Age")
	}
	if !strings.Contains(info, Biodata.Address) {
		t.Errorf("GetBiodataInfo() output does not contain Address")
	}
	if !strings.Contains(info, Biodata.Email) {
		t.Errorf("GetBiodataInfo() output does not contain Email")
	}
}
