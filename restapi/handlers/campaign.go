package handlers

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/jobdoner/vtm-refactor/models"
	"github.com/jobdoner/vtm-refactor/restapi/operations"
)

//TODO find what for this func exist?
func GetCampaignType(T string) []string {
	switch T {
	case "Skippable":
		return []string{"Video", "InStream", "", "Target CPM"}
	case "Bumper":
		return []string{"Video", "Bumper", "Bumper ad []", ""}
	case "Non-skipp":
		return []string{"Video - Non-skippable", "Non-skippable", "Nonskippable ad []", "Target CPM", "Target CPM"}
	case "Out stream":
		return []string{"Video - Outstream", "Outstream", ""}
	}
	return []string{"", "", "", ""}
}

//Search
//Display
//Display - Smart
//Display - Gmail
//Display - Mobile app installs
//Shopping
//Shopping - Smart
//Video
//Video - Drive convesions
//Video - Non-skippable
//Video - Outstream
//Universal app\
func (h *Handler) CreateCampaignFunc(params operations.CreateCampaignParams) middleware.Responder {

	ct := GetCampaignType(params.Body.CampaignType)

	newCampaign := models.Campaign{
		FolderID:     int(params.Body.FolderID),
		Account:      params.Body.Account,
		Campaign:     params.Body.Campaign,
		Budget:       params.Body.Budget,
		BudgetType:   params.Body.BudgetType,
		CampaignType: ct[0],
		Networks:     params.Body.Networks,
		Languages:    params.Body.Languages,
		BidStrategyType: func() string {
			if ct[3] == "" {
				return params.Body.BidStrategyType
			}
			return ct[3]
		}(),
		BidStrategyName:    params.Body.BidStrategyName,
		StartDate:          params.Body.StartDate,
		EndDate:            params.Body.EndDate,
		AdSchedule:         params.Body.AdSchedule,
		AdRotation:         params.Body.AdRotation,
		InventoryType:      params.Body.InventoryType,
		DeliveryMethod:     params.Body.DeliveryMethod,
		TargetingMethod:    params.Body.TargetingMethod,
		ExclusionMethod:    params.Body.ExclusionMethod,
		DSAWebsite:         params.Body.DsaWebsite,
		DSALanguage:        params.Body.DsaLanguage,
		DSATargetingSource: params.Body.DsaTargetingSource,
		DSAPageFeeds:       params.Body.DsaPageFeeds,
		FlexibleReach:      params.Body.FlexibleReach,
		Location:           params.Body.Location,
	}

	_, err := h.DB.Model(&newCampaign).Insert()
	if err != nil {
		return operations.NewCreateCampaignInternalServerError()
		// ErrorCode: string(http.StatusInternalServerError),
		// ErrorMess: err.Error(),

	}

	//Set defoult values for TechAdGroup
	newTechAdGroup := models.TechAdGroup{
		CampaignID:  int64(newCampaign.ID),
		MaxCPC:      params.Body.MaxCPC,
		MaxCPM:      params.Body.MaxCPM,
		TargetCPA:   params.Body.TargetCPA,
		MaxCPV:      params.Body.MaxCPV,
		TargetCPM:   params.Body.TargetCPM,
		AdGroupType: ct[1],
	}

	_, err = h.DB.Model(&newTechAdGroup).Insert()

	if err != nil {
		return operations.NewCreateTechAdGroupOK()
	}

	return operations.NewCreateCampaignOK().WithPayload(&models.CreateCampaignResp{
		Data: &models.CreateCampaignRespData{
			CampaignID:    int64(newCampaign.ID),
			TechAdGroupID: int64(newTechAdGroup.ID),
		},
	})
}

func (h *Handler) GetCampaignFunc(params operations.GetCampaignParams) middleware.Responder {
	campaigns := []models.Campaign{}
	err := h.DB.Model(&campaigns).Select()
	if err != nil {
		return operations.NewGetCampaignOK()
	}
	resp := &models.GetCampaignResp{}
	resp.Data = make([]*models.GetCampaignRespDataItems0, len(campaigns))
	for key := range campaigns {
		resp.Data[key] = &models.GetCampaignRespDataItems0{}
		resp.Data[key].ID = int64(campaigns[key].ID)
		resp.Data[key].Name = campaigns[key].Campaign
	}

	return operations.NewGetCampaignOK().WithPayload(resp)
}

func (h *Handler) UpdateCampaign(params operations.PatchCampaignIDParams) middleware.Responder {
	UpdatedCampaign := models.Campaign{}
	UpdatedCampaign.ID = int(params.ID)

	updateQ := h.DB.Model(&UpdatedCampaign)

	if params.Body.Account != "" {
		updateQ.Column("account")
		UpdatedCampaign.Account = params.Body.Account
	}
	if params.Body.Campaign != "" {
		updateQ.Column("campaign")
		UpdatedCampaign.Campaign = params.Body.Campaign
	}
	if params.Body.Budget != 0 {
		updateQ.Column("budget")
		UpdatedCampaign.Budget = params.Body.Budget
	}
	if params.Body.BudgetType != "" {
		updateQ.Column("budget_type")
		UpdatedCampaign.BudgetType = params.Body.BudgetType
	}
	if params.Body.CampaignType != "" {
		updateQ.Column("campaign_type")
		UpdatedCampaign.CampaignType = params.Body.CampaignType
	}
	if params.Body.Networks != "" {
		updateQ.Column("networks")
		UpdatedCampaign.Networks = params.Body.Networks
	}
	if params.Body.Languages != "" {
		updateQ.Column("languages")
		UpdatedCampaign.Languages = params.Body.Languages
	}
	if params.Body.BidStrategyType != "" {
		updateQ.Column("bid_strategy_type")
		UpdatedCampaign.BidStrategyType = params.Body.BidStrategyType
	}
	if params.Body.BidStrategyName != "" {
		updateQ.Column("bid_strategy_name")
		UpdatedCampaign.BidStrategyName = params.Body.BidStrategyName
	}
	if params.Body.StartDate != "" {
		updateQ.Column("start_date")
		UpdatedCampaign.StartDate = params.Body.StartDate
	}
	if params.Body.EndDate != "" {
		updateQ.Column("end_date")
		UpdatedCampaign.EndDate = params.Body.EndDate
	}
	if params.Body.AdSchedule != "" {
		updateQ.Column("ad_schedule")
		UpdatedCampaign.AdSchedule = params.Body.AdSchedule
	}
	if params.Body.AdRotation != "" {
		updateQ.Column("ad_rotation")
		UpdatedCampaign.AdRotation = params.Body.AdRotation
	}
	if params.Body.InventoryType != "" {
		updateQ.Column("inventory_type")
		UpdatedCampaign.InventoryType = params.Body.InventoryType
	}
	if params.Body.DeliveryMethod != "" {
		updateQ.Column("delivery_method")
		UpdatedCampaign.DeliveryMethod = params.Body.DeliveryMethod
	}
	if params.Body.TargetingMethod != "" {
		updateQ.Column("targeting_method")
		UpdatedCampaign.TargetingMethod = params.Body.TargetingMethod
	}
	if params.Body.ExclusionMethod != "" {
		updateQ.Column("exclusion_method")
		UpdatedCampaign.ExclusionMethod = params.Body.ExclusionMethod
	}
	if params.Body.DsaWebsite != "" {
		updateQ.Column("dsa_website")
		UpdatedCampaign.DSAWebsite = params.Body.DsaWebsite
	}
	if params.Body.DsaLanguage != "" {
		updateQ.Column("dsa_language")
		UpdatedCampaign.DSALanguage = params.Body.DsaLanguage
	}
	if params.Body.DsaTargetingSource != "" {
		updateQ.Column("dsa_targeting_source")
		UpdatedCampaign.DSATargetingSource = params.Body.DsaTargetingSource
	}
	if params.Body.DsaPageFeeds != "" {
		updateQ.Column("dsa_page_feeds")
		UpdatedCampaign.DSAPageFeeds = params.Body.DsaPageFeeds
	}
	if params.Body.FlexibleReach != "" {
		updateQ.Column("flexible_reach")
		UpdatedCampaign.FlexibleReach = params.Body.FlexibleReach
	}
	_, err := updateQ.Returning("id").WherePK().Update()

	if err != nil {
		return operations.NewPatchCampaignIDOK()
	}

	return operations.NewPatchCampaignIDOK().WithPayload(&models.CreateCampaignResp{Data: &models.CreateCampaignRespData{}})
}

func (h *Handler) ShowCampaign(params operations.GetCampaignIDParams) middleware.Responder {
	Campaign := models.Campaign{}

	err := h.DB.Model(&Campaign).Where("id = ?", params.ID).Select()
	if err != nil {
		return operations.NewGetCampaignIDOK()
	}
	return operations.NewGetCampaignIDOK().WithPayload(&models.ShowCampaignResp{
		Data: &models.ShowCampaignRespData{
			Account:            Campaign.Account,
			AdRotation:         Campaign.AdRotation,
			AdSchedule:         Campaign.AdSchedule,
			BidStrategyName:    Campaign.BidStrategyName,
			BidStrategyType:    Campaign.BidStrategyType,
			Budget:             Campaign.Budget,
			BudgetType:         Campaign.BudgetType,
			Campaign:           Campaign.Campaign,
			CampaignType:       Campaign.CampaignType,
			DeliveryMethod:     Campaign.DeliveryMethod,
			DsaLanguage:        Campaign.DSALanguage,
			DsaPageFeeds:       Campaign.DSAPageFeeds,
			DsaTargetingSource: Campaign.DSATargetingSource,
			DsaWebsite:         Campaign.DSAWebsite,
			EndDate:            Campaign.EndDate,
			ExclusionMethod:    Campaign.ExclusionMethod,
			FlexibleReach:      Campaign.FlexibleReach,
			ID:                 int64(Campaign.ID),
			InventoryType:      Campaign.InventoryType,
			Languages:          Campaign.Languages,
			Networks:           Campaign.Networks,
			StartDate:          Campaign.StartDate,
			TargetingMethod:    Campaign.TargetingMethod,
		},
	})
}

func (h *Handler) DeleteCampaign(params operations.DeleteCampaignIDParams) middleware.Responder {
	DeletedCampaign := models.Campaign{ID: int(params.ID)}
	_, err := h.DB.Model(&DeletedCampaign).Returning("id").WherePK().Delete()

	if err != nil {
		return operations.NewDeleteCampaignIDOK()
	}
	return operations.NewDeleteCampaignIDOK()
}
