package service 

type videoservice interface {

    saveVideos(video entity.Video) (entity.Video, error)
	getVideos() ([]entity.Video, error)

}