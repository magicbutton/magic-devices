// -------------------------------------------------------------------
// Generated by 365admin-publish
// -------------------------------------------------------------------
/*
---
title: Import Sheet FullKpi
---
*/
package cmds

import (
	"context"

	"github.com/magicbutton/magic-devices/execution"
	"github.com/magicbutton/magic-devices/utils"
)

func ImportFullkpiPost(ctx context.Context, args []string) (*string, error) {

	result, pwsherr := execution.ExecutePowerShell("john", "*", "magic-devices", "20-import", "fullkpi.ps1", "")
	if pwsherr != nil {
		return nil, pwsherr
	}
	utils.PrintSkip2FirstAnd2LastLines(string(result))
	return nil, nil

}