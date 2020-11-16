package models

type Campaign struct {
	tableName          struct{}    `sql:"campaign"`
	Tad                TechAdGroup `sql:"-"`
	ID                 int         `sql:"id"`
	FolderID           int         `sql:"folder_id"`
	Account            string      `sql:"account"`
	Campaign           string      `sql:"campaign"`
	Budget             float64     `sql:"budget"`
	BudgetType         string      `sql:"budget_type"`
	CampaignType       string      `sql:"campaign_type"`
	Networks           string      `sql:"networks"`
	Languages          string      `sql:"languages"`
	BidStrategyType    string      `sql:"bid_strategy_type"`
	BidStrategyName    string      `sql:"bid_strategy_name"`
	StartDate          string      `sql:"start_date"`
	EndDate            string      `sql:"end_date"`
	AdSchedule         string      `sql:"ad_schedule"`
	AdRotation         string      `sql:"ad_rotation"`
	InventoryType      string      `sql:"inventory_type"`
	DeliveryMethod     string      `sql:"delivery_method"`
	TargetingMethod    string      `sql:"targeting_method"`
	ExclusionMethod    string      `sql:"exclusion_method"`
	DSAWebsite         string      `sql:"dsa_website"`
	DSALanguage        string      `sql:"dsa_language"`
	DSATargetingSource string      `sql:"dsa_targeting_source"`
	DSAPageFeeds       string      `sql:"dsa_page_feeds"`
	FlexibleReach      string      `sql:"flexible_reach"`
	Location           string      `sql:"location"`
}
