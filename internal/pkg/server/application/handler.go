/*
 * Copyright (C)  2018 Nalej - All Rights Reserved
 */

package application

import (
	"context"
	"github.com/nalej/grpc-application-go"
	"github.com/nalej/grpc-common-go"
	"github.com/nalej/grpc-organization-go"
	"github.com/nalej/grpc-utils/pkg/conversions"
	"github.com/nalej/system-model/internal/pkg/entities"
	"github.com/rs/zerolog/log"
)

// Handler structure for the application requests.
type Handler struct {
	Manager Manager
}

// NewHandler creates a new Handler with a linked manager.
func NewHandler(manager Manager) *Handler{
	return &Handler{manager}
}

// AddAppDescriptor adds a new application descriptor to a given organization.
func (h *Handler) AddAppDescriptor(ctx context.Context, addRequest *grpc_application_go.AddAppDescriptorRequest) (*grpc_application_go.AppDescriptor, error) {
	err := entities.ValidAddAppDescriptorRequest(addRequest)
	if err != nil {
		return nil, conversions.ToGRPCError(err)
	}
	log.Debug().Interface("addRequest", addRequest).Msg("Adding application descriptor")
	added, err := h.Manager.AddAppDescriptor(addRequest)
	if err != nil {
		return nil, conversions.ToGRPCError(err)
	}
	return added.ToGRPC(), nil
}

// ListAppDescriptors retrieves a list of application descriptors.
func (h *Handler) ListAppDescriptors(ctx context.Context, orgID *grpc_organization_go.OrganizationId) (*grpc_application_go.AppDescriptorList, error) {
	descriptors, err := h.Manager.ListDescriptors(orgID)
	if err != nil{
		return nil, conversions.ToGRPCError(err)
	}

	toReturn := make([]*grpc_application_go.AppDescriptor, 0)
	for _, d := range descriptors {
		toReturn = append(toReturn, d.ToGRPC())
	}
	result := &grpc_application_go.AppDescriptorList{
		Descriptors:          toReturn,
	}
	return result, nil
}

// GetAppDescriptor retrieves a given application descriptor.
func (h *Handler) GetAppDescriptor(ctx context.Context, appDescID *grpc_application_go.AppDescriptorId) (*grpc_application_go.AppDescriptor, error) {
	descriptor, err := h.Manager.GetDescriptor(appDescID)
	if err != nil {
	    return nil, conversions.ToGRPCError(err)
	}
	return descriptor.ToGRPC(), nil
}

// RemoveAppDescriptor removes an application descriptor.
func (h *Handler) RemoveAppDescriptor(ctx context.Context, appDescID *grpc_application_go.AppDescriptorId) (*grpc_common_go.Success, error){
	err := h.Manager.RemoveAppDescriptor(appDescID)
	if err != nil {
		return nil, conversions.ToGRPCError(err)
	}
	return &grpc_common_go.Success{},nil
}


// AddAppInstance adds a new application instance to a given organization.
func (h *Handler) AddAppInstance(ctx context.Context, addInstanceRequest *grpc_application_go.AddAppInstanceRequest) (*grpc_application_go.AppInstance, error) {
	err := entities.ValidAddAppInstanceRequest(addInstanceRequest)
	if err != nil {
		return nil, conversions.ToGRPCError(err)
	}
	log.Debug().Interface("addAppInstance", addInstanceRequest).Msg("Adding application instance")
	added, err := h.Manager.AddAppInstance(addInstanceRequest)
	if err != nil {
		return nil, conversions.ToGRPCError(err)
	}
	return added.ToGRPC(), nil
}

// ListAppInstances retrieves a list of application instances.
func (h *Handler) ListAppInstances(ctx context.Context, orgID *grpc_organization_go.OrganizationId) (*grpc_application_go.AppInstanceList, error) {
	instances, err := h.Manager.ListInstances(orgID)
	if err != nil{
		return nil, conversions.ToGRPCError(err)
	}

	toReturn := make([]*grpc_application_go.AppInstance, 0)
	for _, inst := range instances {
		toReturn = append(toReturn, inst.ToGRPC())
	}
	result := &grpc_application_go.AppInstanceList{
		Instances:          toReturn,
	}
	return result, nil
}

// GetAppInstance retrieves a given application instance.
func (h *Handler) GetAppInstance(ctx context.Context, appInstID *grpc_application_go.AppInstanceId) (*grpc_application_go.AppInstance, error) {
	instance, err := h.Manager.GetInstance(appInstID)
	if err != nil {
		return nil, conversions.ToGRPCError(err)
	}
	return instance.ToGRPC(), nil
}

// UpdateAppStatus updates the status of an application instance.
func (h *Handler) UpdateAppStatus(ctx context.Context, updateAppStatus *grpc_application_go.UpdateAppStatusRequest) (*grpc_common_go.Success, error) {
	err := entities.ValidUpdateAppStatusRequest(updateAppStatus)
	if err != nil {
		return nil, conversions.ToGRPCError(err)
	}

	derr := h.Manager.UpdateInstance(updateAppStatus)
	if derr != nil {
		return nil, derr
	}
	return &grpc_common_go.Success{},nil
}

// UpdateServiceStatus updates the status of an application instance service.
func (h *Handler) UpdateServiceStatus(ctx context.Context, updateServiceStatus *grpc_application_go.UpdateServiceStatusRequest) (*grpc_common_go.Success, error) {
    err := entities.ValidUpdateServiceStatusRequest(updateServiceStatus)
	if err != nil {
	    return nil, conversions.ToGRPCError(err)
    }
    derr := h.Manager.UpdateService(updateServiceStatus)
    if derr != nil {
        return nil, derr
    }
    return &grpc_common_go.Success{},nil
}

// RemoveAppInstance removes an application instance
func (h *Handler) RemoveAppInstance(ctx context.Context, appInstID *grpc_application_go.AppInstanceId) (*grpc_common_go.Success, error){
	err := h.Manager.RemoveAppInstance(appInstID)
	if err != nil {
		return nil, conversions.ToGRPCError(err)
	}
	return &grpc_common_go.Success{},nil
}