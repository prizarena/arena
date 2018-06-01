package arena

import (
	"testing"
)

func TestCommaSeparatedUniqueValuesList_Add(t *testing.T) {
	var csvl CommaSeparatedUniqueValuesList
	if csvl = csvl.Add("v1"); csvl != "v1" {
		t.Error("unexpected: " + csvl)
	}
	if csvl = csvl.Add("v3"); csvl != "v1,v3" {
		t.Error("unexpected")
	}
	if csvl = csvl.Add("v2"); csvl != "v1,v3,v2" {
		t.Error("unexpected")
	}
}
func TestCommaSeparatedUniqueValuesList_Remove(t *testing.T) {
	var csvl CommaSeparatedUniqueValuesList
	csvl = CommaSeparatedUniqueValuesList("")
	if csvl = csvl.Remove("v1"); csvl != "" {
		t.Error("unexpected: " + csvl)
	}
	csvl = CommaSeparatedUniqueValuesList("v1")
	if csvl = csvl.Remove("v1"); csvl != "" {
		t.Error("unexpected: " + csvl)
	}
	csvl = CommaSeparatedUniqueValuesList("v2")
	if csvl = csvl.Remove("v1"); csvl != "v2" {
		t.Error("unexpected: " + csvl)
	}
	csvl = CommaSeparatedUniqueValuesList("v1,v3,v2,v4")
	if csvl = csvl.Remove("v1"); csvl != "v3,v2,v4" {
		t.Error("unexpected: " + csvl)
	}
	csvl = CommaSeparatedUniqueValuesList("v1,v3,v2,v4")
	if csvl = csvl.Remove("v2"); csvl != "v1,v3,v4" {
		t.Error("unexpected: " + csvl)
	}
	csvl = CommaSeparatedUniqueValuesList("v1,v3,v2,v4")
	if csvl = csvl.Remove("v3"); csvl != "v1,v2,v4" {
		t.Error("unexpected: " + csvl)
	}
	csvl = CommaSeparatedUniqueValuesList("v1,v3,v2,v4")
	if csvl = csvl.Remove("v4"); csvl != "v1,v3,v2" {
		t.Error("unexpected: " + csvl)
	}
}

func TestCommaSeparatedUniqueValuesList_Contains(t *testing.T) {
	if CommaSeparatedUniqueValuesList("").Contains("v1") {
		t.Error("unexpected true")
	}
	if !CommaSeparatedUniqueValuesList("v1").Contains("v1") {
		t.Error("unexpected false")
	}
	if !CommaSeparatedUniqueValuesList("v1,v2").Contains("v1") {
		t.Error("unexpected false")
	}
	if !CommaSeparatedUniqueValuesList("v1,v2").Contains("v2") {
		t.Error("unexpected false")
	}
	if !CommaSeparatedUniqueValuesList("v1,v2,v3").Contains("v1") {
		t.Error("unexpected false")
	}
	if !CommaSeparatedUniqueValuesList("v1,v2,v3").Contains("v2") {
		t.Error("unexpected false")
	}
	if !CommaSeparatedUniqueValuesList("v1,v2,v3").Contains("v3") {
		t.Error("unexpected false")
	}
}
