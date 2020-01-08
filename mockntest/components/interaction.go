package components

func MakeSpeak(s Speaker) int {
	p := s.Speak("Hello1")
	return len(p)
}
