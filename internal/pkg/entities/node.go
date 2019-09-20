/*
 * Copyright (C)  2018 Nalej - All Rights Reserved
 */

package entities

import (
	"github.com/nalej/derrors"
	"github.com/nalej/grpc-infrastructure-go"
)

type NodeState int

const (
	// Unregistered nodes are those whose details are in the platform but we have not perform any action.
	Unregistered NodeState = iota + 1
	// Unassigned nodes are those that have been prepared but are not assigned to a given cluster.
	Unassigned
	// Assigned nodes are those that have been installed and are part of a cluster.
	Assigned
)

var NodeStateToGRPC = map[NodeState]grpc_infrastructure_go.NodeState{
	Unregistered: grpc_infrastructure_go.NodeState_UNREGISTERED,
	Unassigned: grpc_infrastructure_go.NodeState_UNASSIGNED,
	Assigned: grpc_infrastructure_go.NodeState_ASSIGNED,
}

var NodeStateFromGRPC = map[grpc_infrastructure_go.NodeState]NodeState{
	grpc_infrastructure_go.NodeState_UNREGISTERED: Unregistered,
	grpc_infrastructure_go.NodeState_UNASSIGNED: Unassigned,
	grpc_infrastructure_go.NodeState_ASSIGNED: Assigned,
}

// Node entity representing a single node of the architecture that executes application instances.
type Node struct {
	// OrganizationId with the organization identifier.
	OrganizationId string `json:"organization_id,omitempty"`
	// ClusterId with the associated cluster identifier the node is assigned to.
	ClusterId string `json:"cluster_id,omitempty"`
	// ClusterId with the node identifier.
	NodeId string `json:"node_id,omitempty"`
	// Ip with the node IP.
	Ip string `json:"ip,omitempty"`
	// Labels for the node.
	Labels map[string]string `json:"labels,omitempty"`
	// Status of the node based on monitoring information.
	Status InfraStatus `json:"status,omitempty"`
	// State of assignation of the node.
	State                NodeState `json:"state,omitempty"`
}

func NewNodeFromGRPC(addNodeRequest *grpc_infrastructure_go.AddNodeRequest) * Node {
	uuid := GenerateUUID()

	return &Node{
		OrganizationId: addNodeRequest.OrganizationId,
		ClusterId:      "",
		NodeId:         uuid,
		Ip:             addNodeRequest.Ip,
		Labels:         addNodeRequest.Labels,
		Status:         InfraStatusInstalling,
		State:          Unregistered,
	}
}

func (n * Node) ToGRPC() * grpc_infrastructure_go.Node {
	status := InfraStatusToGRPC[n.Status]
	state := NodeStateToGRPC[n.State]
	return &grpc_infrastructure_go.Node{
		OrganizationId:       n.OrganizationId,
		ClusterId:            n.ClusterId,
		NodeId:               n.NodeId,
		Ip:                   n.Ip,
		Labels:               n.Labels,
		Status:               status,
		State:                state,
	}
}

func (n * Node) ApplyUpdate(updateRequest grpc_infrastructure_go.UpdateNodeRequest){
	if updateRequest.AddLabels {
		for k, v := range updateRequest.Labels {
			n.Labels[k] = v
		}
	}
	if updateRequest.RemoveLabels {
		for k, _ := range updateRequest.Labels {
			delete(n.Labels, k)
		}
	}
	if updateRequest.UpdateStatus{
		n.Status = InfraStatusFromGRPC[updateRequest.Status]
	}
	if updateRequest.UpdateState {
		n.State = NodeStateFromGRPC[updateRequest.State]
	}
}

func ValidAddNodeRequest(addNodeRequest *grpc_infrastructure_go.AddNodeRequest) derrors.Error {
	if addNodeRequest.RequestId == "" {
		return derrors.NewInvalidArgumentError(emptyRequestId)
	}
	if addNodeRequest.OrganizationId == "" {
		return derrors.NewInvalidArgumentError(emptyOrganizationId)
	}
	if addNodeRequest.NodeId != "" {
		return derrors.NewInvalidArgumentError("node_id must be empty, and generated by this component")
	}
	if addNodeRequest.Ip == "" {
		return derrors.NewInvalidArgumentError("ip must not be empty")
	}
	return nil
}

func ValidUpdateNodeRequest(updateNodeRequest *grpc_infrastructure_go.UpdateNodeRequest) derrors.Error {
	if updateNodeRequest.OrganizationId == "" {
		return derrors.NewInvalidArgumentError(emptyOrganizationId)
	}
	if updateNodeRequest.NodeId == "" {
		return derrors.NewInvalidArgumentError(emptyNodeId)
	}
	return nil
}


func ValidAttachNodeRequest(attachNodeRequest *grpc_infrastructure_go.AttachNodeRequest) derrors.Error {
	if attachNodeRequest.RequestId == "" {
		return derrors.NewInvalidArgumentError(emptyRequestId)
	}
	if attachNodeRequest.OrganizationId == "" {
		return derrors.NewInvalidArgumentError(emptyOrganizationId)
	}
	if attachNodeRequest.NodeId == "" {
		return derrors.NewInvalidArgumentError(emptyNodeId)
	}
	if attachNodeRequest.ClusterId == "" {
		return derrors.NewInvalidArgumentError(emptyClusterId)
	}
	return nil
}

func ValidRemoveNodesRequest(removeNodesRequest *grpc_infrastructure_go.RemoveNodesRequest) derrors.Error {
	if removeNodesRequest.RequestId == "" {
		return derrors.NewInvalidArgumentError(emptyRequestId)
	}
	if removeNodesRequest.OrganizationId == "" {
		return derrors.NewInvalidArgumentError(emptyOrganizationId)
	}
	if len(removeNodesRequest.Nodes) == 0 {
		return derrors.NewInvalidArgumentError("nodes must not be empty")
	}
	return nil
}
