package impl

import (
	"github.com/bgildson/ifood_backend_challenge/base"
	"github.com/bgildson/ifood_backend_challenge/utils"
)

type playlistsService struct {
	TemperatureRepository base.TemperatureRepository
	PlaylistsRepository   base.PlaylistsRepository
}

// NewPlaylistsService : factory used to create a PlaylistsService
func NewPlaylistsService(temperatureRepository base.TemperatureRepository,
	playlistsRepository base.PlaylistsRepository) base.PlaylistsService {
	return playlistsService{
		TemperatureRepository: temperatureRepository,
		PlaylistsRepository:   playlistsRepository,
	}
}

func (s playlistsService) GetPlaylist(city string, latitude float64, longitude float64) (base.Playlist, error) {
	temperature, err := s.TemperatureRepository.GetTemperature(city, latitude, longitude)
	if err != nil {
		return nil, err
	}

	genre := utils.ParseTemperatureToGenre(temperature)

	playlist, err := s.PlaylistsRepository.GetByGenre(genre)
	if err != nil {
		return nil, err
	}

	return playlist, nil
}
