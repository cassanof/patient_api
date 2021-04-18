package models

import "github.com/lib/pq"

type Patient struct {
	ID           uint64         `gorm:"primary_key;auto_increment" json:"id"`
	Pathogenic   bool           `gorm:"default:false" json:"pathogenic"`
	Gene         string         `gorm:"size:100" json:"gene"`
	Ethnicity    string         `gorm:"size:255" json:"ethnicity"`
	CancerDx     string         `gorm:"size:100" json:"cancer_dx"`
	CancerDxType string         `gorm:"size:255" json:"cancer_dx_type"`
	CancerDxAge  int64          `json:"cancer_dx_age"`
	KnownBrca    string         `gorm:"size:255" json:"known_brca"`
	KnownCancer  string         `gorm:"size:100" json:"known_cancer"`
	RelRelation  pq.StringArray `gorm:"type:string[]" json:"rel_relation"`
	RelCancer    pq.StringArray `gorm:"type:string[]" json:"rel_cancer"`
	RelAge       pq.StringArray `gorm:"type:string[]" json:"rel_age"`
	Prediction   float64        `gorm:"" json:"prediction"`
	User         User           `json:"user"`
	UserID       uint32         `gorm:"not null" json:"user_id"`
}
