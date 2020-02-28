/*
 * Copyright (c) 2020. The StudyGolang Authors. All rights reserved.
 * Use of this source code is governed by a Apache 2.0
 * license that can be found in the LICENSE file.
 * https://studygolang.com
 * Author:polaris	polaris@studygolang.com
 */

package validator

import (
	"sync"

	"github.com/go-playground/validator"
)

type CustomValidator struct {
	once     sync.Once
	validate *validator.Validate
}

func (c *CustomValidator) Validate(i interface{}) error {
	c.lazyInit()
	return c.validate.Struct(i)
}

func (c *CustomValidator) lazyInit() {
	c.once.Do(func() {
		c.validate = validator.New()
	})
}
