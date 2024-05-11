package validator

import proto "github.com/wdevarshi/InternalTransfersSystem/proto"

type Validator interface {
	ValidateCreateAccountRequest(*proto.CreateAccountRequest) error
	ValidateGetAccountRequest(*proto.GetAccountRequest) error
}
