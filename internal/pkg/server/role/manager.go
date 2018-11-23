/*
 * Copyright (C)  2018 Nalej - All Rights Reserved
 */

package role

import (
	"github.com/nalej/derrors"
	"github.com/nalej/grpc-organization-go"
	"github.com/nalej/grpc-role-go"
	"github.com/nalej/system-model/internal/pkg/entities"
	"github.com/nalej/system-model/internal/pkg/provider/organization"
	"github.com/nalej/system-model/internal/pkg/provider/role"
)

// Manager structure with the required providers for role operations.
type Manager struct {
	OrgProvider organization.Provider
	RoleProvider role.Provider
}

// NewManager creates a Manager using a set of providers.
func NewManager(orgProvider organization.Provider, roleProvider role.Provider) Manager{
	return Manager{orgProvider, roleProvider}
}

// AddRole adds a new role to a given organization.
func (m * Manager) AddRole(addRoleRequest *grpc_role_go.AddRoleRequest) (*entities.Role, derrors.Error){
	exists := m.OrgProvider.Exists(addRoleRequest.OrganizationId)
	if !exists{
		return nil, derrors.NewNotFoundError("organizationID").WithParams(addRoleRequest.OrganizationId)
	}
	toAdd := entities.NewRoleFromGRPC(addRoleRequest)
	err := m.RoleProvider.Add(*toAdd)
	if err != nil {
		return nil, err
	}
	err = m.OrgProvider.AddRole(toAdd.OrganizationId, toAdd.RoleId)
	if err != nil {
		return nil, err
	}
	return toAdd, nil
}

// GetRole returns an existing role.
func (m * Manager) GetRole(roleID *grpc_role_go.RoleId) (*entities.Role, derrors.Error){
	if ! m.OrgProvider.Exists(roleID.OrganizationId){
		return nil, derrors.NewNotFoundError("organizationID").WithParams(roleID.OrganizationId)
	}

	if !m.OrgProvider.RoleExists(roleID.OrganizationId, roleID.RoleId){
		return nil, derrors.NewNotFoundError("roleID").WithParams(roleID.OrganizationId, roleID.RoleId)
	}
	return m.RoleProvider.Get(roleID.RoleId)
}

// ListRoles retrieves the list of roles of a given organization.
func (m * Manager) ListRoles(organizationID *grpc_organization_go.OrganizationId) ([]entities.Role, derrors.Error){
	if !m.OrgProvider.Exists(organizationID.OrganizationId){
		return nil, derrors.NewNotFoundError("organizationID").WithParams(organizationID.OrganizationId)
	}
	roles, err := m.OrgProvider.ListRoles(organizationID.OrganizationId)
	if err != nil {
		return nil, err
	}
	result := make([] entities.Role, 0)
	for _, rID := range roles {
		toAdd, err := m.RoleProvider.Get(rID)
		if err != nil {
			return nil, err
		}
		result = append(result, *toAdd)
	}
	return result, nil
}

// RemoveRole removes a given role from an organization.
func (m * Manager) RemoveRole(removeRoleRequest *grpc_role_go.RemoveRoleRequest) derrors.Error{
	if ! m.OrgProvider.Exists(removeRoleRequest.OrganizationId){
		return derrors.NewNotFoundError("organizationID").WithParams(removeRoleRequest.OrganizationId)
	}

	if !m.OrgProvider.RoleExists(removeRoleRequest.OrganizationId, removeRoleRequest.RoleId){
		return derrors.NewNotFoundError("roleID").WithParams(removeRoleRequest.OrganizationId, removeRoleRequest.RoleId)
	}

	err := m.OrgProvider.DeleteRole(removeRoleRequest.OrganizationId, removeRoleRequest.RoleId)
	if err != nil {
		return err
	}
	return m.RoleProvider.Remove(removeRoleRequest.RoleId)
}