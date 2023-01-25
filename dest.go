package go_cups_mod

type Dest struct {
	Name       string
	Instance   string
	IsDefault  bool
	NumOptions int
	Options    map[string]string
}

func (d *Dest) Status() string {
	value, _ := d.Options["printer-state"]
	switch value {
	case PRINTER_STATE_IDLE:
		return "idle"
	case PRINTER_STATE_PRINTING:
		return "printing"
	case PRINTER_STATE_STOPPED:
		return "stopped"
	default:
		return "unknown"
	}
}
