// Copyright 2019 gf Author(https://github.com/gogf/gf). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.

// go test *.go -bench=".*"

package os

import (
	"os"
	"testing"

	"github.com/gogf/gf/crypto/gmd5"
	"github.com/gogf/gf/test/gtest"
)

var (
	s = "pibigstar"
	// online generated MD5 value
	result = "d175a1ff66aedde64344785f7f7a3df8"
)

type user2 struct {
	name     string
	password string
	age      int
}

func TestEncryptmd5(t *testing.T) {
	gtest.Case(t, func() {
		encryptString, _ := Md5加密(s)
		gtest.Assert(encryptString, result)

		result := "1427562bb29f88a1161590b76398ab72"
		encrypt, _ := Md5加密(123456)
		gtest.AssertEQ(encrypt, result)
	})

	gtest.Case(t, func() {
		user2 := &user2{
			name:     "派大星",
			password: "123456",
			age:      23,
		}
		result := "70917ebce8bd2f78c736cda63870fb39"
		encrypt, _ := Md5加密(user2)
		gtest.AssertEQ(encrypt, result)
	})
}

func TestEncryptString(t *testing.T) {
	gtest.Case(t, func() {
		encryptString, _ := gmd5.EncryptString(s)
		gtest.Assert(encryptString, result)
	})
}

func TestEncryptFilemd5(t *testing.T) {
	path := "test.text"
	errorPath := "err.txt"
	result := "e6e6e1cd41895beebff16d5452dfce12"
	gtest.Case(t, func() {
		file, err := os.Create(path)
		defer os.Remove(path)
		defer file.Close()
		gtest.Assert(err, nil)
		_, _ = file.Write([]byte("Hello Go Frame"))
		encryptFile, _ := Md5加密从文件(path)
		gtest.AssertEQ(encryptFile, result)
		// when the file is not exist,encrypt will return empty string
		errEncrypt, _ := Md5加密从文件(errorPath)
		gtest.AssertEQ(errEncrypt, "")
	})

}
