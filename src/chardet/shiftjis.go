package chardet

var dict_sj = map[uint32]int{
	0x000082CC: 0x00, // の
	0x0000815B: 0x01, // ー
	0x000082C9: 0x02, // に
	0x00008393: 0x03, // ン
	0x000082F0: 0x04, // を
	0x000082A2: 0x05, // い
	0x000082C5: 0x06, // で
	0x000082BD: 0x07, // た
	0x000082AA: 0x08, // が
	0x000082B5: 0x09, // し
	0x000082E9: 0x0A, // る
	0x000082CD: 0x0B, // は
	0x000093FA: 0x0C, // 日
	0x000082C8: 0x0D, // な
	0x00008358: 0x0E, // ス
	0x000082C4: 0x0F, // て
	0x000082C6: 0x10, // と
	0x000082B7: 0x11, // す
	0x000082DC: 0x12, // ま
	0x82B582BD: 0x13, // した
	0x82B78142: 0x14, // す。
	0x82DC82B7: 0x15, // ます
	0x82C482A2: 0x16, // てい
	0x82B782E9: 0x17, // する
	0x82B582C4: 0x18, // して
	0x82A982E7: 0x19, // から
	0x905695B7: 0x1A, // 新聞
	0x82C182BD: 0x1B, // った
	0x82BD8142: 0x1C, // た。
	0x93FA967B: 0x1D, // 日本
	0x8385815B: 0x1E, // ュー
	0x82B382EA: 0x1F, // され
	0x815B8358: 0x20, // ース
	0x82C182C4: 0x21, // って
	0x82B382F1: 0x22, // さん
	0x836A8385: 0x23, // ニュ
	0x83438393: 0x24, // イン
	0x82C882A2: 0x25, // ない
	0x82A282E9: 0x26, // いる
	0x83938367: 0x27, // ント
	0x82A282BD: 0x28, // いた
	0x82C882C7: 0x29, // など
	0x815B8367: 0x2A, // ート
	0x8B4C8E96: 0x2B, // 記事
	0x8362834E: 0x2C, // ック
	0x835E815B: 0x2D, // ター
	0x815B838B: 0x2E, // ール
	0x83528393: 0x2F, // コン
	0x835E838B: 0x30, // タル
	0x83668357: 0x31, // デジ
}

// [\x00-\x7F]
// [\xA1-\xDF]
// [\x81-\x9F\xE0-\xEF][\x40-\x7E\x80-\xFC]
type shiftJIS struct {
	byte
	last rune
	curr rune
	hold [50]int
	ttls int
	ttld int
}

func (s shiftJIS) String() string {
	return "shift-jis"
}

func (s *shiftJIS) Feed(x byte) bool {
	defer func() {
		if s.byte != 0 {
			return
		}
		if i, ok := dict_sj[uint32(s.curr)]; ok {
			s.hold[i]++
			s.ttls++
		}
		if s.last > 0 {
			if i, ok := dict_sj[uint32(s.last<<16)|uint32(s.curr)]; ok {
				s.hold[i]++
				s.ttld++
			}
		}
		s.last = s.curr
	}()
	if s.byte == 0 {
		if (x >= 0x00 && x <= 0x7F) || (x >= 0xA1 && x <= 0xDF) {
			s.curr = rune(x)
			return true
		}
		if (x >= 0x81 && x <= 0x9F) || (x >= 0xE0 && x <= 0xEF) {
			s.byte = 1
			s.curr = rune(x) << 8
			return true
		}
	} else {
		if (x >= 0x40 && x <= 0x7E) || (x >= 0x80 && x <= 0xFC) {
			s.byte = 0
			s.curr |= rune(x)
			return true
		}
	}
	return false
}

func (s *shiftJIS) Priority() float64 {
	f := 0.0
	if s.ttls > 0 {
		for i, x := range s.hold[:19] {
			k := 100*float64(x)/float64(s.ttls) - freq_jp[i]
			if k >= 0 {
				f += k
			} else {
				f -= k
			}
		}
	} else {
		f += 100
	}
	if s.ttld > 0 {
		for i, x := range s.hold[19:] {
			k := 100*float64(x)/float64(s.ttls) - freq_jp[i+19]
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
