package app

import (
	"github.com/jedib0t/go-pretty/v6/table"
)

type ArchiveFormatter interface {
	Format(archives []*Archive) string
}

type PrettyArchiveFormatter struct{}

func (p *PrettyArchiveFormatter) Format(archives []*Archive) string {
	if len(archives) == 0 {
		return ""
	}

	t := table.NewWriter()

	t.AppendHeader(table.Row{"ID", "Path", "Size", "Bucket"})

	for _, a := range archives {
		t.AppendRow(table.Row{a.Id, a.Path, a.Size, a.Bucket})
	}

	t.SetStyle(table.StyleLight)
	return t.Render()
}

type JsonArchiveFormatter struct{}

func (p *JsonArchiveFormatter) Format(archives []*Archive) string {
	return ""
}
