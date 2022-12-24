package app

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/photowey/popctl/configs"
	"github.com/photowey/popctl/internal/home"
	"github.com/photowey/popctl/pkg/filez"
)

func onInit() {
	home.PopctlHome()
	configLoad()
}

func configLoad() {
	popctlHome := home.Dir
	popctlConfigFile := filepath.Join(popctlHome, strings.ToLower(home.PopctlConfig))
	if filez.FileExists(popctlHome, home.PopctlConfig) {
		configs.Init(popctlConfigFile)
	} else {
		fmt.Printf("the popctl config file not exists")
	}
}
