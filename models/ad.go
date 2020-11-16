package models

type Ad struct {
	tableName      string `sql:"ad"`
	ID             *int64  `sql:"id"`
	AdGroupID      int64  `sql:"adgroup_id"`
	AdName         string `sql:"ad_name"`
	VideoID        string `sql:"video_id"`
	DisplayURL     string `sql:"display_url"`
	CampaignStatus string `sql:"campaign_status"`
	AdGroupStatus  string `sql:"ad_group_status"`
	Status         string `sql:"status"`
	ApprovalStatus string `sql:"approval_status"`
	Comment        string `sql:"comment"`
	FinalURL       string `sql:"final_url"`
	FinalMobileURL string `sql:"final_mobile_url"`
	BumperAd       string `sql:"bumper_ad"`
	Networks       string `sql:"-"`
	//NonskippableAd string `sql:"nonskippable_ad"`
}
