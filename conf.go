package hlog

import (
	"errors"
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

const (
	HLOG_ACTION_FILE   = "file"
	HLOG_ACTION_STDERR = "stderr"
	HLOG_ACTION_STDOUT = "stdout"
	HLOG_ACTION_SYSLOG = "syslog"
)

type FileConfig struct {
	Path      string `yaml:"path"`
	MaxSize   string `yaml:"maxsize"`
	SyncSize  string `yaml:"syncsize"`
	MaxNumber int    `yaml:"number"`
	SyncTimes int    `yaml:"synctimes"`
}

type DefaultConfig struct {
	Format string     `yaml:"format"`
	File   FileConfig `yaml:"file"`
}

type FormatConfig struct {
	Name   string `yaml:"name"`
	Format string `yaml:"format"`
}

type LevelConfig struct {
	Name   string `yaml:"name"`
	Level  int    `yaml:"level"`
	Syslog string `yaml:"syslog"`
}

type RuleConfig struct {
	Type   string     `yaml:"type"`
	Level  string     `yaml:"level"`
	Action string     `yaml:"action"`
	Syslog string     `yaml:"syslog"`
	File   FileConfig `yaml:"file"`
}

type HLogConfig struct {
	Default DefaultConfig  `yaml:"default"`
	Levels  []LevelConfig  `yaml:"levels"`
	Formats []FormatConfig `yaml:"formats"`
	Rules   []RuleConfig   `yaml:"rules"`
}

var g_hlogConfig *HLogConfig

func loadConfig(filename string) error {

	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}

	config := new(HLogConfig)
	config.Levels = make([]LevelConfig, 0)
	config.Formats = make([]FormatConfig, 0)
	config.Rules = make([]RuleConfig, 0)

	err = yaml.Unmarshal(body, config)
	if err != nil {
		return err
	}

	if !config.verify() {
		return errors.New("This config is verity failed!")
	}

	g_hlogConfig = config

	return nil
}

func (c *HLogConfig) verify() bool {

	return true
}

func init() {
	err := loadConfig("./default.yaml")
	if err != nil {
		log.Fatal(err.Error())
	}
}
