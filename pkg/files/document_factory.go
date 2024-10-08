package files

import (
	"fmt"

	"github.com/danvergara/morphos/pkg/files/documents"
	"github.com/danvergara/morphos/pkg/files/ebooks"
)

// DocumentFactory implements the FileFactory interface.
type DocumentFactory struct {
	filename string
}

func NewDocumentFactory(filename string) *DocumentFactory {
	return &DocumentFactory{filename: filename}
}

// NewFile method returns an object that implements the File interface,
// given a document format as input.
// If not supported, it will error out.
func (d *DocumentFactory) NewFile(f string) (File, error) {
	switch f {
	case documents.PDF:
		return documents.NewPdf(d.filename), nil
	case documents.DOCX, documents.DOCXMIMEType:
		return documents.NewDocx(d.filename), nil
	case documents.XLSX, documents.XLSXMIMEType:
		return documents.NewXlsx(d.filename), nil
	case documents.CSV:
		return documents.NewCsv(d.filename), nil
	case ebooks.EpubMimeType, ebooks.EPUB:
		return ebooks.NewEpub(d.filename), nil
	case ebooks.MobiMimeType, ebooks.MOBI:
		return ebooks.NewMobi(d.filename), nil
	default:
		return nil, fmt.Errorf("type file  %s not recognized", f)
	}
}
