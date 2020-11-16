package handlers

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/jobdoner/vtm-refactor/models"
	"github.com/jobdoner/vtm-refactor/restapi/operations"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-pg/pg"
	"github.com/gocarina/gocsv"
)

func createFile(f multipart.File, path, name string) {

	buf := bytes.NewBuffer(nil)

	_, err := io.Copy(buf, f)

	err = ioutil.WriteFile(path+name, buf.Bytes(), os.ModePerm)
	if err != nil {
		fmt.Println(err)
	}
}

func (h *Handler) UploadFiles(video operations.PostVideoParams) middleware.Responder {
	err := os.Mkdir("./folders", os.ModePerm)
	err = os.Mkdir("./folders/"+strconv.Itoa(int(*video.FolderID)), os.ModePerm)
	//err = os.Mkdir("./folders/"+strconv.Itoa(int(*video.FolderID)), os.ModePerm)
	if err != nil {
		fmt.Println(err)
	}
	resp := make([]*models.ListVideosRespItems0, len(video.HTTPRequest.MultipartForm.File["uploadFile[]"]))
	for key := range video.HTTPRequest.MultipartForm.File["uploadFile[]"] {
		resp[key] = &models.ListVideosRespItems0{}
		var f multipart.File
		f, err = video.HTTPRequest.MultipartForm.File["uploadFile[]"][key].Open()
		if err != nil {
			fmt.Println(err)
		}

		createFile(f, "./folders/"+strconv.Itoa(int(*video.FolderID))+"/", video.HTTPRequest.MultipartForm.File["uploadFile[]"][key].Filename)

		//TODO: change to array creative ID, filepath make static without filename

		v := models.Video{
			FolderID:    *video.FolderID,
			CreativeID:  strings.ReplaceAll(video.HTTPRequest.MultipartForm.File["uploadFile[]"][key].Filename, ".mp4", ""),
			FilePath:    "./folders/" + strconv.Itoa(int(*video.FolderID)) + "/" + video.HTTPRequest.MultipartForm.File["uploadFile[]"][key].Filename,
			Title:       *video.Title + "_" + strconv.Itoa(key),
			Description: *video.Description + "_" + strconv.Itoa(key),
		}
		_, err := h.DB.Model(&v).Insert()
		{
			resp[key].Description = v.Description
			resp[key].CreativeID = v.CreativeID
			resp[key].Title = v.Title
			resp[key].YoutubeID = v.YouTubeID
			resp[key].FolderID = v.FolderID
		}

		if err != nil {
			fmt.Println(err)
		}
	}
	//TODO: upload to YT

	h.UploadToYoutube(operations.UploadVideoParams{FolderID: *video.FolderID})

	return operations.NewPostVideoOK().WithPayload(models.ListVideosResp(resp))
}

func (h *Handler) GetTargetsFile(params operations.GetTargetsParams) middleware.Responder {

	createTargetExport(*params.FolderID, *params.CampaignID, h.DB)

	f1, err := os.Open("./csv/" + strconv.Itoa(int(*params.FolderID)) + "/" + strconv.Itoa(int(*params.CampaignID)) + "_targets.csv")

	if err != nil {
		fmt.Println(err)
	}

	return NewOctetStream(f1)
}

func (h *Handler) UpdateTargetWithCsv(params operations.PostImportTargetsParams) middleware.Responder {
	//for key:= range params.HTTPRequest.MultipartForm.File["uploadFile"]{
	//	f,err :=params.HTTPRequest.MultipartForm.File["uploadFile"][key].Open()
	//	if err != nil {
	//		fmt.Println(err)
	//	}

	imports := make([]*models.ExportAdgroup, 0, 0)
	//var f multipart.File
	//f:=
	//if err != nil {
	//	fmt.Println(err)
	//}
	f, err := params.HTTPRequest.MultipartForm.File["uploadFile"][0].Open()
	if err != nil {
		fmt.Println(err)
	}
	gocsv.SetCSVReader(func(in io.Reader) gocsv.CSVReader {
		r := csv.NewReader(in)
		r.Comma = ';'
		return r // Allows use tab as delimiter in CSV
	})

	err = gocsv.Unmarshal(f, &imports)
	newAdgroups := make([]models.AdGroup, len(imports))

	for key := range imports {
		fmt.Println(*imports[key])

		newAdgroups[key].ID = imports[key].ID
		newAdgroups[key].FolderID = imports[key].FolderID
		newAdgroups[key].TechAdgroupId = imports[key].TechAdgroupId
		newAdgroups[key].AdGroupID = strings.Split(imports[key].AdGroupID, ",")
		newAdgroups[key].Keyword = strings.Split(imports[key].Keyword, ",")
		newAdgroups[key].Gender = strings.Split(imports[key].Gender, ",")
		newAdgroups[key].Age = strings.Split(imports[key].Age, ",")
		newAdgroups[key].ChannelID = strings.Split(imports[key].ChannelID, ",")
		newAdgroups[key].Languages = strings.Split(imports[key].Languages, ",")
		newAdgroups[key].Audience = strings.Split(imports[key].Audience, ",")
		newAdgroups[key].CustomAudience = strings.Split(imports[key].CustomAudience, ",")
		newAdgroups[key].Devices = strings.Split(imports[key].Devices, ",")
		newAdgroups[key].Placement = strings.Split(imports[key].Placement, ",")
		newAdgroups[key].Location = strings.Split(imports[key].Location, ",")
		newAdgroups[key].Topics = strings.Split(imports[key].Topics, ",")
		newAdgroups[key].Interest = strings.Split(imports[key].Interest, ",")
		newAdgroups[key].AdSchedule = strings.Split(imports[key].AdSchedule, ",")

	}
	_, err = h.DB.Model(&newAdgroups).OnConflict("(id) DO UPDATE").Insert()

	if err != nil {
		fmt.Println(err)
	}
	//}
	return nil
}

func createTargetExport(fid, cid int64, db *pg.DB) {
	var targetsModels []models.AdGroup
	err := db.Model(&targetsModels).Where("folder_id=?", fid).Order("id ASC").Select()
	if err != nil {
		fmt.Println(err)
	}

	err = os.Mkdir("./csv/", os.ModePerm)

	err = os.Mkdir("./csv/"+strconv.Itoa(int(fid)), os.ModePerm)
	if err != nil {
		fmt.Println(err)
	}
	os.Remove("./csv/" + strconv.Itoa(int(fid)) + "/" + strconv.Itoa(int(cid)) + "_targets.csv")

	clientsFile, err := os.OpenFile("./csv/"+strconv.Itoa(int(fid))+"/"+strconv.Itoa(int(cid))+"_targets.csv", os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		panic(err)
	}
	defer clientsFile.Close()

	targets := make([]models.ExportAdgroup, len(targetsModels))
	for key := range targetsModels {
		targets[key].ID = targetsModels[key].ID
		targets[key].CustomAudience = strings.Join(targetsModels[key].CustomAudience, ",")
		targets[key].FolderID = targetsModels[key].FolderID
		targets[key].TechAdgroupId = targetsModels[key].TechAdgroupId
		targets[key].AdGroupID = strings.Join(targetsModels[key].AdGroupID, ",")
		targets[key].Keyword = strings.Join(targetsModels[key].Keyword, ",")
		targets[key].Gender = strings.Join(targetsModels[key].Gender, ",")
		targets[key].Age = strings.Join(targetsModels[key].Age, ",")
		targets[key].ChannelID = strings.Join(targetsModels[key].ChannelID, ",")
		targets[key].Languages = strings.Join(targetsModels[key].Languages, ",")
		targets[key].Audience = strings.Join(targetsModels[key].Audience, ",")
		targets[key].CustomAudience = strings.Join(targetsModels[key].CustomAudience, ",")
		targets[key].Devices = strings.Join(targetsModels[key].Devices, ",")
		targets[key].Placement = strings.Join(targetsModels[key].Placement, ",")
		targets[key].Location = strings.Join(targetsModels[key].Location, ",")
		targets[key].Topics = strings.Join(targetsModels[key].Topics, ",")
		targets[key].Interest = strings.Join(targetsModels[key].Interest, ",")
		targets[key].AdSchedule = strings.Join(targetsModels[key].AdSchedule, ",")
	}
	err = gocsv.MarshalFile(targets, clientsFile)
	if err != nil {
		fmt.Println(err)
	}
}

func NewOctetStream(f *os.File) middleware.Responder {
	return middleware.ResponderFunc(func(w http.ResponseWriter, _ runtime.Producer) {
		fn := filepath.Base(f.Name())
		w.Header().Set("Content-Type", "application/octet-stream")
		w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=%q", fn))
		io.Copy(w, f)
		f.Close()
	})
}
