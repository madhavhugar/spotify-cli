package main

import (
	"flag"
	"fmt"
	spotify "spotify-cli/spotify"
)

var album = flag.String("album", "", "album name")
var artist = flag.String("artist", "", "artist name")
var track = flag.String("track", "", "track name")
var all = flag.String("all", "", "discover track, artist, album or playlist name")

var playlist = flag.String("playlist", "", "playlist name")
var playlistID = flag.String("playlistId", "", "playlist ID")
var showPlaylistTracks = flag.Bool("showTracks", false, "show playlist tracks, playlist ID must be provided")

var id = flag.String("id", "", "spotify track ID")
var audioFeatures = flag.Bool("features", false, "print track features, track ID must be provided")
var audioAnalysis = flag.Bool("analysis", false, "print track analysis, track ID must be provided")

func main() {
	flag.Parse()

	if *all != "" {
		track = all
		artist = all
		playlist = all
		album = all
	}

	if *track != "" {
		fmt.Println("Tracks:")
		for _, item := range spotify.Tracks(*track) {
			fmt.Printf("\t %s: %s - %s\n", item.ID, item.Artists[0].Name, item.Name)
		}
	}

	if *album != "" {
		fmt.Println("Albums:")
		for _, item := range spotify.Albums(*album) {
			fmt.Printf("\t %s - %s\n", item.Artists[0].Name, item.Name)
		}
	}

	if *playlist != "" {
		fmt.Println("Playlists:")
		for _, item := range spotify.Playlists(*playlist) {
			fmt.Printf("\t%s - %s\n", item.ID, item.Name)
		}
	}

	if *artist != "" {
		fmt.Println("Artists:")
		for _, item := range spotify.Artists(*artist) {
			fmt.Printf("\t %s - %d - %d\n", item.Name, item.Popularity, item.Followers.Count)
		}
	}

	if *audioFeatures && *id != "" {
		fmt.Println("Audio track features:", len(spotify.TrackAudioFeatures(*id)))
		for _, item := range spotify.TrackAudioFeatures(*id) {
			fmt.Printf("\t Dancability: %f - Energy: %f - Instrumentalness: %f - Duration: %d\n",
				item.Danceability, item.Energy, item.Instrumentalness, item.Duration)
		}
	}

	if *audioAnalysis && *id != "" {
		fmt.Println("Audio track analysis")
		analysis := spotify.TrackAudioAnalysis(*id)
		fmt.Println("Bars", analysis.Bars)
		fmt.Println("Beats", analysis.Beats)
		fmt.Println("Meta", analysis.Meta)
	}

	if *showPlaylistTracks && *playlistID != "" {
		fmt.Println("Playlist tracks")
		name, tracks := spotify.PlaylistTracks(*playlistID)
		fmt.Println(name)
		for _, item := range tracks {
			// ensure tracks with multiple artists are rendered properly
			if len(item.Track.Artists) > 1 {
				fmt.Printf("\t%s", item.Track.Artists[0].Name)
				for _, artist := range item.Track.Artists[1:] {
					fmt.Printf(", %s ", artist.Name)
				}
				fmt.Printf("- %s\n", item.Track.Name)
			} else {
				fmt.Printf("\t%s- %s\n", item.Track.Artists[0].Name, item.Track.Name)
			}
		}
	}
}
