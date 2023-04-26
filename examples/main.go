/*
 * Copyright 2023 CloudWeGo Authors
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package main

import (
	"context"
	"fmt"
	"time"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/client"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/protocol"
	"github.com/hertz-contrib/ctxtimeout"
)

func main() {
	c, _ := client.NewClient()
	c.Use(ctxtimeout.CtxTimeout)

	h := server.New()
	h.GET("/ping", func(c context.Context, ctx *app.RequestContext) {
		time.Sleep(time.Second * 2)
		ctx.WriteString("hello, world") //nolint:errcheck
	})
	go h.Run()
	time.Sleep(time.Second)

	ctxA, cancelA := context.WithTimeout(context.Background(), time.Second*3)
	defer func() {
		cancelA()
	}()
	req, resp := protocol.AcquireRequest(), protocol.AcquireResponse()
	req.SetRequestURI("http://127.0.0.1:8888/ping")
	err := c.Do(ctxA, req, resp)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(resp.Body()))

	ctxB, cancelB := context.WithTimeout(context.Background(), time.Second)
	defer func() {
		cancelB()
	}()
	err = c.Do(ctxB, req, resp)
	if err != ctxtimeout.ErrTimeout {
		panic(err)
	}
}
