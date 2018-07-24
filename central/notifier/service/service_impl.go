package service

import (
	"fmt"
	"sort"

	"bitbucket.org/stack-rox/apollo/central/notifier/processor"
	"bitbucket.org/stack-rox/apollo/central/notifier/store"
	"bitbucket.org/stack-rox/apollo/central/role/resources"
	"bitbucket.org/stack-rox/apollo/central/service"
	"bitbucket.org/stack-rox/apollo/generated/api/v1"
	"bitbucket.org/stack-rox/apollo/pkg/auth/permissions"
	"bitbucket.org/stack-rox/apollo/pkg/errorhelpers"
	"bitbucket.org/stack-rox/apollo/pkg/grpc/authz"
	"bitbucket.org/stack-rox/apollo/pkg/grpc/authz/perrpc"
	"bitbucket.org/stack-rox/apollo/pkg/grpc/authz/user"
	"bitbucket.org/stack-rox/apollo/pkg/notifications/notifiers"
	"bitbucket.org/stack-rox/apollo/pkg/secrets"
	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	authorizer = perrpc.FromMap(map[authz.Authorizer][]string{
		user.With(permissions.View(resources.Notifier)): {
			"/v1.NotifierService/GetNotifier",
			"/v1.NotifierService/GetNotifiers",
		},
		user.With(permissions.Modify(resources.Notifier)): {
			"/v1.NotifierService/PutNotifier",
			"/v1.NotifierService/PostNotifier",
			"/v1.NotifierService/TestNotifier",
			"/v1.NotifierService/DeleteNotifier",
		},
	})
)

// ClusterService is the struct that manages the cluster API
type serviceImpl struct {
	storage   store.Store
	processor processor.Processor
	detector  policyDetector
}

// RegisterServiceServer registers this service with the given gRPC Server.
func (s *serviceImpl) RegisterServiceServer(grpcServer *grpc.Server) {
	v1.RegisterNotifierServiceServer(grpcServer, s)
}

// RegisterServiceHandlerFromEndpoint registers this service with the given gRPC Gateway endpoint.
func (s *serviceImpl) RegisterServiceHandlerFromEndpoint(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	return v1.RegisterNotifierServiceHandlerFromEndpoint(ctx, mux, endpoint, opts)
}

// AuthFuncOverride specifies the auth criteria for this API.
func (s *serviceImpl) AuthFuncOverride(ctx context.Context, fullMethodName string) (context.Context, error) {
	return ctx, service.ReturnErrorCode(authorizer.Authorized(ctx, fullMethodName))
}

// GetNotifier retrieves all registries that matches the request filters
func (s *serviceImpl) GetNotifier(ctx context.Context, request *v1.ResourceByID) (*v1.Notifier, error) {
	if request.GetId() == "" {
		return nil, status.Error(codes.InvalidArgument, "Notifier id must be provided")
	}
	notifier, exists, err := s.storage.GetNotifier(request.GetId())
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	if !exists {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("Notifier %v not found", request.GetId()))
	}
	secrets.ScrubSecretsFromStruct(notifier)
	s.populatePolicies(notifier)
	return notifier, nil
}

// GetNotifiers retrieves all notifiers that match the request filters
func (s *serviceImpl) GetNotifiers(ctx context.Context, request *v1.GetNotifiersRequest) (*v1.GetNotifiersResponse, error) {
	notifiers, err := s.storage.GetNotifiers(request)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	for _, n := range notifiers {
		secrets.ScrubSecretsFromStruct(n)
		s.populatePolicies(n)
	}
	return &v1.GetNotifiersResponse{Notifiers: notifiers}, nil
}

func validateNotifier(notifier *v1.Notifier) error {
	errorList := errorhelpers.NewErrorList("Validation")
	if notifier.GetName() == "" {
		errorList.AddString("Notifier name must be defined")
	}
	if notifier.GetType() == "" {
		errorList.AddString("Notifier type must be defined")
	}
	if notifier.GetUiEndpoint() == "" {
		errorList.AddString("Notifier UI endpoint must be defined")
	}
	return errorList.ToError()
}

// PutNotifier updates a notifier in the system
func (s *serviceImpl) PutNotifier(ctx context.Context, request *v1.Notifier) (*empty.Empty, error) {
	if err := validateNotifier(request); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	notifierCreator, ok := notifiers.Registry[request.Type]
	if !ok {
		return nil, status.Error(codes.InvalidArgument, fmt.Sprintf("Notifier type %v is not a valid notifier type", request.Type))
	}
	notifier, err := notifierCreator(request)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	if err := s.storage.UpdateNotifier(request); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	s.processor.UpdateNotifier(notifier)
	return &empty.Empty{}, nil
}

// PostNotifier inserts a new registry into the system if it doesn't already exist
func (s *serviceImpl) PostNotifier(ctx context.Context, request *v1.Notifier) (*v1.Notifier, error) {
	if err := validateNotifier(request); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	if request.GetId() != "" {
		return nil, status.Error(codes.InvalidArgument, "Id field should be empty when posting a new notifier")
	}
	notifier, err := notifiers.CreateNotifier(request)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	id, err := s.storage.AddNotifier(request)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	request.Id = id
	s.processor.UpdateNotifier(notifier)
	return request, nil
}

// TestNotifier tests to see if the config is setup properly
func (s *serviceImpl) TestNotifier(ctx context.Context, request *v1.Notifier) (*empty.Empty, error) {
	if err := validateNotifier(request); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	notifier, err := notifiers.CreateNotifier(request)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	if err := notifier.Test(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	return &empty.Empty{}, nil
}

// DeleteNotifier deletes a notifier from the system
func (s *serviceImpl) DeleteNotifier(ctx context.Context, request *v1.DeleteNotifierRequest) (*empty.Empty, error) {
	if request.GetId() == "" {
		return nil, status.Error(codes.InvalidArgument, "Notifier id must be provided")
	}
	n, err := s.GetNotifier(ctx, &v1.ResourceByID{Id: request.GetId()})
	if err != nil {
		return nil, err
	}

	if !request.GetForce() && len(n.Policies) != 0 {
		m := jsonpb.Marshaler{}
		policiesOnly := &v1.Notifier{
			Policies: n.GetPolicies(),
		}
		jsonString, err := m.MarshalToString(policiesOnly)

		if err != nil {
			log.Error(err)
			return nil, status.Error(codes.FailedPrecondition, "Notifier is in use by policies")
		}

		return nil, status.Errorf(codes.FailedPrecondition, "Notifier is in use by policies: %s", jsonString)
	}

	if err := s.storage.RemoveNotifier(request.GetId()); err != nil {
		return nil, service.ReturnErrorCode(err)
	}
	s.processor.RemoveNotifier(request.GetId())
	s.detector.RemoveNotifier(request.GetId())
	return &empty.Empty{}, nil
}

func (s *serviceImpl) populatePolicies(notifier *v1.Notifier) {
	policies := s.processor.GetIntegratedPolicies(notifier.GetId())

	for _, p := range policies {
		notifier.Policies = append(notifier.Policies, &v1.Notifier_Policy{Id: p.GetId(), Name: p.GetName()})
	}

	sort.Slice(notifier.Policies, func(i, j int) bool {
		return notifier.Policies[i].Name < notifier.Policies[j].Name
	})
}
