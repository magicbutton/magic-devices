package main

import (
	"runtime/debug"
	"strings"

	"github.com/magicbutton/magic-devices/magicapp"
	"github.com/magicbutton/magic-devices/utils"
)

func main() {
	info, _ := debug.ReadBuildInfo()

	// split info.Main.Path by / and get the last element
	s1 := strings.Split(info.Main.Path, "/")
	name := s1[len(s1)-1]
	description := `---
title: magic-devices
description: Describe the main purpose of this kitchen
---

# magic-devices
`
	utils.Setup(".env")
	magicapp.RegisterServeCmd("magic-devices", description, "0.0.1", 8080)
	magicapp.RegisterCmds()
	magicapp.RegisterServiceCmd()
	magicapp.RegisterSQLCmd()
	magicapp.RegisterImportCmd()
	utils.RootCmd.PersistentFlags().BoolVarP(&utils.Verbose, "verbose", "v", false, "verbose output")

	magicapp.Execute(name, "magic-devices", "")
}
