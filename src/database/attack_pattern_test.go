package database

import (
	"github.com/stretchr/testify/assert"
	"github.com/zGina/Attack-Seaman/src/model"
)

func (s *DatabaseSuite) TestCreateAttackPattern() {
	s.db.DB.Collection("mitre_attack").Drop(nil)

	killChainPhase := model.KillChainPhase{
		KillChainName: "mitre_attack",
		PhaseName:     "privilege-escalation",
	}

	externalReference := model.ExternalReference{
		SourceName: "mitre-attack",
		ExternalID: "T1546.004",
		URL:        "https://attack.mitre.org/techniques/T1546/004",
	}

	externalReference2 := model.ExternalReference{
		SourceName: "mitre-attack",
		ExternalID: "T1546.004",
		URL:        "https://attack.mitre.org/techniques/T1546/004",
	}

	technique := model.Base{

		STIX_ID: "id287487",
		Type:    "attack-pattern",
	}

	attackPattern := (&model.AttackPattern{
		Base:                 technique,
		ExternalReferences:   []model.ExternalReference{externalReference, externalReference2},
		Name:                 "我是哈哈",
		Description:          "我是嘻嘻",
		KillChainPhases:      []model.KillChainPhase{killChainPhase},
		XMitreIsSubtechnique: true,
	})

	err := s.db.CreateAttackPattern(attackPattern)
	assert.Nil(s.T(), err)
}

func (s *DatabaseSuite) TestGetAttackPatterns() {
	start := int64(0)
	limit := int64(10)
	sort := "_id"
	order := -1

	users := s.db.GetAttackPatterns(&model.Paging{
		Skip:      &start,
		Limit:     &limit,
		SortKey:   sort,
		SortVal:   order,
		Condition: nil,
	})

	assert.Len(s.T(), users, 6)
}

func (s *DatabaseSuite) TestGetAttackByName() {
	user := s.db.GetAttackPatternByName("我是ss哈哈")

	assert.Equal(s.T(), "我是ss哈哈", user.Name)
}

func (s *DatabaseSuite) TestGetAttackByIDs() {
	user := s.db.GetAttackPatternByStixID("id287487")

	println(user.Name)
}

func (s *DatabaseSuite) TestUpdateAttackByIDs() {
	attackPattern := s.db.GetAttackPatternByStixID("id287487")
	attackPattern.Name = "改变厚的哈哈哈哈哈哈哈"
	s.db.UpdateAttackPattern(attackPattern)
	println(attackPattern.Name)
}
