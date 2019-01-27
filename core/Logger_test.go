/*
 * Copyright 2018 log4go@freelists.org
 *
 * Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
package core

import (
	"testing"

	"github.com/log4go/log4go"
	"github.com/log4go/log4go/spi"
)

func TestDefault(t *testing.T) {
	var (
		logger spi.Logger
		err    error
	)
	logger, err = log4go.LogManager.GetLogger("test")
	if err != nil {
		t.Fail()
	}
	err = logger.Trace("this is a", log4go.Level.TRACE.String(), "message")
	if err != nil {
		t.Fail()
	}
	err = logger.Debug("this is a", log4go.Level.DEBUG.String(), "message")
	if err != nil {
		t.Fail()
	}
	err = logger.Info("this is a", log4go.Level.INFO.String(), "message")
	if err != nil {
		t.Fail()
	}
	err = logger.Warn("this is a", log4go.Level.WARN.String(), "message")
	if err != nil {
		t.Fail()
	}
	err = logger.Error("this is a", log4go.Level.ERROR.String(), "message")
	if err != nil {
		t.Fail()
	}
	err = logger.Fatal("this is a", log4go.Level.FATAL.String(), "message")
	if err != nil {
		t.Fail()
	}
}
