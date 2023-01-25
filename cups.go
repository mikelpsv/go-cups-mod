package go_cups_mod

/*
#cgo CFLAGS: -g -Wall
#cgo LDFLAGS: -lcups
#include "cups/cups.h"
#include "cupsmod.h"
#define DEBUG
*/
import "C"
import (
	"unsafe"
)

const (
	PrinterStateIdle     = "3"
	PrinterStatePrinting = "4"
	PrinterStateStopped  = "5"
)

type Connection struct {
	isDefault bool
	address   string
	NumDests  int
	Dests     []Dest
}

func NewConnection() *Connection {
	connection := &Connection{
		isDefault: true,
	}
	return connection
}

// GetOptions возвращает список возможных предопределенных опций
func (c *Connection) GetOptions() []string {
	return []string{
		"auth-info-required",
		"printer-info",
		"printer-is-accepting-jobs",
		"printer-is-shared",
		"printer-location",
		"printer-make-and-model",
		"printer-state",
		"printer-state-change-time",
		"printer-state-reasons",
		"printer-type",
		"printer-uri-supported",
	}
}

// EnumDestinations заполняет список назначений
func (c *Connection) EnumDestinations() (int, error) {
	var dests *C.cups_dest_t
	var t, m C.cups_ptype_t
	n := C.cups_enum_dests(t, m, &dests)
	destPtr0 := uintptr(unsafe.Pointer(dests))

	c.Dests = make([]Dest, 0)

	for i := 0; i < int(n); i++ {
		p := destPtr0 + unsafe.Sizeof(*dests)*uintptr(i)
		dest := (*C.cups_dest_t)(unsafe.Pointer(p))

		d := Dest{
			Name:       C.GoString(dest.name),
			Instance:   C.GoString(dest.instance),
			IsDefault:  int(dest.is_default) > 0,
			NumOptions: int(dest.num_options),
			options:    make(map[string]string, 0),
		}

		for _, option := range c.GetOptions() {
			o := C.cupsGetOption(C.CString(option), dest.num_options, dest.options)
			d.options[option] = C.GoString(o)
		}
		c.Dests = append(c.Dests, d)
	}
	return int(n), nil
}

// https://github.com/OpenPrinting/cups/blob/63890581f643759bd93fa4416ab53e7380c6bd2d/cups/cups.h#L465
func (d *Dest) StartDocument() {

}

// https://github.com/OpenPrinting/cups/blob/63890581f643759bd93fa4416ab53e7380c6bd2d/cups/cups.h#L556
func (d *Dest) StartDestDocument() {

}
