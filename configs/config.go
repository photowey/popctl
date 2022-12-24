package configs

import (
	"fmt"
	"os"

	"github.com/photowey/popctl/pkg/jsonz"
)

var _config Config

type Config struct {
	Project  Project
	Database Database
}

type Project struct {
	Author  string  `toml:"author" json:"author" yaml:"author"`
	Email   string  `toml:"email" json:"email"`
	Version string  `toml:"version" json:"version" yaml:"version"`
	Dto     Dto     `toml:"dto" json:"dto" yaml:"dto"`
	Entity  Entity  `toml:"entity" json:"entity" yaml:"entity"`
	Payload Payload `toml:"payload" json:"payload" yaml:"payload"`
	Query   Query   `toml:"query" json:"query" yaml:"query"`
}

type Database struct {
	Host     string   `toml:"host" json:"host" yaml:"host"`
	Port     int      `toml:"port" json:"port" yaml:"port"`
	Dialect  string   `toml:"dialect" json:"dialect" yaml:"dialect"`
	Driver   string   `toml:"driver" json:"driver" yaml:"driver"`
	Database string   `toml:"database" json:"database" yaml:"database"`
	Username string   `toml:"username" json:"username" yaml:"username"`
	Password string   `toml:"password" json:"password" yaml:"password"`
	Includes []string `toml:"includes" json:"includes" yaml:"includes"`
	Excludes []string `toml:"excludes" json:"excludes" yaml:"excludes"`
	Prefixes []string `toml:"prefixes" json:"prefixes" yaml:"prefixes"`
}

type Dto struct {
	// Dto model excludes fields
	Excludes []string `toml:"excludes" json:"excludes" yaml:"excludes"`
}

type Entity struct {
	// Dto model excludes fields
	Excludes []string `toml:"excludes" json:"excludes" yaml:"excludes"`
}

type Payload struct {
	// Payload model excludes fields
	Add    Condition `toml:"add" json:"add" yaml:"add"`
	Delete Condition `toml:"delete" json:"delete" yaml:"delete"`
	Update Condition `toml:"update" json:"update" yaml:"update"`
}

type Query struct {
	// Query model excludes fields
	Excludes []string `toml:"excludes" json:"excludes" yaml:"excludes"`
}

type Condition struct {
	// Payload model excludes fields
	Includes []string `toml:"includes" json:"includes" yaml:"includes"`
	Excludes []string `toml:"excludes" json:"excludes" yaml:"excludes"`
}

func Init(popctlConfigFile string) {
	conf, err := os.ReadFile(popctlConfigFile)
	if err != nil {
		fmt.Printf("parse the popctl config error:%s", err.Error())
		return
	}
	jsonz.UnmarshalStruct(conf, &_config)
}

func ProjectFunc() Project {
	return _config.Project
}

func DatabaseFunc() Database {
	return _config.Database
}
