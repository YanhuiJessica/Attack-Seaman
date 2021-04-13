package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ExternalReference struct {
	SourceName string `bson:"source_name" json:"source_name"`
	ExternalID string `bson:"external_id" json:"external_id"`
	URL        string `bson:"url" json:"url"`
}

// Base is
type Base struct {
	ID primitive.ObjectID `bson:"_id,omitempty" json:"_id"`
	// ObjectMarkingRefs  []string            `json:"object_marking_refs"`
	// CreatedByRef       string              `bson:"created_by_ref"`
	STIX_ID  string    `bson:"id,omitempty" json:"id,omitempty"`
	Type     string    `bson:"type" json:"type"`
	Created  time.Time `bson:"created" json:"created"`
	Modified time.Time `bson:"modified" json:"modified"`
}

type Relationship struct {
	Base             `bson:",inline"`
	SourceRef        string `bson:"source_ref" json:"source_ref"`
	RelationshipType string `bson:"relationship_type" json:"relationship_type"`
	TargetRef        string `bson:"target_ref" json:"target_ref"`
}

type KillChainPhase struct {
	KillChainName string `bson:"kill_chain_name" json:"kill_chain_name"`
	PhaseName     string `bson:"phase_name" json:"phase_name" `
}

type BaseStix struct {
	Base               `bson:",inline"`
	Name               string              `bson:"name" json:"name"`
	Description        string              `bson:"description" json:"description"`
	ExternalReferences []ExternalReference `bson:"external_references" json:"external_references"`
}

type XMitreTactic struct {
	BaseStix        `bson:",inline"`
	XMitreShortName string `bson:"x_mitre_shortname" json:"x_mitre_shortname"`
}

type AttackPattern struct {
	BaseStix                  `bson:",inline"`
	KillChainPhases           []KillChainPhase `bson:"kill_chain_phases" json:"kill_chain_phases"`
	XMitreIsSubtechnique      bool             `bson:"x_mitre_is_subtechnique" json:"x_mitre_is_subtechnique"`
	XMitreVersion             string           `bson:"x_mitre_version,omitempty" json:"x_mitre_version,omitempty"`
	XMitreDetection           string           `bson:"x_mitre_detection,omitempty" json:"x_mitre_detection,omitempty"`
	XMitrePermissionsRequired []string         `bson:"x_mitre_permissions_required,omitempty" json:"x_mitre_permissions_required,omitempty"`
	XMitreDataSources         []string         `bson:"x_mitre_data_sources,omitempty" json:"x_mitre_data_sources,omitempty"`
	XMitrePlatforms           []string         `bson:"x_mitre_platforms,omitempty" json:"x_mitre_platforms,omitempty"`
}

// New is
func (ap *AttackPattern) New() *AttackPattern {
	return &AttackPattern{
		BaseStix:                  ap.BaseStix,
		KillChainPhases:           ap.KillChainPhases,
		XMitreVersion:             ap.XMitreVersion,
		XMitreIsSubtechnique:      ap.XMitreIsSubtechnique,
		XMitreDetection:           ap.XMitreDetection,
		XMitrePermissionsRequired: ap.XMitrePermissionsRequired,
		XMitreDataSources:         ap.XMitreDataSources,
		XMitrePlatforms:           ap.XMitrePlatforms,
	}
}

// New is
func (r *Relationship) New() *Relationship {
	return &Relationship{
		Base:             r.Base,
		SourceRef:        r.SourceRef,
		RelationshipType: r.RelationshipType,
		TargetRef:        r.TargetRef,
	}
}

// New is
func (b *BaseStix) New() *BaseStix {
	return &BaseStix{
		Base:               b.Base,
		Name:               b.Name,
		Description:        b.Description,
		ExternalReferences: b.ExternalReferences,
	}
}

// New is
func (b *Base) New() *Base {
	return &Base{
		STIX_ID:  b.STIX_ID,
		Type:     b.Type,
		Created:  time.Now(),
		Modified: time.Now(),
	}
}
