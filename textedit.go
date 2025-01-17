package gimu

import (
	"github.com/coalaura/gimu/nk"
)

type TextEdit struct {
	editor *nk.TextEdit
}

func NewTextEdit() *TextEdit {
	edt := &nk.TextEdit{}
	nk.NkTexteditInitDefault(edt)

	return &TextEdit{editor: edt}
}

type EditFilter func(rune) bool

func EditFilterDefault(r rune) bool {
	return true
}

func EditFilterAscii(r rune) bool {
	return nk.NkFilterAscii(nil, toNkRune(r)) > 0
}

func EditFilterFloat(r rune) bool {
	return nk.NkFilterFloat(nil, toNkRune(r)) > 0
}

func EditFilterDecimal(r rune) bool {
	return nk.NkFilterDecimal(nil, toNkRune(r)) > 0
}

func EditFilterHex(r rune) bool {
	return nk.NkFilterHex(nil, toNkRune(r)) > 0
}

func EditFilterOct(r rune) bool {
	return nk.NkFilterOct(nil, toNkRune(r)) > 0
}

func EditFilterBinary(r rune) bool {
	return nk.NkFilterBinary(nil, toNkRune(r)) > 0
}

func (t *TextEdit) Edit(w *Window, flag nk.Flags, filter EditFilter) {
	nk.NkEditBuffer(w.ctx, flag, t.editor, toNkPluginFilter(filter))
}

func (t *TextEdit) GetString() string {
	return t.editor.GetGoString()
}
