package service

func Error(v ...interface{}) error {
	return logger.Error(v...)
}

func Warning(v ...interface{}) error {
	return logger.Warning(v...)
}

func Info(v ...interface{}) error {
	return logger.Info(v...)
}

func Errorf(format string, a ...interface{}) error {
	return logger.Errorf(format, a...)
}

func Warningf(format string, a ...interface{}) error {
	return logger.Warningf(format, a...)
}

func Infof(format string, a ...interface{}) error {
	return logger.Infof(format, a...)
}
