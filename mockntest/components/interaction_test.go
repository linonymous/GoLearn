package components

import (
	"testing"

	gomock "github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestMakeSpeak(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockSpeaker := NewMockSpeaker(mockCtrl)

	mockSpeaker.EXPECT().Speak("Hello1").Return("linonymous")

	s := MakeSpeak(mockSpeaker)
	assert.Equal(t, s, 10)
}
