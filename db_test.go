package main

import (
	"testing"
	"github.com/google/uuid"
	"path/filepath"
)

func TestCreateDB(t *testing.T) {
	tmpDir := t.TempDir()
	testFile := filepath.Join(tmpDir, "testDB.db")
	db, err := CreateDB(testFile)
	if err != nil {
	  db.DBClose()
		t.Fatalf("failed to create db at %v, got %v\n", testFile, err)
	}
	db.DBClose()
}

func TestCreateUser(t *testing.T) {
	tmpDir := t.TempDir()
	testFile := filepath.Join(tmpDir, "testDB.db")
	db, err := CreateDB(testFile)
	if err != nil {
	  db.DBClose()
		t.Fatalf("failed to create db at %v, got %v\n", testFile, err)
	}
	defer db.DBClose()

	tests := []struct {
		name string
		userName, userEmail string
		want bool
		validUUID bool
	}{
		{name: "create user", userName: "test", userEmail: "test@test.com", want: false, validUUID: true},
		{name: "duplicate email", userName: "testDup", userEmail: "test@test.com", want: true, validUUID: false},
		{name: "no username", userName: "", userEmail: "noname@test.com", want: true, validUUID: false},
		{name: "no email", userName: "noemail", userEmail: "", want: true, validUUID: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func (t *testing.T){
			got, err := db.CreateUser(tt.userName, tt.userEmail)
			if err != nil {
				if tt.want != IsEmailUniqueErr(err) && tt.want == IsEmailUniqueErr(err) {
					t.Errorf("failed to create user %v, got %v\n", tt.userName, err)
				}
			}
			validUUID := uuid.Validate(got)
			if tt.validUUID && validUUID != nil {
				t.Errorf("expected valid UUID, got %q (err: %v)", got, validUUID)
			}
		})
	}
}
