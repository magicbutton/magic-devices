// -------------------------------------------------------------------
// Generated by 365admin-publish
// -------------------------------------------------------------------
/*
---
title: Download Excel
---
*/
package cmds

import (
	"context"

	"github.com/magicbutton/magic-devices/execution"
	"github.com/magicbutton/magic-devices/utils"
)

func TransformDownloadExcelPost(ctx context.Context, args []string) (*string, error) {

	result, pwsherr := execution.ExecutePowerShell("john", "*", "magic-devices", "23-transform", "transform.ps1", "")
	if pwsherr != nil {
		return nil, pwsherr
	}
	utils.PrintSkip2FirstAnd2LastLines(string(result))
	return nil, nil

}
