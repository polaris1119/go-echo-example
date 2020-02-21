/*
 * Copyright (c) 2020. The StudyGolang Authors. All rights reserved.
 * Use of this source code is governed by a Apache 2.0
 * license that can be found in the LICENSE file.
 * https://studygolang.com
 * Author:polaris	polaris@studygolang.com
 */

// 进行 Binder 讲解之前，介绍客户端参数服务端如何处理
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		req.ParseMultipartForm(32 << 20)

		data := map[string]interface{}{
			"form":      req.Form,
			"post_form": req.PostForm,
		}

		reqBody, err := ioutil.ReadAll(req.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		data["json_data"] = string(reqBody)

		fmt.Fprintln(w, data)
	})

	log.Fatal(http.ListenAndServe(":2020", nil))
}
