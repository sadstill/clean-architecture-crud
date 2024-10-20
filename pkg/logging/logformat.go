package logging

import (
	"fmt"
	"strings"
)

type LogFormat int

const (
	Text LogFormat = iota
	Json
)

var logFormatMap = map[string]LogFormat{
	"text": Text,
	"json": Json,
}

func (f *LogFormat) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var formatStr string
	if err := unmarshal(&formatStr); err != nil {
		return err
	}

	formatStr = strings.ToLower(formatStr)

	logFormat, exists := logFormatMap[formatStr]
	if !exists {
		return fmt.Errorf("invalid log format in config file: %s", formatStr)
	}

	*f = logFormat
	return nil
}
