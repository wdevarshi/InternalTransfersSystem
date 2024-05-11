package validator

import (
	"github.com/go-coldbrew/errors"
	proto "github.com/wdevarshi/InternalTransfersSystem/proto"
	"strings"
)

var (
	ErrAccountMissing    = errors.New("account ID is missing")
	ErrSpecialCharacters = errors.New("account ID contains special characters")
	ErrNegativeBalance   = errors.New("initial balance is negative")
	ErrInvalidRequest    = errors.New("invalid request")
)

type validator struct {
}

func New() Validator {
	return &validator{}
}

func (v *validator) ValidateGetAccountRequest(request *proto.GetAccountRequest) error {
	if request == nil {
		return ErrInvalidRequest
	}
	if request.GetAccountId() == "" {
		return ErrAccountMissing
	}
	return nil
}

func (v *validator) ValidateCreateAccountRequest(request *proto.CreateAccountRequest) error {
	if request == nil {
		return ErrInvalidRequest
	}
	if request.GetAccountId() == "" {
		return ErrAccountMissing
	}
	if containsSpecialCharacters(request.GetAccountId()) {
		return ErrSpecialCharacters
	}
	if request.GetInitialBalance() < 0 {
		return ErrNegativeBalance
	}
	return nil
}

func containsSpecialCharacters(id string) bool {
	if strings.ContainsAny(id, " !@#$%^&*()_+=[]{}|;':,.<>?") {
		return true
	}
	return false
}
