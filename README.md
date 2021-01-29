## Spotify CLI

Enable track, album, artist and playlist discovery on your terminal

### Usage

You'll need to sign up for spotify API developer credentials and expose environment variables `SPOTIFY_CLIENT_ID` and `SPOTIFY_CLIENT_SECRET`. Below is a sample usage:

```
go run main.go -track "under pressure"
```

```

  -album    	album name
  -artist    	artist name
  -playlist    	playlist name
  -track    	track name
  -all    	    discover track, artist, album or playlist name

```

