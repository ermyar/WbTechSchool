package task8

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTime(t *testing.T) {
	assert.Equal(t, 0, getTime("0.beevik-ntp.pool.ntp.org"))
	assert.Equal(t, 0, getTime("0.ru.pool.ntp.org"))
	assert.Equal(t, 0, getTime("1.ru.pool.ntp.org"))
	assert.Equal(t, 0, getTime("ntp0.NL.net"))
	assert.Equal(t, 0, getTime("ntp2.vniiftri.ru"))
	assert.Equal(t, 0, getTime("ntp.ix.ru"))
	assert.Equal(t, 0, getTime("ntps1-1.cs.tu-berlin.de"))
}
