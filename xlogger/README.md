# xlogger

## Functions

### Logrus Formatter

+ `type CustomFormatter struct {}`
+ `(f *CustomFormatter) Format(entry *logrus.Entry) ([]byte, error)`

### Logrus Rotate Hook

+ `type RotateFileConfig struct {}`
+ `type RotateFileHook struct {}`
+ `NewRotateFileHook(config *RotateFileConfig) logrus.Hook`
+ `type RotateLogConfig struct {}`
+ `type RotateLogHook struct {}`
+ `NewRotateLogHook(config *RotateLogConfig) logrus.Hook`

### StdLogger

+ `type StdLogger struct {}`
+ `NewStdLogger(out io.Writer) *StdLogger`
+ `(l *StdLogger) Writer() io.Writer`
+ `(l *StdLogger) Output(s string)`
+ `(l *StdLogger) Outputf(format string, v ...interface{})`
+ `(l *StdLogger) Outputln(format string)`
+ `Writer() io.Writer`
+ `Output(s string)`
+ `Outputf(format string, v ...interface{})`
+ `Outputln(format string)`
