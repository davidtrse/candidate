package model

import "testing"

func TestUpdateName(t *testing.T) {
	u := User{ID: 1, Name: "Giap"}
	u.UpdateName("John")

	if u.Name != "John" {
		t.Fatalf("expected Name to be %q after UpdateName, got %q", "John", u.Name)
	}
}
