package go_cups_mod

/*
#cgo LDFLAGS: -lcups
#include "cups/cups.h"
*/
import "C"
import (
	"fmt"
)

type Dest struct {
	Name       string
	Instance   string
	IsDefault  bool
	NumOptions int
	options    map[string]string
}

// Status return the status as a string
func (d *Dest) Status() string {
	value, _ := d.options["printer-state"]
	switch value {
	case PrinterStateIdle:
		return "idle"
	case PrinterStatePrinting:
		return "printing"
	case PrinterStateStopped:
		return "stopped"
	default:
		return "unknown"
	}
}

// GetOptions returns destination options
func (d *Dest) GetOptions() map[string]string {
	return d.options
}

// PrintFile send a file to print
// return job id
//
// https://github.com/apple/cups/blob/c9da6f63b263faef5d50592fe8cf8056e0a58aa2/cups/util.c#L696
func (d *Dest) PrintFile(fileName string, jobTitle string) (int, error) {
	jobId := C.cupsPrintFile(C.CString(d.Name), C.CString(fileName), C.CString(jobTitle), C.int(len(d.options)), nil)
	if jobId == 0 {
		errId := int(C.cupsLastError())
		errStr := C.GoString(C.cupsLastErrorString())
		return 0, fmt.Errorf("failed to print with error code: %d %s", errId, string(errStr))
	}
	return int(jobId), nil
}

// TODO: https://github.com/OpenPrinting/cups/blob/63890581f643759bd93fa4416ab53e7380c6bd2d/cups/cups.h#L465
func (d *Dest) StartDocument() {
	//
}

// TODO: https://github.com/OpenPrinting/cups/blob/63890581f643759bd93fa4416ab53e7380c6bd2d/cups/cups.h#L556
func (d *Dest) StartDestDocument() {
	//
}
