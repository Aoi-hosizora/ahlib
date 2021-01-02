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

	// LogSilent logs never, equals to disable the logger.
	LogSilent = LogLevel(0)
)

// Logger represents ModuleContainer's logger.
type Logger interface {
	// LogName invoked by ModuleContainer.ProvideName.
	LogName(name, typ string)

	// LogType invoked by ModuleContainer.ProvideType.
	LogType(typ string)

	// LogImpl invoked by ModuleContainer.ProvideImpl.
	LogImpl(itfTyp, srvTyp string)

	// LogInject invoked by ModuleContainer.Inject.
	LogInject(parentTyp, fieldTyp, fieldName string)
}

// defaultLogger represents a default Logger.
type defaultLogger struct {
	Logger
	level LogLevel
}

// DefaultLogger creates a default Logger instance. Log style see LogName, LogType, LogImpl, LogInject.
// Notice that red color represents the name, yellow represents the type, ~ represents no name.
func DefaultLogger(level LogLevel) Logger {
	xcolor.ForceColor()
	return &defaultLogger{level: level}
}

// LogName logs like:
// 	[XMODULE] Name:    a <-- string
// Here `a` (in red) is the name in ModuleName, `string` (in yellow) is the type of this module.
func (d *defaultLogger) LogName(name, typ string) {
	if d.level&LogName != 0 {
		name = xcolor.Red.Sprint(name)
		typ = xcolor.Yellow.Sprint(typ)
		LogLeftArrow("Name:", name, typ)
	}
}

// LogType logs like:
// 	[XMODULE] Type:    ~ <-- string
// Here `~` (in red) is the flag of no name, `string` (in yellow) is the type of this module.
func (d *defaultLogger) LogType(typ string) {
	if d.level&LogType != 0 {
		auto := xcolor.Red.Sprint("~")
		typ = xcolor.Yellow.Sprint(typ)
		LogLeftArrow("Type:", auto, typ)
	}
}

// LogImpl logs like:
// 	[XMODULE] Impl:    ~ <-- IModule (*Module)
// Here `~` (in red) is the flag of no name, `IModule` (in yellow) is the interface type of module, `Module` (in yellow) is the type of this module.
func (d *defaultLogger) LogImpl(interfaceTyp, implTyp string) {
	if d.level&LogImpl != 0 {
		auto := xcolor.Red.Sprint("~")
		interfaceTyp = xcolor.Yellow.Sprint(interfaceTyp)
		implTyp = xcolor.Yellow.Sprint(implTyp)
		LogLeftArrow("Impl:", auto, fmt.Sprintf("%s (%s)", interfaceTyp, implTyp))
	}
}

// Inject logs like:
// 	[XMODULE] Inject:  int --> (*Module).I
// Here `int` (in yellow) is the type of field, `Module` (in yellow) is the type of struct, `I` (in red) is the name of field.
func (d *defaultLogger) LogInject(parentTyp, fieldTyp, fieldName string) {
	if d.level&LogInject != 0 {
		parentTyp = xcolor.Yellow.Sprint(parentTyp)
		fieldTyp = xcolor.Yellow.Sprint(fieldTyp)
		fieldName = xcolor.Red.Sprint(fieldName)
		LogRightArrow("Inject:", fieldTyp, fmt.Sprintf("(%s).%s", parentTyp, fieldName))
	}
}

// LogLeftArrow is the logger function with <-- (used in LogName, LogType, LogImpl).
// You can overwrite this function.
var LogLeftArrow = func(arg1, arg2, arg3 string) {
	fmt.Printf("[XMODULE] %-8s %-30s <-- %s\n", arg1, arg2, arg3)
}

// LogRightArrow is the logger function with --> (used in LogInject).
// You can overwrite this function.
var LogRightArrow = func(arg1, arg2, arg3 string) {
	fmt.Printf("[XMODULE] %-8s %-30s --> %s\n", arg1, arg2, arg3)
}
