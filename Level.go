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

import "github.com/log4go/log4go/spi"

var Level = initLevel()

type level struct {
	ALL   spi.Level
	TRACE spi.Level
	DEBUG spi.Level
	INFO  spi.Level
	WARN  spi.Level
	ERROR spi.Level
	FATAL spi.Level
	OFF   spi.Level
}

func initLevel() *level {
	lvl := new(level)
	lvl.ALL = spi.NewStandardLevel(70, "ALL")
	lvl.TRACE = spi.NewStandardLevel(60, "TRACE")
	lvl.DEBUG = spi.NewStandardLevel(50, "DEBUG")
	lvl.INFO = spi.NewStandardLevel(40, "INFO")
	lvl.WARN = spi.NewStandardLevel(30, "WARN")
	lvl.ERROR = spi.NewStandardLevel(20, "ERROR")
	lvl.FATAL = spi.NewStandardLevel(10, "FATAL")
	lvl.OFF = spi.NewStandardLevel(0, "OFF")
	return lvl
}
