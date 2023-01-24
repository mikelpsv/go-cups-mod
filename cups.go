package go_cups_mod

/*
#cgo LDFLAGS: -lcups
#include "cups/cups.h"
#include "cupsmod.h"
*/
import "C"
import (
	"fmt"
	"unsafe"
)

type Dest struct {
	Name       string
	Instance   string
	IsDefault  bool
	NumOptions int
	Options    map[string]string
}

type Connection struct {
	isDefault bool
	address   string
	NumDests  int
	Dests     []Dest
}

func NewConnection() *Connection {
	connection := &Connection{
		isDefault: true,
		Dests:     make([]Dest, 0),
	}
	return connection
}

func (c *Connection) EnumDestinations() (int, error) {
	var dests *C.cups_dest_t
	var t, m C.cups_ptype_t
	n := C.cups_enum_dests(t, m, &dests)

	destPtr0 := uintptr(unsafe.Pointer(dests))
	for i := 0; i < int(n); i++ {
		p := destPtr0 + unsafe.Sizeof(*dests)*uintptr(i)
		dest := (*C.cups_dest_t)(unsafe.Pointer(p))
		d := Dest{
			Name:       C.GoString(dest.name),
			Instance:   C.GoString(dest.instance),
			IsDefault:  int(dest.is_default) > 0,
			NumOptions: int(dest.num_options),
			Options:    make(map[string]string, 0),
		}
		c.Dests = append(c.Dests, d)
		fmt.Printf("%v\n", C.GoString(dest.name))
	}
	return int(n), nil
}

// https://github.com/OpenPrinting/cups/blob/63890581f643759bd93fa4416ab53e7380c6bd2d/cups/cups.h#L350
func (d *Dest) PrintFile() int {

	return 0
}

// https://github.com/OpenPrinting/cups/blob/63890581f643759bd93fa4416ab53e7380c6bd2d/cups/cups.h#L465
func (d *Dest) StartDocument() {

}

// https://github.com/OpenPrinting/cups/blob/63890581f643759bd93fa4416ab53e7380c6bd2d/cups/cups.h#L556
func (d *Dest) StartDestDocument() {

}
