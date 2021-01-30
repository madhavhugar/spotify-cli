package spotify

import (
	"context"
	"log"
	"os"

	"github.com/zmb3/spotify"
	"golang.org/x/oauth2/clientcredentials"
)

// new returns an authenticated spotify client
func new() (spotify.Client, error) {
	config := &clientcredentials.Config{
		ClientID:     os.Getenv("SPOTIFY_CLIENT_ID"),
		ClientSecret: os.Getenv("SPOTIFY_CLIENT_SECRET"),
		TokenURL:     spotify.TokenURL,
	}
	token, err := config.Token(context.Background())
	if err != nil {
		return spotify.Client{}, err
	}

	client := spotify.Authenticator{}.NewClient(token)

	return client, nil
}

func Albums(name string) []spotify.SimpleAlbum {
	client, err := new()
	if err != nil {
		log.Fatalf("error instantiating spotify client")
	}
	results, err := client.Search(name, spotify.SearchTypeAlbum)
	if err != nil {
		log.Fatalf("could not search album name %s: %v", name, err)
	}
	return results.Albums.Albums
}

func Artists(name string) []spotify.FullArtist {
	client, err := new()
	if err != nil {
		log.Fatalf("error instantiating spotify client")
	}
	results, err := client.Search(name, spotify.SearchTypeArtist)
	if err != nil {
		log.Fatalf("could not search artist name %s: %v", name, err)
	}
	return results.Artists.Artists
}

func Tracks(name string) []spotify.FullTrack {
	client, err := new()
	if err != nil {
		log.Fatalf("error instantiating spotify client")
	}
	results, err := client.Search(name, spotify.SearchTypeTrack)
	if err != nil {
		log.Fatalf("could not search track name %s: %v", name, err)
	}
	return results.Tracks.Tracks
}

func Playlists(name string) []spotify.SimplePlaylist {
	client, err := new()
	if err != nil {
		log.Fatalf("error instantiating spotify client")
	}
	results, err := client.Search(name, spotify.SearchTypePlaylist)
	if err != nil {
		log.Fatalf("could not search playlist name %s: %v", name, err)
	}
	return results.Playlists.Playlists
}

func TrackAudioFeatures(ID string) []*spotify.AudioFeatures {
	client, err := new()
	if err != nil {
		log.Fatalf("error instantiating spotify client")
	}
	features, err := client.GetAudioFeatures(spotify.ID(ID))
	if err != nil {
		log.Fatalf("could not get audio features ID - %s: %v", ID, err)
	}
	return features
}

func All(name string) ([]spotify.FullTrack, []spotify.SimpleAlbum, []spotify.FullArtist, []spotify.SimplePlaylist) {
	client, err := new()
	if err != nil {
		log.Fatalf("error instantiating spotify client")
	}
	results, err := client.Search(
		name,
		spotify.SearchTypePlaylist|spotify.SearchTypeArtist|spotify.SearchTypeTrack|spotify.SearchTypeAlbum,
	)
	if err != nil {
		log.Fatalf("could not search playlist name %s: %v", name, err)
	}
	return results.Tracks.Tracks, results.Albums.Albums, results.Artists.Artists, results.Playlists.Playlists
}
