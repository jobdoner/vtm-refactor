package youtube

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/jobdoner/vtm-refactor/models"
	youtube "google.golang.org/api/youtube/v3"
)

var (
	category = flag.String("category", "22", "Video category")
	keywords = flag.String("keywords", "test", "Comma separated list of video keywords")
	privacy  = flag.String("privacy", "unlisted", "Video privacy status")
)

func UploadToYT(video *models.Video, id string) error {
	if video.FilePath == "" {
		return fmt.Errorf("You must provide a filename of a video file to upload")
	}

	client := getClient(youtube.YoutubeUploadScope, id)

	service, err := youtube.New(client)
	if err != nil {
		return fmt.Errorf("Error creating YouTube client: %v", err)
	}

	upload := &youtube.Video{
		Snippet: &youtube.VideoSnippet{
			Title:       video.Title,
			Description: video.Description,
			CategoryId:  *category,
		},
		Status: &youtube.VideoStatus{PrivacyStatus: *privacy},
	}

	// The API returns a 400 Bad Request response if tags is an empty string.
	if strings.Trim(*keywords, "") != "" {
		upload.Snippet.Tags = strings.Split(*keywords, ",")
	}

	call := service.Videos.Insert([]string{"snippet,status"}, upload)

	file, err := os.Open(video.FilePath)
	defer file.Close()
	if err != nil {
		return fmt.Errorf("Error opening %v: %v", video.FilePath, err)
	}

	response, err := call.Media(file).Do()
	if err != nil {
		return fmt.Errorf("Error: %v", err)
	}

	fmt.Printf("Upload successful! Video ID: %v\n", response.Id)
	video.YouTubeID = response.Id
	return nil

}
