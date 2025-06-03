package tools

import (
	"github.com/mark3labs/mcp-go/server"
	"github.com/weaveworks/eksctl/cmd/mcp/tools/associate"
	"github.com/weaveworks/eksctl/cmd/mcp/tools/completion"
	"github.com/weaveworks/eksctl/cmd/mcp/tools/create"
	"github.com/weaveworks/eksctl/cmd/mcp/tools/delete"
	"github.com/weaveworks/eksctl/cmd/mcp/tools/deregister"
	"github.com/weaveworks/eksctl/cmd/mcp/tools/disassociate"
	"github.com/weaveworks/eksctl/cmd/mcp/tools/drain"
	"github.com/weaveworks/eksctl/cmd/mcp/tools/enable"
	"github.com/weaveworks/eksctl/cmd/mcp/tools/get"
	"github.com/weaveworks/eksctl/cmd/mcp/tools/help"
	"github.com/weaveworks/eksctl/cmd/mcp/tools/info"
	"github.com/weaveworks/eksctl/cmd/mcp/tools/register"
	"github.com/weaveworks/eksctl/cmd/mcp/tools/scale"
	"github.com/weaveworks/eksctl/cmd/mcp/tools/set"
	"github.com/weaveworks/eksctl/cmd/mcp/tools/unset"
	"github.com/weaveworks/eksctl/cmd/mcp/tools/update"
	"github.com/weaveworks/eksctl/cmd/mcp/tools/upgrade"
	"github.com/weaveworks/eksctl/cmd/mcp/tools/utils"
	"github.com/weaveworks/eksctl/cmd/mcp/tools/version"
)

// RegisterTools registers all eksctl tools with the MCP server
func RegisterTools(s *server.MCPServer) {
	// Register associate commands
	associate.RegisterTools(s)

	// Register completion commands
	completion.RegisterTools(s)

	// Register create commands
	create.RegisterTools(s)

	// Register delete commands
	delete.RegisterTools(s)

	// Register deregister commands
	deregister.RegisterTools(s)

	// Register disassociate commands
	disassociate.RegisterTools(s)

	// Register drain commands
	drain.RegisterTools(s)

	// Register enable commands
	enable.RegisterTools(s)

	// Register get commands
	get.RegisterTools(s)

	// Register help command
	help.RegisterTools(s)

	// Register info command
	info.RegisterTools(s)

	// Register register commands
	register.RegisterTools(s)

	// Register scale commands
	scale.RegisterTools(s)

	// Register set commands
	set.RegisterTools(s)

	// Register unset commands
	unset.RegisterTools(s)

	// Register update commands
	update.RegisterTools(s)

	// Register upgrade commands
	upgrade.RegisterTools(s)

	// Register utils commands
	utils.RegisterTools(s)

	// Register version command
	version.RegisterTools(s)
}
