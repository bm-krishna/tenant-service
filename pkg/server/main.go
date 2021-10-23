package server

import (
	"context"
	"encoding/json"
	"log"

	"github.com/bm-krishna/tenant-service/pkg/api/tenant"
)

type TenantServer struct {
	tenant.UnimplementedTenantServer
}

// API implement all TenantServer api Methods
func (server *TenantServer) API(ctx context.Context, tenantRequest *tenant.TenantRequest) (*tenant.TenantResponse, error) {
	payload := make(map[string]interface{})
	req := tenantRequest.Request

	err := json.Unmarshal(req, &payload)
	if err != nil {
		log.Fatal("Failed to unamrshal the payload")
		return nil, err
	}
	name, ok := payload["name"]
	resp := make(map[string]interface{})
	if !ok {
		resp["err"] = "Failed to Process Request"
	}
	resp["name"] = name.(string) + "krishna"
	respInBytes, err := json.Marshal(resp)
	if err != nil {
		log.Fatal("Failed to marshal the resp")
		return nil, err
	}
	return &tenant.TenantResponse{
		Response: respInBytes,
	}, nil
}

// Create api implementation which is definie in rpc protocals
func (server *TenantServer) Create(ctx context.Context, request *tenant.TenantRequest) (*tenant.TenantResponse, error) {
	return &tenant.TenantResponse{
		Response: []byte{},
	}, nil
}

// Fetch ...
func (server *TenantServer) Fetch(ctx context.Context, request *tenant.TenantRequest) (*tenant.TenantResponse, error) {
	return &tenant.TenantResponse{
		Response: []byte{},
	}, nil
}

// Update ...
func (server *TenantServer) Update(ctx context.Context, request *tenant.TenantRequest) (*tenant.TenantResponse, error) {
	return &tenant.TenantResponse{
		Response: []byte{},
	}, nil
}

// Delete ...
func (server *TenantServer) Delete(ctx context.Context, request *tenant.TenantRequest) (*tenant.TenantResponse, error) {
	return &tenant.TenantResponse{
		Response: []byte{},
	}, nil
}

// GetByID ...
func (server *TenantServer) GetByID(ctx context.Context, request *tenant.TenantRequest) (*tenant.TenantResponse, error) {
	return &tenant.TenantResponse{
		Response: []byte{},
	}, nil
}
