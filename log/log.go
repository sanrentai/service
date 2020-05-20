package log

import (
	"errors"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/gogf/gf/os/glog"
	"github.com/kardianos/service"
)

var logger *MyLog

type MyLog struct {
	Slogger service.Logger
}

func New(slogger service.Logger) *MyLog {
	s, _ := GetCurrentPath()
	glog.SetPath(s + "/log")
	logger = &MyLog{slogger}
	return logger
}

func GetCurrentPath() (string, error) {
	file, err := exec.LookPath(os.Args[0])
	if err != nil {
		return "", err
	}
	path, err := filepath.Abs(file)
	if err != nil {
		return "", err
	}
	i := strings.LastIndex(path, "/")
	if i < 0 {
		i = strings.LastIndex(path, "\\")
	}
	if i < 0 {
		return "", errors.New(`error: Can't find "/" or "\".`)
	}
	return string(path[0 : i+1]), nil
}

func SetStdoutPrint(b bool) {
	glog.SetStdoutPrint(b)
}

func (log *MyLog) Error(v ...interface{}) error {

	log.Slogger.Error(v...)

	glog.Error(v...)

	return nil
}
func (log *MyLog) Warning(v ...interface{}) error {

	log.Slogger.Warning(v...)

	glog.Warning(v...)

	return nil
}
func (log *MyLog) Info(v ...interface{}) error {

	log.Slogger.Info(v...)

	glog.Info(v...)

	return nil
}

func (log *MyLog) Errorf(format string, a ...interface{}) error {

	log.Slogger.Errorf(format, a...)

	glog.Errorf(format, a...)

	return nil
}
func (log *MyLog) Warningf(format string, a ...interface{}) error {

	log.Slogger.Warningf(format, a...)

	glog.Warningf(format, a...)

	return nil
}
func (log *MyLog) Infof(format string, a ...interface{}) error {

	log.Slogger.Infof(format, a...)

	glog.Infof(format, a...)

	return nil
}

func Print(v ...interface{}) {
	glog.Print(v...)
}

// Printf prints <v> with format <format> using fmt.Sprintf.
// The parameter <v> can be multiple variables.
func Printf(format string, v ...interface{}) {
	glog.Printf(format, v...)
}

// See Print.
func Println(v ...interface{}) {
	glog.Println(v...)
}

// Fatal prints the logging content with [FATA] header and newline, then exit the current process.
func Fatal(v ...interface{}) {
	glog.Fatal(v...)
}

// Fatalf prints the logging content with [FATA] header, custom format and newline, then exit the current process.
func Fatalf(format string, v ...interface{}) {
	glog.Fatalf(format, v...)
}

// Panic prints the logging content with [PANI] header and newline, then panics.
func Panic(v ...interface{}) {
	glog.Panic(v...)
}

// Panicf prints the logging content with [PANI] header, custom format and newline, then panics.
func Panicf(format string, v ...interface{}) {
	glog.Panicf(format, v...)
}

// Info prints the logging content with [INFO] header and newline.
func Info(v ...interface{}) {
	logger.Info(v...)
}

// Infof prints the logging content with [INFO] header, custom format and newline.
func Infof(format string, v ...interface{}) {
	logger.Infof(format, v...)
}

// Debug prints the logging content with [DEBU] header and newline.
func Debug(v ...interface{}) {
	glog.Debug(v...)
}

// Debugf prints the logging content with [DEBU] header, custom format and newline.
func Debugf(format string, v ...interface{}) {
	glog.Debugf(format, v...)
}

// Notice prints the logging content with [NOTI] header and newline.
// It also prints caller stack info if stack feature is enabled.
func Notice(v ...interface{}) {
	glog.Notice(v...)
}

// Noticef prints the logging content with [NOTI] header, custom format and newline.
// It also prints caller stack info if stack feature is enabled.
func Noticef(format string, v ...interface{}) {
	glog.Noticef(format, v...)
}

// Warning prints the logging content with [WARN] header and newline.
// It also prints caller stack info if stack feature is enabled.
func Warning(v ...interface{}) {
	logger.Warning(v...)
}

// Warningf prints the logging content with [WARN] header, custom format and newline.
// It also prints caller stack info if stack feature is enabled.
func Warningf(format string, v ...interface{}) {
	logger.Warningf(format, v...)
}

// Error prints the logging content with [ERRO] header and newline.
// It also prints caller stack info if stack feature is enabled.
func Error(v ...interface{}) {
	logger.Error(v...)
}

// Errorf prints the logging content with [ERRO] header, custom format and newline.
// It also prints caller stack info if stack feature is enabled.
func Errorf(format string, v ...interface{}) {
	logger.Errorf(format, v...)
}

// Critical prints the logging content with [CRIT] header and newline.
// It also prints caller stack info if stack feature is enabled.
func Critical(v ...interface{}) {
	glog.Critical(v...)
}

// Criticalf prints the logging content with [CRIT] header, custom format and newline.
// It also prints caller stack info if stack feature is enabled.
func Criticalf(format string, v ...interface{}) {
	glog.Criticalf(format, v...)
}
