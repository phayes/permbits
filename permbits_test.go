package permbits

import (
	"testing"
)

func TestPermissions(t *testing.T) {
	var allTrue PermissionBits

	allTrue.SetSetuid(true)
	allTrue.SetSetgid(true)
	allTrue.SetSticky(true)
	allTrue.SetUserRead(true)
	allTrue.SetUserWrite(true)
	allTrue.SetUserExecute(true)
	allTrue.SetGroupRead(true)
	allTrue.SetGroupWrite(true)
	allTrue.SetGroupExecute(true)
	allTrue.SetOtherRead(true)
	allTrue.SetOtherWrite(true)
	allTrue.SetOtherExecute(true)

	if !allTrue.Setuid() {
		t.Error("allTrue: UserRead returns false")
	}
	if !allTrue.Setgid() {
		t.Error("allTrue: UserRead returns false")
	}
	if !allTrue.Sticky() {
		t.Error("allTrue: UserRead returns false")
	}
	if !allTrue.UserRead() {
		t.Error("allTrue: UserRead returns false")
	}
	if !allTrue.UserWrite() {
		t.Error("allTrue: UserWrite returns false")
	}
	if !allTrue.UserExecute() {
		t.Error("allTrue: UserExecute returns false")
	}
	if !allTrue.GroupRead() {
		t.Error("allTrue: GroupRead returns false")
	}
	if !allTrue.GroupWrite() {
		t.Error("allTrue: GroupWrite returns false")
	}
	if !allTrue.GroupExecute() {
		t.Error("allTrue: GroupExecute returns false")
	}
	if !allTrue.OtherRead() {
		t.Error("allTrue: OtherRead returns false")
	}
	if !allTrue.OtherWrite() {
		t.Error("allTrue: OtherWrite returns false")
	}
	if !allTrue.OtherExecute() {
		t.Error("allTrue: OtherExecute returns false")
	}

	if allTrue.String() != "rwxrwxrwx" {
		t.Error("allTrue: string incorrect")
	}

	allFalse := allTrue
	allFalse.SetSetuid(false)
	allFalse.SetSetgid(false)
	allFalse.SetSticky(false)
	allFalse.SetUserRead(false)
	allFalse.SetUserWrite(false)
	allFalse.SetUserExecute(false)
	allFalse.SetGroupRead(false)
	allFalse.SetGroupWrite(false)
	allFalse.SetGroupExecute(false)
	allFalse.SetOtherRead(false)
	allFalse.SetOtherWrite(false)
	allFalse.SetOtherExecute(false)

	if allFalse.Setuid() {
		t.Error("allFalse: UserRead returns true")
	}
	if allFalse.Setgid() {
		t.Error("allFalse: UserRead returns true")
	}
	if allFalse.Sticky() {
		t.Error("allFalse: UserRead returns true")
	}
	if allFalse.UserRead() {
		t.Error("allFalse: UserRead returns true")
	}
	if allFalse.UserWrite() {
		t.Error("allFalse: UserWrite returns true")
	}
	if allFalse.UserExecute() {
		t.Error("allFalse: UserExecute returns true")
	}
	if allFalse.GroupRead() {
		t.Error("allFalse: GroupRead returns true")
	}
	if allFalse.GroupWrite() {
		t.Error("allFalse: GroupWrite returns true")
	}
	if allFalse.GroupExecute() {
		t.Error("allFalse: GroupExecute returns true")
	}
	if allFalse.OtherRead() {
		t.Error("allFalse: OtherRead returns true")
	}
	if allFalse.OtherWrite() {
		t.Error("allFalse: OtherWrite returns true")
	}
	if allFalse.OtherExecute() {
		t.Error("allFalse: OtherExecute returns true")
	}
	if allFalse.String() != "---------" {
		t.Error("allFalse: string incorrect")
	}

}
