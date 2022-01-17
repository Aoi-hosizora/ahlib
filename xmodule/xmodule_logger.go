package xmodule

import (
	"fmt"
	"github.com/Aoi-hosizora/ahlib/xcolor"
)

// LogLevel represents ModuleContainer's logger level.
type LogLevel uint8

const (
	// LogName logs only when ModuleContainer.ProvideName invoked.
	LogName LogLevel = 1 << iota

	// LogType logs only when ModuleContainer.ProvideType invoked.
	LogType

	// LogImpl logs only when ModuleContainer.ProvideImpl invoked.
	LogImpl

	// LogInject logs only when ModuleContainer.Inject invoked.
	LogInject

	// LogAll logs when ModuleContainer.ProvideName, ModuleContainer.ProvideType, ModuleContainer.ProvideImpl, ModuleContainer.Inject invoked.
	LogAll = LogName | LogType | LogImpl | LogInject

	// LogSilent never logs, equals to disable the logger.
	LogSilent = LogLevel(0)
)

// Logger represents ModuleContainer's logger.
type Logger interface {
	// LogName invoked by ModuleContainer.ProvideName.
	LogName(moduleName, moduleTyp string)

	// LogType invoked by ModuleContainer.ProvideType.
	LogType(moduleTyp string)

	// LogImpl invoked by ModuleContainer.ProvideImpl.
	LogImpl(interfaceTyp, moduleTyp string)

	// LogInjectField invoked by ModuleContainer.Inject.
	LogInjectField(moduleName, structTyp, fieldName, fieldTyp string)

	// LogInject invoked by ModuleContainer.Inject.
	LogInject(structTyp string, num int)
}

// defaultLogger represents a default Logger.
type defaultLogger struct {
	level LogLevel
}

// DefaultLogger creates a default Logger instance. Log style see LogName, LogType, LogImpl, LogInject.
// Note that the red color represents the module and field name (~ represents no module name), and the yellow color represents the module and field type.
func DefaultLogger(level LogLevel) Logger {
	xcolor.ForceColor()
	return &defaultLogger{level: level}
}

// LogName logs like:
// 	[Xmodule] Prov: a <-- string
// 	               ---    ------
// 	               red    yellow
// Here `a` is the module name, `string` is the module type.
func (d *defaultLogger) LogName(moduleName, moduleTyp string) {
	if d.level&LogName != 0 {
		moduleName = xcolor.Red.Sprint(moduleName)
		moduleTyp = xcolor.Yellow.Sprint(moduleTyp)
		logLeftArrow("Prov:", moduleName, moduleTyp)
	}
}

// LogType logs like:
// 	[Xmodule] Prov: ~ <-- string
// 	               ---    ------
// 	               red    yellow
// Here `~` is the flag of no name, `string` is the module type.
func (d *defaultLogger) LogType(moduleTyp string) {
	if d.level&LogType != 0 {
		auto := xcolor.Red.Sprint("~")
		moduleTyp = xcolor.Yellow.Sprint(moduleTyp)
		logLeftArrow("Prov:", auto, moduleTyp)
	}
}

// LogImpl logs like:
// 	[Xmodule] Prov: ~ <-- IModule (*Module)
// 	               ---    -------  -------
// 	               red    yellow   yellow
// Here `~` is the flag of no name, `IModule` is the interface type, `*Module` is the module type.
func (d *defaultLogger) LogImpl(interfaceTyp, moduleTyp string) {
	if d.level&LogImpl != 0 {
		auto := xcolor.Red.Sprint("~")
		interfaceTyp = xcolor.Yellow.Sprint(interfaceTyp)
		moduleTyp = xcolor.Yellow.Sprint(moduleTyp)
		logLeftArrow("Prov:", auto, fmt.Sprintf("%s (%s)", interfaceTyp, moduleTyp))
	}
}

// LogInjectField logs like:
// 	[Xmodule] Inje: a --> (*Struct).Str string
// 	               ---     -------  --- ------
// 	               red     yellow   red yellow
// Here `a` is the module name, `*Struct` is the struct type, `Str` is the field name, `string` is the field type.
func (d *defaultLogger) LogInjectField(moduleName, structTyp, fieldName, fieldTyp string) {
	if d.level&LogInject != 0 {
		moduleName = xcolor.Red.Sprint(moduleName)
		structTyp = xcolor.Yellow.Sprint(structTyp)
		fieldName = xcolor.Red.Sprint(fieldName)
		fieldTyp = xcolor.Yellow.Sprint(fieldTyp)
		logRightArrow("Inje:", moduleName, fmt.Sprintf("(%s).%s %s", structTyp, fieldName, fieldTyp))
	}
}

// LogInject logs like:
// 	[Xmodule] Inje: ... --> (*Struct).(#3)
// 	                         -------
// 	                         yellow
// Here `*Struct` is the struct type, `#0` is the injected field count.
func (d *defaultLogger) LogInject(structTyp string, num int) {
	if d.level&LogInject != 0 {
		auto := xcolor.Default.Sprint("...")
		numStr := xcolor.Default.Sprintf("#%d", num)
		structTyp = xcolor.Yellow.Sprintf(structTyp)
		logRightArrow("Inje:", auto, fmt.Sprintf("(%s).(%s)", structTyp, numStr))
	}
}

// LogLeftArrowFunc is a logger function with left arrow (<--), used in LogName, LogType, LogImpl.
var LogLeftArrowFunc func(arg1, arg2, arg3 string)

// LogRightArrowFunc is a logger function with right arrow (-->), used in LogInject, LogInjectField.
var LogRightArrowFunc func(arg1, arg2, arg3 string)

// logLeftArrow represents the inner logger function with left arrow.
//
// The default format logs like:
// 	[Xmodule] Proj: ~                     <-- error (*errors.errorString)
// 	         |-----|---------------------|   |---------------------------|
// 	            5        30 (colored)                      ...
func logLeftArrow(arg1, arg2, arg3 string) {
	if LogLeftArrowFunc != nil {
		LogLeftArrowFunc(arg1, arg2, arg3)
		return
	}
	fmt.Printf("[Xmodule] %-5s %-30s <-- %s\n", arg1, arg2, arg3)
}

// logLeftArrow represents the inner logger function with right arrow.
//
// The default format logs like:
// 	[Xmodule] Inje: ~                     --> (*xmodule.testStruct).Err error
// 	         |-----|---------------------|   |-------------------------------|
// 	            5        30 (colored)                      ...
func logRightArrow(arg1, arg2, arg3 string) {
	if LogRightArrowFunc != nil {
		LogRightArrowFunc(arg1, arg2, arg3)
		return
	}
	fmt.Printf("[Xmodule] %-5s %-30s --> %s\n", arg1, arg2, arg3)
}
