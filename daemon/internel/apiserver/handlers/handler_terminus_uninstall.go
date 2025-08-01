package handlers

import (
	"net/http"

	"github.com/beclab/Olares/daemon/pkg/commands"
	"github.com/gofiber/fiber/v2"
	"k8s.io/klog/v2"
)

func (h *Handlers) PostTerminusUninstall(ctx *fiber.Ctx, cmd commands.Interface) error {
	// run in background
	_, err := cmd.Execute(h.mainCtx, nil)

	if err != nil {
		klog.Error("execute command error, ", err, ", ", cmd.OperationName().Stirng())
		return h.ErrJSON(ctx, http.StatusBadRequest, err.Error())
	}

	return h.OkJSON(ctx, "start to "+cmd.OperationName().Stirng())
}
