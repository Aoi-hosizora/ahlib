package xdi

import (
	"fmt"
	"github.com/Aoi-hosizora/ahlib/xcolor"
)

type LogLevel int8

const (
	LogName LogLevel = 1 << iota
	LogType
	LogImpl
	LogInject
	LogAll    = LogName | LogType | LogImpl | LogInject
	LogSilent = LogLevel(0)
)

// Logger represents xdi.DiContainer's logger function.
type Logger interface {
	// LogName logs when DiContainer.ProvideName invoked.
	LogName(name, typ string)

	// LogType logs when DiContainer.ProvideType invoked.
	LogType(typ string)

	// LogImpl logs when DiContainer.ProvideImpl invoked.
	LogImpl(interfaceTyp, implTyp string)

	// LogInject logs when DiContainer.Inject invoked.
	LogInject(parentTyp, fieldTyp, fieldName string)
}

// defaultLogger represents the default logger.
type defaultLogger struct {
	level LogLevel
}

// DefaultLogger creates a defaultLogger.
func DefaultLogger(level LogLevel) Logger {
	xcolor.ForceColor()
	return &defaultLogger{
		level: level,
	}
}

// LogName logs like:
// 	[XDI] Name:    a <- *xdi.ServiceA
func (d *defaultLogger) LogName(name, typ string) {
	if d.level&LogName != 0 {
		name = xcolor.Red.Sprint(name)
		typ = xcolor.Yellow.Sprint(typ)
		fmt.Printf("[XDI] %-8s %-30s <-- %s\n", "Name:", name, typ)
	}
}

// LogType logs like:
// 	[XDI] Type:    ~ <- *xdi.ServiceB
func (d *defaultLogger) LogType(typ string) {
	if d.level&LogType != 0 {
		auto := xcolor.Red.Sprint("~")
		typ = xcolor.Yellow.Sprint(typ)
		fmt.Printf("[XDI] %-8s %-30s <-- %s\n", "Type:", auto, typ)
	}
}

// LogImpl logs like:
// 	[XDI] Impl:    ~ <- xdi.IServiceA  (*xdi.ServiceA)
func (d *defaultLogger) LogImpl(interfaceTyp, implTyp string) {
	if d.level&LogImpl != 0 {
		auto := xcolor.Red.Sprint("~")
		interfaceTyp = xcolor.Yellow.Sprint(interfaceTyp)
		implTyp = xcolor.Yellow.Sprint(implTyp)
		fmt.Printf("[XDI] %-8s %-30s <-- %s (%s)\n", "Impl:", auto, interfaceTyp, implTyp)
	}
}

// Inject logs like:
// 	[XDI] Inject:  xdi.IServiceA -> (*xdi.ServiceB).SA
func (d *defaultLogger) LogInject(parentTyp, fieldTyp, fieldName string) {
	if d.level&LogInject != 0 {
		parentTyp = xcolor.Yellow.Sprint(parentTyp)
		fieldTyp = xcolor.Yellow.Sprint(fieldTyp)
		fieldName = xcolor.Red.Sprint(fieldName)
		fmt.Printf("[XDI] %-8s %-30s --> (%s).%s\n", "Inject:", fieldTyp, parentTyp, fieldName)
	}
}
