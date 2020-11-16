package models

type AdGroup struct {
	tableName      struct{} `sql:"adgroup" csv:"-"`
	Ad             []Ad     `sql:"-" csv:"-"`
	ID             int64    `sql:"id" csv:"-"`
	FolderID       int64    `sql:"folder_id" csv:"FolderID"`
	TechAdgroupId  int64    `sql:"tech_adgroup_id" csv:"TechAdGroupID"`
	AdGroupID      []string `sql:"adgroup_id,array" csv:"AdGroupID"`
	Keyword        []string `sql:"keyword,array" csv:"Keyword"`
	Gender         []string `sql:"gender,array" csv:"Gender"`
	Age            []string `sql:"age,array" csv:"Age"`
	ChannelID      []string `sql:"channel_id,array" csv:"ChannelID"`
	Languages      []string `sql:"languages,array" csv:"Language"`
	Audience       []string `sql:"audience,array" csv:"Audience"`
	CustomAudience []string `sql:"custom_audience,array" csv:"CustomAudience"`
	Devices        []string `sql:"devices,array" csv:"Devices"`
	Placement      []string `sql:"placement,array" csv:"Placement"`
	Location       []string `sql:"location,array" csv:"Location"`
	Topics         []string `sql:"topics,array" csv:"Topics"`
	Interest       []string `sql:"interest,array" csv:"Interest"`
	AdSchedule     []string `sql:"ad_schedule,array" csv:"Placement"`
	//LocationIDS []string `sql:"-"`
}
type ExportAdgroup struct {
	ID             int64  `sql:"id" csv:"ID"`
	FolderID       int64  `sql:"folder_id" csv:"FolderID"`
	TechAdgroupId  int64  `sql:"tech_adgroup_id" csv:"TechAdGroupID"`
	AdGroupID      string `sql:"adgroup_id,array" csv:"AdGroupID"`
	Keyword        string `sql:"keyword,array" csv:"Keyword"`
	Gender         string `sql:"gender,array" csv:"Gender"`
	Age            string `sql:"age,array" csv:"Age"`
	ChannelID      string `sql:"channel_id,array" csv:"ChannelID"`
	Languages      string `sql:"languages,array" csv:"Language"`
	Audience       string `sql:"audience,array" csv:"Audience"`
	CustomAudience string `sql:"custom_audience,array" csv:"CustomAudience"`
	Devices        string `sql:"devices,array" csv:"Devices"`
	Placement      string `sql:"placement,array" csv:"Placement"`
	Location       string `sql:"location,array" csv:"Location"`
	Topics         string `sql:"topics,array" csv:"Topics"`
	Interest       string `sql:"interest,array" csv:"Interest"`
	AdSchedule     string `sql:"ad_schedule,array" csv:"AdSchedule"`
}

type TechAdGroup struct {
	tableName             struct{}  `sql:"tech_adgroup"`
	Adg                   []AdGroup `sql:"-"`
	ID                    int       `sql:"id"`
	CampaignID            int64     `sql:"campaign_id"`
	AdGroup               string    `sql:"ad_group"`
	MaxCPC                float64   `sql:"max_cpc"`
	MaxCPM                float64   `sql:"max_cpm"` //- 0.01
	TargetCPA             string    `sql:"target_cpa"`
	MaxCPV                float64   `sql:"max_cpv"`    //- 0.01
	TargetCPM             float64   `sql:"target_cpm"` //- 50.00
	TargetROAS            string    `sql:"target_roas"`
	DesktopBidModifier    string    `sql:"desktop_bid_modifier"`
	MobileBidModifier     string    `sql:"mobile_bid_modifier"`
	TabletBidModifier     string    `sql:"tablet_bid_modifier"`
	TVScreenBidModifier   string    `sql:"tv_screen_bid_modifier"`
	TopContentBidModifier string    `sql:"top_content_bid_modifier"`
	DisplayNetwork        string    `sql:"display_network"` //Keywords
	CustomBidType         string    `sql:"custom_bid_type"`
	FirstPageBid          string    `sql:"first_page_bid"`
	TopOfPageBid          string    `sql:"top_of_page_bid"`
	TargetingOptimization string    `sql:"targeting_optimization"` //- Disabled
	//ContentKeywords       string  `sql:"content_keywords"`       //- Disabled
	AdGroupType      string `sql:"ad_group_type"` //- Non-skippable
	TrackingTemplate string `sql:"tracking_template"`
	FinalURLSuffix   string `sql:"final_url_suffix"`
	CustomParameters string `sql:"custom_parameters"`
	BidModifier      string `sql:"bid_modifier"`
	DestinationURL   string `sql:"destination_url"`
	CriterionType    string `sql:"criterion_type"`
	Reach            string `sql:"reach"`
	Feed             string `sql:"feed"`
	Radius           string `sql:"radius"`
	Unit             string `sql:"unit"`
}
