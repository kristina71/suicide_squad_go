package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGoodMorning(t *testing.T) {
	t.Parallel()
	assert.Equal(t,"Доброе утро, kristina",getTextResultByHour("kristina",5))
	assert.Equal(t,"Доброе утро, kristina",getTextResultByHour("kristina",9))
}

func TestGoodAfternoon(t *testing.T) {
	t.Parallel()
	assert.Equal(t,"Добрый день, kristina",getTextResultByHour("kristina",10))
	assert.Equal(t,"Добрый день, kristina",getTextResultByHour("kristina",16))
}

func TestGoodEvening(t *testing.T) {
	t.Parallel()
	assert.Equal(t,"Добрый вечер, kristina",getTextResultByHour("kristina",17))
	assert.Equal(t,"Добрый вечер, kristina",getTextResultByHour("kristina",23))
}

func TestGoodNight(t *testing.T) {
	t.Parallel()
	assert.Equal(t,"Доброй ночи, kristina",getTextResultByHour("kristina",0))
	assert.Equal(t,"Доброй ночи, kristina",getTextResultByHour("kristina",4))
}