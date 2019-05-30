/*
 * Copyright (C)  2018 Nalej - All Rights Reserved
 */

package application

import (
	"context"
	"github.com/nalej/derrors"
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

// UpdateAppDescriptor allows the user to update the information of a registered descriptor.
func (h *Handler) UpdateAppDescriptor(ctx context.Context, request *grpc_application_go.UpdateAppDescriptorRequest) (*grpc_application_go.AppDescriptor, error){
	err := entities.ValidUpdateAppDescriptorRequest(request)
	if err != nil {
		return nil, conversions.ToGRPCError(err)
	}
	updated, err := h.Manager.UpdateAppDescriptor(request)
	if err != nil {
		return nil, conversions.ToGRPCError(err)
	}
	return updated.ToGRPC(), nil
}

// RemoveAppDescriptor removes an application descriptor.
func (h *Handler) RemoveAppDescriptor(ctx context.Context, appDescID *grpc_application_go.AppDescriptorId) (*grpc_common_go.Success, error){
	err := h.Manager.RemoveAppDescriptor(appDescID)
	if err != nil {
		return nil, conversions.ToGRPCError(err)
	}
	return &grpc_common_go.Success{},nil
}

// GetDescriptorAppParameters retrieves a list of application parameters of a descriptor
func (h *Handler) GetDescriptorAppParameters(ctx context.Context, request *grpc_application_go.AppDescriptorId) (*grpc_application_go.AppParameterList, error){

	err := entities.ValidAppDescriptorId(request)
	if err != nil {
		return nil, conversions.ToGRPCError(err)
	}
	parameters, err := h.Manager.GetDescriptorAppParameters(request)
	if err != nil {
		return nil, conversions.ToGRPCError(err)
	}

	toReturn := make([]*grpc_application_go.AppParameter, 0)
	for _, param := range parameters {
		toReturn = append(toReturn, param.ToGRPC())
	}

	return &grpc_application_go.AppParameterList{
		Parameters: toReturn,
	}, nil
}
// GetInstanceParameters retrieves a list of application parameters of an instance
func (h *Handler) GetInstanceParameters(ctx context.Context, request *grpc_application_go.AppInstanceId) (*grpc_application_go.InstanceParameterList, error) {
	err := entities.ValidAppInstanceId(request)
	if err != nil {
		return nil, conversions.ToGRPCError(err)
	}
	instanceParams, err := h.Manager.GetInstanceParameters(request)
	if err != nil {
		return nil, conversions.ToGRPCError(err)
	}

	toReturn := make([]*grpc_application_go.InstanceParameter, 0)
	for _, param := range instanceParams {
		toReturn = append(toReturn, param.ToGRPC())
	}

	return &grpc_application_go.InstanceParameterList{
		Parameters: toReturn,
	}, nil


	return nil, nil
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

// UpdateAppInstance adds a new application instance to a given organization.
func (h *Handler) UpdateAppInstance(ctx context.Context, appInstance *grpc_application_go.AppInstance) (*grpc_common_go.Success, error) {
	// TODO validate the application instance
	log.Debug().Interface("updatedAppInstance", appInstance).Msg("Update application instance")
	err := h.Manager.UpdateAppInstance(appInstance)
	if err != nil {
		return nil, err
	}
	return &grpc_common_go.Success{},nil
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
		log.Warn().Interface("updateServiceStatus", updateServiceStatus).Msg("error in validation updating ServiceStatus")
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


// AddServiceGroupInstance to an already existing application instance
func (h *Handler) AddServiceGroupInstances(ctx context.Context, addRequest *grpc_application_go.AddServiceGroupInstancesRequest) (*grpc_application_go.ServiceGroupInstancesList, error){
	err := entities.ValidAddServiceGroupInstanceRequest(addRequest)
	if err != nil {
		log.Warn().Interface("addRequest", addRequest).Msg("error in validation adding serviceGroupInstances")
		return nil, conversions.ToGRPCError(err)
	}

	instances, err := h.Manager.AddServiceGroupInstances(addRequest)
	if err != nil {
		return nil, conversions.ToGRPCError(err)
	}

	result := make([]*grpc_application_go.ServiceGroupInstance, len(instances))
	for i, inst := range instances {
		result[i] = inst.ToGRPC()
	}

	toReturn := grpc_application_go.ServiceGroupInstancesList{
		ServiceGroupInstances: result,
	}

	return &toReturn, nil
}


func (h *Handler) RemoveServiceGroupInstances(ctx context.Context, removeRequest *grpc_application_go.RemoveServiceGroupInstancesRequest) (*grpc_common_go.Success, error) {
	err := entities.ValidateRemoveServiceGroupInstancesRequest(removeRequest)
	if err != nil {
		return nil,conversions.ToGRPCError(err)
	}

	err = h.Manager.RemoveServiceGroupInstances(removeRequest)
	if err != nil {
		return nil,conversions.ToGRPCError(err)
	}
	return &grpc_common_go.Success{}, nil
}


// AddServiceInstance to an already existing service group instance
func (h *Handler) AddServiceInstance(ctx context.Context, addRequest *grpc_application_go.AddServiceInstanceRequest) (*grpc_application_go.ServiceInstance, error) {
	err := entities.ValidAddServiceInstanceRequest(addRequest)
	if err != nil {
		return nil, conversions.ToGRPCError(err)
	}

	serviceInstance, err := h.Manager.AddServiceInstance(addRequest)
	if err != nil {
		return nil, conversions.ToGRPCError(err)
	}

	return serviceInstance.ToGRPC(), nil
}


// GetServiceGroupInstanceMetadata returns the metadata entry of an existing ServiceGroupInstance
func (h *Handler) GetServiceGroupInstanceMetadata(ctx context.Context, getRequest *grpc_application_go.GetServiceGroupInstanceMetadataRequest) (*grpc_application_go.InstanceMetadata, error) {
	err := entities.ValidGetServiceGroupInstanceMetadataRequest(getRequest)
	if err != nil {
		return nil, conversions.ToGRPCError(err)
	}

	metadata, err := h.Manager.GetServiceGroupInstanceMetadata(getRequest)
	if err != nil {
		return nil, conversions.ToGRPCError(err)
	}

	return metadata.ToGRPC(), nil
}


// UpdateServiceGroupInstanceMetadata updates the value of an existing metadata instance
func (h *Handler) UpdateServiceGroupInstanceMetadata(ctx context.Context, updateMetadataRequest *grpc_application_go.InstanceMetadata) (*grpc_common_go.Success, error) {
	err := entities.ValidUpdateInstanceMetadata(updateMetadataRequest)
	if err != nil {
		log.Warn().Interface("updateMetadataRequest", updateMetadataRequest).Msg("error in validation updating serviceGroupInstance Metadata")
		return nil, conversions.ToGRPCError(err)
	}

	err = h.Manager.UpdateServiceGroupInstanceMetadata(updateMetadataRequest)
	if err != nil {
		return nil, conversions.ToGRPCError(err)
	}

	return &grpc_common_go.Success{}, nil

}

// AddAppEndPoint adds a new App Endpoint to a given service instance
func (h *Handler) AddAppEndpoint(ctx context.Context, request *grpc_application_go.AddAppEndpointRequest) (*grpc_common_go.Success, error){
	err := entities.ValidAddAppEndpointRequest(request)
	if err != nil {
		return nil, conversions.ToGRPCError(err)
	}
	err = h.Manager.AddAppEndpoint(request)
	if err != nil {
		return nil, conversions.ToGRPCError(err)
	}
	return &grpc_common_go.Success{}, nil
}
// GetAppEndPoint retrieves an appEndpoint
func (h *Handler) GetAppEndpoints(ctx context.Context, request *grpc_application_go.GetAppEndPointRequest) (*grpc_application_go.AppEndpointList, error){
	err := entities.ValidGetAppEndPointRequest(request)
	if err != nil {
		return nil, conversions.ToGRPCError(err)
	}
	endpoint, err := h.Manager.GetAppEndpoint(request)
	if err != nil {
		return nil, conversions.ToGRPCError(err)
	}
	return endpoint, nil
}

func (h *Handler) RemoveAppEndpoints(ctx context.Context, request *grpc_application_go.RemoveAppEndpointRequest) (*grpc_common_go.Success, error){
	err := entities.ValidRemoveEndpointRequest(request)
	if err != nil {
		return nil, conversions.ToGRPCError(err)
	}
	err = h.Manager.RemoveAppEndpoints(request)
	if err != nil {
		return nil, conversions.ToGRPCError(err)
	}
	return &grpc_common_go.Success{}, nil
}

// AddAppZtNetwork add a new zerotier network for an existing app instance
func (h *Handler) AddAppZtNetwork(ctx context.Context, request *grpc_application_go.AddAppZtNetworkRequest) (*grpc_common_go.Success, error) {
	err := entities.ValidAddAppZtNetworkRequest(request)
	if err != nil {
		return nil, err
	}
	err = h.Manager.AddZtNetwork(request)
	if err != nil {
		return nil, conversions.ToGRPCError(err)
	}
	return &grpc_common_go.Success{}, nil
}

// RemoveAppZtNetwork remove the network instance for an application instance
func (h *Handler) RemoveAppZtNetwork(ctx context.Context, request *grpc_application_go.RemoveAppZtNetworkRequest) (*grpc_common_go.Success, error) {
	err := entities.ValidRemoveAppZtNetworkRequest(request)
	if err != nil {
		return nil, err
	}
	err = h.Manager.RemoveZtNetwork(request)
	if err != nil {
		return nil, conversions.ToGRPCError(err)
	}
	return &grpc_common_go.Success{}, nil
}

// GetAppZtnetwork get an existing network instance associated with an application.
func (h *Handler) GetAppZtNetwork(ctx context.Context, request *grpc_application_go.GetAppZtNetworkRequest) (*grpc_application_go.AppZtNetwork, error) {
	err := entities.ValidGetAppZtNetworkRequest(request)
	if err != nil {
		return nil, err
	}
	retrieved, err := h.Manager.GetAppZtNetwork(request)
	if err != nil {
		return nil, conversions.ToGRPCError(err)
	}
	return retrieved.ToGRPC(), nil
}

// AddParametrizedDescriptor adds a parametrized descriptor to a given descriptor
func (h *Handler) AddParametrizedDescriptor(ctx context.Context, request *grpc_application_go.ParametrizedDescriptor) (*grpc_application_go.ParametrizedDescriptor, error) {

	param, err := h.Manager.AddParametrizedDescriptor(request)
	if err != nil {
		return nil, conversions.ToGRPCError(err)
	}

	return param.ToGRPC() ,nil
}
// GetParametrizedDescriptor retrieves the parametrized descriptor associated with an instance
func (h *Handler) GetParametrizedDescriptor(ctx context.Context, instanceID *grpc_application_go.AppInstanceId) (*grpc_application_go.ParametrizedDescriptor, error) {

	err := entities.ValidAppInstanceId(instanceID)
	if err != nil {
		return nil, conversions.ToGRPCError(err)
	}

	descriptor, err := h.Manager.GetParametrizedDescriptor(instanceID)
	if err != nil {
		return nil, conversions.ToGRPCError(err)
	}

	return descriptor.ToGRPC(), nil
}
// RemoveParametrizedDescriptor removes the parametrized descriptor associated with an instance
func (h *Handler) RemoveParametrizedDescriptor(ctx context.Context, instanceID *grpc_application_go.AppInstanceId) (*grpc_common_go.Success, error) {
	err := entities.ValidAppInstanceId(instanceID)
	if err != nil {
		return nil, conversions.ToGRPCError(err)
	}

	err = h.Manager.RemoveParametrizedDescriptor(instanceID)
	if err != nil {
		return nil, conversions.ToGRPCError(err)
	}

	return &grpc_common_go.Success{}, nil
}

// Add Zt member authorization data
func (h *Handler) AddAuthorizedZtNetworkMember(_ context.Context, in *grpc_application_go.AddAuthorizedZtNetworkMemberRequest) (*grpc_application_go.ZtNetworkMember, error) {
	return nil, derrors.NewUnimplementedError("not implemented")
}
// Delete Zt member authorization data
func (h *Handler) RemoveAuthorizedZtNetworkMember(_ context.Context, in *grpc_application_go.RemoveAuthorizedZtNetworkMemberRequest) (*grpc_common_go.Success, error){
	return nil, derrors.NewUnimplementedError("not implemented")
}

// Delete all the members of a Zt network
func (h *Handler) RemoveCompleteAppZtNetworkMemberNet(_ context.Context, in *grpc_application_go.RemoveCompleteAppZtNetworkMemberNetRequest) (*grpc_common_go.Success, error){
	return nil, derrors.NewUnimplementedError("not implemented")
}