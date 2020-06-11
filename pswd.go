// Copyright 2020 Cango Author.

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at

//    http://www.apache.org/licenses/LICENSE-2.0

// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"
)

type char struct {
	chars *[]rune
	start rune
	size  int
}

var numbers = char{chars: new([]rune), start: '0', size: 10}
var alphabets = char{chars: new([]rune), start: 'a', size: 26}
var capitalAlphabets = char{chars: new([]rune), start: 'A', size: 26}
var symbols = char{chars: &[]rune{'%', '(', ')', '-', '=', '@'}}

var passwordLength = flag.Int("l", 8, "密码长度")
var userChars = flag.String("c", "", "自定义字符")
var symbolsFlag = flag.Bool("s", false, "是否使用特殊字符%、(、)、-、=、@")

func main() {
	flag.Parse()
	for _, v := range []char{numbers, alphabets, capitalAlphabets} {
		for i := v.start; i < v.start+rune(v.size); i++ {
			*v.chars = append(*(v.chars), i)
		}
	}
	chars := append(*numbers.chars, append(*alphabets.chars, append(*capitalAlphabets.chars, append(func() []rune {
		if *symbolsFlag {
			return *symbols.chars
		}
		return []rune{}
	}(), []rune(*userChars)...)...)...)...)
	for *passwordLength > len(chars) {
		chars = append(chars, chars...)
	}
	rand.New(rand.NewSource(time.Now().UnixNano())).Shuffle(len(chars), func(i, j int) { chars[i], chars[j] = chars[j], chars[i] })
	fmt.Println(string(chars[0:*passwordLength]))
}
