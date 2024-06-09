package config

import (
	"fmt"
	"io"
	"os"
	"path"
	"runtime"
	"strings"

	formatter "github.com/antonfisher/nested-logrus-formatter"
	"github.com/sirupsen/logrus"
)

func NewLogger() *logrus.Logger {
	f, _ := os.OpenFile("log.txt", os.O_CREATE|os.O_WRONLY, 0777)
	log := logrus.New()
	log.SetLevel(logrus.DebugLevel)
	log.SetFormatter(&formatter.Formatter{
		NoColors:        false,
		TimestampFormat: "02 Jan 06 - 15:04",
		HideKeys:        false,
		CallerFirst:     true,
		CustomCallerFormatter: func(f *runtime.Frame) string {
			s := strings.Split(f.Function, ".")
			funcName := s[len(s)-1]
			return fmt.Sprintf(" \x1b[%dm[%s:%d][%s()]", 34, path.Base(f.File), f.Line, funcName)
		},
	})
	log.Out = io.MultiWriter(os.Stderr, f)
	log.SetReportCaller(true)

	return log
}
