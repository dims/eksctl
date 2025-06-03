// Package common provides shared utilities and standardized patterns for MCP tools
package common

import (
	"context"
	"fmt"
	"github.com/mark3labs/mcp-go/mcp"
)

// StandardParameters provides commonly used parameters for eksctl commands
type StandardParameters struct {
	Name      string
	Region    string
	Cluster   string
	Nodegroup string
}

// GetStandardParameters extracts common parameters from a request
func GetStandardParameters(request mcp.CallToolRequest) StandardParameters {
	return StandardParameters{
		Name:      request.GetString("name", ""),
		Region:    request.GetString("region", ""),
		Cluster:   request.GetString("cluster", ""),
		Nodegroup: request.GetString("nodegroup", ""),
	}
}

// ValidateRequiredParams checks if required parameters are provided and returns an error if not
func ValidateRequiredParams(params map[string]string) error {
	var missingParams []string
	for name, value := range params {
		if value == "" {
			missingParams = append(missingParams, name)
		}
	}

	if len(missingParams) > 0 {
		return fmt.Errorf("missing required parameters: %v", missingParams)
	}

	return nil
}

// CreateStubResponse creates a standardized response for stub implementations
func CreateStubResponse(ctx context.Context, command string, params map[string]string) (*mcp.CallToolResult, error) {
	if err := ValidateRequiredParams(params); err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	// Build a message that includes all parameters
	message := fmt.Sprintf("Command: %s\n", command)
	message += "Parameters:\n"
	for name, value := range params {
		message += fmt.Sprintf("  %s: %s\n", name, value)
	}
	message += "\nThis is a stub implementation. The actual command would perform the operation on AWS resources."

	return mcp.NewToolResultText(message), nil
}
