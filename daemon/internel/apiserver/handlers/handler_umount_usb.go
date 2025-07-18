package handlers

import (
	"net/http"

	"github.com/beclab/Olares/daemon/pkg/commands"
	umountusb "github.com/beclab/Olares/daemon/pkg/commands/umount_usb"
	"github.com/gofiber/fiber/v2"
	"k8s.io/klog/v2"
)

type UmountReq struct {
	Path string ``
}

func (h *Handlers) umountUsbInNode(ctx *fiber.Ctx, cmd commands.Interface, pathInNode string) error {
	_, err := cmd.Execute(ctx.Context(), &umountusb.Param{
		Path: pathInNode,
	})

	if err != nil {
		return h.ErrJSON(ctx, http.StatusInternalServerError, err.Error())
	}

	return h.OkJSON(ctx, "success to umount")
}

func (h *Handlers) PostUmountUsb(ctx *fiber.Ctx, cmd commands.Interface) error {
	var req UmountReq
	if err := h.ParseBody(ctx, &req); err != nil {
		klog.Error("parse request error, ", err)
		return h.ErrJSON(ctx, http.StatusBadRequest, err.Error())
	}
	if req.Path == "" {
		return h.ErrJSON(ctx, http.StatusBadRequest, "ip is empty")
	}

	return h.umountUsbInNode(ctx, cmd, req.Path)
}

func (h *Handlers) PostUmountUsbInCluster(ctx *fiber.Ctx, cmd commands.Interface) error {
	var req UmountReq
	if err := h.ParseBody(ctx, &req); err != nil {
		klog.Error("parse request error, ", err)
		return h.ErrJSON(ctx, http.StatusBadRequest, err.Error())
	}
	if req.Path == "" {
		return h.ErrJSON(ctx, http.StatusBadRequest, "ip is empty")
	}

	nodePath := clusterPathToNodePath(req.Path)

	return h.umountUsbInNode(ctx, cmd, nodePath)
}
