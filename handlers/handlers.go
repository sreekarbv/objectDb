package handlers

import (
	"endor/db"
	"endor/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Handler struct {
	rDb *db.RedisDB
}

func NewHandler(rDb *db.RedisDB) Handler {
	return Handler{rDb: rDb}
}

func (handler *Handler) Store(ctx *gin.Context) {
	obj, err := model.BindObject(ctx)
	if err != nil {
		fmt.Println("Unsupported Object Type")
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err)
	}
	err = handler.rDb.Store(ctx.Request.Context(), obj)
	if err != nil {
		fmt.Println("Error while storing", err)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, err)
	}
	ctx.Status(http.StatusOK)
}

func (handler *Handler) GetObjectById(ctx *gin.Context) {
	id := ctx.Param("id")
	obj, err := handler.rDb.GetObjectByID(ctx.Request.Context(), id)
	if err != nil {
		fmt.Println("Error while getting object by id", err)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, err)
	}
	ctx.JSON(http.StatusOK, obj)
}

func (handler *Handler) GetObjectByName(ctx *gin.Context) {
	name := ctx.Param("name")
	obj, err := handler.rDb.GetObjectByName(ctx.Request.Context(), name)
	if err != nil {
		fmt.Println("Error while getting object by name", err)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, err)
	}
	ctx.JSON(http.StatusOK, obj)
}

func (handler *Handler) ListObjects(ctx *gin.Context) {
	kind := ctx.Param("kind")
	objs, err := handler.rDb.ListObjects(ctx.Request.Context(), kind)
	if err != nil {
		fmt.Println("Error while getting list of objects by kind", err)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, err)
	}
	ctx.JSON(http.StatusOK, objs)
}

func (handler *Handler) DeleteObject(ctx *gin.Context) {
	id := ctx.Param("id")
	err := handler.rDb.DeleteObject(ctx.Request.Context(), id)
	if err != nil {
		fmt.Println("Error while deleting object by id", err)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, err)
	}
	ctx.JSON(http.StatusOK, "")
}
