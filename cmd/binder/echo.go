/*
 * Copyright (c) 2020. The StudyGolang Authors. All rights reserved.
 * Use of this source code is governed by a Apache 2.0
 * license that can be found in the LICENSE file.
 * https://studygolang.com
 * Author:polaris	polaris@studygolang.com
 */

// echo 自定义 Binder 实现
package main

import (
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/vmihailenco/msgpack/v4"
)

type User struct {
	Name string `query:"name" form:"name"  xml:"name"`
	Sex  string `query:"sex" form:"sex" json:"sex" xml:"sex"`
}

func main() {
	e := echo.New()

	e.Binder = new(MsgpackBinder)

	e.Any("/", func(ctx echo.Context) error {
		user := new(User)
		if err := ctx.Bind(user); err != nil {
			return err
		}

		return ctx.JSON(http.StatusOK, user)
	})

	e.Logger.Fatal(e.Start(":2020"))
}

type MsgpackBinder struct{}

func (b *MsgpackBinder) Bind(i interface{}, ctx echo.Context) (err error) {
	// 也支持默认 Binder 相关的绑定
	db := new(echo.DefaultBinder)
	if err = db.Bind(i, ctx); err != echo.ErrUnsupportedMediaType {
		return
	}

	req := ctx.Request()
	ctype := req.Header.Get(echo.HeaderContentType)
	if strings.HasPrefix(ctype, echo.MIMEApplicationMsgpack) {
		if err = msgpack.NewDecoder(req.Body).Decode(i); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error()).SetInternal(err)
		}

		return
	}

	return echo.ErrUnsupportedMediaType
}
