package logger


import "github.com/sirupsen/logrus"

// Panic Level, highest level of severity. Panic logs and then calls panic with the
// message passed to Debug, Info, ...
func (p *TDLog) Panicf(format string, args ...interface{}) {
	p.log.Panicf(format, args...)
}

// Panic Level, highest level of severity. Panic logs and then calls panic with the
// message passed to Debug, Info, ...
func (p *TDLog) Panic(args ...interface{}) {
	p.panic(args...)
}

func (p *TDLog) panic(args ...interface{}) {
	p.log.Panic(args...)
}

// Fatal Level. Fatal logs and then calls `os.Exit(1)`. It will exit even if the
// logging level is set to Panic.
func (p *TDLog) Fatalf(format string, args ...interface{}) {
	p.log.Fatalf(format, args...)
}

// Fatal Level. Fatal logs and then calls `os.Exit(1)`. It will exit even if the
// logging level is set to Panic.
func (p *TDLog) Fatal(args ...interface{}) {
	p.fatal(args...)
}

func (p *TDLog) fatal(args ...interface{}) {
	p.log.Fatal(args...)
}

// Error Level. Error logs and is used for errors that should be definitely noted.
// Commonly used for hooks to send errors to an error tracking service.
func (p *TDLog) Errorf(format string, args ...interface{}) {
	p.log.Errorf(format, args...)
}

// Error Level. Error logs and is used for errors that should be definitely noted.
// Commonly used for hooks to send errors to an error tracking service.
func (p *TDLog) Error(args ...interface{}) {
	p.error(args...)
}

func (p *TDLog) error(args ...interface{}) {
	p.log.Error(args...)
}

// Warn Level. Non-critical entries that deserve eyes.
func (p *TDLog) Warnf(format string, args ...interface{}) {
	p.log.Warnf(format, args...)
}

// Warn Level. Non-critical entries that deserve eyes.
func (p *TDLog) Warn(args ...interface{}) {
	p.warn(args...)
}

func (p *TDLog) warn(args ...interface{}) {
	p.log.Warn(args...)
}

// Info Level. General operational entries about what's going on inside the
// application.
func (p *TDLog) Infof(format string, args ...interface{}) {
	p.log.Infof(format, args...)
}

// Info Level. General operational entries about what's going on inside the
// application.
func (p *TDLog) Info(args ...interface{}) {
	p.info(args...)
}

func (p *TDLog) info(args ...interface{}) {
	p.log.Info(args...)
}


// Debug Level. Usually only enabled when debugging. Very verbose logging.
func (p *TDLog) Debugf(format string, args ...interface{}) {
	p.log.Debugf(format, args...)
}

// Debug Level. Usually only enabled when debugging. Very verbose logging.
func (p *TDLog) Debug(args ...interface{}) {
	p.debug(args...)
}

func (p *TDLog) debug(args ...interface{}) {
	p.log.Debug(args...)
}

// SetLevel set the log level
func (p *TDLog) SetLevel(level logrus.Level) {
	p.log.SetLevel(level)
}

// GetLevel get the log level
func (p *TDLog) GetLevel() logrus.Level {
	return p.log.Level
}