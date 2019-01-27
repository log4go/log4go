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
package spi

type StandardLevel struct {
	int uint8
	str string
}

func (stdlvl *StandardLevel) Int() uint8 {
	return stdlvl.int
}

func (stdlvl *StandardLevel) String() string {
	return stdlvl.str
}

func (stdlvl *StandardLevel) Equal(other Level) bool {
	if stdlvl.Int() == other.Int() {
		return true
	} else {
		return false
	}
}

func (stdlvl *StandardLevel) GreaterOrEqual(other Level) bool {
	if stdlvl.Int() >= other.Int() {
		return true
	} else {
		return false
	}
}

func NewStandardLevel(int uint8, str string) *StandardLevel {
	stdlvl := new(StandardLevel)
	stdlvl.int = int
	stdlvl.str = str
	return stdlvl
}
