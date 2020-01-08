package components

// mockgen -source=speaker.go -destination=speaker_mock.go -package=components

type Speaker interface {
	Speak(string) string
}
