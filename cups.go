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

const (
	CupsCopies = C.CUPS_COPIES

	CupsMedia        = C.CUPS_MEDIA
	CupsMedia3x5     = C.CUPS_MEDIA_3X5
	CupsMedia4x6     = C.CUPS_MEDIA_4X6
	CupsMedia5x7     = C.CUPS_MEDIA_5X7
	CupsMedia8x10    = C.CUPS_MEDIA_8X10
	CupsMediaA3      = C.CUPS_MEDIA_A3
	CupsMediaA4      = C.CUPS_MEDIA_A4
	CupsMediaA5      = C.CUPS_MEDIA_A5
	CupsMediaA6      = C.CUPS_MEDIA_A6
	CupsMediaEnv10   = C.CUPS_MEDIA_ENV10
	CupsMediaEnvDL   = C.CUPS_MEDIA_ENVDL
	CupsMediaLegal   = C.CUPS_MEDIA_LEGAL
	CupsMediaLetter  = C.CUPS_MEDIA_LETTER
	CupsMediaPhotoL  = C.CUPS_MEDIA_PHOTO_L
	CupsMediaTabloid = C.CUPS_MEDIA_TABLOID

	CupsMediaSource       = C.CUPS_MEDIA_SOURCE
	CupsMediaSourceAuto   = C.CUPS_MEDIA_SOURCE_AUTO
	CupsMediaSourceManual = C.CUPS_MEDIA_SOURCE_MANUAL
	CupsFinishings        = C.CUPS_FINISHINGS

	CupsMediaType             = C.CUPS_MEDIA_TYPE
	CupsMediaTypeAuto         = C.CUPS_MEDIA_TYPE_AUTO
	CupsMediaTypeEnvelope     = C.CUPS_MEDIA_TYPE_ENVELOPE
	CupsMediaTypeLabels       = C.CUPS_MEDIA_TYPE_LABELS
	CupsMediaTypeLetterhead   = C.CUPS_MEDIA_TYPE_LETTERHEAD
	CupsMediaTypePhoto        = C.CUPS_MEDIA_TYPE_PHOTO
	CupsMediaTypePhotoGlossy  = C.CUPS_MEDIA_TYPE_PHOTO_GLOSSY
	CupsMediaTypePhotoMatte   = C.CUPS_MEDIA_TYPE_PHOTO_MATTE
	CupsMediaTypePlain        = C.CUPS_MEDIA_TYPE_PLAIN
	CupsMediaTypeTransparency = C.CUPS_MEDIA_TYPE_TRANSPARENCY

	CupsNumberUp = C.CUPS_NUMBER_UP
)

//CUPS_ORIENTATION: Controls the orientation of document pages placed on the media: CUPS_ORIENTATION_PORTRAIT or CUPS_ORIENTATION_LANDSCAPE.
//CUPS_PRINT_COLOR_MODE: Controls whether the output is in color (CUPS_PRINT_COLOR_MODE_COLOR), grayscale (CUPS_PRINT_COLOR_MODE_MONOCHROME), or either (CUPS_PRINT_COLOR_MODE_AUTO).
//CUPS_PRINT_QUALITY: Controls the generate quality of the output: CUPS_PRINT_QUALITY_DRAFT, CUPS_PRINT_QUALITY_NORMAL, or CUPS_PRINT_QUALITY_HIGH.
//CUPS_SIDES: Controls whether prints are placed on one or both sides of the media: CUPS_SIDES_ONE_SIDED, CUPS_SIDES_TWO_SIDED_PORTRAIT, or CUPS_SIDES_TWO_SIDED_LANDSCAPE

type Connection struct {
	isDefault bool
	address   string
	NumDests  int
	Dests     []Dest
}

// NewConnection creates a connection object
func NewConnection() *Connection {
	connection := &Connection{
		isDefault: true,
	}
	return connection
}

// GetOptions returns a list of possible predefined options
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

// EnumDestinations updates destinations list
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
