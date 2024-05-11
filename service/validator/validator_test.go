package validator

import (
	proto "github.com/wdevarshi/InternalTransfersSystem/proto"
	"testing"
)

func TestContainsSpecialCharacters(t *testing.T) {
	if !containsSpecialCharacters("abc!def") {
		t.Errorf("Expected special characters to be found, but not found")
	}

	if !containsSpecialCharacters("abc def") {
		t.Errorf("Expected special characters to be found, but not found")
	}

	if containsSpecialCharacters("abcdefgh") {
		t.Errorf("Expected no special characters to be found, but found")
	}
}

func TestValidateCreateAccountRequest(t *testing.T) {
	v := &validator{}

	// Test case: nil request
	err := v.ValidateCreateAccountRequest(nil)
	if err != ErrInvalidRequest {
		t.Errorf("Expected error: %v, got: %v", ErrInvalidRequest, err)
	}

	// Test case: missing account ID
	err = v.ValidateCreateAccountRequest(&proto.CreateAccountRequest{})
	if err != ErrAccountMissing {
		t.Errorf("Expected error: %v, got: %v", ErrAccountMissing, err)
	}

	// Test case: account ID with special characters
	err = v.ValidateCreateAccountRequest(&proto.CreateAccountRequest{
		AccountId:      "!@#$%^&*()",
		InitialBalance: 100,
	})
	if err != ErrSpecialCharacters {
		t.Errorf("Expected error: %v, got: %v", ErrSpecialCharacters, err)
	}

	// Test case: negative initial balance
	err = v.ValidateCreateAccountRequest(&proto.CreateAccountRequest{
		AccountId:      "test",
		InitialBalance: -100,
	})
	if err != ErrNegativeBalance {
		t.Errorf("Expected error: %v, got: %v", ErrNegativeBalance, err)
	}

	// Test case: valid request
	err = v.ValidateCreateAccountRequest(&proto.CreateAccountRequest{
		AccountId:      "test",
		InitialBalance: 100,
	})
	if err != nil {
		t.Errorf("Expected no error, got: %v", err)
	}
}

func TestValidateGetAccountRequest(t *testing.T) {
	v := &validator{}

	// Test case: nil request
	err := v.ValidateGetAccountRequest(nil)
	if err != ErrInvalidRequest {
		t.Errorf("Expected error: %v, got: %v", ErrInvalidRequest, err)
	}

	// Test case: missing account ID
	err = v.ValidateGetAccountRequest(&proto.GetAccountRequest{})
	if err != ErrAccountMissing {
		t.Errorf("Expected error: %v, got: %v", ErrAccountMissing, err)
	}

	// Test case: valid request
	err = v.ValidateGetAccountRequest(&proto.GetAccountRequest{AccountId: "test"})
	if err != nil {
		t.Errorf("Expected no error, got: %v", err)
	}
}
