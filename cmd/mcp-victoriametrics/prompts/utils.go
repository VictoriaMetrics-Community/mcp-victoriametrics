package prompts

import (
	"fmt"

	"github.com/mark3labs/mcp-go/mcp"
)

func GetPromptReqParam(gpr mcp.GetPromptRequest, param string, required bool) (string, error) {
	value, ok := gpr.Params.Arguments[param]
	if !ok && required {
		return value, fmt.Errorf("%s param is required", param)
	}
	return value, nil
}
