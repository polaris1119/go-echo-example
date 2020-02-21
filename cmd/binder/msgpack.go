/*
 * Copyright (c) 2020. The StudyGolang Authors. All rights reserved.
 * Use of this source code is governed by a Apache 2.0
 * license that can be found in the LICENSE file.
 * https://studygolang.com
 * Author:polaris	polaris@studygolang.com
 */

// 简单 msgpack 请求客户端
package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/vmihailenco/msgpack"
)

func main() {
	type User struct {
		Name string
		Sex  string
	}

	b, err := msgpack.Marshal(&User{Name: "xuxinhua", Sex: "male"})
	if err != nil {
		panic(err)
	}

	resp, err := http.DefaultClient.Post("http://localhost:2020/", "application/msgpack", bytes.NewReader(b))
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	result, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", result)
}
