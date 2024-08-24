package main

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
	"github.com/vipulvpatil/fplpodcast/internal/data"
	"github.com/vipulvpatil/fplpodcast/internal/downloader"
	"github.com/vipulvpatil/fplpodcast/internal/files"
	"github.com/vipulvpatil/fplpodcast/internal/model"
	"github.com/vipulvpatil/fplpodcast/internal/transcriber"
	"github.com/vipulvpatil/fplpodcast/services/openai"
)

func main() {
	viper.AutomaticEnv()
	episodes := data.LoadEpisodes()

	for _, episode := range episodes {
		processEpisode(episode)
	}
}

func processEpisode(episode model.Episode) {
	fmt.Println("main processing episode: ", episode.Title)
	err := downloader.DownloadPodcastEpisode(episode)
	if err != nil {
		fmt.Println(err)
		return
	}

	file := transcriber.GetTranscripionFilepath(episode.Filename())
	transcriptionText, err := os.ReadFile(file)
	if err != nil {
		fmt.Println(err)
		return
	}

	analysisFilepath := files.PathForAnalysisFile(episode.Filename())
	_, err = os.Stat(analysisFilepath)
	if err == nil {
		fmt.Println("Analysis file exists: ", analysisFilepath)
		return
	}

	analysis, err := openai.GetAnalysis(string(transcriptionText))
	if err != nil {
		fmt.Println(err)
		return
	}
	out, err := os.Create(analysisFilepath)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer out.Close()
	_, err = out.WriteString(analysis)
	if err != nil {
		fmt.Println(err)
		return
	}
}
