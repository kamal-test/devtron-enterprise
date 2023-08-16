/*
 * Copyright (c) 2020 Devtron Labs
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *    http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

package router

import (
	"github.com/devtron-labs/devtron/api/logger"
	user "github.com/devtron-labs/devtron/api/restHandler"
	"github.com/gorilla/mux"
)

type UserAttributesRouter interface {
	InitUserAttributesRouter(helmRouter *mux.Router)
}

type UserAttributesRouterImpl struct {
	userAttributesRestHandler user.UserAttributesRestHandler
	userAuth                  logger.UserAuth
}

func NewUserAttributesRouterImpl(userAttributesRestHandler user.UserAttributesRestHandler, userAuth logger.UserAuth) *UserAttributesRouterImpl {
	router := &UserAttributesRouterImpl{
		userAttributesRestHandler: userAttributesRestHandler,
		userAuth:                  userAuth,
	}
	return router
}

func (router UserAttributesRouterImpl) InitUserAttributesRouter(attributesRouter *mux.Router) {
	attributesRouter.Use(router.userAuth.LoggingMiddleware)
	attributesRouter.Path("/update").
		HandlerFunc(router.userAttributesRestHandler.UpdateUserAttributes).Methods("POST")
	attributesRouter.Path("/get").
		HandlerFunc(router.userAttributesRestHandler.GetUserAttribute).Queries("key", "{key}").Methods("GET")
}
