package mcp

import (
	"context"
	"encoding/json"
	"fmt"
	"os/exec"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

// AutoModeEnableParams defines parameters for enabling Auto Mode
type AutoModeEnableParams struct {
	ClusterName          string `json:"clusterName"`
	Region               string `json:"region,omitempty"`
	ComputeConfiguration string `json:"computeConfiguration,omitempty"`
	StorageConfiguration string `json:"storageConfiguration,omitempty"`
	NetworkConfiguration string `json:"networkConfiguration,omitempty"`
}

// AutoModeConfigureParams defines parameters for configuring Auto Mode
type AutoModeConfigureParams struct {
	ClusterName          string `json:"clusterName"`
	Region               string `json:"region,omitempty"`
	ComputeConfiguration string `json:"computeConfiguration,omitempty"`
	StorageConfiguration string `json:"storageConfiguration,omitempty"`
	NetworkConfiguration string `json:"networkConfiguration,omitempty"`
}

func registerAutoModeTools(s *server.MCPServer) {
	// Register automode_enable tool
	s.AddTool(
		mcp.NewTool("automode_enable", mcp.WithDescription("Enable EKS Auto Mode for a cluster")),
		func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			var p AutoModeEnableParams
			if err := json.Unmarshal([]byte(fmt.Sprintf("%v", request.Params.Arguments)), &p); err != nil {
				return mcp.NewToolResultError("Failed to parse parameters"), err
			}

			args := []string{"enable", "automode", "--cluster", p.ClusterName}

			if p.Region != "" {
				args = append(args, "--region", p.Region)
			}
			if p.ComputeConfiguration != "" {
				args = append(args, "--compute-configuration", p.ComputeConfiguration)
			}
			if p.StorageConfiguration != "" {
				args = append(args, "--storage-configuration", p.StorageConfiguration)
			}
			if p.NetworkConfiguration != "" {
				args = append(args, "--network-configuration", p.NetworkConfiguration)
			}

			cmd := exec.Command("eksctl", args...)
			output, err := cmd.CombinedOutput()
			if err != nil {
				return mcp.NewToolResultError(fmt.Sprintf("Failed to enable Auto Mode: %s\n%s", err, string(output))), nil
			}

			return mcp.NewToolResultText(string(output)), nil
		},
	)

	// Register automode_configure tool
	s.AddTool(
		mcp.NewTool("automode_configure", mcp.WithDescription("Configure EKS Auto Mode settings")),
		func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			var p AutoModeConfigureParams
			if err := json.Unmarshal([]byte(fmt.Sprintf("%v", request.Params.Arguments)), &p); err != nil {
				return mcp.NewToolResultError("Failed to parse parameters"), err
			}

			args := []string{"update", "automode", "--cluster", p.ClusterName}

			if p.Region != "" {
				args = append(args, "--region", p.Region)
			}
			if p.ComputeConfiguration != "" {
				args = append(args, "--compute-configuration", p.ComputeConfiguration)
			}
			if p.StorageConfiguration != "" {
				args = append(args, "--storage-configuration", p.StorageConfiguration)
			}
			if p.NetworkConfiguration != "" {
				args = append(args, "--network-configuration", p.NetworkConfiguration)
			}

			cmd := exec.Command("eksctl", args...)
			output, err := cmd.CombinedOutput()
			if err != nil {
				return mcp.NewToolResultError(fmt.Sprintf("Failed to configure Auto Mode: %s\n%s", err, string(output))), nil
			}

			return mcp.NewToolResultText(string(output)), nil
		},
	)
}
