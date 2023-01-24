package go_cups_mod

/*
#cgo LDFLAGS: -lcups
#include "cups/cups.h"
#include <cupsmod.h>
*/

import (
	"C"
	"fmt"
	"unsafe"
)

type Dest struct {
	Name      string
	IsDefault bool
	Options   map[string]string
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
	}
	return connection
}

func (c *Connection) EnumDests() (int, error) {
	var dests *C.cups_dest_t
	var userData C.user_data_t

	C.cupsEnumDests(C.CUPS_DEST_FLAGS_NONE, 1000, C.NULL, 0, 0, (C.cups_dest_cb_t)&cb, &userData)

	C.cupsFreeDests(userData.NumDests, userData.Dests)
	*dests = C.NULL

	return 0, nil
}

func cb(userData *C.user_data_t, flags int, dest unsafe.Pointer) {
	fmt.Println(dest)
}
