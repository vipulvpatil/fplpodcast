package downloader

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/vipulvpatil/fplpodcast/internal/files"
	"github.com/vipulvpatil/fplpodcast/internal/model"
)

func DownloadPodcastEpisode(episode model.Episode) error {
	fmt.Println("start downloading podcast: ", episode.Title)
	filepath := files.PathForAudioFile(episode.Filename())
	if _, err := os.Stat(filepath); err == nil {
		fmt.Println("podcast already exists")
		return nil
	}
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Get the data
	resp, err := http.Get(episode.AudioUrl)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Check server response
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("bad status: %s", resp.Status)
	}

	// Writer the body to file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	fmt.Println("podcast downloaded: ", episode.Title)

	return nil
}
