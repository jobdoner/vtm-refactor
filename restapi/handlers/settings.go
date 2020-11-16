package handlers

import (
	"context"

	"github.com/go-openapi/runtime/middleware"
	"github.com/jobdoner/vtm-refactor/models"
	"github.com/jobdoner/vtm-refactor/restapi/operations"
)

func (h *Handler) GetSettings(params operations.GetSettingsParams) middleware.Responder {
	resp, err := h.RedisConn.LRange(context.Background(), params.ListName, 0, -1).Result()
	if err != nil {
		return operations.NewGetSettingsOK()
	}
	return operations.NewGetSettingsOK().WithPayload(models.GetSettingsResp(resp))
}

func (h *Handler) GetSettingObject(params operations.GetSettingsObjectParams) middleware.Responder {
	data, _ := h.RedisConn.HGetAll(context.Background(), "adgroup:settings").Result()
	respmap := make([]*models.GetSettingsObjectRespItems0, len(data))
	for key := range respmap {
		respmap[key] = &models.GetSettingsObjectRespItems0{}
	}
	for key := range data {
		respmap[adGroupsParameters[key]] = &models.GetSettingsObjectRespItems0{}
		result, err := h.RedisConn.HGetAll(context.Background(), "adgroup:"+key).Result()
		if len(result) > 0 {
			respmap[adGroupsParameters[key]].Name = data[key]
			respmap[adGroupsParameters[key]].Key = key
			respmap[adGroupsParameters[key]].Options = result
		} else {
			respmap[adGroupsParameters[key]].Name = data[key]
			respmap[adGroupsParameters[key]].Key = key
		}

		if err != nil {
			return operations.NewGetSettingsObjectOK().WithPayload([]*models.GetSettingsObjectRespItems0{
				{
					Options: err,
				},
			})
		}
	}
	return operations.NewGetSettingsObjectOK().WithPayload(respmap)
}

func (h *Handler) GetTechObject(params operations.GetSettingsTechParams) middleware.Responder {
	respMap := make([]*models.GetSettingsObjectRespItems0, len(techSettings))
	var err error
	i := 0
	for key := range techSettings {
		respMap[i] = &models.GetSettingsObjectRespItems0{}
		respMap[i].Key = key
		respMap[i].Name = techSettings[key]
		respMap[i].Options, err = h.RedisConn.HGetAll(context.Background(), "adgroup:"+key).Result()
		if err != nil {
			return operations.NewGetSettingsTechOK().WithPayload([]*models.GetSettingsObjectRespItems0{
				{
					Options: err,
				},
			})
		}
		i++
	}
	return operations.NewGetSettingsTechOK().WithPayload(respMap)
}

var techSettings = map[string]string{
	"budgetType":   "Budget Type",
	"bidStrategy":  "Bid Strategy",
	"campaignType": "Campaign Type",
}
var adGroupsParameters = map[string]int{
	"location":       0,
	"age":            1,
	"gender":         2,
	"topics":         3,
	"languages":      4,
	"keyword":        5,
	"channelID":      6,
	"adSchedule":     7,
	"audience":       8,
	"customAudience": 9,
	"devices":        10,
	"placement":      11,
	//"BLSSplit":       13,
	//"frequency":      14,
}
