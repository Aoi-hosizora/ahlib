package xmodule

import (
	"fmt"
	"github.com/Aoi-hosizora/ahlib/xcolor"
)

// LogLevel represents ModuleContainer's logger level.
type LogLevel uint8

const (
	// LogPrvName logs when ModuleContainer.ProvideName invoked.
	LogPrvName LogLevel = 1 << iota

	// LogPrvType logs when ModuleContainer.ProvideType invoked.
	LogPrvType

	// LogPrvImpl logs when ModuleContainer.ProvideImpl invoked.
	LogPrvImpl

	// LogInjField logs when ModuleContainer.Inject invoked and some fields are injected.
	LogInjField

	// LogInjFinish logs when ModuleContainer.Inject invoked and injecting is finished.
	LogInjFinish

	// LogAll logs when ModuleContainer.ProvideName, ModuleContainer.ProvideType, ModuleContainer.ProvideImpl and ModuleContainer.Inject invoked.
	LogAll = LogPrvName | LogPrvType | LogPrvImpl | LogInjField | LogInjFinish

	// LogSilent never logs, means disable the logger.
	LogSilent = LogLevel(0)
)

// Logger represents ModuleContainer's logger interface, a default logger can be created by DefaultLogger.
type Logger interface {
	// PrvName logs when ModuleContainer.ProvideName invoked, can be enabled by LogPrvName flag.
	PrvName(moduleName, moduleType string)

	// PrvType logs when ModuleContainer.ProvideType invoked, can be enabled by LogPrvType flag.
	PrvType(moduleType string)

	// PrvImpl logs when ModuleContainer.ProvideImpl invoked, can be enabled by LogInjField flag.
	PrvImpl(interfaceTyp, moduleType string)

	// InjField logs when ModuleContainer.Inject invoked and some fields are injected, can be enabled by LogInjField flag.
	InjField(moduleName, structType, fieldName, fieldType string)

	// InjFinish logs when ModuleContainer.Inject invoked and injecting is finished, can be enabled by LogInjFinish flag.
	InjFinish(structTyp string, count int, allInjected bool)
}

// defaultLogger represents a default Logger.
type defaultLogger struct {
	level      LogLevel
	logPrvFunc func(moduleName, moduleType string)
	logInjFunc func(moduleName, structName, addition string)
}

var _ Logger = (*defaultLogger)(nil)

// DefaultLogger creates a default Logger instance, with given LogLevel, nillable logPrjFunc and nillable logInjFunc.
//
// The default format for providing logs like:
// 	[Xmodule] Prv: int                  <-- int
// 	[Xmodule] Prv: ~                    <-- float64
// 	[Xmodule] Prv: ~                    <-- error (*errors.errorString)
// 	              |--------------------|   |---------------------------|
// 	                        20                           ...
//
// The default format for injecting logs like:
// 	[Xmodule] Inj: uint                 --> *xmodule.testStruct (Uint uint)
// 	[Xmodule] Inj: ~                    --> *xmodule.testStruct (String string)
// 	[Xmodule] Inj: ~                    --> *xmodule.testStruct (Error error)
// 	[Xmodule] Inj: ...                  --> *xmodule.testStruct (#=6, all injected)
// 	[Xmodule] Inj: ...                  --> *xmodule.testStruct (#=4, not all injected)
// 	              |--------------------|   |-------------------|-----------------------|
// 	                        20                      ...                   ...
func DefaultLogger(level LogLevel, logPrvFunc func(moduleName, moduleType string), logInjFunc func(moduleName, structName, addition string)) Logger {
	xcolor.ForceColor()
	if logPrvFunc == nil {
		logPrvFunc = func(moduleName, moduleType string) {
			fmt.Printf("[Xmodule] Prv: %s <-- %s\n", xcolor.Red.AlignedSprint(-20, moduleName), xcolor.Yellow.Sprint(moduleType))
		}
	}
	if logInjFunc == nil {
		logInjFunc = func(moduleName, structName, addition string) {
			fmt.Printf("[Xmodule] Inj: %s --> %s %s\n", xcolor.Red.AlignedSprint(-20, moduleName), structName, xcolor.Yellow.Sprintf("(%s)", addition))
		}
	}
	return &defaultLogger{level: level, logPrvFunc: logPrvFunc, logInjFunc: logInjFunc}
}

// PrvName logs when ModuleContainer.ProvideName invoked, can be enabled by LogPrvName flag.
//
// The default format logs like:
// 	[Xmodule] Prv: xxx-tag <-- *Module
// 	              |-------|   |-------|
// 	                 red       yellow
// Here `xxx-tag` is module name, `*Module` is module type.
func (d *defaultLogger) PrvName(moduleName, moduleType string) {
	if d.level&LogPrvName != 0 {
		d.logPrvFunc(moduleName, moduleType)
	}
}

// PrvType logs when ModuleContainer.ProvideType invoked, can be enabled by LogPrvType flag.
//
// The default format logs like:
// 	[Xmodule] Prv: ~ <-- *Module
// 	              |-|   |-------|
// 	              red    yellow
// Here `~` is module name (means provide by type or impl), `*Module` is module type.
func (d *defaultLogger) PrvType(moduleType string) {
	if d.level&LogPrvType != 0 {
		d.logPrvFunc("~", moduleType)
	}
}

// PrvImpl logs when ModuleContainer.ProvideImpl invoked, can be enabled by LogInjField flag.
//
// The default format logs like:
// 	[Xmodule] Prv: ~ <-- IModule (*Module)
// 	              |-|   |-----------------|
// 	              red         yellow
// Here `~` is module name (means provide by type or impl), `IModule` is interface type, `*Module` is module type.
func (d *defaultLogger) PrvImpl(interfaceType, moduleType string) {
	if d.level&LogPrvImpl != 0 {
		d.logPrvFunc("~", fmt.Sprintf("%s (%s)", interfaceType, moduleType))
	}
}

// InjField logs when ModuleContainer.Inject invoked and some fields are injected, can be enabled by LogInjField flag.
//
// The default format logs like:
// 	[Xmodule] Inj: xxx-tag --> *Struct (Field string)
// 	[Xmodule] Inj: ~       --> *Struct (Field string)
// 	              |-------|   |-------|--------------|
// 	                 red       default     yellow
// Here `xxx-tag` or `~` is module name, `*Struct` is struct type, `Field` is field name, `string` is field type.
func (d *defaultLogger) InjField(moduleName, structType, fieldName, fieldType string) {
	if d.level&LogInjField != 0 {
		d.logInjFunc(moduleName, structType, fmt.Sprintf("%s %s", fieldName, fieldType))
	}
}

// InjFinish logs when ModuleContainer.Inject invoked and injecting is finished, can be enabled by LogInjFinish flag.
//
// The default format logs like:
// 	[Xmodule] Inj: ... --> *Struct (#=3, all injected)
// 	[Xmodule] Inj: ... --> *Struct (#=2, not all injected)
// 	              |---|   |-------|-----------------------|
// 	               red     default         yellow
// Here `...` means injecting is finished, `*Struct` is struct type, `#=3` is injected fields count, `all injected` and
// `not all injected` means whether all fields with `module` are injected or not.
func (d *defaultLogger) InjFinish(structType string, count int, allInjected bool) {
	if d.level&LogInjFinish != 0 {
		flag := "all injected"
		if !allInjected {
			flag = "not all injected"
		}
		d.logInjFunc("...", structType, fmt.Sprintf("#=%d, %s", count, flag))
	}
}
