package models

type Patient struct {
	ID        		uint64    `gorm:"primary_key;auto_increment" json:"id"`
	Pathogenic  	bool	  `gorm:"" json:"pathogenic"`
	Gene			string	  ``
	Ethnicity		string	  ``
	CancerDx		bool
	CancerDxType	string
	CancerDxAge		int
	KnownCancer		string
	RelRelation		[]string
	RelCancer		[]string
	RelAge
	IndexRel
	Prediction
	User      User      `json:"user"`
	UserID    uint32    `gorm:"not null" json:"user_id"`
}