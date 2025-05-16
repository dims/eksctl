package mcp

import (
	"context"
	"encoding/json"
	"fmt"
	"os/exec"
	"strconv"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

// NodeGroupCreateParams defines parameters for nodegroup creation
type NodeGroupCreateParams struct {
	ClusterName string            `json:"clusterName"`
	Name        string            `json:"name"`
	Region      string            `json:"region,omitempty"`
	NodeType    string            `json:"nodeType,omitempty"`
	Nodes       int               `json:"nodes,omitempty"`
	MinNodes    int               `json:"minNodes,omitempty"`
	MaxNodes    int               `json:"maxNodes,omitempty"`
	AmiFamily   string            `json:"amiFamily,omitempty"`
	Labels      map[string]string `json:"labels,omitempty"`
	Tags        map[string]string `json:"tags,omitempty"`
}

// NodeGroupDeleteParams defines parameters for nodegroup deletion
type NodeGroupDeleteParams struct {
	ClusterName string `json:"clusterName"`
	Name        string `json:"name"`
	Region      string `json:"region,omitempty"`
	Force       bool   `json:"force,omitempty"`
	Drain       bool   `json:"drain,omitempty"`
}

// NodeGroupScaleParams defines parameters for scaling a nodegroup
type NodeGroupScaleParams struct {
	ClusterName string `json:"clusterName"`
	Name        string `json:"name"`
	Region      string `json:"region,omitempty"`
	Nodes       int    `json:"nodes,omitempty"`
	MinNodes    int    `json:"minNodes,omitempty"`
	MaxNodes    int    `json:"maxNodes,omitempty"`
}

// NodeGroupUpdateParams defines parameters for updating a nodegroup
type NodeGroupUpdateParams struct {
	ClusterName    string            `json:"clusterName"`
	Name           string            `json:"name"`
	Region         string            `json:"region,omitempty"`
	MaxUnavailable int               `json:"maxUnavailable,omitempty"`
	Labels         map[string]string `json:"labels,omitempty"`
	Tags           map[string]string `json:"tags,omitempty"`
}

// NodeGroupListParams defines parameters for listing nodegroups
type NodeGroupListParams struct {
	ClusterName string `json:"clusterName"`
	Region      string `json:"region,omitempty"`
	Output      string `json:"output,omitempty"`
}

func registerNodeGroupTools(s *server.MCPServer) {
	// Register nodegroup_create tool
	s.AddTool(
		mcp.NewTool("nodegroup_create", mcp.WithDescription("Create a new nodegroup in an EKS cluster")),
		func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			var p NodeGroupCreateParams
			if err := json.Unmarshal([]byte(fmt.Sprintf("%v", request.Params.Arguments)), &p); err != nil {
				return mcp.NewToolResultError("Failed to parse parameters"), err
			}

			args := []string{"create", "nodegroup", "--cluster", p.ClusterName, "--name", p.Name}

			if p.Region != "" {
				args = append(args, "--region", p.Region)
			}
			if p.NodeType != "" {
				args = append(args, "--node-type", p.NodeType)
			}
			if p.Nodes > 0 {
				args = append(args, "--nodes", strconv.Itoa(p.Nodes))
			}
			if p.MinNodes > 0 {
				args = append(args, "--nodes-min", strconv.Itoa(p.MinNodes))
			}
			if p.MaxNodes > 0 {
				args = append(args, "--nodes-max", strconv.Itoa(p.MaxNodes))
			}
			if p.AmiFamily != "" {
				args = append(args, "--node-ami-family", p.AmiFamily)
			}

			// Handle labels
			for k, v := range p.Labels {
				args = append(args, "--node-labels", fmt.Sprintf("%s=%s", k, v))
			}

			// Handle tags
			for k, v := range p.Tags {
				args = append(args, "--tags", fmt.Sprintf("%s=%s", k, v))
			}

			cmd := exec.Command("eksctl", args...)
			output, err := cmd.CombinedOutput()
			if err != nil {
				return mcp.NewToolResultError(fmt.Sprintf("Failed to create nodegroup: %s\n%s", err, string(output))), nil
			}

			return mcp.NewToolResultText(string(output)), nil
		},
	)

	// Register nodegroup_delete tool
	s.AddTool(
		mcp.NewTool("nodegroup_delete", mcp.WithDescription("Delete a nodegroup from an EKS cluster")),
		func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			var p NodeGroupDeleteParams
			if err := json.Unmarshal([]byte(fmt.Sprintf("%v", request.Params.Arguments)), &p); err != nil {
				return mcp.NewToolResultError("Failed to parse parameters"), err
			}

			args := []string{"delete", "nodegroup", "--cluster", p.ClusterName, "--name", p.Name}

			if p.Region != "" {
				args = append(args, "--region", p.Region)
			}
			if p.Force {
				args = append(args, "--force")
			}
			if p.Drain {
				args = append(args, "--drain")
			}

			cmd := exec.Command("eksctl", args...)
			output, err := cmd.CombinedOutput()
			if err != nil {
				return mcp.NewToolResultError(fmt.Sprintf("Failed to delete nodegroup: %s\n%s", err, string(output))), nil
			}

			return mcp.NewToolResultText(string(output)), nil
		},
	)

	// Register nodegroup_scale tool
	s.AddTool(
		mcp.NewTool("nodegroup_scale", mcp.WithDescription("Scale a nodegroup in an EKS cluster")),
		func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			var p NodeGroupScaleParams
			if err := json.Unmarshal([]byte(fmt.Sprintf("%v", request.Params.Arguments)), &p); err != nil {
				return mcp.NewToolResultError("Failed to parse parameters"), err
			}

			args := []string{"scale", "nodegroup", "--cluster", p.ClusterName, "--name", p.Name}

			if p.Region != "" {
				args = append(args, "--region", p.Region)
			}
			if p.Nodes > 0 {
				args = append(args, "--nodes", strconv.Itoa(p.Nodes))
			}
			if p.MinNodes > 0 {
				args = append(args, "--nodes-min", strconv.Itoa(p.MinNodes))
			}
			if p.MaxNodes > 0 {
				args = append(args, "--nodes-max", strconv.Itoa(p.MaxNodes))
			}

			cmd := exec.Command("eksctl", args...)
			output, err := cmd.CombinedOutput()
			if err != nil {
				return mcp.NewToolResultError(fmt.Sprintf("Failed to scale nodegroup: %s\n%s", err, string(output))), nil
			}

			return mcp.NewToolResultText(string(output)), nil
		},
	)

	// Register nodegroup_update tool
	s.AddTool(
		mcp.NewTool("nodegroup_update", mcp.WithDescription("Update a nodegroup in an EKS cluster")),
		func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			var p NodeGroupUpdateParams
			if err := json.Unmarshal([]byte(fmt.Sprintf("%v", request.Params.Arguments)), &p); err != nil {
				return mcp.NewToolResultError("Failed to parse parameters"), err
			}

			args := []string{"update", "nodegroup", "--cluster", p.ClusterName, "--name", p.Name}

			if p.Region != "" {
				args = append(args, "--region", p.Region)
			}
			if p.MaxUnavailable > 0 {
				args = append(args, "--max-pods-unavailable", strconv.Itoa(p.MaxUnavailable))
			}

			// Handle labels
			for k, v := range p.Labels {
				args = append(args, "--node-labels", fmt.Sprintf("%s=%s", k, v))
			}

			// Handle tags
			for k, v := range p.Tags {
				args = append(args, "--tags", fmt.Sprintf("%s=%s", k, v))
			}

			cmd := exec.Command("eksctl", args...)
			output, err := cmd.CombinedOutput()
			if err != nil {
				return mcp.NewToolResultError(fmt.Sprintf("Failed to update nodegroup: %s\n%s", err, string(output))), nil
			}

			return mcp.NewToolResultText(string(output)), nil
		},
	)

	// Register nodegroup_list tool
	s.AddTool(
		mcp.NewTool("nodegroup_list", mcp.WithDescription("List all nodegroups in an EKS cluster")),
		func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			var p NodeGroupListParams
			if err := json.Unmarshal([]byte(fmt.Sprintf("%v", request.Params.Arguments)), &p); err != nil {
				return mcp.NewToolResultError("Failed to parse parameters"), err
			}

			args := []string{"get", "nodegroup", "--cluster", p.ClusterName}

			if p.Region != "" {
				args = append(args, "--region", p.Region)
			}
			if p.Output != "" {
				args = append(args, "--output", p.Output)
			}

			cmd := exec.Command("eksctl", args...)
			output, err := cmd.CombinedOutput()
			if err != nil {
				return mcp.NewToolResultError(fmt.Sprintf("Failed to list nodegroups: %s\n%s", err, string(output))), nil
			}

			return mcp.NewToolResultText(string(output)), nil
		},
	)
}
