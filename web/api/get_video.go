package api

import (
	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
	"github.com/valyala/fasthttp"
	"google.golang.org/api/youtube/v3"
)

type VideoResponse struct {
	// VideoListResponseはyoutube-gen.go内に定義されている
	VideoList *youtube.VideoListResponse `json:"video_list"`
}

func GetVideo() echo.HandlerFunc {
	return func(c echo.Context) error {
		yts := c.Get("yts").(*youtube.Service)

		videoId := c.Param("id")
		call := yts.Videos.List("id,snippet").Id(videoId)

		res, err := call.Do()

		if err != nil {
			logrus.Fatalf("Error calling Youtube API: %v", err)
		}

		v := VideoResponse{
			VideoList: res,
		}

		return c.JSON(fasthttp.StatusOK, v)
	}
}