package diff

import (
	"os"
	"testing"
)

func TestIntegration_SimpleFixtures(t *testing.T) {
	// Load fixture files
	simpleLockA, err := os.ReadFile("../testdata/composer-simple.lock")
	if err != nil {
		t.Skipf("fixture file not found: %v", err)
	}

	simpleJsonA, err := os.ReadFile("../testdata/composer-simple.json")
	if err != nil {
		t.Skipf("fixture file not found: %v", err)
	}

	// Test 1: Diff simple.lock with itself (no changes)
	out, err := Diff(simpleLockA, simpleLockA, simpleJsonA, simpleJsonA)
	if err != nil {
		t.Errorf("Diff() error = %v", err)
		return
	}

	if len(out.Packages) != 0 {
		t.Errorf("Diff() got %d packages, want 0 (identical files)", len(out.Packages))
	}
}

func TestIntegration_InvalidJSON(t *testing.T) {
	invalidJson, err := os.ReadFile("../testdata/composer-invalid.json")
	if err != nil {
		t.Skipf("fixture file not found: %v", err)
	}

	lock := []byte(`{"packages": []}`)

	_, err = Diff(lock, lock, invalidJson, invalidJson)
	if err == nil {
		t.Errorf("Diff() expected error for invalid JSON, got nil")
	}
}

func TestIntegration_InvalidFormat(t *testing.T) {
	invalidFormat, err := os.ReadFile("../testdata/composer-invalid-format.json")
	if err != nil {
		t.Skipf("fixture file not found: %v", err)
	}

	lock := []byte(`{"packages": []}`)
	validJson := []byte(`{}`)

	// composer-invalid-format.json is valid JSON but not a valid lock file
	// It should fail when used as a lock file, not as a json file
	_, err = Diff(invalidFormat, lock, validJson, validJson)
	if err == nil {
		t.Errorf("Diff() expected error for invalid format lock file, got nil")
	}
}
