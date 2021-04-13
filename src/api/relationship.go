package api

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"github.com/zGina/Attack-Seaman/src/model"
	"go.mongodb.org/mongo-driver/bson"
)

// The RelationshipDatabase interface for encapsulating database access.
type RelationshipDatabase interface {
	GetRelationshipByIDs(ids []string) []*model.Relationship
	GetRelationshipByID(id string) *model.Relationship
	DeleteRelationshipByID(id string) error
	CreateRelationship(relationship *model.Relationship) *model.Relationship
	GetRelationships(paging *model.Paging) []*model.Relationship
	UpdateRelationship(relationship *model.Relationship) *model.Relationship
	CountRelationship(condition interface{}) string
}

// The RelationshipAPI provides handlers for managing relationships.
type RelationshipAPI struct {
	DB RelationshipDatabase
}

// GetRelationshipByIDs returns the relationship by id
func (a *RelationshipAPI) GetRelationshipByIDs(ctx *gin.Context) {
	withIDs(ctx, "id", func(ids []string) {
		ctx.JSON(200, a.DB.GetRelationshipByIDs(ids))
	})
}

// DeleteRelationshipByID deletes the relationship by id
func (a *RelationshipAPI) DeleteRelationshipByID(ctx *gin.Context) {
	withID(ctx, "id", func(id string) {
		if err := a.DB.DeleteRelationshipByID(id); err == nil {
			ctx.JSON(200, http.StatusOK)
		} else {
			if err != nil {
				ctx.AbortWithError(500, err)
			} else {
				ctx.AbortWithError(404, errors.New("relationship does not exist"))
			}
		}
	})
}

// GetRelationships returns all the relationships
// _end=5&_order=DESC&_sort=id&_start=0 adapt react-admin
func (a *RelationshipAPI) GetRelationships(ctx *gin.Context) {
	var (
		start int64
		end   int64
		sort  string
		order int
	)
	id := ctx.DefaultQuery("id", "")
	if id != "" {
		a.GetRelationshipByIDs(ctx)
		return
	}
	start, _ = strconv.ParseInt(ctx.DefaultQuery("_start", "0"), 10, 64)
	end, _ = strconv.ParseInt(ctx.DefaultQuery("_end", "10"), 10, 64)
	// sort = ctx.DefaultQuery("_sort", "_id")
	relationshipType := ctx.DefaultQuery("relationship_type", "")

	sort = "modified"
	order = 1

	condition := bson.M{}
	if relationshipType != "" {
		condition = bson.M{
			"relationship_type": relationshipType,
			"type":              "relationship",
			"revoked":           bson.M{"$ne": true}}
	} else {
		condition = bson.M{
			"type":    "relationship",
			"revoked": bson.M{"$ne": true}}
	}

	if ctx.DefaultQuery("_order", "DESC") == "DESC" {
		order = -1
	}

	limit := end - start
	relationships := a.DB.GetRelationships(
		&model.Paging{
			Skip:      &start,
			Limit:     &limit,
			SortKey:   sort,
			SortVal:   order,
			Condition: condition,
		})

	ctx.Header("X-Total-Count", a.DB.CountRelationship(condition))
	ctx.JSON(200, relationships)
}

// CreateRelationship creates a relationship.
func (a *RelationshipAPI) CreateRelationship(ctx *gin.Context) {
	var relationship = model.Relationship{}

	if err := ctx.ShouldBind(&relationship); err == nil {
		if relationship.STIX_ID != "" {
			if res := a.DB.GetRelationshipByID(relationship.STIX_ID); res != nil {
				ctx.AbortWithError(403, errors.New("Repeated relationship ID"))
				return
			}
		} else {
			myuuid := uuid.NewV4()
			relationship.STIX_ID = "relationship--" + myuuid.String()
		}
		if result := a.DB.CreateRelationship(relationship.New()); result != nil {
			ctx.JSON(201, result)
		} else {
			ctx.AbortWithError(500, errors.New("CreateRelationship error"))
		}
	} else {
		ctx.AbortWithError(500, errors.New("ShouldBind error"))
	}
}

// GetRelationshipByID returns the relationship by id
func (a *RelationshipAPI) GetRelationshipByID(ctx *gin.Context) {
	withID(ctx, "id", func(id string) {
		if relationship := a.DB.GetRelationshipByID(id); relationship != nil {
			ctx.JSON(200, relationship)
		} else {
			ctx.AbortWithError(404, errors.New("relationship does not exist"))
		}
	})
}

// UpdateRelationshipByID returns the relationship by id
func (a *RelationshipAPI) UpdateRelationshipByID(ctx *gin.Context) {
	withID(ctx, "id", func(id string) {
		var relationship = model.Relationship{}
		abort := errors.New("relationship does not exist")
		if err := ctx.ShouldBind(&relationship); err == nil {
			if result := a.DB.UpdateRelationship(&relationship); result != nil {
				ctx.JSON(200, result)
			} else {
				ctx.AbortWithError(404, abort)
			}
		} else {
			ctx.AbortWithError(404, abort)
		}
	})
}
