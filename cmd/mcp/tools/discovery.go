package tools

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
	"sync"
	"time"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

// ParameterInfo stores information about a command parameter
type ParameterInfo struct {
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Type        string   `json:"type"` // "string", "boolean", etc.
	Required    bool     `json:"required"`
	Choices     []string `json:"choices,omitempty"` // Possible values for enum-like parameters
}

// CommandInfo stores information about an eksctl command
type CommandInfo struct {
	Command     string          `json:"command"`
	Description string          `json:"description"`
	Parameters  []ParameterInfo `json:"parameters"`
	Timestamp   time.Time       `json:"timestamp"` // When this info was last updated
}

// CommandRegistry stores information about all discovered commands
type CommandRegistry struct {
	Commands     map[string]CommandInfo `json:"commands"`
	CompletionTS time.Time              `json:"completion_timestamp"` // When completion script was last fetched
	mutex        sync.RWMutex           // For thread safety
}

var (
	// Global registry instance
	registry     *CommandRegistry
	registryOnce sync.Once
	// Cache file path
	cacheFile = filepath.Join(os.TempDir(), "eksctl-mcp-command-cache.json")
	// Cache expiration (24 hours)
	cacheExpiration = 24 * time.Hour
)

// getRegistry returns the singleton registry instance
func getRegistry() *CommandRegistry {
	registryOnce.Do(func() {
		registry = &CommandRegistry{
			Commands: make(map[string]CommandInfo),
		}
		// Try to load from cache
		registry.loadFromCache()
	})
	return registry
}

// saveToCache persists the registry to disk
func (r *CommandRegistry) saveToCache() error {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	data, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(cacheFile, data, 0644)
}

// loadFromCache loads the registry from disk if available and not expired
func (r *CommandRegistry) loadFromCache() error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	data, err := os.ReadFile(cacheFile)
	if err != nil {
		// It's okay if the cache doesn't exist yet
		return nil
	}

	err = json.Unmarshal(data, r)
	if err != nil {
		// If cache is corrupted, start fresh
		r.Commands = make(map[string]CommandInfo)
		return nil
	}

	// Check if cache is expired
	if time.Since(r.CompletionTS) > cacheExpiration {
		// Cache expired, clear it
		r.Commands = make(map[string]CommandInfo)
	}

	return nil
}

// getCompletionScript fetches and parses the bash completion script
func getCompletionScript() (string, error) {
	reg := getRegistry()
	reg.mutex.RLock()

	// Check if we have a recent completion script
	if time.Since(reg.CompletionTS) < cacheExpiration {
		reg.mutex.RUnlock()
		// Run command to get completion script
		completionCmd := exec.Command("eksctl", "completion", "bash")
		var completionOut bytes.Buffer
		completionCmd.Stdout = &completionOut

		if err := completionCmd.Run(); err != nil {
			return "", fmt.Errorf("failed to get completion script: %w", err)
		}

		script := completionOut.String()

		// Update timestamp and save
		reg.mutex.Lock()
		reg.CompletionTS = time.Now()
		reg.mutex.Unlock()
		go reg.saveToCache() // Save asynchronously

		return script, nil
	}

	// We have a cached completion script
	reg.mutex.RUnlock()
	return "", nil // No need to fetch again
}

// extractCommandsFromCompletion parses bash completion script to find commands
func extractCommandsFromCompletion(completionScript string) []string {
	var commands []string

	// Look for command patterns in the completion script
	cmdPattern := regexp.MustCompile(`(?m)commands=\("([^"]+)"\)`)
	matches := cmdPattern.FindAllStringSubmatch(completionScript, -1)

	for _, match := range matches {
		if len(match) >= 2 {
			cmds := strings.Split(match[1], " ")
			commands = append(commands, cmds...)
		}
	}

	// Also look for subcommands
	subcmdPattern := regexp.MustCompile(`(?m)subcmds=\("([^"]+)"\)`)
	submatches := subcmdPattern.FindAllStringSubmatch(completionScript, -1)

	for _, match := range submatches {
		if len(match) >= 2 {
			cmds := strings.Split(match[1], " ")
			commands = append(commands, cmds...)
		}
	}

	return commands
}

// extractFlagsFromCompletion parses bash completion script to find flags and their choices
func extractFlagsFromCompletion(completionScript string) (map[string][]string, map[string][]string) {
	// Map to store command -> flags
	commandFlags := make(map[string][]string)

	// Map to store flag -> choices
	flagChoices := make(map[string][]string)

	// Look for flag patterns in the completion script
	// This regex looks for lines that define flags for specific commands
	cmdFlagPattern := regexp.MustCompile(`(?m)case \$prev in\s+(--[a-zA-Z0-9-]+)\)\s+case \$\{words\[0\]\} in\s+(.*?)\s+;;`)
	matches := cmdFlagPattern.FindAllStringSubmatch(completionScript, -1)

	for _, match := range matches {
		if len(match) >= 3 {
			flag := match[1]
			commands := strings.Split(match[2], "|")

			for _, cmd := range commands {
				cmd = strings.TrimSpace(cmd)
				if cmd != "" {
					if _, exists := commandFlags[cmd]; !exists {
						commandFlags[cmd] = []string{}
					}
					commandFlags[cmd] = append(commandFlags[cmd], flag)
				}
			}
		}
	}

	// Also look for flag choices (enum values)
	flagChoicesPattern := regexp.MustCompile(`(?m)case \$\{words\[c\]\} in\s+(--[a-zA-Z0-9-]+)\)\s+COMPREPLY=\(\$\(compgen -W "(.*?)" -- \$cur\)\)`)
	choicesMatches := flagChoicesPattern.FindAllStringSubmatch(completionScript, -1)

	for _, match := range choicesMatches {
		if len(match) >= 3 {
			flag := match[1]
			choices := strings.Split(match[2], " ")
			flagChoices[flag] = choices
		}
	}

	// Also look for boolean flags (flags that don't expect values)
	boolFlagPattern := regexp.MustCompile(`(?m)flags=\("([^"]+)"\)`)
	boolMatches := boolFlagPattern.FindAllStringSubmatch(completionScript, -1)

	for _, match := range boolMatches {
		if len(match) >= 2 {
			flags := strings.Split(match[1], " ")
			for _, flag := range flags {
				if strings.HasPrefix(flag, "--") && !strings.Contains(flag, "=") {
					// This is likely a boolean flag
					for cmd := range commandFlags {
						commandFlags[cmd] = append(commandFlags[cmd], flag)
					}
				}
			}
		}
	}

	return commandFlags, flagChoices
}

// extractCommandDescriptionFromHelp gets the command description from help output
func extractCommandDescriptionFromHelp(helpOutput string) string {
	// Try to find a description line
	descPattern := regexp.MustCompile(`(?m)^(.+)$`)
	matches := descPattern.FindAllStringSubmatch(helpOutput, 3) // Look at first few lines

	for _, match := range matches {
		if len(match) >= 2 {
			line := strings.TrimSpace(match[1])
			// Skip usage lines and empty lines
			if !strings.HasPrefix(line, "Usage:") && line != "" {
				return line
			}
		}
	}

	return ""
}

// DiscoverCommandParameters extracts parameter information from eksctl help and completion
// with an optional recursion depth limit
func DiscoverCommandParameters(command string, depth ...int) (CommandInfo, error) {
	// Default depth is 0, meaning no recursion limit
	currentDepth := 0
	maxDepth := 2 // Default max depth of 2 (eksctl -> create -> cluster)

	// If depth is provided, use it
	if len(depth) > 0 {
		currentDepth = depth[0]
	}
	if len(depth) > 1 {
		maxDepth = depth[1]
	}
	reg := getRegistry()

	// Check if we already have cached info for this command
	reg.mutex.RLock()
	if info, exists := reg.Commands[command]; exists && time.Since(info.Timestamp) < cacheExpiration {
		reg.mutex.RUnlock()
		return info, nil
	}
	reg.mutex.RUnlock()

	// Get bash completion script
	completionScript, err := getCompletionScript()
	if err != nil {
		return CommandInfo{}, err
	}

	commandFlags, flagChoices := extractFlagsFromCompletion(completionScript)

	// Run eksctl help for the command
	helpCmd := exec.Command("eksctl", append(strings.Split(command, " "), "--help")...)
	var helpOut, helpErr bytes.Buffer
	helpCmd.Stdout = &helpOut
	helpCmd.Stderr = &helpErr

	if err := helpCmd.Run(); err != nil {
		return CommandInfo{}, fmt.Errorf("failed to get help for command %s: %w", command, err)
	}

	helpOutput := helpOut.String()

	// Extract command description
	description := extractCommandDescriptionFromHelp(helpOutput)

	// Parse the help output to extract parameters
	info := CommandInfo{
		Command:     command,
		Description: description,
		Timestamp:   time.Now(),
	}

	// Regular expression to match parameter descriptions in help output
	flagPattern := regexp.MustCompile(`(?m)^\s+--([a-zA-Z0-9-]+)\s+(.+?)(?:\s+\(required\))?$`)
	matches := flagPattern.FindAllStringSubmatch(helpOutput, -1)

	// Create a map to store parameter info
	paramMap := make(map[string]ParameterInfo)

	// Process flags from help output
	for _, match := range matches {
		if len(match) >= 3 {
			name := match[1]
			description := strings.TrimSpace(match[2])
			required := strings.Contains(match[0], "(required)")

			// Determine parameter type based on name or description
			paramType := "string"
			if strings.Contains(name, "enable") ||
				strings.Contains(name, "disable") ||
				strings.Contains(name, "force") ||
				strings.Contains(name, "approve") ||
				strings.Contains(description, "Toggle") ||
				strings.Contains(description, "Enable") ||
				strings.Contains(description, "Disable") {
				paramType = "boolean"
			}

			paramMap[name] = ParameterInfo{
				Name:        name,
				Description: description,
				Type:        paramType,
				Required:    required,
			}
		}
	}

	// Enhance with information from completion script
	// This can help identify boolean flags and enum values
	cmdParts := strings.Split(command, " ")
	cmdName := cmdParts[len(cmdParts)-1]

	if flags, exists := commandFlags[cmdName]; exists {
		for _, flag := range flags {
			// Strip leading "--"
			flagName := strings.TrimPrefix(flag, "--")

			// If we already have this parameter from help, enhance it
			if param, ok := paramMap[flagName]; ok {
				// Check if this flag has predefined choices
				if choices, hasChoices := flagChoices[flag]; hasChoices {
					param.Choices = choices
					paramMap[flagName] = param
				}
			} else {
				// This flag wasn't in the help output, add it
				// Assume it's a boolean flag if we don't know otherwise
				paramMap[flagName] = ParameterInfo{
					Name:        flagName,
					Description: "Flag for " + flagName,
					Type:        "boolean",
					Required:    false,
				}
			}
		}
	}

	// Look for subcommands in the help output
	subcommandPattern := regexp.MustCompile(`(?m)^\s+([a-zA-Z0-9-]+)\s+(.+?)$`)
	subcommandMatches := subcommandPattern.FindAllStringSubmatch(helpOutput, -1)

	// Process subcommands and recursively get their parameters
	for _, match := range subcommandMatches {
		if len(match) >= 3 {
			subcommandName := match[1]
			// Skip if it's a flag (starts with -) or common help text
			if strings.HasPrefix(subcommandName, "-") ||
				subcommandName == "help" ||
				subcommandName == "version" ||
				subcommandName == "flags:" ||
				subcommandName == "commands:" ||
				subcommandName == "args:" {
				continue
			}

			// Form the full subcommand
			subcommand := command + " " + subcommandName

			// Only recurse if we haven't reached the maximum depth
			if currentDepth < maxDepth {
				// Recursively get parameters for this subcommand with incremented depth
				subcommandInfo, err := DiscoverCommandParameters(subcommand, currentDepth+1, maxDepth)
				if err == nil {
					// Enhance the description with subcommand information
					if info.Description == "" {
						info.Description = "Parent command for various subcommands"
					}

					// Add subcommand parameters to the parent command's description
					info.Description += fmt.Sprintf("\n\nSubcommand '%s': %s",
						subcommandName, subcommandInfo.Description)

					// Optionally, we could also merge parameters, but that might be confusing
					// Instead, we're just enhancing the description
				}
			}
		}
	}

	// Convert map to slice
	for _, param := range paramMap {
		info.Parameters = append(info.Parameters, param)
	}

	// Sort parameters for consistent output
	sort.Slice(info.Parameters, func(i, j int) bool {
		return info.Parameters[i].Name < info.Parameters[j].Name
	})

	// Cache the result
	reg.mutex.Lock()
	reg.Commands[command] = info
	reg.mutex.Unlock()

	// Save cache asynchronously
	go reg.saveToCache()

	return info, nil
}

// RegisterDynamicTool registers a tool with auto-discovered parameters
func RegisterDynamicTool(s *server.MCPServer, toolName, description, command string) error {
	// Discover parameters
	info, err := DiscoverCommandParameters(command, 0, 2)
	if err != nil {
		return err
	}

	// If no description was provided, use the one from the command
	if description == "" {
		description = info.Description
	}

	// Create the tool with basic description
	toolOptions := []mcp.ToolOption{mcp.WithDescription(description)}

	// Add parameters based on their type
	for _, param := range info.Parameters {
		if param.Type == "boolean" {
			toolOptions = append(toolOptions, mcp.WithBoolean(param.Name, mcp.Description(param.Description)))
		} else {
			// Default to string type
			stringOpts := []mcp.PropertyOption{mcp.Description(param.Description)}
			if param.Required {
				stringOpts = append(stringOpts, mcp.Required())
			}
			if len(param.Choices) > 0 {
				stringOpts = append(stringOpts, mcp.Enum(param.Choices...))
			}
			toolOptions = append(toolOptions, mcp.WithString(param.Name, stringOpts...))
		}
	}

	// Create the tool with all options
	tool := mcp.NewTool(toolName, toolOptions...)

	// Register the tool
	s.AddTool(tool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		return ExecuteEksctlCommandFromRequest(ctx, command, request)
	})

	return nil
}

// ListAvailableCommands returns a list of all available eksctl commands
func ListAvailableCommands() ([]string, error) {
	// Get bash completion script
	completionScript, err := getCompletionScript()
	if err != nil {
		return nil, err
	}

	return extractCommandsFromCompletion(completionScript), nil
}

// RefreshCommandCache forces a refresh of the command cache
func RefreshCommandCache() error {
	reg := getRegistry()

	// Clear existing cache
	reg.mutex.Lock()
	reg.Commands = make(map[string]CommandInfo)
	reg.CompletionTS = time.Time{}
	reg.mutex.Unlock()

	// Get fresh completion script
	_, err := getCompletionScript()
	if err != nil {
		return err
	}

	// Get list of commands
	commands, err := ListAvailableCommands()
	if err != nil {
		return err
	}

	// Discover parameters for each command
	for _, cmd := range commands {
		_, err := DiscoverCommandParameters(cmd, 0, 2)
		if err != nil {
			// Log error but continue with other commands
			fmt.Printf("Error discovering parameters for %s: %v\n", cmd, err)
		}
	}

	return reg.saveToCache()
}

// GetCachedCommands returns all commands from the cache
func GetCachedCommands() []CommandInfo {
	reg := getRegistry()
	reg.mutex.RLock()
	defer reg.mutex.RUnlock()

	var commands []CommandInfo
	for _, info := range reg.Commands {
		commands = append(commands, info)
	}

	return commands
}
