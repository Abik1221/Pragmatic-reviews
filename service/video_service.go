package service 




type VideoService interface {
    saveVideos(video entity.Video) (entity.Video, error)
	getVideos() ([]entity.Video, error)

}

type videoService struct {
      videos []entity.Video
}

func NewVideoService() 