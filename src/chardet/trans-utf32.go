package chardet

import (
	"golang.org/x/text/transform"
	. "unicode/utf8"
)

// 提供给用户使用的并发安全的编解码器接口
var (
	UTF32BE = &utf32be{}
	UTF32LE = &utf32le{}
)

type u32bt8 struct {
	hold []byte
}

// UTF-16BE => UTF-8
func (u *u32bt8) Transform(dst, src []byte, atEOF bool) (nDst, nSrc int, err error) {
	ld, ls := len(dst), len(src)
	us := []rune{}
	for nSrc < ls {
		if nSrc+3 >= ls {
			if !atEOF {
				err = transform.ErrShortSrc
				break
			}
			us = append(us, 0xFFFD)
			nSrc++
		} else {
			a := src[nSrc]
			nSrc++
			b := src[nSrc]
			nSrc++
			c := src[nSrc]
			nSrc++
			d := src[nSrc]
			nSrc++
			us = append(us, (rune(a)<<24)|(rune(b)<<16)|(rune(c)<<8)|rune(d))
		}
	}
	if l := len(us); l > 0 {
		u.hold = append(u.hold, string(us)...)
	}
	lu := len(u.hold)
	copy(dst, u.hold)
	if lu > ld {
		u.hold = u.hold[ld:]
		return ld, nSrc, transform.ErrShortDst
	}
	nDst = lu
	return
}

func (u *u32bt8) Reset() {
	u.hold = nil
}

type u32lt8 struct {
	hold []byte
}

// UTF-16BE => UTF-8
func (u *u32lt8) Transform(dst, src []byte, atEOF bool) (nDst, nSrc int, err error) {
	ld, ls := len(dst), len(src)
	us := []rune{}
	for nSrc < ls {
		if nSrc+3 >= ls {
			if !atEOF {
				err = transform.ErrShortSrc
				break
			}
			us = append(us, 0xFFFD)
			nSrc++
		} else {
			a := src[nSrc]
			nSrc++
			b := src[nSrc]
			nSrc++
			c := src[nSrc]
			nSrc++
			d := src[nSrc]
			nSrc++
			us = append(us, (rune(d)<<24)|(rune(c)<<16)|(rune(b)<<8)|rune(a))
		}
	}
	if l := len(us); l > 0 {
		u.hold = append(u.hold, string(us)...)
	}
	lu := len(u.hold)
	copy(dst, u.hold)
	if lu > ld {
		u.hold = u.hold[ld:]
		return ld, nSrc, transform.ErrShortDst
	}
	nDst = lu
	return
}

func (u *u32lt8) Reset() {
	u.hold = nil
}

type u8t32b struct {
	hold []byte
}

// UTF-8 => UTF-16BE
func (u *u8t32b) Transform(dst, src []byte, atEOF bool) (nDst, nSrc int, err error) {
	ld, ls := len(dst), len(src)
	rs := []rune{}
	for nSrc < ls {
		r, n := DecodeRune(src[nSrc:])
		if n == 1 && r == 0xFFFD {
			if nSrc+1 >= ls {
				if !atEOF {
					err = transform.ErrShortSrc
					break
				}
			}
		}
		rs = append(rs, r)
		nSrc += n
	}
	for _, r := range rs {
		a := byte((r >> 0x18) & 255)
		b := byte((r >> 0x10) & 255)
		c := byte((r >> 0x08) & 255)
		d := byte((r >> 0x00) & 255)
		u.hold = append(u.hold, a, b, c, d)
	}
	lu := len(u.hold)
	copy(dst, u.hold)
	if lu > ld {
		u.hold = u.hold[ld:]
		return ld, nSrc, transform.ErrShortDst
	}
	nDst = lu
	return
}

func (u *u8t32b) Reset() {
	u.hold = nil
}

type u8t32l struct {
	hold []byte
}

// UTF-8 => UTF-16BE
func (u *u8t32l) Transform(dst, src []byte, atEOF bool) (nDst, nSrc int, err error) {
	ld, ls := len(dst), len(src)
	rs := []rune{}
	for nSrc < ls {
		r, n := DecodeRune(src[nSrc:])
		if n == 1 && r == 0xFFFD {
			if nSrc+1 >= ls {
				if !atEOF {
					err = transform.ErrShortSrc
					break
				}
			}
		}
		rs = append(rs, r)
		nSrc += n
	}
	for _, r := range rs {
		a := byte((r >> 0x18) & 255)
		b := byte((r >> 0x10) & 255)
		c := byte((r >> 0x08) & 255)
		d := byte((r >> 0x00) & 255)
		u.hold = append(u.hold, d, c, b, a)
	}
	lu := len(u.hold)
	copy(dst, u.hold)
	if lu > ld {
		u.hold = u.hold[ld:]
		return ld, nSrc, transform.ErrShortDst
	}
	nDst = lu
	return
}

func (u *u8t32l) Reset() {
	u.hold = nil
}

type utf32le struct{}

// 返回解码小端在前的UTF-16到UTF-8的解码器
func (u utf32le) NewDecoder() transform.Transformer {
	return new(u32lt8)
}

// 返回编码UTF-8到小端在前的UTF-16的编码器
func (u utf32le) NewEncoder() transform.Transformer {
	return new(u8t32l)
}

type utf32be struct{}

// 返回编码UTF-8到大端在前的UTF-16的解码器
func (u utf32be) NewDecoder() transform.Transformer {
	return new(u32bt8)
}

// 返回解码大端在前的UTF-16到UTF-8的解码器
func (u utf32be) NewEncoder() transform.Transformer {
	return new(u8t32b)
}
