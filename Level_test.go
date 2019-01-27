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
package log4go

import (
	"testing"

	"github.com/log4go/log4go/spi"
)

func TestAll(t *testing.T) {
	l := Level.ALL
	if l.Int() != 70 {
		t.Fail()
	}
}

func TestTrace(t *testing.T) {
	l := Level.TRACE
	if l.Int() != 60 {
		t.Fail()
	}
}

func TestDebug(t *testing.T) {
	l := Level.DEBUG
	if l.Int() != 50 {
		t.Fail()
	}
}

func TestInfo(t *testing.T) {
	l := Level.INFO
	if l.Int() != 40 {
		t.Fail()
	}
}

func TestWarn(t *testing.T) {
	l := Level.WARN
	if l.Int() != 30 {
		t.Fail()
	}
}

func TestError(t *testing.T) {
	l := Level.ERROR
	if l.Int() != 20 {
		t.Fail()
	}
}

func TestFatal(t *testing.T) {
	l := Level.FATAL
	if l.Int() != 10 {
		t.Fail()
	}
}

func TestOff(t *testing.T) {
	l := Level.OFF
	if l.Int() != 0 {
		t.Fail()
	}
}

func TestEqual(t *testing.T) {
	l := Level.INFO
	l2 := spi.NewStandardLevel(l.Int(), l.String())
	if !l.Equal(l2) {
		t.Fail()
	}
}
