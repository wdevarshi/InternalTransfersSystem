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
	ErrSameAccountId     = errors.New("source and destination accounts are same")
)

type validator struct {
}

func New() Validator {
	return &validator{}
}

func (v *validator) ValidateTransactionSubmissionRequest(request *proto.TransactionSubmissionRequest) error {
	if request == nil {
		return ErrInvalidRequest
	}
	if request.GetFromAccountId() == "" {
		return ErrAccountMissing
	}
	if request.GetToAccountId() == "" {
		return ErrAccountMissing
	}
	if request.GetAmount() < 0 {
		return ErrNegativeBalance
	}
	if request.GetFromAccountId() == request.GetToAccountId() {
		return ErrSameAccountId
	}
	return nil
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
