## Spotify CLI

Enable track, album, artist and playlist discovery on your terminal

### Usage

You'll need to sign up for spotify API developer credentials and expose environment variables `SPOTIFY_CLIENT_ID` and `SPOTIFY_CLIENT_SECRET`. Below is a sample usage:

```
go build
./spotify-cli -track "under pressure"
```

```

Usage of ./spotify-cli:
  -album string
    	album name
  -all string
    	discover track, artist, album or playlist name
  -analysis
    	print track analysis, track ID must be provided
  -artist string
    	artist name
  -features
    	print track features, track ID must be provided
  -id string
    	spotify track ID
  -playlist string
    	playlist name
  -track string
    	track name

```

