package mcp

// This file contains type definitions and interfaces for the MCP server implementation
// It can be expanded as needed to support additional functionality

// ToolRegistrationError represents an error during tool registration
type ToolRegistrationError struct {
	CommandPath string
	Err         error
}

// Error implements the error interface
func (e *ToolRegistrationError) Error() string {
	return "failed to register tool for command " + e.CommandPath + ": " + e.Err.Error()
}

// Unwrap returns the underlying error
func (e *ToolRegistrationError) Unwrap() error {
	return e.Err
}
