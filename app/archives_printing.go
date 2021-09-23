package app

type ArchivePrinter interface {
	print(archives []Archive) string
}

type PrettyArchivePrinter struct{}

func (p *PrettyArchivePrinter) print(archives []Archive) string {
	return ""
}

type JsonArchivePrinter struct{}

func (p *JsonArchivePrinter) print(archives []Archive) string {
	return ""
}
