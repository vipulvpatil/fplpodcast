package files

import "fmt"

func PathForAudioFile(filename string) string {
	return fmt.Sprintf("internal/files/audio/%s.mp3", filename)
}

func PathForTranscriptionFile(filename string) string {
	return fmt.Sprintf("internal/files/transcripts/%s.txt", filename)
}
