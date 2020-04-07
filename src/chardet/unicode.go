package chardet

// [\x00-\x7F]
// [\xC0-\xDF][\x80-\xBF]
// [\xE0-\xEF][\x80-\xBF]{2}
// [\xF0-\xF7][\x80-\xBF]{3}
type utf8 struct {
	byte
}

func (u utf8) String() string {
	return "utf-8"
}

func (u *utf8) Feed(x byte) bool {
	if u.byte == 0 {
		if x >= 0x00 && x <= 0x7F {
			return true
		}
		if x >= 0xC0 && x <= 0xDF {
			u.byte = 1
			return true
		}
		if x >= 0xE0 && x <= 0xEF {
			u.byte = 2
			return true
		}
		if x >= 0xF0 && x <= 0xF7 {
			u.byte = 3
			return true
		}
	} else {
		if x >= 0x80 && x <= 0xBF {
			u.byte -= 1
			return true
		}
	}
	return false
}

func (u utf8) Priority() float64 {
	return 0
}

// [\x00-\xD7\xE0-\xFF][\x00-\xFF]
// [\xD8-\xDB][\x00-\xFF][\xDC-\DF][\x00-\xFF]
type utf16BE struct {
	byte
}

func (u utf16BE) String() string {
	return "utf-16be"
}

func (u *utf16BE) Feed(x byte) bool {
	switch u.byte {
	case 0:
		if (x >= 0x00 && x <= 0xD7) || (x >= 0xE0 && x <= 0xFF) {
			u.byte = 1
			return true
		}
		if x >= 0xD8 && x <= 0xDB {
			u.byte = 2
			return true
		}
	case 1:
		u.byte = 0
		return true
	case 2:
		u.byte = 3
		return true
	default:
		if x >= 0xDC && x <= 0xDF {
			u.byte = 1
			return true
		}
	}
	return false
}

func (u utf16BE) Priority() float64 {
	return 0
}

// [\x00-\xFF][\x00-\xD7\xE0-\xFF]
// [\x00-\xFF][\xD8-\xDB][\x00-\xFF][\xDC-\DF]
type utf16LE struct {
	byte
}

func (u utf16LE) String() string {
	return "utf-16le"
}

func (u *utf16LE) Feed(x byte) bool {
	switch u.byte {
	case 0:
		u.byte = 1
		return true
	case 1:
		if (x >= 0x00 && x <= 0xD7) || (x >= 0xE0 && x <= 0xFF) {
			u.byte = 0
			return true
		}
		if x >= 0xD8 && x <= 0xDB {
			u.byte = 2
			return true
		}
	case 2:
		u.byte = 3
		return true
	default:
		if x >= 0xDC && x <= 0xDF {
			u.byte = 0
			return true
		}
	}
	return false
}

func (u utf16LE) Priority() float64 {
	return 0
}

// \x00[\x00-\x0F][\x00-\xFF]{2}
type utf32BE struct {
	byte
}

func (u utf32BE) String() string {
	return "utf-32be"
}

func (u *utf32BE) Feed(x byte) bool {
	switch u.byte {
	case 0:
		if x == 0x00 {
			u.byte = 1
			return true
		}
	case 1:
		if x >= 0x00 && x <= 0x1F {
			u.byte = 2
			return true
		}
	case 2:
		u.byte = 3
		return true
	default:
		u.byte = 0
		return true
	}
	return false
}

func (u utf32BE) Priority() float64 {
	return 0
}

// [\x00-\xFF]{2}[\x00-\x0F]\x00
type utf32LE struct {
	byte
}

func (u utf32LE) String() string {
	return "utf-32le"
}

func (u *utf32LE) Feed(x byte) bool {
	switch u.byte {
	case 0:
		u.byte = 1
		return true
	case 1:
		u.byte = 2
		return true
	case 2:
		if x >= 0x00 && x <= 0x1F {
			u.byte = 3
			return true
		}
	default:
		if x == 0x00 {
			u.byte = 0
			return true
		}
	}
	return false
}

func (u utf32LE) Priority() float64 {
	return 0
}
