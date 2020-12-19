// Code generated by "stringer -type=PenGraphic"; DO NOT EDIT.

package state

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[PenNormal-0]
	_ = x[PenBold-2]
	_ = x[PenFaint-4]
	_ = x[PenBlink-8]
	_ = x[PenConceal-16]
	_ = x[PenItalic-32]
	_ = x[PenFraktur-64]
	_ = x[PenUnderlineSingle-128]
	_ = x[PenUnderlineDouble-256]
	_ = x[PenReverse-512]
	_ = x[PenStrikeThrough-1024]
	_ = x[PenFramed-2048]
	_ = x[PenEncircled-4096]
	_ = x[PenOverlined-8192]
	_ = x[PenIntensity-6]
	_ = x[PenStyle-96]
	_ = x[PenUnderlineCurly-384]
	_ = x[PenUnderline-384]
	_ = x[PenWrapper-6144]
}

const _PenGraphic_name = "PenNormalPenBoldPenFaintPenIntensityPenBlinkPenConcealPenItalicPenFrakturPenStylePenUnderlineSinglePenUnderlineDoublePenUnderlineCurlyPenReversePenStrikeThroughPenFramedPenEncircledPenWrapperPenOverlined"

var _PenGraphic_map = map[PenGraphic]string{
	0:    _PenGraphic_name[0:9],
	2:    _PenGraphic_name[9:16],
	4:    _PenGraphic_name[16:24],
	6:    _PenGraphic_name[24:36],
	8:    _PenGraphic_name[36:44],
	16:   _PenGraphic_name[44:54],
	32:   _PenGraphic_name[54:63],
	64:   _PenGraphic_name[63:73],
	96:   _PenGraphic_name[73:81],
	128:  _PenGraphic_name[81:99],
	256:  _PenGraphic_name[99:117],
	384:  _PenGraphic_name[117:134],
	512:  _PenGraphic_name[134:144],
	1024: _PenGraphic_name[144:160],
	2048: _PenGraphic_name[160:169],
	4096: _PenGraphic_name[169:181],
	6144: _PenGraphic_name[181:191],
	8192: _PenGraphic_name[191:203],
}

func (i PenGraphic) String() string {
	if str, ok := _PenGraphic_map[i]; ok {
		return str
	}
	return "PenGraphic(" + strconv.FormatInt(int64(i), 10) + ")"
}
