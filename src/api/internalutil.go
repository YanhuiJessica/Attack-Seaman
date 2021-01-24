package api

import (
	"errors"

	"github.com/gin-gonic/gin"
)

func withID(ctx *gin.Context, name string, f func(id string)) {
	// if id, err := primitive.ObjectIDFromHex(ctx.Param(name)); err == nil {
	f(ctx.Param(name))
	// } else {
	// ctx.AbortWithError(400, errors.New("invalid id"))
	// }
}

func withIDs(ctx *gin.Context, name string, f func(id []string)) {
	ids, b := ctx.GetQueryArray(name)
	objectIds := []string{}
	abort := errors.New("invalid id")
	if b {
		for _, id := range ids {
			// if objID, err := id, 1; err == nil {
			objectIds = append(objectIds, id)
			// } else {
			// 	ctx.AbortWithError(400, abort)
			// }
		}
		f(objectIds)
	} else {
		ctx.AbortWithError(400, abort)
	}
}
