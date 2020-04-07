package chardet

type hzgb2312 struct {
	byte
	rune
	hold [80]int
	ttls int
}

func (h hzgb2312) String() string {
	return "hz-gb2312"
}

func (h *hzgb2312) Feed(x byte) bool {
	switch h.byte {
	case 0:
		if x == '~' {
			h.byte = 1
			return true
		}
		if x >= 0x00 && x <= 0x7F {
			return true
		}
	case 1:
		if x == '~' {
			h.byte = 0
			return true
		}
		if x == '{' {
			h.byte = 2
			return true
		}
	case 2:
		if x == '~' {
			h.byte = 3
			return true
		}
		if x >= 0x21 && x <= 0x77 {
			h.byte = 4
			h.rune = (rune(x) | 0x80) << 8
			return true
		}
	case 3:
		if x == '}' {
			h.byte = 0
			return true
		}
	case 4:
		if x >= 0x21 && x <= 0x7E {
			h.byte = 2
			h.rune |= rune(x) | 0x80
			h.count()
			return true
		}
	}
	return false
}

func (h *hzgb2312) Priority() float64 {
	if h.ttls == 0 {
		return 0
	}
	f := 0.0
	for i, x := range h.hold {
		k := 100*float64(x)/float64(h.ttls) - freq_ch[i]
		if k >= 0 {
			f += k
		} else {
			f -= k
		}
	}
	return 100 - f
}

func (h *hzgb2312) count() {
	if i, ok := dict_gb[uint32(h.rune)]; ok {
		h.hold[i]++
		h.ttls++
	}
}
