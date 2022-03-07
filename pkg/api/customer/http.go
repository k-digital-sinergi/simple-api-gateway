package customer

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"simple-api-gateway/pkg/api/customer/model"
	"simple-api-gateway/pkg/util"
)

type HTTP struct {
	svc Service
}

func NewHTTP(svc Service, group *gin.RouterGroup) {
	h := HTTP{svc: svc}

	g := group.Group("/customer")

	g.GET("/list", h.list)
	g.GET("/:id", h.get)
	g.POST("", h.create)
	g.PUT("/:id", h.update)
	g.DELETE("/:id", h.delete)
}

func (h *HTTP) list(ctx *gin.Context) {
	customers, err := h.svc.List(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, util.ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, customers)
}

func (h *HTTP) get(ctx *gin.Context) {
	var (
		err error
		req model.GetRequest
	)

	if err = ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, util.ErrorResponse(err))
		return
	}

	customer, err := h.svc.Get(ctx, &req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, util.ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, customer)
}

func (h *HTTP) create(ctx *gin.Context) {
	var (
		err error
		req model.CreateRequest
	)

	if err = ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, util.ErrorResponse(err))
		return
	}

	custID, err := h.svc.Create(ctx, &req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, util.ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusCreated, custID)
}

func (h *HTTP) update(ctx *gin.Context) {
	var (
		err error
		req model.UpdateRequest
	)

	if err = ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, util.ErrorResponse(err))
		return
	}

	if err = ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, util.ErrorResponse(err))
		return
	}

	err = h.svc.Update(ctx, &req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, util.ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, "data updated")
}

func (h *HTTP) delete(ctx *gin.Context) {
	var (
		err error
		req model.DeleteRequest
	)

	if err = ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, util.ErrorResponse(err))
		return
	}

	err = h.svc.Delete(ctx, &req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, util.ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, "data deleted")
}
