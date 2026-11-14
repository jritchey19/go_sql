package main

import (
	"testing"
	"path/filepath"
)

func TestCreateDB(t *testing.T) {
	tmpDir := t.TempDir()
	testFile := filepath.Join(tmpDir, "testDB.db")
	_, err := CreateDB(testFile)
	if err != nil {
		t.Fatalf("failed to create db at %v, got %v\n", testFile, err)
	}
}
