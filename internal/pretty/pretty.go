package pretty

import "fmt"

type PrettyPrintOption func(p *prettyPrinter)

type PrettyPrinter interface {
	Print(s string)
	Printf(format string, a ...any)
	PrintWarning(s string)
	PrintWarningf(format string, a ...any)
	PrintError(s string)
	PrintErrorf(format string, a ...any)
}

type prettyPrinter struct {
	InfoColor int32 `json:"infoColor"`
	WarnColor int32 `json:"warnColor"`
	ErrColor  int32 `json:"errorColor"`
}

func WithInfoColor(c int32) PrettyPrintOption {
	return func(p *prettyPrinter) {
		p.InfoColor = c
	}
}

func WithWarnColor(c int32) PrettyPrintOption {
	return func(p *prettyPrinter) {
		p.WarnColor = c
	}
}

func WithErrColor(c int32) PrettyPrintOption {
	return func(p *prettyPrinter) {
		p.ErrColor = c
	}
}

func NewPrettyPrinter(opts ...PrettyPrintOption) *prettyPrinter {
	const (
		infoColor = int32(92)
		warnColor = int32(93)
		errColor  = int32(91)
	)

	printer := &prettyPrinter{
		InfoColor: infoColor,
		WarnColor: warnColor,
		ErrColor:  errColor,
	}
	for _, opt := range opts {
		opt(printer)
	}
	return printer
}

func (p *prettyPrinter) Print(s string) {
	fmt.Printf("\x1b[1;%dm%s\x1b[0m\n", p.InfoColor, s)
}

func (p *prettyPrinter) Printf(format string, a ...any) {
	fstring := fmt.Sprintf(format, a)
	p.Print(fstring)
}

func (p *prettyPrinter) PrintWarning(s string) {
	fmt.Printf("\x1b[1;%dm%s\x1b[0m\n", p.WarnColor, s)
}

func (p *prettyPrinter) PrintWarningf(format string, a ...any) {
	fstring := fmt.Sprintf(format, a)
	p.PrintWarning(fstring)
}

func (p *prettyPrinter) PrintError(s string) {
	fmt.Printf("\x1b[1;%dm%s\x1b[0m\n", p.ErrColor, s)
}

func (p *prettyPrinter) PrintErrorf(format string, a ...any) {
	fstring := fmt.Sprintf(format, a)
	p.PrintError(fstring)
}

func Print(s string) {
	printer := NewPrettyPrinter()
	printer.Print(s)
}

func Printf(format string, a ...any) {
	p := NewPrettyPrinter()
	fstring := fmt.Sprintf(format, a)
	p.Print(fstring)
}

func PrintWarning(s string) {
	printer := NewPrettyPrinter()
	printer.PrintWarning(s)
}

func PrintWarningf(format string, a ...any) {
	p := NewPrettyPrinter()
	fstring := fmt.Sprintf(format, a)
	p.PrintWarning(fstring)
}

func PrintError(s string) {
	printer := NewPrettyPrinter()
	printer.PrintError(s)
}

func PrintErrorf(format string, a ...any) {
	printer := NewPrettyPrinter()
	fstring := fmt.Sprintf(format, a)
	printer.PrintError(fstring)
}
