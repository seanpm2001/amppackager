// Code generated by protoc-gen-goext. DO NOT EDIT.

package kafka

import (
	fieldmaskpb "google.golang.org/protobuf/types/known/fieldmaskpb"
)

func (m *GetConnectorRequest) SetClusterId(v string) {
	m.ClusterId = v
}

func (m *GetConnectorRequest) SetConnectorName(v string) {
	m.ConnectorName = v
}

func (m *ListConnectorsRequest) SetClusterId(v string) {
	m.ClusterId = v
}

func (m *ListConnectorsRequest) SetPageSize(v int64) {
	m.PageSize = v
}

func (m *ListConnectorsRequest) SetPageToken(v string) {
	m.PageToken = v
}

func (m *ListConnectorsResponse) SetConnectors(v []*Connector) {
	m.Connectors = v
}

func (m *ListConnectorsResponse) SetNextPageToken(v string) {
	m.NextPageToken = v
}

func (m *CreateConnectorRequest) SetClusterId(v string) {
	m.ClusterId = v
}

func (m *CreateConnectorRequest) SetConnectorSpec(v *ConnectorSpec) {
	m.ConnectorSpec = v
}

func (m *CreateConnectorMetadata) SetClusterId(v string) {
	m.ClusterId = v
}

func (m *CreateConnectorMetadata) SetConnectorName(v string) {
	m.ConnectorName = v
}

func (m *UpdateConnectorRequest) SetClusterId(v string) {
	m.ClusterId = v
}

func (m *UpdateConnectorRequest) SetConnectorName(v string) {
	m.ConnectorName = v
}

func (m *UpdateConnectorRequest) SetUpdateMask(v *fieldmaskpb.FieldMask) {
	m.UpdateMask = v
}

func (m *UpdateConnectorRequest) SetConnectorSpec(v *UpdateConnectorSpec) {
	m.ConnectorSpec = v
}

func (m *UpdateConnectorMetadata) SetClusterId(v string) {
	m.ClusterId = v
}

func (m *UpdateConnectorMetadata) SetConnectorName(v string) {
	m.ConnectorName = v
}

func (m *DeleteConnectorRequest) SetClusterId(v string) {
	m.ClusterId = v
}

func (m *DeleteConnectorRequest) SetConnectorName(v string) {
	m.ConnectorName = v
}

func (m *DeleteConnectorMetadata) SetClusterId(v string) {
	m.ClusterId = v
}

func (m *DeleteConnectorMetadata) SetConnectorName(v string) {
	m.ConnectorName = v
}

func (m *ResumeConnectorRequest) SetClusterId(v string) {
	m.ClusterId = v
}

func (m *ResumeConnectorRequest) SetConnectorName(v string) {
	m.ConnectorName = v
}

func (m *ResumeConnectorMetadata) SetClusterId(v string) {
	m.ClusterId = v
}

func (m *ResumeConnectorMetadata) SetConnectorName(v string) {
	m.ConnectorName = v
}

func (m *PauseConnectorRequest) SetClusterId(v string) {
	m.ClusterId = v
}

func (m *PauseConnectorRequest) SetConnectorName(v string) {
	m.ConnectorName = v
}

func (m *PauseConnectorMetadata) SetClusterId(v string) {
	m.ClusterId = v
}

func (m *PauseConnectorMetadata) SetConnectorName(v string) {
	m.ConnectorName = v
}