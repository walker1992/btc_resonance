package log

import (
	"fmt"
	"os"
	"sync/atomic"
	"time"

	"github.com/go-stack/stack"
)

const skipLevel = 2

// swapHandler wraps another handler that may be swapped out
// dynamically at runtime in a thread-safe fashion.
type swapHandler struct {
	handler atomic.Value
}

func (h *swapHandler) Log(r *Record) error {
	return (*h.handler.Load().(*Handler)).Log(r)
}

func (h *swapHandler) Swap(newHandler Handler) {
	h.handler.Store(&newHandler)
}

func (h *swapHandler) Get() Handler {
	return *h.handler.Load().(*Handler)
}

type Lvl int

const (
	LvlAll Lvl = iota
	LvlDebug
	LvlInfo
	LvlWarn
	LvlError
	LvlOff
	LvlCrit
)

// AlignedString returns a 5-character string containing the name of a Lvl.
func (l Lvl) AlignedString() string {
	switch l {
	case LvlAll:
		return "all"
	case LvlDebug:
		return "debug"
	case LvlInfo:
		return "info"
	case LvlWarn:
		return "warn"
	case LvlError:
		return "error"
	case LvlOff:
		return "off"
	case LvlCrit:
		return "crit"
	default:
		panic("Unknown level")
	}
}

// Strings returns the name of a Lvl.
func (l Lvl) String() string {
	switch l {
	case LvlAll:
		return "all   "
	case LvlDebug:
		return "debug "
	case LvlInfo:
		return "info  "
	case LvlWarn:
		return "warn  "
	case LvlError:
		return "error "
	case LvlOff:
		return "off   "
	case LvlCrit:
		return "lvlcrit"
	default:
		panic("Unknown level")
	}
}

//LvlFromString returns the appropriate Lvl from a string name.
//Useful for parsing command line args and configuration files.
func LvlFromString(lvlString string) (Lvl, error) {
	switch lvlString {
	case "all":
		return LvlAll, nil
	case "debug", "dbug":
		return LvlDebug, nil
	case "info":
		return LvlInfo, nil
	case "warn":
		return LvlWarn, nil
	case "error", "eror":
		return LvlError, nil
	case "off":
		return LvlOff, nil
	case "crit":
		return LvlCrit, nil
	default:
		return LvlDebug, fmt.Errorf("Unknown level: %v", lvlString)
	}
}

// A Record is what a Logger asks its handler to write
type Record struct {
	Name string
	Time time.Time
	Lvl  Lvl
	Msg  string
	Call stack.Call
}

func (r Record) String() string {
	return r.Call.String()
}

// A Logger writes key/value pairs to a Handler
type Logger interface {
	// New returns a new Logger that has this logger's context plus the given context
	New(name string) Logger

	// GetHandler gets the handler associated with the logger.
	GetHandler() Handler

	// SetHandler updates the logger to write records to the specified handler.
	SetHandler(h Handler)

	// SetEnable enable output without handler.
	SetEnable(e bool)

	// Log a message at the given level with context key/value pairs
	Debug(format string, arg ...interface{})
	Info(format string, arg ...interface{})
	Warn(format string, arg ...interface{})
	Error(format string, arg ...interface{})
	Crit(format string, arg ...interface{})
}

type logger struct {
	name   string
	h      *swapHandler
	enable bool
}

func (l *logger) New(name string) Logger {
	child := &logger{name: name, h: new(swapHandler), enable: true}
	child.SetHandler(l.h)
	return child
}
func (l *logger) SetEnable(e bool) {
	l.enable = e
}
func (l *logger) Debug(format string, arg ...interface{}) {
	if l != nil && l.enable {
		l.write(LvlDebug, fmt.Sprintf(format, arg...), skipLevel)
	}
}
func (l *logger) Info(format string, arg ...interface{}) {
	if l != nil && l.enable {
		l.write(LvlInfo, fmt.Sprintf(format, arg...), skipLevel)
	}
}

func (l *logger) Warn(format string, arg ...interface{}) {
	if l != nil && l.enable {
		l.write(LvlWarn, fmt.Sprintf(format, arg...), skipLevel)
	}
}
func (l *logger) Error(format string, arg ...interface{}) {
	if l != nil && l.enable {
		l.write(LvlError, fmt.Sprintf(format, arg...), skipLevel)
	}
}

func (l *logger) Crit(format string, arg ...interface{}) {
	if l != nil && l.enable {
		l.write(LvlCrit, fmt.Sprintf(format, arg...), skipLevel)
		os.Exit(1)
	}
}

func (l *logger) GetHandler() Handler {
	return l.h.Get()
}
func (l *logger) SetHandler(h Handler) {
	l.h.Swap(h)
}

func (l *logger) write(lvl Lvl, msg string, skip int) {
	l.h.Log(&Record{
		Name: l.name,
		Time: time.Now(),
		Lvl:  lvl,
		Msg:  msg,
		Call: stack.Caller(skip),
	})
}

// Lazy allows you to defer calculation of a logged value that is expensive
// to compute until it is certain that it must be evaluated with the given filters.
//
// Lazy may also be used in conjunction with a Logger's New() function
// to generate a child logger which always reports the current value of changing
// state.
//
// You may wrap any function which takes no arguments to Lazy. It may return any
// number of values of any type.
type Lazy struct {
	Fn interface{}
}
