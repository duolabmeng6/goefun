package chardet

var dict_kr = map[uint32]int{
	0x0000C0CC: 0x00, // 이
	0x0000B4D9: 0x01, // 다
	0x0000B4C2: 0x02, // 는
	0x0000C0C7: 0x03, // 의
	0x0000BFA1: 0x04, // 에
	0x0000B0A1: 0x05, // 가
	0x0000C1F6: 0x06, // 지
	0x0000B0ED: 0x07, // 고
	0x0000B1E2: 0x08, // 기
	0x0000C7D1: 0x09, // 한
	0x0000C0BB: 0x0A, // 을
	0x0000B7CE: 0x0B, // 로
	0x0000C7CF: 0x0C, // 하
	0x0000BBE7: 0x0D, // 사
	0x0000BEC6: 0x0E, // 아
	0x0000C0CE: 0x0F, // 인
	0x0000B4EB: 0x10, // 대
	0x0000BCAD: 0x11, // 서
	0xB4CFB4D9: 0x12, // 니다
	0xC0B8B7CE: 0x13, // 으로
	0xBFA1BCAD: 0x14, // 에서
	0xC7CFB4C2: 0x15, // 하는
	0xC7DFB4D9: 0x16, // 했다
	0xBDC0B4CF: 0x17, // 습니
	0xC0D6B4D9: 0x18, // 있다
	0xC7D1B1B9: 0x19, // 한국
	0xC7CFB0ED: 0x1A, // 하고
	0xC0CCB6F3: 0x1B, // 이라
	0xC0CCB4D9: 0x1C, // 이다
	0xB4EBC5EB: 0x1D, // 대통
	0xC5EBB7C9: 0x1E, // 통령
	0xB1EEC1F6: 0x1F, // 까지
	0xC1F6B8B8: 0x20, // 지만
	0xC7D1B4D9: 0x21, // 한다
	0xC0D6B4C2: 0x22, // 있는
	0xB5E9C0CC: 0x23, // 들이
	0xBFF9C8A3: 0x24, // 월호
	0xBEC6C0CC: 0x25, // 아이
	0xBCBCBFF9: 0x26, // 세월
	0xB4D9B0ED: 0x27, // 다고
	0xB6F3B0ED: 0x28, // 라고
	0xBACEC5CD: 0x29, // 부터
	0xC7CFC1F6: 0x2A, // 하지
	0xB4D9B4C2: 0x2B, // 다는
	0xB1E2C0DA: 0x2C, // 기자
	0xBEFAB4D9: 0x2D, // 었다
	0xC7D4B2B2: 0x2E, // 함께
	0xB6F3B4C2: 0x2F, // 라는
	0xBFA1B0D4: 0x30, // 에게
	0xBCADBFEF: 0x31, // 서울
}

// [\x00-\x7F]
// [\xA1-\xFE]{2}
type eucKR struct {
	byte
	last rune
	curr rune
	hold [50]int
	ttls int
	ttld int
}

func (e eucKR) String() string {
	return "euc-kr"
}

func (e *eucKR) Feed(x byte) bool {
	if e.byte == 0 {
		if x >= 0x00 && x <= 0x7F {
			return true
		}
		if x >= 0xA1 && x <= 0xFE {
			e.byte = 1
			e.curr = rune(x) << 8
			return true
		}
	} else {
		if x >= 0xA1 && x <= 0xFE {
			e.byte = 0
			e.curr |= rune(x)
			e.count()
			return true
		}
	}
	return false
}

func (e *eucKR) Priority() float64 {
	f := 0.0
	if e.ttls > 0 {
		for i, x := range e.hold[:18] {
			k := 100*float64(x)/float64(e.ttls) - freq_kr[i]
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
		for i, x := range e.hold[18:] {
			k := 100*float64(x)/float64(e.ttls) - freq_kr[i+18]
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

func (e *eucKR) count() {
	if i, ok := dict_kr[uint32(e.curr)]; ok {
		e.hold[i]++
		e.ttls++
	}
	if e.last > 0 {
		if i, ok := dict_kr[uint32(e.last<<16)|uint32(e.curr)]; ok {
			e.hold[i]++
			e.ttld++
		}
	}
	e.last = e.curr
}
