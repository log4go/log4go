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
	"errors"
	"fmt"
	"reflect"
	"runtime"
	"sync"

	"github.com/log4go/log4go/spi"
)

var LogManager = initLogManager()

type logManager struct {
	FQCN                  string
	ROOT_LOGGER_NAME      string
	FACTORY_PROPERTY_NAME string
	Logger                spi.Logger
	factory               spi.LoggerContextFactory
	factory_repository    *sync.Map
}

func initLogManager() *logManager {
	lm := new(logManager)
	lm.FQCN = fmt.Sprint(reflect.TypeOf(lm))
	lm.FACTORY_PROPERTY_NAME = "log4go.loggerContextFactory"
	lm.factory_repository = new(sync.Map)
	return lm
}

func (lm *logManager) GetContext() (spi.LoggerContext, error) {
	var (
		err  error
		ctxt spi.LoggerContext
	)
	_, file, _, ok := runtime.Caller(2)
	if ok {
		ctxt, err = lm.factory.GetContext(file)
	} else {
		err = errors.New("Not able to identify the caller")
	}
	return ctxt, err
}

func (lm *logManager) GetLogger(name string) (spi.Logger, error) {
	var (
		ctxt spi.LoggerContext
		logr spi.Logger
		err  error
	)
	ctxt, err = lm.GetContext()
	if err == nil {
		logr, err = ctxt.GetLogger(name)
	}
	return logr, err
}

func (lm *logManager) AddLoggerContextFactory(name string, factory spi.LoggerContextFactory) {
	lm.factory_repository.Store(name, factory)
}

func (lm *logManager) GetLoggerContextFactory(name string) (spi.LoggerContextFactory, error) {
	var (
		lcf spi.LoggerContextFactory
		err error
	)
	ret, ok := lm.factory_repository.Load(name)
	if ok {
		lcf = ret.(spi.LoggerContextFactory)
	} else {
		err = errors.New(fmt.Sprint("LoggerContextFactory ", name, " not found in repository"))
	}
	return lcf, err
}

func (lm *logManager) SetLoggerContextFactory(name string) (err error) {
	lm.factory, err = lm.GetLoggerContextFactory(name)
	return
}
