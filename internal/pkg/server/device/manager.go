package device

import (
	"github.com/nalej/derrors"
	"github.com/nalej/grpc-device-go"
	"github.com/nalej/grpc-organization-go"
	devEnt "github.com/nalej/system-model/internal/pkg/entities/device"
	"github.com/nalej/system-model/internal/pkg/provider/device"
	"github.com/nalej/system-model/internal/pkg/provider/organization"
)

// Manager structure with the required providers for application operations.
type Manager struct {
	DevProvider device.Provider
	OrgProvider organization.Provider
}


// NewManager creates a Manager using a set of providers.
func NewManager(devProvider device.Provider, orgProvider organization.Provider) Manager {
	return Manager{devProvider, orgProvider}
}

// ---------------------------------------------------------------------------------------------------------
func (m *Manager) AddDeviceGroup(addRequest *grpc_device_go.AddDeviceGroupRequest) (* devEnt.DeviceGroup, derrors.Error){

	exists, err := m.OrgProvider.Exists(addRequest.OrganizationId)
	if err != nil{
		return nil, err
	}
	if !exists{
		return nil, derrors.NewNotFoundError("organizationID").WithParams(addRequest.OrganizationId)
	}
	group := devEnt.NewDeviceGroupFromGRPC(addRequest)
	err = m.DevProvider.AddDeviceGroup(*group)
	if err != nil {
		return nil, err
	}

	return group, nil
}
// ListDeviceGroups obtains a list of device groups in an organization.
func (m *Manager) ListDeviceGroups(organizationID *grpc_organization_go.OrganizationId) ([] devEnt.DeviceGroup, derrors.Error){
	exists, err := m.OrgProvider.Exists(organizationID.OrganizationId)
	if err != nil{
		return nil, err
	}
	if !exists{
		return nil, derrors.NewNotFoundError("organizationID").WithParams(organizationID.OrganizationId)
	}
	groups, err := m.DevProvider.ListDeviceGroups(organizationID.OrganizationId)
	if err != nil {
		return nil, err
	}
	return groups, nil

}
// GetDeviceGroup retrieves a given device group in an organization.
func (m *Manager) GetDeviceGroup(deviceGroupID *grpc_device_go.DeviceGroupId) (* devEnt.DeviceGroup, derrors.Error){
	exists, err := m.OrgProvider.Exists(deviceGroupID.OrganizationId)
	if err != nil {
		return nil, err
	}
	if ! exists{
		return nil, derrors.NewNotFoundError("organizationID").WithParams(deviceGroupID.OrganizationId)
	}

	exists, err = m.DevProvider.ExistsDeviceGroup(deviceGroupID.OrganizationId, deviceGroupID.DeviceGroupId)
	if err != nil {
		return nil, err
	}
	if !exists{
		return nil, derrors.NewNotFoundError("DeviceGroup").WithParams(deviceGroupID.OrganizationId, deviceGroupID.DeviceGroupId)
	}
	return m.DevProvider.GetDeviceGroup(deviceGroupID.OrganizationId, deviceGroupID.DeviceGroupId)
}
// RemoveDeviceGroup removes a device group
func (m *Manager) RemoveServiceGroup(removeRequest *grpc_device_go.RemoveDeviceGroupRequest) (derrors.Error){

	exists, err := m.OrgProvider.Exists(removeRequest.OrganizationId)
	if err != nil {
		return err
	}
	if ! exists{
		return derrors.NewNotFoundError("organizationID").WithParams(removeRequest.OrganizationId)
	}

	exists, err = m.DevProvider.ExistsDeviceGroup(removeRequest.OrganizationId, removeRequest.DeviceGroupId)
	if err != nil {
		return err
	}
	if !exists{
		return derrors.NewNotFoundError("device group").WithParams(removeRequest.OrganizationId, removeRequest.DeviceGroupId)
	}

	err = m.DevProvider.RemoveDeviceGroup(removeRequest.OrganizationId, removeRequest.DeviceGroupId)
	if err != nil {
		return err
	}

	return nil
}

// ---------------------------------------------------------------------------------------------------------
// AddDevice adds a new group to the system
func (m *Manager) AddDevice(addRequest *grpc_device_go.AddDeviceRequest) (* devEnt.Device, derrors.Error){

	exists, err := m.OrgProvider.Exists(addRequest.OrganizationId)
	if err != nil{
		return nil, err
	}
	if !exists{
		return nil, derrors.NewNotFoundError("organizationID").WithParams(addRequest.OrganizationId)
	}

	exists, err = m.DevProvider.ExistsDeviceGroup(addRequest.OrganizationId, addRequest.DeviceGroupId)
	if err != nil{
		return nil, err
	}
	if !exists{
		return nil, derrors.NewNotFoundError("deviceGroup").WithParams(addRequest.OrganizationId, addRequest.DeviceGroupId)
	}

	device := devEnt.NewDeviceFromGRPC(addRequest)
	err = m.DevProvider.AddDevice(*device)
	if err != nil {
		return nil, err
	}

	return device, nil

}
// ListDevice obtains a list of devices in a device_group
func (m *Manager) ListDevices(deviceGroupID *grpc_device_go.DeviceGroupId) ([] devEnt.Device, derrors.Error){

	exists, err := m.OrgProvider.Exists(deviceGroupID.OrganizationId)
	if err != nil{
		return nil, err
	}
	if !exists{
		return nil, derrors.NewNotFoundError("organizationID").WithParams(deviceGroupID.OrganizationId)
	}
	exists, err = m.DevProvider.ExistsDeviceGroup(deviceGroupID.OrganizationId, deviceGroupID.DeviceGroupId)
	if err != nil{
		return nil, err
	}
	if !exists{
		return nil, derrors.NewNotFoundError("device group").WithParams(deviceGroupID.OrganizationId, deviceGroupID.DeviceGroupId)
	}

	groups, err := m.DevProvider.ListDevices(deviceGroupID.OrganizationId, deviceGroupID.DeviceGroupId)
	if err != nil {
		return nil, err
	}
	return groups, nil
}
// GetDevice retrieves a given device in an organization.
func (m *Manager) GetDevice(deviceID *grpc_device_go.DeviceId) (* devEnt.Device, derrors.Error){

	exists, err := m.OrgProvider.Exists(deviceID.OrganizationId)
	if err != nil {
		return nil, err
	}
	if ! exists{
		return nil, derrors.NewNotFoundError("organizationID").WithParams(deviceID.OrganizationId)
	}

	exists, err = m.DevProvider.ExistsDeviceGroup(deviceID.OrganizationId, deviceID.DeviceGroupId)
	if err != nil {
		return nil, err
	}
	if !exists{
		return nil, derrors.NewNotFoundError("DeviceGroup").WithParams(deviceID.OrganizationId, deviceID.DeviceGroupId)
	}

	exists, err = m.DevProvider.ExistsDevice(deviceID.OrganizationId, deviceID.DeviceGroupId, deviceID.DeviceId)
	if err != nil {
		return nil, err
	}
	if !exists{
		return nil, derrors.NewNotFoundError("Device").WithParams(deviceID.OrganizationId, deviceID.DeviceGroupId, deviceID.DeviceId)
	}


	return m.DevProvider.GetDevice(deviceID.OrganizationId, deviceID.DeviceGroupId, deviceID.DeviceId)
}
// RemoveDevice removes a given device
func (m *Manager) RemoveDevice(removeRequest *grpc_device_go.RemoveDeviceRequest) (derrors.Error) {

	exists, err := m.OrgProvider.Exists(removeRequest.OrganizationId)
	if err != nil {
		return err
	}
	if ! exists{
		return derrors.NewNotFoundError("organizationID").WithParams(removeRequest.OrganizationId)
	}

	exists, err = m.DevProvider.ExistsDeviceGroup(removeRequest.OrganizationId, removeRequest.DeviceGroupId)
	if err != nil {
		return err
	}
	if !exists{
		return derrors.NewNotFoundError("device group").WithParams(removeRequest.OrganizationId, removeRequest.DeviceGroupId)
	}

	exists, err = m.DevProvider.ExistsDevice(removeRequest.OrganizationId, removeRequest.DeviceGroupId, removeRequest.DeviceId)
	if err != nil {
		return err
	}
	if !exists{
		return derrors.NewNotFoundError("Device").WithParams(removeRequest.OrganizationId, removeRequest.DeviceGroupId, removeRequest.DeviceId)
	}

	err = m.DevProvider.RemoveDevice(removeRequest.OrganizationId, removeRequest.DeviceGroupId, removeRequest.DeviceId)
	if err != nil {
		return err
	}

	return nil
}