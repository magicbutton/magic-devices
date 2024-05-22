package magicapp

import (
	"os"
	"testing"

	"github.com/magicbutton/magic-devices/utils"
)

func TestMain(m *testing.M) {
	//cwd, _ := os.Getwd()

	//utils.MakeEnvFile(cwd)
	utils.Setup("./.env")
	code := m.Run()

	os.Exit(code)
}
