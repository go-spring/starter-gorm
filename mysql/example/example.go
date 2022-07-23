/*
 * Copyright 2012-2019 the original author or authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      https://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package main

import (
	"fmt"

	"github.com/go-spring/spring-base/log"
	"github.com/go-spring/spring-core/gs"
	"gorm.io/gorm"

	_ "github.com/go-spring/starter-gorm/mysql"
)

var (
	logger = log.GetLogger()
)

type runner struct {
	DB *gorm.DB `autowire:""`
}

func (r *runner) Run(ctx gs.Context) {
	var engines []string
	r.DB.Raw("select engine from engines").Scan(&engines)
	logger.Infof("got mysql engines %v", engines)
	go gs.ShutDown()
}

func main() {
	gs.Object(&runner{}).Export((*gs.AppRunner)(nil))
	fmt.Printf("program exited %v\n", gs.Web(false).Run())
}
