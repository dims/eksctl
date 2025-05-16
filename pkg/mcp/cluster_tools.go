package mcp

import (
	"context"
	"encoding/json"
	"fmt"
	"os/exec"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

// ClusterCreateParams defines parameters for cluster creation
type ClusterCreateParams struct {
	Name            string `json:"name,omitempty"`
	Region          string `json:"region,omitempty"`
	Version         string `json:"version,omitempty"`
	NodeType        string `json:"nodeType,omitempty"`
	Nodes           int    `json:"nodes,omitempty"`
	WithoutNodeGroup bool   `json:"withoutNodeGroup,omitempty"`
	ConfigFile      string `json:"configFile,omitempty"`
}

// ClusterDeleteParams defines parameters for cluster deletion
type ClusterDeleteParams struct {
	Name   string `json:"name"`
	Region string `json:"region,omitempty"`
	Force  bool   `json:"force,omitempty"`
	Wait   bool   `json:"wait,omitempty"`
}

// ClusterGetParams defines parameters for getting cluster info
type ClusterGetParams struct {
	Name   string `json:"name,omitempty"`
	Region string `json:"region,omitempty"`
	Output string `json:"output,omitempty"`
}

// ClusterUpdateParams defines parameters for updating a cluster
type ClusterUpdateParams struct {
	Name    string `json:"name"`
	Region  string `json:"region,omitempty"`
	Version string `json:"version,omitempty"`
	Timeout string `json:"timeout,omitempty"`
}

// ClusterDescribeParams defines parameters for describing a cluster
type ClusterDescribeParams struct {
	Name   string `json:"name"`
	Region string `json:"region,omitempty"`
	Output string `json:"output,omitempty"`
}

func registerClusterTools(s *server.MCPServer) {
	// Register cluster_create tool
	s.AddTool(
		mcp.NewTool("cluster_create", mcp.WithDescription("Create a new EKS cluster")),
		func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			var p ClusterCreateParams
			if err := json.Unmarshal([]byte(fmt.Sprintf("%v", request.Params.Arguments)), &p); err != nil {
				return mcp.NewToolResultError("Failed to parse parameters"), err
			}

			args := []string{"create", "cluster"}

			if p.Name != "" {
				args = append(args, "--name", p.Name)
			}
			if p.Region != "" {
				args = append(args, "--region", p.Region)
			}
			if p.Version != "" {
				args = append(args, "--version", p.Version)
			}
			if p.NodeType != "" {
				args = append(args, "--node-type", p.NodeType)
			}
			if p.Nodes > 0 {
				args = append(args, "--nodes", fmt.Sprintf("%d", p.Nodes))
			}
			if p.WithoutNodeGroup {
				args = append(args, "--without-nodegroup")
			}
			if p.ConfigFile != "" {
				args = append(args, "--config-file", p.ConfigFile)
			}

			cmd := exec.Command("eksctl", args...)
			output, err := cmd.CombinedOutput()
			if err != nil {
				return mcp.NewToolResultError(fmt.Sprintf("Failed to create cluster: %s\n%s", err, string(output))), nil
			}

			return mcp.NewToolResultText(string(output)), nil
		},
	)

	// Register cluster_delete tool
	s.AddTool(
		mcp.NewTool("cluster_delete", mcp.WithDescription("Delete an EKS cluster")),
		func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			var p ClusterDeleteParams
			if err := json.Unmarshal([]byte(fmt.Sprintf("%v", request.Params.Arguments)), &p); err != nil {
				return mcp.NewToolResultError("Failed to parse parameters"), err
			}

			args := []string{"delete", "cluster", "--name", p.Name}

			if p.Region != "" {
				args = append(args, "--region", p.Region)
			}
			if p.Force {
				args = append(args, "--force")
			}
			if p.Wait {
				args = append(args, "--wait")
			}

			cmd := exec.Command("eksctl", args...)
			output, err := cmd.CombinedOutput()
			if err != nil {
				return mcp.NewToolResultError(fmt.Sprintf("Failed to delete cluster: %s\n%s", err, string(output))), nil
			}

			return mcp.NewToolResultText(string(output)), nil
		},
	)

	// Register cluster_get tool
	s.AddTool(
		mcp.NewTool("cluster_get", mcp.WithDescription("Get information about EKS clusters")),
		func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			var p ClusterGetParams
			if err := json.Unmarshal([]byte(fmt.Sprintf("%v", request.Params.Arguments)), &p); err != nil {
				return mcp.NewToolResultError("Failed to parse parameters"), err
			}

			args := []string{"get", "cluster"}

			if p.Name != "" {
				args = append(args, "--name", p.Name)
			}
			if p.Region != "" {
				args = append(args, "--region", p.Region)
			}
			if p.Output != "" {
				args = append(args, "--output", p.Output)
			}

			cmd := exec.Command("eksctl", args...)
			output, err := cmd.CombinedOutput()
			if err != nil {
				return mcp.NewToolResultError(fmt.Sprintf("Failed to get cluster info: %s\n%s", err, string(output))), nil
			}

			return mcp.NewToolResultText(string(output)), nil
		},
	)

	// Register cluster_update tool
	s.AddTool(
		mcp.NewTool("cluster_update", mcp.WithDescription("Update an EKS cluster")),
		func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			var p ClusterUpdateParams
			if err := json.Unmarshal([]byte(fmt.Sprintf("%v", request.Params.Arguments)), &p); err != nil {
				return mcp.NewToolResultError("Failed to parse parameters"), err
			}

			args := []string{"update", "cluster", "--name", p.Name}

			if p.Region != "" {
				args = append(args, "--region", p.Region)
			}
			if p.Version != "" {
				args = append(args, "--version", p.Version)
			}
			if p.Timeout != "" {
				args = append(args, "--timeout", p.Timeout)
			}

			cmd := exec.Command("eksctl", args...)
			output, err := cmd.CombinedOutput()
			if err != nil {
				return mcp.NewToolResultError(fmt.Sprintf("Failed to update cluster: %s\n%s", err, string(output))), nil
			}

			return mcp.NewToolResultText(string(output)), nil
		},
	)

	// Register cluster_describe tool
	s.AddTool(
		mcp.NewTool("cluster_describe", mcp.WithDescription("Get detailed information about an EKS cluster")),
		func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			var p ClusterDescribeParams
			if err := json.Unmarshal([]byte(fmt.Sprintf("%v", request.Params.Arguments)), &p); err != nil {
				return mcp.NewToolResultError("Failed to parse parameters"), err
			}

			// eksctl doesn't have a direct "describe cluster" command, so we use AWS CLI
			args := []string{"eks", "describe-cluster", "--name", p.Name}

			if p.Region != "" {
				args = append(args, "--region", p.Region)
			}
			if p.Output != "" && (p.Output == "json" || p.Output == "yaml" || p.Output == "text") {
				args = append(args, "--output", p.Output)
			}

			cmd := exec.Command("aws", args...)
			output, err := cmd.CombinedOutput()
			if err != nil {
				return mcp.NewToolResultError(fmt.Sprintf("Failed to describe cluster: %s\n%s", err, string(output))), nil
			}

			return mcp.NewToolResultText(string(output)), nil
		},
	)
}
