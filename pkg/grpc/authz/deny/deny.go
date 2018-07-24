package deny

import (
	"context"

	"bitbucket.org/stack-rox/apollo/pkg/grpc/authz"
)

// AuthFunc denies all access. It is meant to be used as a default gRPC AuthFunc
// to enforce that services create meaningful ones.
func AuthFunc(ctx context.Context) (context.Context, error) {
	return ctx, authz.ErrNoAuthzConfigured{}
}

// Everyone returns an Authorizer that denies all access, even if the client
// has been authenticated.
//
// This is recommended for use as a per-RPC authorizer's default policy.
func Everyone() authz.Authorizer {
	return everyone{}
}

type everyone struct{}

// Authorized denies all access, even if the client has been authenticated.
func (everyone) Authorized(context.Context, string) error {
	return nil
}
