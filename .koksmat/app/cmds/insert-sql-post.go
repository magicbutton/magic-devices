// -------------------------------------------------------------------
// Generated by 365admin-publish
// -------------------------------------------------------------------
/*
---
title: Insert SQL
---
*/
package cmds

import (
	"context"

	"github.com/magicbutton/magic-devices/execution"
	"github.com/magicbutton/magic-devices/utils"
)

func InsertSqlPost(ctx context.Context, args []string) (*string, error) {

	result, pwsherr := execution.ExecutePowerShell("john", "*", "magic-devices", "22-insertsql", "10-insertsql.ps1", "")
	if pwsherr != nil {
		return nil, pwsherr
	}
	utils.PrintSkip2FirstAnd2LastLines(string(result))
	return nil, nil

}