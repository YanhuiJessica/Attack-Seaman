package router

import (
	"fmt"
	"log"
	"net/http"
	"os/exec"
	"strings"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/zGina/Attack-Seaman/src/api"
	"github.com/zGina/Attack-Seaman/src/config"
	"github.com/zGina/Attack-Seaman/src/database"
	"github.com/zGina/Attack-Seaman/src/error"
	"github.com/zGina/Attack-Seaman/src/model"
)

// Create creates the gin engine with all routes.
func Create(db *database.TenDatabase, vInfo *model.VersionInfo, conf *config.Configuration) *gin.Engine {
	g := gin.New()

	g.Use(gin.Logger(), gin.Recovery(), error.Handler())
	g.NoRoute(error.NotFound())

	g.Use(func(ctx *gin.Context) {
		ctx.Header("Content-Type", "application/json")
		origin := ctx.Request.Header.Get("Origin")
		for header, value := range conf.Server.ResponseHeaders {
			if origin == "http://attack-seaman.com:3000" && header == "Access-Control-Allow-Origin" {
				ctx.Header("Access-Control-Allow-Origin", "http://attack-seaman.com:3000")
			} else {
				ctx.Header(header, value)
			}
		}
		if ctx.Request.Method == "OPTIONS" {
			ctx.AbortWithStatus(http.StatusNoContent)
		}
	})

	attackPatternAPIHandler := api.AttackPatternAPI{DB: db}
	relationshipAPIHandler := api.RelationshipAPI{DB: db}

	postAP := g.Group("/attackPatterns")
	{
		postAP.GET("", attackPatternAPIHandler.GetAttackPatterns)
		postAP.POST("/save", attackPatternAPIHandler.SaveAttackPatternByID)
		postAP.POST("", attackPatternAPIHandler.CreateAttackPattern)
		postAP.GET(":id", attackPatternAPIHandler.GetAttackPatternByID)
		postAP.PUT(":id", attackPatternAPIHandler.UpdateAttackPatternByID)
		postAP.DELETE(":id", attackPatternAPIHandler.DeleteAttackPatternByID)
	}

	postR := g.Group("/relationships")
	{
		postR.GET("", relationshipAPIHandler.GetRelationships)
		postR.POST("", relationshipAPIHandler.CreateRelationship)
		postR.GET(":id", relationshipAPIHandler.GetRelationshipByID)
		postR.PUT(":id", relationshipAPIHandler.UpdateRelationshipByID)
		postR.DELETE(":id", relationshipAPIHandler.DeleteRelationshipByID)
	}

	g.POST("/upload", func(c *gin.Context) {

		file, err := c.FormFile("File")
		if err != nil {
			c.AbortWithError(403, errors.New("Bad file post"))
			return
		}
		log.Println(file.Filename)
		FILE_UPLOAD_PATH := conf.Utils.Filepath
		
		err = c.SaveUploadedFile(file, FILE_UPLOAD_PATH+file.Filename)
		if err != nil {
			log.Fatal(err)
		}
		c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))

		filepath := FILE_UPLOAD_PATH + file.Filename
		basename := ".json"
		name := strings.TrimSuffix(file.Filename, basename)

		args := []string{"/app/tools/import.sh", name, filepath}
		fmt.Println(args)

		_, err = exec.Command("/bin/sh", args...).Output()

		if err != nil {
			fmt.Println(err)
			log.Fatal(err)
		}
		fmt.Println("import success")
		conf.Database.Tbname = name
		config.Save(conf)

	})

	g.GET("version", func(ctx *gin.Context) {
		ctx.JSON(200, vInfo)
	})

	return g
}
