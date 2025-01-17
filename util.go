package gimu

import (
	"image"
	"image/color"
	"image/draw"
	"unsafe"

	"github.com/coalaura/gimu/nk"
)

func toNkFlag(align string) nk.Flags {
	var a nk.Flags
	switch align {
	case "RT":
		a = nk.TextAlignRight | nk.TextAlignTop
	case "RC":
		a = nk.TextAlignRight | nk.TextAlignMiddle
	case "RB":
		a = nk.TextAlignRight | nk.TextAlignBottom
	case "CT":
		a = nk.TextAlignCentered | nk.TextAlignTop
	case "CC":
		a = nk.TextAlignCentered | nk.TextAlignMiddle
	case "CB":
		a = nk.TextAlignCentered | nk.TextAlignBottom
	case "LT":
		a = nk.TextAlignLeft | nk.TextAlignTop
	case "LB":
		a = nk.TextAlignLeft | nk.TextAlignBottom
	case "LC":
		a = nk.TextAlignLeft | nk.TextAlignMiddle
	default:
		a = nk.TextAlignLeft | nk.TextAlignMiddle
	}

	return a
}

func toNkColor(c color.RGBA) nk.Color {
	nc := nk.NewColor()
	nc.SetRGBA(nk.Byte(c.R), nk.Byte(c.G), nk.Byte(c.B), nk.Byte(c.A))
	return *nc
}

func toNkRune(r rune) nk.Rune {
	return *(*nk.Rune)(unsafe.Pointer(&r))
}

func toGoRune(r nk.Rune) rune {
	return *(*rune)(unsafe.Pointer(&r))
}

func toNkPluginFilter(f EditFilter) func(*nk.TextEdit, nk.Rune) int32 {
	return func(nt *nk.TextEdit, r nk.Rune) int32 {
		result := f(toGoRune(r))
		if result {
			return 1
		}
		return 0
	}
}

func toInt32(b bool) int32 {
	if b {
		return 1
	}
	return 0
}

func getDynamicWidth(ctx *nk.Context) float32 {
	bounds := nk.NkLayoutWidgetBounds(ctx)
	padding := ctx.Style().Window().Padding().X()
	return bounds.W() - (padding * 2)
}

func ImgToRgba(img image.Image) *image.RGBA {
	switch trueim := img.(type) {
	case *image.RGBA:
		return trueim
	default:
		copy := image.NewRGBA(trueim.Bounds())
		draw.Draw(copy, trueim.Bounds(), trueim, image.Pt(0, 0), draw.Src)
		return copy
	}
}

type Texture struct {
	image nk.Image
}

func RgbaToTexture(rgba *image.RGBA) *Texture {
	var textureId uint32
	return &Texture{
		image: nk.RgbaToNkImage(&textureId, rgba),
	}
}

func LoadDefaultFont() *nk.Font {
	atlas := nk.NewFontAtlas()
	nk.NkFontStashBegin(&atlas)
	nk.NkFontStashEnd()

	return atlas.DefaultFont()
}

func LoadFontFromFile(filePath string, size float32, config *nk.FontConfig) *nk.Font {
	atlas := nk.NewFontAtlas()
	nk.NkFontStashBegin(&atlas)
	f := nk.NkFontAtlasAddFromFile(atlas, filePath, size, config)
	nk.NkFontStashEnd()

	return f
}

func SetFont(ctx *nk.Context, font *nk.Font) {
	nk.NkStyleSetFont(ctx, font.Handle())
}
