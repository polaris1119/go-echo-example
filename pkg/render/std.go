/*
 * Copyright (c) 2020. The StudyGolang Authors. All rights reserved.
 * Use of this source code is governed by a Apache 2.0
 * license that can be found in the LICENSE file.
 * https://studygolang.com
 * Author:polaris	polaris@studygolang.com
 */

package render

import (
	"html/template"
	"io"
	"net/http"

	"github.com/labstack/echo/v4"
)

type layoutTemplate struct{}

var LayoutTemplate = &layoutTemplate{}

func (l *layoutTemplate) Render(w io.Writer, contentTpl string, data interface{}, ctx echo.Context) error {
	layout := "layout.html"

	// if data != nil {
	// 	if dataMap, ok := data.(map[string]interface{}); ok {
	// 		if layoutInter, ok := dataMap["layout"]; ok {
	// 			layout = layoutInter.(string)
	// 		}
	// 	}
	// }
	layoutInter := ctx.Get("layout")
	if layoutInter != nil {
		layout = layoutInter.(string)
	}

	tpl, err := template.New(layout).ParseFiles("template/common/"+layout, "template/"+contentTpl)
	if err != nil {
		return err
	}

	return tpl.Execute(w, data)
}

// NoNavRender 没有导航的 layout html 输出
func NoNavRender(ctx echo.Context, contentTpl string, data interface{}) error {
	// if data == nil {
	// 	data = make(map[string]interface{})
	// }
	// data["layout"] = "nonav_layout.html"

	ctx.Set("layout", "nonav_layout.html")

	return ctx.Render(http.StatusOK, contentTpl, data)
}
