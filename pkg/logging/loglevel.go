package logging

import (
	"fmt"
	"strings"
)

type LogLevel int

const (
	Debug LogLevel = iota
	Info
	Warn
	Error
)

var logLevelMap = map[string]LogLevel{
	"debug": Debug,
	"info":  Info,
	"warn":  Warn,
	"error": Error,
}

func (l *LogLevel) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var levelStr string
	if err := unmarshal(&levelStr); err != nil {
		return err
	}

	levelStr = strings.ToLower(levelStr)

	logLevel, exists := logLevelMap[levelStr]
	if !exists {
		return fmt.Errorf("invalid log level in config file")
	}

	*l = logLevel
	return nil
}
