package models

type Patient struct {
	ID        		uint64    `gorm:"primary_key;auto_increment" json:"id"`
	Pathogenic  	bool	  `gorm:"" json:"pathogenic"`
	Gene			string	  `gorm:"" json:"gene"`
	Ethnicity		string	  `gorm:"" json:"ethnicity"`
	CancerDx		bool	  `gorm:"" json:"cancer_dx"`
	CancerDxType	string	  `gorm:"" json:"cancer_dx_type"`
	CancerDxAge		int		  `gorm:"" json:"cancer_dx_age"`
	KnownCancer		string	  `gorm:"" json:"known_cancer"`
	RelRelation		[]string  `gorm:"" json:"rel_relation"`
	RelCancer		[]string  `gorm:"" json:"rel_cancer"`
	RelAge			[]string  `gorm:"" json:"rel_age"`
	IndexRel		float64   `gorm:"" json:"index_rel"`
	Prediction		float64	  `gorm:"" json:"prediction"`
	User      		User      `json:"user"`
	UserID    		uint32    `gorm:"not null" json:"user_id"`
}