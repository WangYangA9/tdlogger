# TDLogger

```go
func TestLog(t *testing.T) {
	// should read in viper
	LogConfiguration.IsDebug = true   //default true
	LogConfiguration.PrintLog = true //default true, Output to terminal
	LogConfiguration.MaxAge = 24 * time.Hour
	LogConfiguration.RotationTime = 1 * time.Hour
	LogConfiguration.DataDir = "./log"
	//GetLogger after set LogConfiguration
	var log = GetLogger("main")
	log.Debug("debug")
	log.Info("info")
	log.Error("error")
}
```
update tag