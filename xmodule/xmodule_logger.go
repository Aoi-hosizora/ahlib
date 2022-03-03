package xmodule

import (
	"fmt"
	"github.com/Aoi-hosizora/ahlib/xcolor"
)

// LogLevel represents ModuleContainer's logger level.
type LogLevel uint8

const (
	// LogPrvName logs when ModuleContainer.ProvideByName invoked.
	LogPrvName LogLevel = 1 << iota

	// LogPrvType logs when ModuleContainer.ProvideByType invoked.
	LogPrvType

	// LogPrvIntf logs when ModuleContainer.ProvideByIntf invoked.
	LogPrvIntf

	// LogInjField logs when ModuleContainer.Inject invoked and some fields are injected.
	LogInjField

	// LogInjFinish logs when ModuleContainer.Inject invoked and injecting is finished.
	LogInjFinish

	// LogAll logs when ModuleContainer.ProvideByName, ModuleContainer.ProvideByType, ModuleContainer.ProvideByIntf and ModuleContainer.Inject invoked.
	LogAll = LogPrvName | LogPrvType | LogPrvIntf | LogInjField | LogInjFinish

	// LogSilent never logs, means disable the logger.
	LogSilent = LogLevel(0)
)

// Logger represents ModuleContainer's logger interface, a default logger can be created by DefaultLogger.
type Logger interface {
	// PrvName logs when ModuleContainer.ProvideByName invoked, can be enabled by LogPrvName flag.
	PrvName(moduleName, moduleType string)

	// PrvType logs when ModuleContainer.ProvideByType invoked, can be enabled by LogPrvType flag.
	PrvType(moduleType string)

	// PrvIntf logs when ModuleContainer.ProvideByIntf invoked, can be enabled by LogPrvIntf flag.
	PrvIntf(interfaceType, moduleType string)

	// InjField logs when ModuleContainer.Inject invoked and some fields are injected, can be enabled by LogInjField flag.
	InjField(moduleName, injecteeType, fieldName, moduleType string)

	// InjFinish logs when ModuleContainer.Inject invoked and injecting is finished, can be enabled by LogInjFinish flag.
	InjFinish(injecteeType string, count int, allInjected bool)
}

// defaultLogger represents an unexported default Logger type.
type defaultLogger struct {
	level      LogLevel
	logPrvFunc func(moduleName, moduleType string)
	logInjFunc func(moduleName, injecteeType, addition string)
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
func DefaultLogger(level LogLevel, logPrvFunc func(moduleName, moduleType string), logInjFunc func(moduleName, injecteeType, addition string)) Logger {
	xcolor.ForceColor()
	if logPrvFunc == nil {
		logPrvFunc = func(moduleName, moduleType string) {
			fmt.Printf("[Xmodule] Prv: %s <-- %s\n", xcolor.Red.AlignedSprint(-20, moduleName), xcolor.Yellow.Sprint(moduleType))
		}
	}
	if logInjFunc == nil {
		logInjFunc = func(moduleName, injecteeType, addition string) {
			fmt.Printf("[Xmodule] Inj: %s --> %s %s\n", xcolor.Red.AlignedSprint(-20, moduleName), xcolor.Blue.Sprint(injecteeType), xcolor.Yellow.Sprintf("(%s)", addition))
		}
	}
	return &defaultLogger{level: level, logPrvFunc: logPrvFunc, logInjFunc: logInjFunc}
}

// PrvName logs when ModuleContainer.ProvideByName invoked, can be enabled by LogPrvName flag.
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

// PrvType logs when ModuleContainer.ProvideByType invoked, can be enabled by LogPrvType flag.
//
// The default format logs like:
// 	[Xmodule] Prv: ~ <-- *Module
// 	              |-|   |-------|
// 	              red    yellow
// Here `~` is module name (means provide by type or intf), `*Module` is module type.
func (d *defaultLogger) PrvType(moduleType string) {
	if d.level&LogPrvType != 0 {
		d.logPrvFunc("~", moduleType)
	}
}

// PrvIntf logs when ModuleContainer.ProvideByIntf invoked, can be enabled by LogInjField flag.
//
// The default format logs like:
// 	[Xmodule] Prv: ~ <-- IModule (*Module)
// 	              |-|   |-----------------|
// 	              red         yellow
// Here `~` is module name (means provide by type or intf), `IModule` is interface type, `*Module` is module type.
func (d *defaultLogger) PrvIntf(interfaceType, moduleType string) {
	if d.level&LogPrvIntf != 0 {
		d.logPrvFunc("~", fmt.Sprintf("%s (%s)", interfaceType, moduleType))
	}
}

// InjField logs when ModuleContainer.Inject invoked and some fields are injected, can be enabled by LogInjField flag.
//
// The default format logs like:
// 	[Xmodule] Inj: xxx-tag --> *Struct (Field string)
// 	[Xmodule] Inj: ~       --> *Struct (Field string)
// 	              |-------|   |-------|--------------|
// 	                 red        blue       yellow
// Here `xxx-tag` or `~` is module name, `*Struct` is injectee type, `Field` is field name, `string` is module type.
func (d *defaultLogger) InjField(moduleName, injecteeType, fieldName, moduleType string) {
	if d.level&LogInjField != 0 {
		d.logInjFunc(moduleName, injecteeType, fmt.Sprintf("%s %s", fieldName, moduleType))
	}
}

// InjFinish logs when ModuleContainer.Inject invoked and injecting is finished, can be enabled by LogInjFinish flag.
//
// The default format logs like:
// 	[Xmodule] Inj: ... --> *Struct (#=3, all injected)
// 	[Xmodule] Inj: ... --> *Struct (#=2, not all injected)
// 	              |---|   |-------|-----------------------|
// 	               red      blue           yellow
// Here `...` means injecting is finished, `*Struct` is injectee type, `#=3` is injected fields count, `all injected` and
// `not all injected` means whether all fields with `module` are injected or not.
func (d *defaultLogger) InjFinish(injecteeType string, count int, allInjected bool) {
	if d.level&LogInjFinish != 0 {
		flag := "all injected"
		if !allInjected {
			flag = "not all injected"
		}
		d.logInjFunc("...", injecteeType, fmt.Sprintf("#=%d, %s", count, flag))
	}
}
