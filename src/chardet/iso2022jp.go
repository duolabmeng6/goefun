package chardet

var dict_ij = map[uint32]int{
	0x0000244E: 0x00, // の
	0x0000213C: 0x01, // ー
	0x0000244B: 0x02, // に
	0x00002573: 0x03, // ン
	0x00002472: 0x04, // を
	0x00002424: 0x05, // い
	0x00002447: 0x06, // で
	0x0000243F: 0x07, // た
	0x0000242C: 0x08, // が
	0x00002437: 0x09, // し
	0x0000246B: 0x0A, // る
	0x0000244F: 0x0B, // は
	0x0000467C: 0x0C, // 日
	0x0000244A: 0x0D, // な
	0x00002539: 0x0E, // ス
	0x00002446: 0x0F, // て
	0x00002448: 0x10, // と
	0x00002439: 0x11, // す
	0x0000245E: 0x12, // ま
	0x2437243F: 0x13, // した
	0x24392123: 0x14, // す。
	0x245E2439: 0x15, // ます
	0x24462424: 0x16, // てい
	0x2439246B: 0x17, // する
	0x24372446: 0x18, // して
	0x242B2469: 0x19, // から
	0x3F374A39: 0x1A, // 新聞
	0x2443243F: 0x1B, // った
	0x243F2123: 0x1C, // た。
	0x467C4B5C: 0x1D, // 日本
	0x2565213C: 0x1E, // ュー
	0x2435246C: 0x1F, // され
	0x213C2539: 0x20, // ース
	0x24432446: 0x21, // って
	0x24352473: 0x22, // さん
	0x254B2565: 0x23, // ニュ
	0x25242573: 0x24, // イン
	0x244A2424: 0x25, // ない
	0x2424246B: 0x26, // いる
	0x25732548: 0x27, // ント
	0x2424243F: 0x28, // いた
	0x244A2449: 0x29, // など
	0x213C2548: 0x2A, // ート
	0x352D3B76: 0x2B, // 記事
	0x2543252F: 0x2C, // ック
	0x253F213C: 0x2D, // ター
	0x213C256B: 0x2E, // ール
	0x25332573: 0x2F, // コン
	0x253F256B: 0x30, // タル
	0x25472538: 0x31, // デジ
}

type iso2022JP struct {
	byte
	last rune
	curr rune
	hold [50]int
	ttls int
	ttld int
}

func (i iso2022JP) String() string {
	return "iso-2022-jp"
}

func (i *iso2022JP) Feed(x byte) bool {
	if i.byte&15 == 0 {
		if x == 0x1B {
			i.byte = 4 << 4
			return true
		}
		switch i.byte >> 4 {
		case 0:
			if (x >= 0x00 && x <= 0x7F) || (x >= 0xA1 && x <= 0xDF) {
				return true
			}
		case 1:
			if x >= 0x21 && x <= 0x7E {
				i.byte++
				return true
			}
		case 2:
			if x >= 0x21 && x <= 0x7E {
				i.byte++
				i.curr = rune(x) << 8
				return true
			}
		case 4:
			if x == '(' {
				i.byte |= 1
				return true
			}
			if x == '$' {
				i.byte |= 2
				return true
			}
		}
	} else {
		switch i.byte >> 4 {
		case 1:
			if x >= 0x21 && x <= 0x7E {
				i.byte = 1 << 4
				return true
			}
		case 2:
			if x >= 0x21 && x <= 0x7E {
				i.byte = 2 << 4
				i.curr |= rune(x)
				i.count()
				return true
			}
		case 4:
			switch i.byte & 15 {
			case 1:
				if x == 'B' {
					i.byte = 0
					return true
				}
				if x == 'J' {
					i.byte = 1 << 4
					return true
				}
			case 2:
				if x == '@' {
					i.byte = 1 << 4
					return true
				}
				if x == 'B' {
					i.byte = 2 << 4
					return true
				}
			}
		}
	}
	return false
}

func (i *iso2022JP) Priority() float64 {
	f := 0.0
	if i.ttls > 0 {
		for t, x := range i.hold[:19] {
			k := 100*float64(x)/float64(i.ttls) - freq_jp[t]
			if k >= 0 {
				f += k
			} else {
				f -= k
			}
		}
	} else {
		f += 100
	}
	if i.ttld > 0 {
		for t, x := range i.hold[19:] {
			k := 100*float64(x)/float64(i.ttls) - freq_jp[t+19]
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

func (i *iso2022JP) count() {
	if j, ok := dict_ij[uint32(i.curr)]; ok {
		i.hold[j]++
		i.ttls++
	}
	if i.last > 0 {
		if j, ok := dict_ij[uint32(i.last<<16)|uint32(i.curr)]; ok {
			i.hold[j]++
			i.ttld++
		}
	}
	i.last = i.curr
}
