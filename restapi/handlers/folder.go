package handlers

import (
	"fmt"
	"strings"

	"github.com/go-openapi/runtime/middleware"
	"github.com/go-pg/pg"
	"github.com/jobdoner/vtm-refactor/models"
	"github.com/jobdoner/vtm-refactor/restapi/operations"
)

func (h *Handler) CreateFolderFunc(params operations.CreateFoldersParams) middleware.Responder {
	folder := models.Folder{
		Name:      params.Folder.Name,
		Type:      params.Folder.Type,
		RulesMask: params.Folder.RulesMask,
		BLSSplit:  params.Folder.BLSSplit,
		Frequency: params.Folder.Frequency,
	}
	_, err := h.DB.Model(&folder).Insert()
	if err != nil {
		return operations.NewCreateFoldersCreated().WithPayload(&models.CreateFolderResp{
			// ErrorCode: string(http.StatusInternalServerError),
			// ErrorMess: err.Error(),
		})
	}

	return operations.NewCreateFoldersCreated().WithPayload(&models.CreateFolderResp{
		Data: &models.CreateFolderRespData{
			ID:        int64(folder.ID),
			Type:      folder.Type,
			Name:      folder.Name,
			RulesMask: folder.RulesMask,
			BLSSplit:  params.Folder.BLSSplit,
			Frequency: params.Folder.Frequency,
		},
	})
}

func (h *Handler) GetFoldersFunc(params operations.GetFoldersParams) middleware.Responder {
	var Folders []models.Folder
	folderQ := h.DB.Model(&Folders)
	if params.Type != nil {
		folderQ.Where("folder_type=?", params.Type)
	}
	folderQ.Order("id DESC")
	err := folderQ.Select()

	if err != nil {
		return operations.NewGetFoldersOK().WithPayload(&models.GetFoldersResp{
			// ErrorCode: "internal",
			// ErrorMess: err.Error(),
		})
	}
	respFolders := &models.GetFoldersResp{}
	respFolders.Data = make([]*models.GetFoldersRespDataItems0, len(Folders))

	for key := range Folders {
		respFolders.Data[key] = &models.GetFoldersRespDataItems0{}
		respFolders.Data[key].ID = int64(Folders[key].ID)
		respFolders.Data[key].Name = Folders[key].Name
		respFolders.Data[key].Type = Folders[key].Type
		respFolders.Data[key].RulesMask = Folders[key].RulesMask
		respFolders.Data[key].Frequency = Folders[key].Frequency
		respFolders.Data[key].BLSSplit = Folders[key].BLSSplit
		respFolders.Data[key].CreatedAt = Folders[key].CreatedAt.Format("2006-01-02")
	}

	return operations.NewGetFoldersOK().WithPayload(respFolders)
}

func (h *Handler) UpdateFolder(params operations.PatchFoldersIDParams) middleware.Responder {
	UpdatedFolder := &models.Folder{}
	UpdatedFolder.ID = int(params.ID)
	updateQ := h.DB.Model(UpdatedFolder)

	if params.Body.Name != "" {
		UpdatedFolder.Name = params.Body.Name
		updateQ.Column("name")
	}

	_, err := updateQ.WherePK().Returning("id").UpdateNotZero()

	if err != nil {
		return operations.NewPatchFoldersIDOK().WithPayload(&models.CreateFolderResp{
			// ErrorCode: "iternal",
			// ErrorMess: err.Error(),
		})
	}

	return operations.NewPatchFoldersIDOK().WithPayload(&models.CreateFolderResp{Data: &models.CreateFolderRespData{
		ID: int64(UpdatedFolder.ID),
	}})
}

func (h *Handler) ShowFolder(params operations.GetFoldersIDParams) middleware.Responder {
	Folder := models.Folder{}
	err := h.DB.Model(&Folder).Where("id=?", params.ID).Select()

	campaign := models.Campaign{}
	err = h.DB.Model(&campaign).Where("folder_id=?", Folder.ID).Select()

	techAdGroup := models.TechAdGroup{}
	err = h.DB.Model(&techAdGroup).Where("campaign_id=?", campaign.ID).Select()

	if err != nil {
		return operations.NewGetFoldersIDOK().WithPayload(&models.ShowFolderResp{
			// ErrorCode: "internal",
			// ErrorMess: err.Error(),
		})
	}

	return operations.NewGetFoldersIDOK().WithPayload(&models.ShowFolderResp{Data: &models.ShowFolderRespData{
		Type:          Folder.Type,
		Name:          Folder.Name,
		RulesMask:     Folder.RulesMask,
		ID:            int64(Folder.ID),
		TechAdGroupID: int64(techAdGroup.ID),
		CampaignID:    int64(campaign.ID),
		CreatedAt:     Folder.CreatedAt.Format("2006-01-02"),
	}})
}

func (h *Handler) DeleteFolder(params operations.DeleteFoldersIDParams) middleware.Responder {
	DeletedFolder := models.Folder{ID: int(params.ID)}

	_, err := h.DB.Model(&DeletedFolder).WherePK().Delete()
	// fmt.Println(DeletedFolder, r)
	if err != nil {
		return operations.NewDeleteFoldersIDOK().WithPayload(&models.ShowFolderResp{
			// ErrorCode: "iternal",
			// ErrorMess: err.Error(),
		})
	}

	return operations.NewDeleteFoldersIDOK().WithPayload(&models.ShowFolderResp{
		Data: &models.ShowFolderRespData{
			ID:        int64(DeletedFolder.ID),
			Type:      DeletedFolder.Type,
			Name:      DeletedFolder.Name,
			RulesMask: DeletedFolder.RulesMask,
			CreatedAt: DeletedFolder.CreatedAt.Format("2006-01-02"),
		},
	})
}

func (h *Handler) MatchingFolder(params operations.MatchFoldersParams) middleware.Responder {
	MatchingFolders := models.MathingFolders{}
	MatchingFolders.Name = params.Body.Name
	MatchingFolders.AudienceFolderID = params.Body.AudienceFolderID
	MatchingFolders.VideoFolderID = params.Body.VideoFolderID
	_, err := h.DB.Model(&MatchingFolders).Insert()
	if err != nil {
		return operations.NewMatchFoldersOK().WithPayload(&models.MatchFoldersResp{
			// ErrorCode: "internal",
			// ErrorMess: err.Error(),
		})
	}

	matchVideoWithAudience(params.Body.AudienceFolderID, params.Body.VideoFolderID, h.DB)

	return operations.NewMatchFoldersOK().WithPayload(&models.MatchFoldersResp{
		Data: &models.MatchFoldersRespData{
			ID:                 MatchingFolders.ID,
			Name:               MatchingFolders.Name,
			VideoFolderID:      MatchingFolders.VideoFolderID,
			AudienceFolderID:   MatchingFolders.AudienceFolderID,
			AudienceFolderName: getFolderName(MatchingFolders.AudienceFolderID, h.DB),
			VideoFolderName:    getFolderName(MatchingFolders.VideoFolderID, h.DB),
		},
	})
}

func matchVideoWithAudience(folderID, videoFolderID int64, db *pg.DB) {
	a := []models.AdGroup{}

	err := db.Model(&a).Where("folder_id=?", folderID).Select()
	if err != nil {
		fmt.Println(err)
	}

	video := []models.Video{}

	err = db.Model(&video).Where("folder_id=?", videoFolderID).Select()
	if err != nil {
		fmt.Println(err)
	}

	for key := range a {
		ad := models.Ad{}
		err := db.Model(&ad).Where("adgroup_id=?", a[key].ID).Select()
		if err != nil {
			fmt.Println("select match err", err)
		}
		containingId := ad.AdName
		for key := range video {
			//if video[key].CreativeID == containingId {
			//	ad.VideoID = video[key].YouTubeID
			//	_, err = db.Model(&ad).WherePK().Update()
			//	if err != nil {
			//		fmt.Println(err)
			//	}
			//	continue
			//}
			if strings.Contains(video[key].CreativeID, containingId) {
				newAd := new(models.Ad)
				_, err = db.Model(&ad).Where("id=?", ad.ID).Delete()
				if err != nil {
					fmt.Println("match delete err", err)
				}
				newAd.AdGroupID = ad.AdGroupID
				newAd.AdName = video[key].CreativeID
				newAd.VideoID = video[key].YouTubeID

				_, err = db.Model(newAd).Insert()
				if err != nil {
					fmt.Println("match insert err", err)
				}
			}
		}
	}

}

func getFolderName(id int64, db *pg.DB) string {
	folder := &models.Folder{ID: int(id)}
	err := db.Select(folder)
	if err != nil {
		fmt.Println(err)
	}
	return folder.Name
}

func (h *Handler) ListMatchingFolder(params operations.ListMatchingFoldersParams) middleware.Responder {
	listMatches := []models.MathingFolders{}
	err := h.DB.Model(&listMatches).Select()
	if err != nil {

	}
	resp := make([]*models.ListMatchingFoldersRespItems0, len(listMatches))
	for key := range listMatches {
		resp[key] = &models.ListMatchingFoldersRespItems0{}
		resp[key].ID = listMatches[key].ID
		resp[key].Name = listMatches[key].Name
		resp[key].AudienceFolderID = listMatches[key].AudienceFolderID
		resp[key].VideoFolderID = listMatches[key].VideoFolderID
		resp[key].VideoFolderName = getFolderName(resp[key].VideoFolderID, h.DB)
		resp[key].AudienceFolderName = getFolderName(resp[key].AudienceFolderID, h.DB)
	}
	return operations.NewListMatchingFoldersOK().WithPayload(resp)
}
