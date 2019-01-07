package packages

import (
	"testing"
)

func TestParseYumLine(t *testing.T) {
	testData := "CharLS.x86_64                                     1.0-17.fc29                            @fedora"
	testPackage := ParseYumLine(testData)
	if testPackage.Name != "CharLS.x86_64" {
		t.Errorf("Received value was incorrect, got: %s, want: %s", testPackage.Name, "CharLS.x86_64")
	}
	if testPackage.Version != "1.0-17.fc29" {
		t.Errorf("Received value was incorrect, got: %s, want: %s", testPackage.Version, "1.0-17.fc29")
	}
}
