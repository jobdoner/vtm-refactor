package handlers

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/go-openapi/runtime/middleware"
	"github.com/jobdoner/vtm-refactor/models"
	"github.com/jobdoner/vtm-refactor/restapi/operations"
	"github.com/jobdoner/vtm-refactor/restapi/youtube"
)

func (h *Handler) ListVideo(params operations.ListVideosParams) middleware.Responder {
	newVideos := []models.Video{}

	err := h.DB.Model(&newVideos).Where("folder_id=?", params.FolderID).Select()
	if err != nil {
		fmt.Println(err)
	}
	resp := make([]*models.ListVideosRespItems0, len(newVideos))
	for key := range newVideos {
		resp[key] = &models.ListVideosRespItems0{}
		resp[key].FolderID = newVideos[key].FolderID
		resp[key].CreativeID = newVideos[key].CreativeID
		resp[key].YoutubeID = newVideos[key].YouTubeID
		resp[key].Description = newVideos[key].Description
		resp[key].Title = newVideos[key].Title
	}
	return operations.NewListVideosOK().WithPayload(resp)
}

func (h *Handler) UploadToYoutube(params operations.UploadVideoParams) middleware.Responder {
	vids := []models.Video{}
	err := h.DB.Model(&vids).Where("folder_id=?", params.FolderID).Select()
	if err != nil {
		fmt.Println(err)
	}
	limit, err := h.RedisConn.HGetAll(context.Background(), "videos:limit").Result()

	if err != nil {
		fmt.Println(err)
	}

	i := 0

	for ; i < len(vids); i++ {
		for _, value := range limit {
			fmt.Println(value, vids[i], i)

			err := youtube.UploadToYT(&vids[i], value)
			if err != nil {
				fmt.Println(err, vids)
				break
			}
			_, err = h.DB.Model(&vids[i]).WherePK().Update()
			if err != nil {
				fmt.Println(err, vids)
				return middleware.NotImplemented(strconv.Itoa(http.StatusInternalServerError))
			}
			i++
			time.Sleep(time.Minute * 2)
		}

	}
	return nil
}
