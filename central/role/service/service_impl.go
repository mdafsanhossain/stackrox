package service

import (
	"context"
	"fmt"
	"sort"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/pkg/errors"
	clusterDS "github.com/stackrox/rox/central/cluster/datastore"
	clusterMappings "github.com/stackrox/rox/central/cluster/index/mappings"
	namespaceDS "github.com/stackrox/rox/central/namespace/datastore"
	namespaceMappings "github.com/stackrox/rox/central/namespace/index/mappings"
	rolePkg "github.com/stackrox/rox/central/role"
	"github.com/stackrox/rox/central/role/datastore"
	"github.com/stackrox/rox/central/role/resources"
	v1 "github.com/stackrox/rox/generated/api/v1"
	"github.com/stackrox/rox/generated/storage"
	"github.com/stackrox/rox/pkg/auth/permissions"
	"github.com/stackrox/rox/pkg/env"
	"github.com/stackrox/rox/pkg/errox"
	"github.com/stackrox/rox/pkg/grpc/authn"
	"github.com/stackrox/rox/pkg/grpc/authz"
	"github.com/stackrox/rox/pkg/grpc/authz/allow"
	"github.com/stackrox/rox/pkg/grpc/authz/perrpc"
	"github.com/stackrox/rox/pkg/grpc/authz/user"
	"github.com/stackrox/rox/pkg/logging"
	"github.com/stackrox/rox/pkg/postgres/schema"
	"github.com/stackrox/rox/pkg/sac"
	"github.com/stackrox/rox/pkg/sac/effectiveaccessscope"
	"github.com/stackrox/rox/pkg/search"
	"google.golang.org/grpc"
)

var (
	authorizer = perrpc.FromMap(map[authz.Authorizer][]string{
		user.With(permissions.View(resources.Role)): {
			"/v1.RoleService/GetRoles",
			"/v1.RoleService/GetRole",
			"/v1.RoleService/ListPermissionSets",
			"/v1.RoleService/GetPermissionSet",
			"/v1.RoleService/ListSimpleAccessScopes",
			"/v1.RoleService/GetSimpleAccessScope",
		},
		user.With(permissions.View(resources.Role), permissions.View(resources.Cluster), permissions.View(resources.Namespace)): {
			"/v1.RoleService/ComputeEffectiveAccessScope",
		},
		user.With(permissions.Modify(resources.Role)): {
			"/v1.RoleService/CreateRole",
			"/v1.RoleService/SetDefaultRole",
			"/v1.RoleService/UpdateRole",
			"/v1.RoleService/DeleteRole",
			"/v1.RoleService/PostPermissionSet",
			"/v1.RoleService/PutPermissionSet",
			"/v1.RoleService/DeletePermissionSet",
			"/v1.RoleService/PostSimpleAccessScope",
			"/v1.RoleService/PutSimpleAccessScope",
			"/v1.RoleService/DeleteSimpleAccessScope",
		},
		allow.Anonymous(): {
			"/v1.RoleService/GetResources",
			"/v1.RoleService/GetMyPermissions",
			"/v1.RoleService/GetClustersForPermission",
			"/v1.RoleService/GetNamespacesForClusterAndPermission",
		},
	})
)

var (
	log = logging.LoggerForModule()
)

type serviceImpl struct {
	v1.UnimplementedRoleServiceServer

	roleDataStore      datastore.DataStore
	clusterDataStore   clusterDS.DataStore
	namespaceDataStore namespaceDS.DataStore
}

func (s *serviceImpl) RegisterServiceServer(grpcServer *grpc.Server) {
	v1.RegisterRoleServiceServer(grpcServer, s)
}

func (s *serviceImpl) RegisterServiceHandler(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error {
	return v1.RegisterRoleServiceHandler(ctx, mux, conn)
}

func (*serviceImpl) AuthFuncOverride(ctx context.Context, fullMethodName string) (context.Context, error) {
	return ctx, authorizer.Authorized(ctx, fullMethodName)
}

func (s *serviceImpl) GetRoles(ctx context.Context, _ *v1.Empty) (*v1.GetRolesResponse, error) {
	roles, err := s.roleDataStore.GetAllRoles(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "failed to retrieve roles")
	}

	// List roles in the same order for consistency across requests.
	sort.Slice(roles, func(i, j int) bool {
		return roles[i].GetName() < roles[j].GetName()
	})

	return &v1.GetRolesResponse{Roles: roles}, nil
}

func (s *serviceImpl) GetRole(ctx context.Context, id *v1.ResourceByID) (*storage.Role, error) {
	role, found, err := s.roleDataStore.GetRole(ctx, id.GetId())
	if err != nil {
		return nil, errors.Wrapf(err, "failed to retrieve role %q", id.GetId())
	}
	if !found {
		return nil, errors.Wrapf(errox.NotFound, "failed to retrieve role %q", id.GetId())
	}
	return role, nil
}

func (s *serviceImpl) GetMyPermissions(ctx context.Context, _ *v1.Empty) (*v1.GetPermissionsResponse, error) {
	return GetMyPermissions(ctx)
}

func (s *serviceImpl) CreateRole(ctx context.Context, roleRequest *v1.CreateRoleRequest) (*v1.Empty, error) {
	role := roleRequest.GetRole()

	// Check role request correctness.
	if role.GetName() != "" && role.GetName() != roleRequest.GetName() {
		return nil, errox.InvalidArgs.CausedBy("different role names in path and body")
	}
	role.Name = roleRequest.GetName()

	// Empty access scope ID is deprecated. Fill the default during the adoption
	// period.
	// TODO(ROX-9510): remove this block.
	if role.GetAccessScopeId() == "" {
		role.AccessScopeId = rolePkg.AccessScopeIncludeAll.GetId()
	}
	err := s.roleDataStore.AddRole(ctx, role)
	if err != nil {
		return nil, err
	}
	return &v1.Empty{}, nil
}

func (s *serviceImpl) UpdateRole(ctx context.Context, role *storage.Role) (*v1.Empty, error) {
	// Empty access scope ID is deprecated. Fill the default during the adoption
	// period.
	// TODO(ROX-9510): remove this block.
	if role.GetAccessScopeId() == "" {
		role.AccessScopeId = rolePkg.AccessScopeIncludeAll.GetId()
	}
	err := s.roleDataStore.UpdateRole(ctx, role)
	if err != nil {
		return nil, err
	}
	return &v1.Empty{}, nil
}

func (s *serviceImpl) DeleteRole(ctx context.Context, id *v1.ResourceByID) (*v1.Empty, error) {
	err := s.roleDataStore.RemoveRole(ctx, id.GetId())
	if err != nil {
		return nil, errors.Wrapf(err, "failed to delete role %q", id.GetId())
	}
	return &v1.Empty{}, nil
}

// GetResources returns all the possible resources in the system
func (s *serviceImpl) GetResources(context.Context, *v1.Empty) (*v1.GetResourcesResponse, error) {
	resourceList := resources.ListAll()
	resources := make([]string, 0, len(resourceList))
	for _, r := range resourceList {
		resources = append(resources, string(r))
	}
	return &v1.GetResourcesResponse{
		Resources: resources,
	}, nil
}

// GetMyPermissions returns the permissions for a user based on the context.
func GetMyPermissions(ctx context.Context) (*v1.GetPermissionsResponse, error) {
	// Get the perms from the current user context.
	id, err := authn.IdentityFromContext(ctx)
	if err != nil {
		return nil, err
	}
	return &v1.GetPermissionsResponse{
		ResourceToAccess: id.Permissions(),
	}, nil
}

////////////////////////////////////////////////////////////////////////////////
// Permission sets                                                            //
//                                                                            //

func (s *serviceImpl) GetPermissionSet(ctx context.Context, id *v1.ResourceByID) (*storage.PermissionSet, error) {
	permissionSet, found, err := s.roleDataStore.GetPermissionSet(ctx, id.GetId())
	if err != nil {
		return nil, errors.Wrapf(err, "failed to retrieve permission set %s", id.GetId())
	}
	if !found {
		return nil, errors.Wrapf(errox.NotFound, "failed to retrieve permission set %s", id.GetId())
	}

	return permissionSet, nil
}

func (s *serviceImpl) ListPermissionSets(ctx context.Context, _ *v1.Empty) (*v1.ListPermissionSetsResponse, error) {
	permissionSets, err := s.roleDataStore.GetAllPermissionSets(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "failed to retrieve permission sets")
	}

	// List permission sets in the same order for consistency across requests.
	sort.Slice(permissionSets, func(i, j int) bool {
		return permissionSets[i].GetName() < permissionSets[j].GetName()
	})

	return &v1.ListPermissionSetsResponse{PermissionSets: permissionSets}, nil
}

func (s *serviceImpl) PostPermissionSet(ctx context.Context, permissionSet *storage.PermissionSet) (*storage.PermissionSet, error) {
	if permissionSet.GetId() != "" {
		return nil, errox.InvalidArgs.CausedBy("setting id field is not allowed")
	}
	permissionSet.Id = rolePkg.GeneratePermissionSetID()

	// Store the augmented permission set; report back on error. Note the
	// permission set is referenced by its name because that's what the caller
	// knows.
	err := s.roleDataStore.AddPermissionSet(ctx, permissionSet)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to store permission set %q", permissionSet.GetName())
	}

	// Assume AddPermissionSet() does not make modifications to the protobuf.
	return permissionSet, nil
}

func (s *serviceImpl) PutPermissionSet(ctx context.Context, permissionSet *storage.PermissionSet) (*v1.Empty, error) {
	err := s.roleDataStore.UpdatePermissionSet(ctx, permissionSet)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to update permission set %s", permissionSet.GetId())
	}

	return &v1.Empty{}, nil
}

func (s *serviceImpl) DeletePermissionSet(ctx context.Context, id *v1.ResourceByID) (*v1.Empty, error) {
	err := s.roleDataStore.RemovePermissionSet(ctx, id.GetId())
	if err != nil {
		return nil, errors.Wrapf(err, "failed to delete permission set %s", id.GetId())
	}

	return &v1.Empty{}, nil
}

////////////////////////////////////////////////////////////////////////////////
// Access scopes                                                              //
//                                                                            //

func (s *serviceImpl) GetSimpleAccessScope(ctx context.Context, id *v1.ResourceByID) (*storage.SimpleAccessScope, error) {
	scope, found, err := s.roleDataStore.GetAccessScope(ctx, id.GetId())
	if err != nil {
		return nil, errors.Wrapf(err, "failed to retrieve access scope %s", id.GetId())
	}
	if !found {
		return nil, errors.Wrapf(errox.NotFound, "failed to retrieve access scope %s", id.GetId())
	}

	return scope, nil
}

func (s *serviceImpl) ListSimpleAccessScopes(ctx context.Context, _ *v1.Empty) (*v1.ListSimpleAccessScopesResponse, error) {
	scopes, err := s.roleDataStore.GetAllAccessScopes(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "failed to retrieve access scopes")
	}

	// List access scopes in the same order for consistency across requests.
	sort.Slice(scopes, func(i, j int) bool {
		return scopes[i].GetName() < scopes[j].GetName()
	})

	return &v1.ListSimpleAccessScopesResponse{AccessScopes: scopes}, nil
}

func (s *serviceImpl) PostSimpleAccessScope(ctx context.Context, scope *storage.SimpleAccessScope) (*storage.SimpleAccessScope, error) {
	if scope.GetId() != "" {
		return nil, errox.InvalidArgs.CausedBy("setting id field is not allowed")
	}
	scope.Id = rolePkg.GenerateAccessScopeID()

	// Store the augmented access scope; report back on error. Note the access
	// scope is referenced by its name because that's what the caller knows.
	err := s.roleDataStore.AddAccessScope(ctx, scope)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to store access scope %q", scope.GetName())
	}

	// Assume AddAccessScope() does not make modifications to the protobuf.
	return scope, nil
}

func (s *serviceImpl) PutSimpleAccessScope(ctx context.Context, scope *storage.SimpleAccessScope) (*v1.Empty, error) {
	err := s.roleDataStore.UpdateAccessScope(ctx, scope)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to update access scope %s", scope.GetId())
	}

	return &v1.Empty{}, nil
}

func (s *serviceImpl) DeleteSimpleAccessScope(ctx context.Context, id *v1.ResourceByID) (*v1.Empty, error) {
	err := s.roleDataStore.RemoveAccessScope(ctx, id.GetId())
	if err != nil {
		return nil, errors.Wrapf(err, "failed to delete access scope %s", id.GetId())
	}

	return &v1.Empty{}, nil
}

func (s *serviceImpl) ComputeEffectiveAccessScope(ctx context.Context, req *v1.ComputeEffectiveAccessScopeRequest) (*storage.EffectiveAccessScope, error) {
	// If we're here, service-level authz has already verified that the caller
	// has at least READ permission on the Role resource.
	err := rolePkg.ValidateSimpleAccessScopeRules(req.GetAccessScope().GetSimpleRules())
	if err != nil {
		return nil, errors.Wrap(err, "failed to compute effective access scope")
	}

	// ctx might not have access to all known clusters and namespaces and hence
	// the resulting effective access scope might not include all known scopes,
	//
	// Imagine Alice has write access to Role and read access to scoped Cluster
	// resources. She can create access scopes that will apply to all clusters
	// but while she is creating them she would only see a sliced view.
	readScopesCtx := ctx

	clusters, err := s.clusterDataStore.GetClusters(readScopesCtx)
	if err != nil {
		return nil, errors.Errorf("failed to compute effective access scope: %v", err)
	}

	namespaces, err := s.namespaceDataStore.GetAllNamespaces(readScopesCtx)
	if err != nil {
		return nil, errors.Errorf("failed to compute effective access scope: %v", err)
	}

	response, err := effectiveAccessScopeForSimpleAccessScope(req.GetAccessScope().GetSimpleRules(), clusters, namespaces, req.GetDetail())
	if err != nil {
		return nil, errors.Errorf("failed to compute effective access scope: %v", err)
	}

	return response, nil
}

func (s *serviceImpl) GetClustersForPermission(ctx context.Context, req *v1.GetClustersForPermissionRequest) (*v1.GetClustersForPermissionResponse, error) {
	if req == nil {
		return nil, errox.InvalidArgs
	}
	response := &v1.GetClustersForPermissionResponse{}
	targetResource := permissions.Resource(req.GetResource())
	targetResourceMetadata, found := resources.MetadataForResource(targetResource)
	if !found {
		return response, nil
	}
	targetAccess := req.GetAccess()
	targetResourceWithAccess := permissions.ResourceWithAccess{
		Resource: targetResourceMetadata,
		Access:   targetAccess,
	}
	scopeChecker := sac.ForResource(targetResourceMetadata).ScopeChecker(ctx, targetAccess)
	scope, err := scopeChecker.EffectiveAccessScope(targetResourceWithAccess)
	if err != nil {
		return nil, err
	}
	if scope != nil {
		if scope.State == effectiveaccessscope.Included {
			elevatedCtx := sac.WithGlobalAccessScopeChecker(ctx,
				sac.AllowFixedScopes(
					sac.AccessModeScopeKeys(storage.Access_READ_ACCESS),
					sac.ResourceScopeKeys(resources.Cluster),
				),
			)
			query := search.NewQueryBuilder().AddStringsHighlighted(search.Cluster, search.WildcardString).ProtoQuery()
			results, err := s.clusterDataStore.Search(elevatedCtx, query)
			if err != nil {
				return nil, err
			}
			var clusterOptionsMap search.OptionsMap
			if env.PostgresDatastoreEnabled.BooleanSetting() {
				clusterOptionsMap = schema.ClustersSchema.OptionsMap
			} else {
				clusterOptionsMap = clusterMappings.OptionsMap
			}
			targetField, fieldFound := clusterOptionsMap.Get(search.Cluster.String())
			for _, r := range results {
				clusterID := r.ID
				clusterName := ""
				if fieldFound {
					for _, v := range r.Matches[targetField.GetFieldPath()] {
						if len(v) > 0 {
							clusterName = v
							break
						}
					}
				} else {
					clusterName = fmt.Sprintf("Cluster with ID %q", clusterID)
				}
				clusterInfo := &v1.ScopeElementForPermission{
					Id:   clusterID,
					Name: clusterName,
				}
				response.Clusters = append(response.Clusters, clusterInfo)
			}
		} else if scope.State != effectiveaccessscope.Excluded {
			for _, clusterID := range scope.GetClusterIDs() {
				clusterName := scope.GetClusterNameForID(clusterID)
				clusterData := scope.GetClusterByID(clusterID)
				if clusterData == nil {
					continue
				}
				if clusterData.State == effectiveaccessscope.Excluded {
					continue
				}
				hasIncludedNamespace := false
				for _, nsData := range clusterData.Namespaces {
					if nsData == nil {
						continue
					}
					if nsData.State != effectiveaccessscope.Excluded {
						hasIncludedNamespace = true
						break
					}
				}
				if hasIncludedNamespace {
					clusterInfo := &v1.ScopeElementForPermission{
						Id:   clusterID,
						Name: clusterName,
					}
					response.Clusters = append(response.Clusters, clusterInfo)
				}
			}
		}
	}
	return response, nil
}

func (s *serviceImpl) GetNamespacesForClusterAndPermission(ctx context.Context, req *v1.GetNamespacesForPermissionAndClusterRequest) (*v1.GetNamespacesForPermissionAndClusterResponse, error) {
	if req == nil {
		return nil, errox.InvalidArgs
	}
	log.Info(req)
	response := &v1.GetNamespacesForPermissionAndClusterResponse{}
	targetResource := permissions.Resource(req.GetResource())
	targetResourceMetadata, found := resources.MetadataForResource(targetResource)
	if !found {
		return response, nil
	}
	targetAccess := req.GetAccess()
	targetResourceWithAccess := permissions.ResourceWithAccess{
		Resource: targetResourceMetadata,
		Access:   targetAccess,
	}
	scopeChecker := sac.ForResource(targetResourceMetadata).ScopeChecker(ctx, targetAccess)
	scope, err := scopeChecker.EffectiveAccessScope(targetResourceWithAccess)
	if err != nil {
		return nil, err
	}
	log.Info(scope)
	if scope != nil {
		isClusterFullyIncluded := false
		if scope.State == effectiveaccessscope.Included {
			isClusterFullyIncluded = true
		} else {
			clusterNode := scope.GetClusterByID(req.GetClusterId())
			if clusterNode == nil || clusterNode.State == effectiveaccessscope.Excluded {
				return response, nil
			}
			if clusterNode.State == effectiveaccessscope.Included || targetResourceMetadata.GetScope() == permissions.ClusterScope {
				isClusterFullyIncluded = true
			}
		}
		log.Info(isClusterFullyIncluded)
		var elevatedCtx context.Context
		if isClusterFullyIncluded {
			elevatedCtx = sac.WithGlobalAccessScopeChecker(ctx,
				sac.AllowFixedScopes(
					sac.AccessModeScopeKeys(storage.Access_READ_ACCESS),
					sac.ResourceScopeKeys(resources.Namespace),
					sac.ClusterScopeKeys(req.GetClusterId()),
				),
			)
		} else {
			clusterNode := scope.GetClusterByID(req.GetClusterId())
			namespaceList := make([]string, 0, len(clusterNode.Namespaces))
			for name, node := range clusterNode.Namespaces {
				if node == nil || node.State != effectiveaccessscope.Included {
					continue
				}
				namespaceList = append(namespaceList, name)
			}
			elevatedCtx = sac.WithGlobalAccessScopeChecker(ctx,
				sac.AllowFixedScopes(
					sac.AccessModeScopeKeys(storage.Access_READ_ACCESS),
					sac.ResourceScopeKeys(resources.Namespace),
					sac.ClusterScopeKeys(req.GetClusterId()),
					sac.NamespaceScopeKeys(namespaceList...),
				),
			)
		}

		query := search.NewQueryBuilder().
			AddStringsHighlighted(search.Namespace, search.WildcardString).
			ProtoQuery()
		log.Info(query)
		results, err := s.namespaceDataStore.Search(elevatedCtx, query)
		if err != nil {
			log.Info("Search error")
			return nil, err
		}
		log.Info(len(results))
		log.Info(results)
		var namespaceOptionsMap search.OptionsMap
		if env.PostgresDatastoreEnabled.BooleanSetting() {
			namespaceOptionsMap = schema.NamespacesSchema.OptionsMap
		} else {
			namespaceOptionsMap = namespaceMappings.OptionsMap
		}
		targetField, fieldFound := namespaceOptionsMap.Get(search.Namespace.String())
		for _, r := range results {
			namespaceID := r.ID
			namespaceName := ""
			if fieldFound {
				for _, v := range r.Matches[targetField.GetFieldPath()] {
					if len(v) > 0 {
						namespaceName = v
						break
					}
				}
			} else {
				namespaceName = fmt.Sprintf("Namespace with ID %q", namespaceID)
			}
			clusterInfo := &v1.ScopeElementForPermission{
				Id:   namespaceID,
				Name: namespaceName,
			}
			response.Namespaces = append(response.Namespaces, clusterInfo)
		}
	}
	return response, nil
}

////////////////////////////////////////////////////////////////////////////////
// Helpers                                                                    //
//                                                                            //

// effectiveAccessScopeForSimpleAccessScope computes the effective access scope
// for the given rules and converts it to the desired response.
func effectiveAccessScopeForSimpleAccessScope(scopeRules *storage.SimpleAccessScope_Rules, clusters []*storage.Cluster, namespaces []*storage.NamespaceMetadata, detail v1.ComputeEffectiveAccessScopeRequest_Detail) (*storage.EffectiveAccessScope, error) {
	tree, err := effectiveaccessscope.ComputeEffectiveAccessScope(scopeRules, clusters, namespaces, detail)
	if err != nil {
		return nil, err
	}

	response, err := effectiveaccessscope.ToEffectiveAccessScope(tree)
	if err != nil {
		return nil, err
	}

	return response, nil
}
