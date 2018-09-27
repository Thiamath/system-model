/*
 * Copyright (C)  2018 Nalej - All Rights Reserved
 */

package organization

import (
	"context"
	"github.com/nalej/derrors"
	"github.com/nalej/grpc-common-go"
	"github.com/nalej/grpc-organization-go"
	"github.com/nalej/grpc-utils/pkg/conversions"
)

type Handler struct{
	Manager Manager
}

// NewHandler creates a new Handler.
func NewHandler(manager Manager) *Handler {
	return &Handler{manager}
}

func (h * Handler) validAddOrganization(toAdd * grpc_organization_go.AddOrganizationRequest) derrors.Error {
	if toAdd.Name != "" {
		return nil
	}
	return derrors.NewInvalidArgumentError("organization required fields missing")
}

func (h * Handler) validOrganizationId(orgId * grpc_organization_go.OrganizationId) derrors.Error {
	if orgId.OrganizationId != "" {
		return nil
	}
	return derrors.NewInvalidArgumentError("organization id is not valid")
}

func (h * Handler) validUpdateOrganization(toUpdate * grpc_organization_go.UpdateOrganizationRequest) derrors.Error {
	return nil
}

func (h *Handler) AddOrganization(ctx context.Context, addOrganizationRequest *grpc_organization_go.AddOrganizationRequest) (*grpc_organization_go.Organization, error) {
	err := h.validAddOrganization(addOrganizationRequest)
	if err != nil {
		return nil, conversions.ToGRPCError(err)
	}
	org, err := h.Manager.AddOrganization(*addOrganizationRequest)
	if err != nil {
		return nil, conversions.ToGRPCError(err)
	}
	return org.ToGRPC(), nil
}
// GetOrganization retrieves the profile information of a given organization.
func (h *Handler) GetOrganization(ctx context.Context, organizationId *grpc_organization_go.OrganizationId) (*grpc_organization_go.Organization, error) {
	retrieved, err := h.Manager.GetOrganization(*organizationId)
	if err != nil {
		return nil, conversions.ToGRPCError(err)
	}
	return retrieved.ToGRPC(), nil
}
// UpdateOrganization updates the public information of an organization.
func (h *Handler) UpdateOrganization(ctx context.Context, updateOrganizationRequest *grpc_organization_go.UpdateOrganizationRequest) (*grpc_common_go.Success, error) {
	notImplemented := derrors.NewUnimplementedError("update organization")

	return nil, conversions.ToGRPCError(notImplemented)
}

