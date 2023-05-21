package chardet

var dict_ej = map[uint32]int{
	0x0000A4CE: 0x00, // の
	0x0000A1BC: 0x01, // ー
	0x0000A4CB: 0x02, // に
	0x0000A5F3: 0x03, // ン
	0x0000A4F2: 0x04, // を
	0x0000A4A4: 0x05, // い
	0x0000A4C7: 0x06, // で
	0x0000A4BF: 0x07, // た
	0x0000A4AC: 0x08, // が
	0x0000A4B7: 0x09, // し
	0x0000A4EB: 0x0A, // る
	0x0000A4CF: 0x0B, // は
	0x0000C6FC: 0x0C, // 日
	0x0000A4CA: 0x0D, // な
	0x0000A5B9: 0x0E, // ス
	0x0000A4C6: 0x0F, // て
	0x0000A4C8: 0x10, // と
	0x0000A4B9: 0x11, // す
	0x0000A4DE: 0x12, // ま
	0xA4B7A4BF: 0x13, // した
	0xA4B9A1A3: 0x14, // す。
	0xA4DEA4B9: 0x15, // ます
	0xA4C6A4A4: 0x16, // てい
	0xA4B9A4EB: 0x17, // する
	0xA4B7A4C6: 0x18, // して
	0xA4ABA4E9: 0x19, // から
	0xBFB7CAB9: 0x1A, // 新聞
	0xA4C3A4BF: 0x1B, // った
	0xA4BFA1A3: 0x1C, // た。
	0xC6FCCBDC: 0x1D, // 日本
	0xA5E5A1BC: 0x1E, // ュー
	0xA4B5A4EC: 0x1F, // され
	0xA1BCA5B9: 0x20, // ース
	0xA4C3A4C6: 0x21, // って
	0xA4B5A4F3: 0x22, // さん
	0xA5CBA5E5: 0x23, // ニュ
	0xA5A4A5F3: 0x24, // イン
	0xA4CAA4A4: 0x25, // ない
	0xA4A4A4EB: 0x26, // いる
	0xA5F3A5C8: 0x27, // ント
	0xA4A4A4BF: 0x28, // いた
	0xA4CAA4C9: 0x29, // など
	0xA1BCA5C8: 0x2A, // ート
	0xB5ADBBF6: 0x2B, // 記事
	0xA5C3A5AF: 0x2C, // ック
	0xA5BFA1BC: 0x2D, // ター
	0xA1BCA5EB: 0x2E, // ール
	0xA5B3A5F3: 0x2F, // コン
	0xA5BFA5EB: 0x30, // タル
	0xA5C7A5B8: 0x31, // デジ
}

// [\x00-\x7F]
// [\xA1-\xFE]{2}
// \x8F[\xA1-\xFE]{2}
// \x8E[\xA1-\xDF]
type eucJP struct {
	byte
	last rune
	curr rune
	hold [50]int
	ttls int
	ttld int
}

func (e eucJP) String() string {
	return "euc-jp"
}

func (e *eucJP) Feed(x byte) bool {
	switch e.byte {
	case 0:
		if x >= 0x00 && x <= 0x7F {
			return true
		}
		if x >= 0xA1 && x <= 0xFE {
			e.byte = 1
			e.curr = rune(x) << 8
			return true
		}
		if x == 0x8F {
			e.byte = 2
			return true
		}
		if x == 0x8E {
			e.byte = 4
			return true
		}
	case 1:
		if x >= 0xA1 && x <= 0xFE {
			e.byte = 0
			e.curr |= rune(x)
			e.count()
			return true
		}
	case 2:
		if x >= 0xA1 && x <= 0xFE {
			e.byte = 3
			return true
		}
	case 3:
		if x >= 0xA1 && x <= 0xFE {
			e.byte = 0
			return true
		}
	default:
		if x >= 0xA1 && x <= 0xDE {
			e.byte = 0
			return true
		}
	}
	return false
}

func (e *eucJP) Priority() float64 {
	f := 0.0
	if e.ttls > 0 {
		for i, x := range e.hold[:19] {
			k := 100*float64(x)/float64(e.ttls) - freq_jp[i]
			if k >= 0 {
				f += k
			} else {
				f -= k
			}
		}
	} else {
		f += 100
	}
	if e.ttld > 0 {
		for i, x := range e.hold[19:] {
			k := 100*float64(x)/float64(e.ttls) - freq_jp[i+19]
			if k >= 0 {
				f += k
			} else {
				f -= k
			}
		}
	} else {
		f += 100
	}
	return 100 - f/2
}

func (e *eucJP) count() {
	if i, ok := dict_ej[uint32(e.curr)]; ok {
		e.hold[i]++
		e.ttls++
	}
	if e.last > 0 {
		if i, ok := dict_ej[uint32(e.last<<16)|uint32(e.curr)]; ok {
			e.hold[i]++
			e.ttld++
		}
	}
	e.last = e.curr
}
