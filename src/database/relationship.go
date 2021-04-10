package database

import (
	"context"
	"strconv"
	"time"

	"github.com/zGina/Attack-Seaman/src/config"
	"github.com/zGina/Attack-Seaman/src/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// GetRelationships returns all relationships.
// start, end int, order, sort string
func (d *TenDatabase) GetRelationships(paging *model.Paging) []*model.Relationship {
	conf := config.Get()
	relationships := []*model.Relationship{}
	condition := bson.M{}
	if paging.Condition != nil {
		condition = (paging.Condition).(bson.M)
	}
	cursor, err := d.DB.Collection(conf.Database.Tbname).
		Find(context.Background(), condition,
			&options.FindOptions{
				Skip:  paging.Skip,
				Sort:  bson.D{bson.E{Key: paging.SortKey, Value: paging.SortVal}},
				Limit: paging.Limit,
			})
	if err != nil {
		return nil
	}
	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		relationship := &model.Relationship{}
		if err := cursor.Decode(relationship); err != nil {
			return nil
		}
		relationships = append(relationships, relationship)
	}

	return relationships
}

// CreateRelationship creates a relationship.
func (d *TenDatabase) CreateRelationship(relationship *model.Relationship) *model.Relationship {
	conf := config.Get()
	relationship.Created = time.Now()
	relationship.Modified = time.Now()
	_, result := d.DB.Collection(conf.Database.Tbname).
		InsertOne(context.Background(), relationship)
	if result != nil {
		return relationship
	}
	return relationship
}

// GetRelationshipByName returns the relationship by the given name or nil.
func (d *TenDatabase) GetRelationshipByName(name string) *model.Relationship {
	conf := config.Get()
	var relationship *model.Relationship
	err := d.DB.Collection(conf.Database.Tbname).
		FindOne(context.Background(), bson.D{{Key: "name", Value: name}}).
		Decode(&relationship)
	if err != nil {
		return nil
	}
	return relationship
}

// GetRelationshipByStixID returns the user by the given name or nil.
func (d *TenDatabase) GetRelationshipByStixID(id string) *model.Relationship {
	conf := config.Get()
	var relationship *model.Relationship
	err := d.DB.Collection(conf.Database.Tbname).
		FindOne(context.Background(), bson.M{"id": id}).
		Decode(&relationship)
	if err != nil {
		return nil
	}
	return relationship
}

// GetRelationshipByIDs returns the relationship by the given id or nil.
func (d *TenDatabase) GetRelationshipByIDs(ids []string) []*model.Relationship {
	conf := config.Get()
	var relationships []*model.Relationship
	cursor, err := d.DB.Collection(conf.Database.Tbname).
		Find(context.Background(), bson.D{{
			Key: "id",
			Value: bson.D{{
				Key:   "$in",
				Value: ids,
			}},
		}})
	if err != nil {
		return nil
	}
	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		relationship := &model.Relationship{}
		if err := cursor.Decode(relationship); err != nil {
			return nil
		}
		relationships = append(relationships, relationship)
	}

	return relationships
}

// CountRelationship returns the relationship count
func (d *TenDatabase) CountRelationship(condition interface{}) string {
	conf := config.Get()
	total, err := d.DB.Collection(conf.Database.Tbname).CountDocuments(context.Background(), condition, &options.CountOptions{})
	if err != nil {
		return "0"
	}
	return strconv.Itoa(int(total))
}

// DeleteRelationshipByID deletes a relationship by its id.
// func (d *TenDatabase) DeleteRelationshipByID(id string) error {
// 	if d.CountRelationship(bson.D{{Key: "id", Value: id}}) == "0" {
// 		_, err := d.DB.Collection(conf.Database.Tbname).DeleteOne(context.Background(), bson.M{"id": id})
// 		return err
// 	}
// 	return errors.New("the current relationship has no posts published")
// }
func (d *TenDatabase) DeleteRelationshipByID(id string) error {
	conf := config.Get()
	_, err := d.DB.Collection(conf.Database.Tbname).DeleteOne(context.Background(), bson.M{"id": id})
	return err
}

// GetRelationshipByID get a relationship by its id.
func (d *TenDatabase) GetRelationshipByID(id string) *model.Relationship {
	conf := config.Get()
	var relationship *model.Relationship
	err := d.DB.Collection(conf.Database.Tbname).
		FindOne(context.Background(), bson.M{"id": id}).
		Decode(&relationship)
	if err != nil {
		return nil
	}
	return relationship
}

// UpdateRelationship updates a relationship.
func (d *TenDatabase) UpdateRelationship(relationship *model.Relationship) *model.Relationship {
	conf := config.Get()
	relationship.Modified = time.Now()
	result := d.DB.Collection(conf.Database.Tbname).
		FindOneAndReplace(context.Background(),
			bson.D{{Key: "id", Value: relationship.STIX_ID}},
			relationship,
			&options.FindOneAndReplaceOptions{},
		)
	if result != nil {
		return relationship
	}
	return nil
}
