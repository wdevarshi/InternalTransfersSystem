package service

import (
	"context"
	"fmt"
	"github.com/go-coldbrew/errors"
	"github.com/go-coldbrew/log"
	"github.com/wdevarshi/InternalTransfersSystem/config"
	"github.com/wdevarshi/InternalTransfersSystem/database"
	proto "github.com/wdevarshi/InternalTransfersSystem/proto"
	"github.com/wdevarshi/InternalTransfersSystem/service/validator"
	"google.golang.org/genproto/googleapis/api/httpbody"
	"google.golang.org/grpc/health"
	"google.golang.org/protobuf/types/known/emptypb"
	"time"
)

// svc should implement the service interface defined in the proto file
var _ proto.InternalTransfersSystemServer = (*svc)(nil)

// Service interface for the service
type svc struct {
	// health server for the service
	*health.Server
	// TODO: remove this, since this is just to demonstrate how to use config
	// prefix to be added to the message in the response
	prefix    string
	store     database.InternalTransferSystemStore
	validator validator.Validator
}

// ReadinessProbe for the service
// This is called by the kubernetes readiness probe
func (s *svc) ReadyCheck(ctx context.Context, _ *emptypb.Empty) (*httpbody.HttpBody, error) {
	return GetReadyState(ctx)
}

// LivenessProbe for the service
// This is called by the kubernetes liveness probe
func (s *svc) HealthCheck(ctx context.Context, _ *emptypb.Empty) (*httpbody.HttpBody, error) {
	return GetHealthCheck(ctx), nil
}

// Echo returns the message with the prefix added
// TODO: remove this, since this is just to demonstrate how to use endpoints and config
func (s *svc) Echo(_ context.Context, req *proto.EchoRequest) (*proto.EchoResponse, error) {
	return &proto.EchoResponse{
		Msg: fmt.Sprintf("%s: %s", s.prefix, req.GetMsg()),
	}, nil
}

func (s *svc) CreateAccount(ctx context.Context, req *proto.CreateAccountRequest) (*proto.CreateAccountResponse, error) {
	err := s.validator.ValidateCreateAccountRequest(req)
	if err != nil {
		return nil, err
	}
	err = s.store.CreateAccount(ctx, &database.Account{
		ID:           req.GetAccountId(),
		Balance:      req.GetInitialBalance(),
		TimeCreated:  time.Now(),
		LastModified: time.Now(),
	})
	if err != nil {
		return nil, err
	}
	return &proto.CreateAccountResponse{}, nil
}

func (s *svc) GetAccount(ctx context.Context, req *proto.GetAccountRequest) (*proto.GetAccountResponse, error) {
	err := s.validator.ValidateGetAccountRequest(req)
	if err != nil {
		return nil, err
	}
	account, err := s.store.GetAccount(ctx, req.GetAccountId())
	if err != nil {
		return nil, err
	}
	return &proto.GetAccountResponse{
		AccountId: account.ID,
		Balance:   account.Balance,
	}, nil
}

// Error returns an error to the client
// TODO: remove this, since this is just to demonstrate how to use endpoints and config
func (s *svc) Error(ctx context.Context, req *proto.EchoRequest) (*proto.EchoResponse, error) {
	err := errors.New("This is an Error")
	log.Info(ctx, "error requested")
	return nil, errors.Wrap(err, "endpoint error")
}

// Creates a new Service instance and returns it
func New(cfg config.Config, store database.InternalTransferSystemStore, validator validator.Validator) (*svc, error) {
	// TODO: Application should validate the config here and return an error if it is invalid or missing
	s := &svc{
		// This is the health server for the service that is used for grpc
		Server: GetHealthCheckServer(),
		// TODO: remove this, since this is just to demonstrate how to use config
		prefix:    cfg.Prefix,
		store:     store,
		validator: validator,
	}
	// TODO: Application should initialize the service here and return an error if it fails to initialize

	// we call SetReady() here to indicate that the service is ready to serve traffic
	// service will not receive any traffic until this is called
	SetReady()
	return s, nil
}
