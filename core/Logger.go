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
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/log4go/log4go"
	"github.com/log4go/log4go/spi"
)

type Logger struct {
	mutex sync.RWMutex
	level spi.Level
	name  string
}

func (l *Logger) Log(level spi.Level, messages ...string) (err error) {
	l.mutex.RLock()
	defer l.mutex.RUnlock()
	prefix := []string{
		time.Now().Format(time.RFC3339Nano),
		"[" + level.String() + "]",
		"|"}
	parts := []string{strings.Join(prefix, "\t")}
	parts = append(parts, messages...)
	fmt.Println(strings.Join(parts, " "))
	return
}

func (l *Logger) IsEnabled(level spi.Level) (en bool) {
	l.mutex.RLock()
	defer l.mutex.RUnlock()
	en = l.level.GreaterOrEqual(level)
	return
}

func (l *Logger) Trace(messages ...string) (err error) {
	if l.IsTraceEnabled() {
		err = l.Log(log4go.Level.TRACE, messages...)
	}
	return
}

func (l *Logger) IsTraceEnabled() (en bool) {
	en = l.IsEnabled(log4go.Level.TRACE)
	return
}

func (l *Logger) Debug(messages ...string) (err error) {
	if l.IsDebugEnabled() {
		err = l.Log(log4go.Level.DEBUG, messages...)
	}
	return
}

func (l *Logger) IsDebugEnabled() (en bool) {
	en = l.IsEnabled(log4go.Level.DEBUG)
	return
}

func (l *Logger) Info(messages ...string) (err error) {
	if l.IsInfoEnabled() {
		err = l.Log(log4go.Level.INFO, messages...)
	}
	return
}

func (l *Logger) IsInfoEnabled() (en bool) {
	en = l.IsEnabled(log4go.Level.INFO)
	return
}

func (l *Logger) Warn(messages ...string) (err error) {
	if l.IsWarnEnabled() {
		err = l.Log(log4go.Level.WARN, messages...)
	}
	return
}

func (l *Logger) IsWarnEnabled() (en bool) {
	en = l.IsEnabled(log4go.Level.WARN)
	return
}

func (l *Logger) Error(messages ...string) (err error) {
	if l.IsErrorEnabled() {
		err = l.Log(log4go.Level.ERROR, messages...)
	}
	return
}

func (l *Logger) IsErrorEnabled() (en bool) {
	en = l.IsEnabled(log4go.Level.ERROR)
	return
}

func (l *Logger) Fatal(messages ...string) (err error) {
	if l.IsFatalEnabled() {
		err = l.Log(log4go.Level.FATAL, messages...)
	}
	return
}

func (l *Logger) IsFatalEnabled() (en bool) {
	en = l.IsEnabled(log4go.Level.FATAL)
	return
}
