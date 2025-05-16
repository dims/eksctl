package mcp

import (
	"context"
	"encoding/json"
	"fmt"
	"os/exec"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

// AddonInstallParams defines parameters for addon installation
type AddonInstallParams struct {
	ClusterName           string `json:"clusterName"`
	Name                  string `json:"name"`
	Version               string `json:"version,omitempty"`
	Region                string `json:"region,omitempty"`
	ServiceAccountRoleARN string `json:"serviceAccountRoleARN,omitempty"`
	Force                 bool   `json:"force,omitempty"`
}

// AddonUpdateParams defines parameters for addon update
type AddonUpdateParams struct {
	ClusterName           string `json:"clusterName"`
	Name                  string `json:"name"`
	Version               string `json:"version,omitempty"`
	Region                string `json:"region,omitempty"`
	ServiceAccountRoleARN string `json:"serviceAccountRoleARN,omitempty"`
	Force                 bool   `json:"force,omitempty"`
}

// AddonDeleteParams defines parameters for addon deletion
type AddonDeleteParams struct {
	ClusterName string `json:"clusterName"`
	Name        string `json:"name"`
	Region      string `json:"region,omitempty"`
	Preserve    bool   `json:"preserve,omitempty"`
}

// AddonListParams defines parameters for listing addons
type AddonListParams struct {
	ClusterName string `json:"clusterName"`
	Region      string `json:"region,omitempty"`
	Output      string `json:"output,omitempty"`
}

func registerAddonTools(s *server.MCPServer) {
	// Register addon_install tool
	s.AddTool(
		mcp.NewTool("addon_install", mcp.WithDescription("Install an addon to an EKS cluster")),
		func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			var p AddonInstallParams
			if err := json.Unmarshal([]byte(fmt.Sprintf("%v", request.Params.Arguments)), &p); err != nil {
				return mcp.NewToolResultError("Failed to parse parameters"), err
			}

			args := []string{"create", "addon", "--cluster", p.ClusterName, "--name", p.Name}

			if p.Version != "" {
				args = append(args, "--version", p.Version)
			}
			if p.Region != "" {
				args = append(args, "--region", p.Region)
			}
			if p.ServiceAccountRoleARN != "" {
				args = append(args, "--service-account-role-arn", p.ServiceAccountRoleARN)
			}
			if p.Force {
				args = append(args, "--force")
			}

			cmd := exec.Command("eksctl", args...)
			output, err := cmd.CombinedOutput()
			if err != nil {
				return mcp.NewToolResultError(fmt.Sprintf("Failed to install addon: %s\n%s", err, string(output))), nil
			}

			return mcp.NewToolResultText(string(output)), nil
		},
	)

	// Register addon_update tool
	s.AddTool(
		mcp.NewTool("addon_update", mcp.WithDescription("Update an addon in an EKS cluster")),
		func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			var p AddonUpdateParams
			if err := json.Unmarshal([]byte(fmt.Sprintf("%v", request.Params.Arguments)), &p); err != nil {
				return mcp.NewToolResultError("Failed to parse parameters"), err
			}

			args := []string{"update", "addon", "--cluster", p.ClusterName, "--name", p.Name}

			if p.Version != "" {
				args = append(args, "--version", p.Version)
			}
			if p.Region != "" {
				args = append(args, "--region", p.Region)
			}
			if p.ServiceAccountRoleARN != "" {
				args = append(args, "--service-account-role-arn", p.ServiceAccountRoleARN)
			}
			if p.Force {
				args = append(args, "--force")
			}

			cmd := exec.Command("eksctl", args...)
			output, err := cmd.CombinedOutput()
			if err != nil {
				return mcp.NewToolResultError(fmt.Sprintf("Failed to update addon: %s\n%s", err, string(output))), nil
			}

			return mcp.NewToolResultText(string(output)), nil
		},
	)

	// Register addon_delete tool
	s.AddTool(
		mcp.NewTool("addon_delete", mcp.WithDescription("Delete an addon from an EKS cluster")),
		func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			var p AddonDeleteParams
			if err := json.Unmarshal([]byte(fmt.Sprintf("%v", request.Params.Arguments)), &p); err != nil {
				return mcp.NewToolResultError("Failed to parse parameters"), err
			}

			args := []string{"delete", "addon", "--cluster", p.ClusterName, "--name", p.Name}

			if p.Region != "" {
				args = append(args, "--region", p.Region)
			}
			if p.Preserve {
				args = append(args, "--preserve")
			}

			cmd := exec.Command("eksctl", args...)
			output, err := cmd.CombinedOutput()
			if err != nil {
				return mcp.NewToolResultError(fmt.Sprintf("Failed to delete addon: %s\n%s", err, string(output))), nil
			}

			return mcp.NewToolResultText(string(output)), nil
		},
	)

	// Register addon_list tool
	s.AddTool(
		mcp.NewTool("addon_list", mcp.WithDescription("List all addons in an EKS cluster")),
		func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			var p AddonListParams
			if err := json.Unmarshal([]byte(fmt.Sprintf("%v", request.Params.Arguments)), &p); err != nil {
				return mcp.NewToolResultError("Failed to parse parameters"), err
			}

			args := []string{"get", "addon", "--cluster", p.ClusterName}

			if p.Region != "" {
				args = append(args, "--region", p.Region)
			}
			if p.Output != "" {
				args = append(args, "--output", p.Output)
			}

			cmd := exec.Command("eksctl", args...)
			output, err := cmd.CombinedOutput()
			if err != nil {
				return mcp.NewToolResultError(fmt.Sprintf("Failed to list addons: %s\n%s", err, string(output))), nil
			}

			return mcp.NewToolResultText(string(output)), nil
		},
	)
}
