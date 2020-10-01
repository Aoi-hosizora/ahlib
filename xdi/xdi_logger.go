package xdi

import (
	"github.com/Aoi-hosizora/ahlib/xcolor"
	"log"
)

// Logger represents xdi.DiContainer's logger function.
type Logger interface {
	// ProvideName logs when DiContainer.ProvideName invoked.
	ProvideName(typ, name string)

	// ProvideType logs when DiContainer.ProvideType invoked.
	ProvideType(typ string)

	// ProvideImpl logs when DiContainer.ProvideImpl invoked.
	ProvideImpl(interfaceTyp, implTyp string)

	// Inject logs when DiContainer.Inject invoked.
	Inject(parentTyp, fieldTyp, fieldName string)
}

// defaultLogger represents the default logger.
type defaultLogger struct{}

// DefaultLogger creates a defaultLogger.
func DefaultLogger() Logger {
	xcolor.ForceColor()
	return &defaultLogger{}
}

// ProvideName logs like:
// 	[XDI] Name:    a <- *xdi.ServiceA
func (d *defaultLogger) ProvideName(typ, name string) {
	typ = xcolor.Yellow.Sprint(typ)
	name = xcolor.Red.Sprint(name)
	log.Printf("[XDI] %-8s %s <- %s", "Name:", name, typ)
}

// ProvideType logs like:
// 	[XDI] Type:    ~ <- *xdi.ServiceB
func (d *defaultLogger) ProvideType(typ string) {
	typ = xcolor.Yellow.Sprint(typ)
	log.Printf("[XDI] %-8s ~ <- %s", "Type:", typ)
}

// ProvideImpl logs like:
// 	[XDI] Impl:    ~ <- xdi.IServiceA  (*xdi.ServiceA)
func (d *defaultLogger) ProvideImpl(interfaceTyp, implTyp string) {
	interfaceTyp = xcolor.Yellow.Sprint(interfaceTyp)
	implTyp = xcolor.Yellow.Sprint(implTyp)
	log.Printf("[XDI] %-8s ~ <- %s (%s）", "Impl:", interfaceTyp, implTyp)
}

// Inject logs like:
// 	[XDI] Inject:  xdi.IServiceA -> (*xdi.ServiceB).SA
func (d *defaultLogger) Inject(parentTyp, fieldTyp, fieldName string) {
	parentTyp = xcolor.Yellow.Sprint(parentTyp)
	fieldTyp = xcolor.Yellow.Sprint(fieldTyp)
	fieldName = xcolor.Red.Sprint(fieldName)
	log.Printf("[XDI] %-8s %s -> (%s).%s", "Inject:", fieldTyp, parentTyp, fieldName)
}

// silentLogger represents the no logger as silent.
type silentLogger struct{}

// SilentLogger creates a silentLogger.
func SilentLogger() Logger {
	return &silentLogger{}
}

// ProvideName does nothing for logger.
func (s *silentLogger) ProvideName(string, string) {}

// ProvideType does nothing for logger.
func (s *silentLogger) ProvideType(string) {}

// ProvideImpl does nothing for logger.
func (s *silentLogger) ProvideImpl(string, string) {}

// Inject does nothing for logger.
func (s *silentLogger) Inject(string, string, string) {}
