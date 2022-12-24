package home

import (
	"bytes"
	"fmt"
	"os"
	"os/user"
	"path/filepath"
	"strings"

	"github.com/photowey/popctl/pkg/filez"
)

const (
	PopctlConfig = "popctl.json"
)

var (
	Home   = ".popctl"
	Usr, _ = user.Current()
	Dir    = filepath.Join(Usr.HomeDir, string(os.PathSeparator), Home)
)

func PopctlHome() {
	popctlHome := Dir
	if ok := filez.DirExists(popctlHome); !ok {
		if err := os.MkdirAll(popctlHome, os.ModePerm); err != nil {
			panic(fmt.Sprintf("mkdir popctl home dir:%s error:%v", popctlHome, err))
		}
	}

	if filez.FileNotExists(popctlHome, PopctlConfig) {
		buf := bytes.NewBufferString(popctlConfigContent)
		popctlConfigFile := filepath.Join(popctlHome, strings.ToLower(PopctlConfig))
		if err := os.WriteFile(popctlConfigFile, buf.Bytes(), 0o644); err != nil {
			panic(fmt.Sprintf("writing file %s: %v", popctlConfigFile, err))
		}
	}
}
