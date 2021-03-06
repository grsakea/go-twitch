package twitch

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/grafov/m3u8"
	"github.com/stretchr/testify/assert"
)

func TestGetHLSAccessToken(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, r.RequestURI, "/twitch/access_token")
		assert.Equal(t, r.Header.Get("Client-ID"), "foobar")
		fmt.Fprintln(w, "{}")
	}))
	defer ts.Close()

	testSession := Session{ClientID: "foobar"}
	testSession.getHLSAccessToken("twitch", ts.URL+"/")
}

func TestGetChannelM3UPlaylist(t *testing.T) {
	ts := emptyHTTPServer(t, "/twitch.m3u8?allow_source=true&p=499379&sig=&token=")
	defer ts.Close()

	getChannelM3U8Playlist("twitch", accessToken{}, ts.URL+"/")
}

func TestParsePlaylist(t *testing.T) {
	pl := m3u8.NewMasterPlaylist()
	pl.Append("test", &m3u8.MediaPlaylist{}, m3u8.VariantParams{})
	pl.Append("test2", &m3u8.MediaPlaylist{}, m3u8.VariantParams{})

	out, _ := parsePlaylist(pl.Encode())

	if out[0].URL != "test" || out[1].URL != "test2" {
		t.Fatal("Failed parsing playlist")
	}
}
