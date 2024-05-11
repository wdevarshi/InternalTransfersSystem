package validator

import (
	"fmt"
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

func (v *validator) ValidateCreateAccountRequest(request *proto.CreateAccountRequest) error {
	if request == nil {
		return ErrInvalidRequest
	}
	fmt.Println("1")
	if request.GetAccountId() == "" {
		return ErrAccountMissing
	}
	fmt.Println("2")
	if containsSpecialCharacters(request.GetAccountId()) {
		return ErrSpecialCharacters
	}
	fmt.Println("3")
	if request.GetInitialBalance() < 0 {
		return ErrNegativeBalance
	}
	fmt.Println("4")
	return nil
}

func containsSpecialCharacters(id string) bool {
	if strings.ContainsAny(id, " !@#$%^&*()_+=[]{}|;':,.<>?") {
		return true
	}
	return false
}
