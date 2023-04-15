/*
 * Copyright (c) 2023 a-clap. All rights reserved.
 * Use of this source code is governed by a MIT-style license that can be found in the LICENSE file.
 */

package distillation

import (
	"errors"

	"github.com/gin-gonic/gin"
)

const (
	RoutesEnableHeater          = "/api/heater"
	RoutesGetAllHeaters         = "/api/heater"
	RoutesGetEnabledHeaters     = "/api/heater/enabled"
	RoutesConfigureHeater       = "/api/heater/enabled"
	RoutesGetDS                 = "/api/onewire"
	RoutesConfigureDS           = "/api/onewire"
	RoutesGetDSTemperatures     = "/api/onewire/temperatures"
	RoutesGetPT                 = "/api/pt100"
	RoutesConfigurePT           = "/api/pt100"
	RoutesGetPTTemperatures     = "/api/pt100/temperatures"
	RoutesGetGPIO               = "/api/gpio"
	RoutesConfigureGPIO         = "/api/gpio"
	RoutesProcessPhases         = "/api/phases"
	RoutesProcessConfigPhase    = "/api/phases/:id"
	RoutesProcessConfigValidate = "/api/process/validate"
	RoutesProcess               = "/api/process"
	RoutesProcessStatus         = "/api/process/status"
)

var (
	ErrNotImplemented = errors.New("not implemented")
)

// routes configures default handlers for paths above
func (h *Handler) routes() {
	h.GET(RoutesGetAllHeaters, h.getAllHeaters())
	h.GET(RoutesGetEnabledHeaters, h.getEnabledHeaters())
	h.PUT(RoutesEnableHeater, h.enableHeater())
	h.PUT(RoutesConfigureHeater, h.configEnabledHeater())

	h.GET(RoutesGetDS, h.getDS())
	h.GET(RoutesGetDSTemperatures, h.getDSTemperatures())
	h.PUT(RoutesConfigureDS, h.configureDS())

	h.GET(RoutesGetPT, h.getPT())
	h.GET(RoutesGetPTTemperatures, h.getPTTemperatures())
	h.PUT(RoutesConfigurePT, h.configurePT())

	h.GET(RoutesGetGPIO, h.getGPIO())
	h.PUT(RoutesConfigureGPIO, h.configureGPIO())

	h.GET(RoutesProcessPhases, h.getPhaseCount())
	h.PUT(RoutesProcessPhases, h.configurePhaseCount())
	h.GET(RoutesProcessConfigPhase, h.getProcessConfig())
	h.PUT(RoutesProcessConfigPhase, h.setProcessConfig())

	h.GET(RoutesProcessConfigValidate, h.getConfigValidation())
	h.GET(RoutesProcessStatus, h.getProcessStatus())
	h.PUT(RoutesProcess, h.configureProcess())

}

// common respond for whole rest API
func (*Handler) respond(ctx *gin.Context, code int, obj any) {
	ctx.JSON(code, obj)
}