package mocks

// MockLogger объект логгера, который необходим для моков в тестах.
type MockLogger struct{}

func (m MockLogger) Debug(args ...interface{}) {
	return
}

func (m MockLogger) Info(args ...interface{}) {
	return
}

func (m MockLogger) Warn(args ...interface{}) {
	return
}

func (m MockLogger) Error(args ...interface{}) {
	return
}

func (m MockLogger) Debugf(template string, args ...interface{}) {
	return
}

func (m MockLogger) Infof(template string, args ...interface{}) {
	return
}

func (m MockLogger) Warnf(template string, args ...interface{}) {
	return
}

func (m MockLogger) Errorf(template string, args ...interface{}) {
	return
}

func (m MockLogger) Debugw(msg string, keysAndValues ...interface{}) {
	return
}

func (m MockLogger) Infow(msg string, keysAndValues ...interface{}) {
	return
}

func (m MockLogger) Warnw(msg string, keysAndValues ...interface{}) {
	return
}

func (m MockLogger) Errorw(msg string, keysAndValues ...interface{}) {
	return
}

func (m MockLogger) Print(args ...interface{}) {
	return
}

func (m MockLogger) Printf(format string, args ...interface{}) {
	return
}
