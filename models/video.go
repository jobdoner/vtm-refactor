package models

type Video struct {
	tableName   struct{} `sql:"video"`
	ID          int64    `sql:"id"`
	FolderID    int64    `sql:"folder_id"`
	CreativeID  string   `sql:"creative_id"`
	YouTubeID   string   `sql:"youtube_id"`
	FilePath    string   `sql:"filepath"`
	Title       string   `sql:"title"`
	Description string   `sql:"description"`
}
