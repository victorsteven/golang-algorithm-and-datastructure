package encrypt_this

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEncryptThis(t *testing.T) {

	assert.Equal(t, EncryptThis("Hello"), "72olle")

	assert.Equal(t, EncryptThis("The more he saw the less he spoke"), "84eh 109ero 104e 115wa 116eh 108sse 104e 115eokp")

	assert.Equal(t, EncryptThis("Why can we not all be like that wise old bird"), "87yh 99na 119e 110to 97ll 98e 108eki 116tah 119esi 111dl 98dri")

	assert.Equal(t, EncryptThis("The more he saw the less he spoke"), "84eh 109ero 104e 115wa 116eh 108sse 104e 115eokp")

	assert.Equal(t, EncryptThis("Thank you Piotr for all your help"), "84kanh 121uo 80roti 102ro 97ll 121ruo 104ple")

	assert.Equal(t, EncryptThis("The less he spoke the more he heard"), "84eh 108sse 104e 115eokp 116eh 109ero 104e 104dare")

}
