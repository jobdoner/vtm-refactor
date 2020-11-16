package models

import "time"

type Folder struct {
	tableName struct{}  `sql:"folder"`
	ID        int       `sql:"id"`
	Name      string    `sql:"name,notnull"`
	Type      string    `sql:"folder_type,notnull"`
	RulesMask string    `sql:"rules_mask"`
	Frequency string    `sql:"frequency"`
	BLSSplit  bool      `sql:"bls_split"`
	DeletedAt time.Time `pg:",soft_delete"`
	CreatedAt time.Time `sql:"created_at"`
}

type MathingFolders struct {
	tableName        struct{}  `sql:"matching_folders"`
	ID               int64     `sql:"id"`
	VideoFolderID    int64     `sql:"video_folder_id"`
	AudienceFolderID int64     `sql:"audience_folder_id"`
	Name             string    `sql:"name,notnull"`
	DeletedAt        time.Time `pg:",soft_delete"`
	CreatedAt        time.Time `sql:"created_at"`
}
