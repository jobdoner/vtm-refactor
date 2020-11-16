package handlers

import (
	"fmt"
	"strconv"

	"github.com/go-pg/pg"
	"github.com/jobdoner/vtm-refactor/models"
	"github.com/jobdoner/vtm-refactor/restapi/operations"

	"github.com/go-openapi/runtime/middleware"
)

func (h *Handler) GetAdGroupFunc(params operations.GetAdgroupParams) middleware.Responder {
	adGroups := []models.AdGroup{}
	q := h.DB.Model(&adGroups)
	if params.FolderID != nil {
		q.Where("folder_id=?", params.FolderID)
	}
	err := q.Select()
	if err != nil {
		return operations.NewGetAdgroupOK()
	}

	respAdGroups := &models.GetAdGroupResp{}
	respAdGroups.Data = make([]*models.GetAdGroupRespDataItems0, len(adGroups))
	for key := range adGroups {
		respAdGroups.Data[key] = &models.GetAdGroupRespDataItems0{}
		respAdGroups.Data[key].ID = adGroups[key].ID
		respAdGroups.Data[key].TechAdGroupID = adGroups[key].TechAdgroupId
		respAdGroups.Data[key].FolderID = adGroups[key].FolderID
		respAdGroups.Data[key].AdGroupID = iflen0([]string{})
		respAdGroups.Data[key].Keyword = iflen0(adGroups[key].Keyword)
		respAdGroups.Data[key].Gender = iflen0(adGroups[key].Gender)
		respAdGroups.Data[key].Age = iflen0(adGroups[key].Age)
		respAdGroups.Data[key].ChannelID = iflen0(adGroups[key].ChannelID)
		respAdGroups.Data[key].Audience = iflen0(adGroups[key].Audience)
		respAdGroups.Data[key].CustomAudience = iflen0(adGroups[key].CustomAudience)
		respAdGroups.Data[key].Location = iflen0(adGroups[key].Location)
		respAdGroups.Data[key].Devices = iflen0(adGroups[key].Devices)
		respAdGroups.Data[key].Placement = iflen0(adGroups[key].Placement)
		respAdGroups.Data[key].Topics = iflen0(adGroups[key].Topics)
		respAdGroups.Data[key].Interest = iflen0(adGroups[key].Interest)

		respAdGroups.Data[key].Frequency = iflen0([]string{})
		respAdGroups.Data[key].Languages = iflen0(adGroups[key].Languages)
		respAdGroups.Data[key].AdSchedule = iflen0(adGroups[key].AdSchedule)
	}

	return operations.NewGetAdgroupOK().WithPayload(respAdGroups)
}

func iflen0(s []string) []string {
	if len(s) == 0 {
		return []string{}
	}
	return s
}
func (h *Handler) CreateAd(params operations.CreateAdParams) middleware.Responder {
	newAd := models.Ad{}
	newAd.AdGroupID = params.Body.Data.AdGroupID
	newAd.AdName = params.Body.Data.AdName
	newAd.VideoID = params.Body.Data.VideoID
	newAd.DisplayURL = params.Body.Data.DisplayURL
	newAd.CampaignStatus = params.Body.Data.CampaignStatus
	newAd.AdGroupStatus = params.Body.Data.AdGroupStatus
	newAd.Status = params.Body.Data.Status
	newAd.ApprovalStatus = params.Body.Data.ApprovalStatus
	newAd.Comment = params.Body.Data.Comment
	newAd.FinalMobileURL = params.Body.Data.FinalMobileURL
	newAd.FinalURL = params.Body.Data.FinalURL
	newAd.BumperAd = params.Body.Data.BumperAd

	_, err := h.DB.Model(&newAd).Insert()

	if err != nil {
		return operations.NewCreateAdOK().WithPayload(&models.CreateAdResp{
			// ErrorCode: "iternal",
			// ErrorMess: err.Error(),
		})
	}

	return operations.NewCreateAdOK().WithPayload(&models.CreateAdResp{
		Data: &models.CreateAdRespData{
			ID:             *newAd.ID,
			AdGroupID:      newAd.AdGroupID,
			AdName:         newAd.AdName,
			VideoID:        newAd.VideoID,
			DisplayURL:     newAd.DisplayURL,
			CampaignStatus: newAd.CampaignStatus,
			AdGroupStatus:  newAd.AdGroupStatus,
			Status:         newAd.Status,
			ApprovalStatus: newAd.ApprovalStatus,
			Comment:        newAd.Comment,
			FinalMobileURL: newAd.FinalMobileURL,
			FinalURL:       newAd.FinalURL,
			BumperAd:       newAd.BumperAd,
		},
	})
}

func (h *Handler) GetAd(params operations.GetAdParams) middleware.Responder {
	getAds := []models.Ad{}
	err := h.DB.Model(&getAds).Select()
	if err != nil {
		return operations.NewGetAdOK().WithPayload(&models.GetAdResp{
			// ErrorCode: "iternal",
			// ErrorMess: err.Error(),
		})
	}
	resp := &models.GetAdResp{}
	resp.Data = make([]*models.GetAdRespDataItems0, len(getAds))
	for key := range getAds {
		resp.Data[key] = &models.GetAdRespDataItems0{}
		resp.Data[key].ID = *getAds[key].ID
		resp.Data[key].AdGroupID = getAds[key].AdGroupID
		resp.Data[key].AdName = getAds[key].AdName
		resp.Data[key].VideoID = getAds[key].VideoID
		resp.Data[key].DisplayURL = getAds[key].DisplayURL
		resp.Data[key].CampaignStatus = getAds[key].CampaignStatus
		resp.Data[key].AdGroupStatus = getAds[key].AdGroupStatus
		resp.Data[key].Status = getAds[key].Status
		resp.Data[key].ApprovalStatus = getAds[key].ApprovalStatus
		resp.Data[key].Comment = getAds[key].Comment
		resp.Data[key].FinalURL = getAds[key].FinalURL
		resp.Data[key].FinalMobileURL = getAds[key].FinalMobileURL
		resp.Data[key].BumperAd = getAds[key].BumperAd
	}
	return operations.NewGetAdOK().WithPayload(resp)
}

func (h *Handler) ShowAdGroup(params operations.ShowAdgroupParams) middleware.Responder {
	AdGroup := models.AdGroup{}
	err := h.DB.Model(&AdGroup).Where("id = ?", params.ID).Select()

	if err != nil {
		return operations.NewShowAdgroupOK().WithPayload(&models.ShowAdGroupResp{
			// ErrorCode: "internal",
			// ErrorMess: err.Error(),
		})
	}

	return operations.NewShowAdgroupOK().WithPayload(&models.ShowAdGroupResp{Data: &models.ShowAdGroupRespData{
		ID:       int64(AdGroup.ID),
		FolderID: AdGroup.FolderID,
	}})
}

func (h *Handler) CreateTechAdGroup(params operations.CreateTechAdGroupParams) middleware.Responder {
	return middleware.NotImplemented("not yet")
}

func (h *Handler) CreateAdGroupFunc(params operations.CreateAdGroupParams) middleware.Responder {
	newAdGroup := make([]models.AdGroup, len(params.Body))
	for key := range params.Body {
		newAdGroup[key].ID = params.Body[key].ID
		newAdGroup[key].FolderID = params.Body[key].FolderID
		newAdGroup[key].TechAdgroupId = params.Body[key].TechAdGroupID
		newAdGroup[key].Keyword = params.Body[key].Keyword
		newAdGroup[key].Gender = params.Body[key].Gender
		newAdGroup[key].Age = params.Body[key].Age
		newAdGroup[key].ChannelID = params.Body[key].ChannelID
		newAdGroup[key].Audience = params.Body[key].Audience
		newAdGroup[key].CustomAudience = params.Body[key].CustomAudience
		newAdGroup[key].Devices = params.Body[key].Devices
		newAdGroup[key].Placement = params.Body[key].Placement
		//TODO: if location > 1 => create campaign for that ?
		newAdGroup[key].Location = params.Body[key].Location
		newAdGroup[key].Topics = params.Body[key].Topics
		newAdGroup[key].Interest = params.Body[key].Interest
		newAdGroup[key].Languages = params.Body[key].Languages
		newAdGroup[key].AdSchedule = params.Body[key].AdSchedule

	}

	_, err := h.DB.Model(&newAdGroup).OnConflict("(id) DO UPDATE").
		Set("keyword=EXCLUDED.keyword").
		Set("folder_id=EXCLUDED.folder_id").
		Set("tech_adgroup_id=EXCLUDED.tech_adgroup_id").
		Set("gender=EXCLUDED.gender").
		Set("age=EXCLUDED.age").
		Set("channel_id=EXCLUDED.channel_id").
		Set("audience=EXCLUDED.audience").
		Set("custom_audience=EXCLUDED.custom_audience").
		Set("devices=EXCLUDED.devices").
		Set("placement=EXCLUDED.placement").
		Set("location=EXCLUDED.location").
		Set("languages=EXCLUDED.languages").
		Set("topics=EXCLUDED.topics").
		Set("ad_schedule=EXCLUDED.ad_schedule").
		Set("interest=EXCLUDED.interest").Insert()

	for key := range newAdGroup {
		newAdGroup[key].Ad = []models.Ad{}
		err := h.DB.Model(&newAdGroup[key].Ad).Where("adgroup_id=?", newAdGroup[key].ID).Select()
		if err != nil {
			fmt.Println(err)
		}
		if len(newAdGroup[key].Ad) == 0 {
			newAdGroup[key].Ad = make([]models.Ad, 1)
			newAdGroup[key].Ad[0] = models.Ad{}
		}

		newAdGroup[key].Ad[0].AdGroupID = newAdGroup[key].ID
		newAdGroup[key].Ad[0].AdName = concatAdGroupIdFromCampaignName(params.Body[key].CampaignID, key, h.DB) // concatAdGroupId(params.Body[key].CampaignID, newAdGroup[key], key)

		_, err = h.DB.Model(&newAdGroup[key].Ad[0]).
			OnConflict("(id) DO update").
			Set("ad_name=EXCLUDED.ad_name").
			Insert()
		if err != nil {
			fmt.Println(err)
		}
	}

	if err != nil {
		return operations.NewCreateAdGroupOK().WithPayload(&models.CreateAdGroupResp{
			// ErrorCode: "internal",
			// ErrorMess: err.Error(),
		})
	}
	resp := &models.CreateAdGroupResp{}
	resp.Data = make([]*models.CreateAdGroupRespDataItems0, len(newAdGroup))
	for key := range newAdGroup {
		resp.Data[key] = &models.CreateAdGroupRespDataItems0{}
		resp.Data[key].ID = newAdGroup[key].ID
		resp.Data[key].FolderID = newAdGroup[key].FolderID
		resp.Data[key].TechAdGroupID = newAdGroup[key].TechAdgroupId
		resp.Data[key].AdGroupID = iflen0(newAdGroup[key].AdGroupID)
		resp.Data[key].Keyword = iflen0(newAdGroup[key].Keyword)
		resp.Data[key].Gender = iflen0(newAdGroup[key].Gender)
		resp.Data[key].Age = iflen0(newAdGroup[key].Age)
		resp.Data[key].ChannelID = iflen0(newAdGroup[key].ChannelID)
		resp.Data[key].Audience = iflen0(newAdGroup[key].Audience)
		resp.Data[key].CustomAudience = iflen0(newAdGroup[key].CustomAudience)
		resp.Data[key].Location = iflen0(newAdGroup[key].Location)
		resp.Data[key].Devices = iflen0(newAdGroup[key].Devices)
		resp.Data[key].Placement = iflen0(newAdGroup[key].Placement)
		resp.Data[key].Topics = iflen0(newAdGroup[key].Topics)
		resp.Data[key].Interest = iflen0(newAdGroup[key].Interest)
		resp.Data[key].Languages = iflen0(newAdGroup[key].Languages)
		resp.Data[key].AdSchedule = iflen0(newAdGroup[key].AdSchedule)
	}
	return operations.NewCreateAdGroupOK().WithPayload(resp)
}

func concatAdGroupIdFromCampaignName(campaignID int64, key int, db *pg.DB) string {
	c := models.Campaign{ID: int(campaignID)}
	err := db.Model(&c).WherePK().Select()
	if err != nil {
		fmt.Println(err)
	}
	return c.Campaign + "AdGroup_" + strconv.Itoa(key)
}
