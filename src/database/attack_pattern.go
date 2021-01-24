package database

import (
	"context"
	"strconv"
	"time"

	"github.com/zGina/Attack-Seaman/src/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// GetAttackPatterns returns all attackPatterns.
// start, end int, order, sort string
func (d *TenDatabase) GetAttackPatterns(paging *model.Paging) []*model.AttackPattern {
	attackPatterns := []*model.AttackPattern{}
	condition := bson.M{}
	if paging.Condition != nil {
		condition = (paging.Condition).(bson.M)
	}
	cursor, err := d.DB.Collection("mitre_attack").
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
		attackPattern := &model.AttackPattern{}
		if err := cursor.Decode(attackPattern); err != nil {
			return nil
		}
		attackPatterns = append(attackPatterns, attackPattern)
	}

	return attackPatterns
}

// CreateAttackPattern creates a attackPattern.
func (d *TenDatabase) CreateAttackPattern(attackPattern *model.AttackPattern) *model.AttackPattern {
	attackPattern.Created = time.Now()
	attackPattern.Modified = time.Now()
	_, result := d.DB.Collection("mitre_attack").
		InsertOne(context.Background(), attackPattern)
	if result != nil {
		return attackPattern
	}
	return attackPattern
}

// GetAttackPatternByName returns the attackPattern by the given name or nil.
func (d *TenDatabase) GetAttackPatternByName(name string) *model.AttackPattern {
	var attackPattern *model.AttackPattern
	err := d.DB.Collection("mitre_attack").
		FindOne(context.Background(), bson.D{{Key: "name", Value: name}}).
		Decode(&attackPattern)
	if err != nil {
		return nil
	}
	return attackPattern
}

// GetAttackPatternByStixID returns the user by the given name or nil.
func (d *TenDatabase) GetAttackPatternByStixID(id string) *model.AttackPattern {
	var attackPattern *model.AttackPattern
	err := d.DB.Collection("mitre_attack").
		FindOne(context.Background(), bson.M{"id": id}).
		Decode(&attackPattern)
	if err != nil {
		return nil
	}
	return attackPattern
}

// GetAttackPatternByIDs returns the attackPattern by the given id or nil.
func (d *TenDatabase) GetAttackPatternByIDs(ids []string) []*model.AttackPattern {
	var attackPatterns []*model.AttackPattern
	cursor, err := d.DB.Collection("mitre_attack").
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
		attackPattern := &model.AttackPattern{}
		if err := cursor.Decode(attackPattern); err != nil {
			return nil
		}
		attackPatterns = append(attackPatterns, attackPattern)
	}

	return attackPatterns
}

// CountAttackPattern returns the attackPattern count
func (d *TenDatabase) CountAttackPattern(condition interface{}) string {
	total, err := d.DB.Collection("mitre_attack").CountDocuments(context.Background(), condition, &options.CountOptions{})
	if err != nil {
		return "0"
	}
	return strconv.Itoa(int(total))
}

// DeleteAttackPatternByID deletes a attackPattern by its id.
func (d *TenDatabase) DeleteAttackPatternByID(id string) error {
	_, err := d.DB.Collection("mitre_attack").DeleteOne(context.Background(), bson.M{"id": id})
	return err
}

// GetAttackPatternByID get a attackPattern by its id.
func (d *TenDatabase) GetAttackPatternByID(id string) *model.AttackPattern {
	var attackPattern *model.AttackPattern
	print(id)
	err := d.DB.Collection("mitre_attack").
		FindOne(context.Background(), bson.M{"id": id}).
		Decode(&attackPattern)
	print(attackPattern)
	if err != nil {
		return nil
	}
	return attackPattern
}

// UpdateAttackPattern updates a attackPattern.
func (d *TenDatabase) UpdateAttackPattern(attackPattern *model.AttackPattern) *model.AttackPattern {
	attackPattern.Modified = time.Now()
	result := d.DB.Collection("mitre_attack").
		FindOneAndReplace(context.Background(),
			bson.D{{Key: "id", Value: attackPattern.STIX_ID}},
			attackPattern,
			&options.FindOneAndReplaceOptions{},
		)
	if result != nil {
		return attackPattern
	}
	return nil
}
