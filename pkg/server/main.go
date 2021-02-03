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

func (tenantServer *TenantServer) API(ctx context.Context, tenantRequest *tenant.TenantRequest) (*tenant.TenantResponse, error) {
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
