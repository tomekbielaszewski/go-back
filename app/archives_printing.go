package app

import (
	"github.com/jedib0t/go-pretty/v6/table"
)

type ArchivePrinter interface {
	Print(archives []*Archive) string
}

type PrettyArchivePrinter struct{}

func (p *PrettyArchivePrinter) Print(archives []*Archive) string {
	t := table.NewWriter()

	t.AppendHeader(table.Row{"ID", "Path", "Size", "Bucket"})

	for _, a := range archives {
		t.AppendRow(table.Row{a.Id, a.Path, a.Size, a.Bucket})
	}

	t.SetStyle(table.StyleLight)
	return t.Render()
}

type JsonArchivePrinter struct{}

func (p *JsonArchivePrinter) Print(archives []*Archive) string {
	return ""
}
