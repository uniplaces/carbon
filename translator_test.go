package carbon

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestLoadResource(t *testing.T) {
	tr := NewTranslator()
	err := tr.loadResource("pt_BR")
	assert.Nil(t, err)
}

func TestTrans(t *testing.T) {
	c, _ := Create(2009, time.November, 10, 23, 0, 0, 0, "UTC")
	c.SetLocale("pt")
	assert.Equal(t, "pt", c.GetLocale())
}
