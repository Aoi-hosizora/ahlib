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
	InjField(moduleName, injecteeType, fieldName, fieldType string)

	// InjFinish logs when ModuleContainer.Inject invoked and injecting is finished, can be enabled by LogInjFinish flag.
	InjFinish(injecteeType string, injectedCount, totalCount int)
}

// defaultLogger represents an unexported default Logger type.
type defaultLogger struct {
	level      LogLevel
	logPrvFunc func(moduleName, moduleType string)
	logInjFunc func(moduleName, injecteeType, addition string)
}

var _ Logger = (*defaultLogger)(nil)

// DefaultLogger creates a default Logger instance, with given LogLevel, nillable logPrjFunc and logInjFunc.
//
// The default format for providing logs like:
// 	[Xmodule] Prv: int                  <-- int
// 	[Xmodule] Prv: ~                    <-- float64
// 	[Xmodule] Prv: ~                    <-- error (*errors.errorString)
// 	              |--------------------|   |---------------------------|
// 	                        20                           ...
//
// The default format for injecting logs like:
// 	[Xmodule] Inj: uint                 --> *xmodule.testStruct (Uint: uint)
// 	[Xmodule] Inj: ~                    --> *xmodule.testStruct (String: string)
// 	[Xmodule] Inj: ~                    --> *xmodule.testStruct (Error: error)
// 	[Xmodule] Inj: ...                  --> *xmodule.testStruct (#=6/6, all injected)
// 	[Xmodule] Inj: ...                  --> *xmodule.testStruct (#=4/6, partially injected)
// 	              |--------------------|   |-------------------|---------------------------|
// 	                        20                      ...                     ...
func DefaultLogger(level LogLevel, logPrvFunc func(moduleName, moduleType string), logInjFunc func(moduleName, injecteeType, addition string)) Logger {
	xcolor.ForceColor()
	if logPrvFunc == nil {
		logPrvFunc = func(moduleName, moduleType string) {
			fmt.Printf("[Xmodule] Prv: %s <-- %s\n", xcolor.Red.ASprint(-20, moduleName), xcolor.Yellow.Sprint(moduleType))
		}
	}
	if logInjFunc == nil {
		logInjFunc = func(moduleName, injecteeType, addition string) {
			fmt.Printf("[Xmodule] Inj: %s --> %s %s\n", xcolor.Red.ASprint(-20, moduleName), xcolor.Blue.Sprint(injecteeType), xcolor.Yellow.Sprintf("(%s)", addition))
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
// Here `~` is module name (means provide by type or intf), `IModule` is interface type, `*Module` is implemented module type.
func (d *defaultLogger) PrvIntf(interfaceType, moduleType string) {
	if d.level&LogPrvIntf != 0 {
		d.logPrvFunc("~", fmt.Sprintf("%s (%s)", interfaceType, moduleType))
	}
}

// InjField logs when ModuleContainer.Inject invoked and some fields are injected, can be enabled by LogInjField flag.
//
// The default format logs like:
// 	[Xmodule] Inj: xxx-tag --> *Struct (Field: string)
// 	[Xmodule] Inj: ~       --> *Struct (Field: string)
// 	              |-------|   |-------|---------------|
// 	                 red        blue       yellow
// Here `xxx-tag` or `~` is module name, `*Struct` is injectee type, `Field` is injectee field name, `string` is field type.
func (d *defaultLogger) InjField(moduleName, injecteeType, fieldName, fieldType string) {
	if d.level&LogInjField != 0 {
		d.logInjFunc(moduleName, injecteeType, fmt.Sprintf("%s: %s", fieldName, fieldType))
	}
}

// InjFinish logs when ModuleContainer.Inject invoked and injecting is finished, can be enabled by LogInjFinish flag.
//
// The default format logs like:
// 	[Xmodule] Inj: ... --> *Struct (#=3/3, all injected)
// 	[Xmodule] Inj: ... --> *Struct (#=2/3, partially injected)
// 	              |---|   |-------|---------------------------|
// 	               red      blue             yellow
// Here `...` means injecting is finished, `*Struct` is injectee type, `#=3/3` is injected and total fields count, `all injected`
// and `partially injected` means whether all module fields are injected or not.
func (d *defaultLogger) InjFinish(injecteeType string, injectedCount, totalCount int) {
	if d.level&LogInjFinish != 0 {
		flag := "all injected"
		if injectedCount != totalCount {
			flag = "partially injected"
		}
		d.logInjFunc("...", injecteeType, fmt.Sprintf("#=%d/%d, %s", injectedCount, totalCount, flag))
	}
}
