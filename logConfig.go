/**
*  @file
*  @copyright defined in go-seele/LICENSE
 */

package logger

import (
	"os"
	"path/filepath"
	"time"
)

// LogConfiguration is the Configuration of log
var LogConfiguration = &LogConfig{
	PrintLog: true,
	IsDebug: true,
	DataDir: filepath.Join(os.TempDir(), "log"),
	MaxAge: 24*7*time.Hour,
	RotationTime: 1*time.Hour,
	LogExtension:".log",
}

// LogConfig is the Configuration of log
type LogConfig struct {
	// If IsDebug is true, the log level will be DebugLevel, otherwise it is InfoLevel
	IsDebug bool `json:"isDebug"`

	// If PrintLog is true, all logs will be printed in the console, otherwise they will be stored in the file.
	PrintLog bool `json:"printLog"`

	// DataDir default log directory in temp folder
	DataDir string `json:"-"`

	// Time to wait until old logs are purged.
	MaxAge time.Duration `json:"-"`

	// Interval between file rotation.
	RotationTime time.Duration `json:"-"`

	// The suffix name of the file requires the user to add the decimal point before the suffix name.
	LogExtension string `json:"-"`
}
