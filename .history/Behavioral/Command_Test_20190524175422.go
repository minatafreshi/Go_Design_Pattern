package ehavioral

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTalk(t *testing.T) {

	bufferOutputWriter := outputWriter
	outputWriter = new(bytes.Buffer)
	defer func() { outputWriter = bufferOutputWriter }()

	wilma := NewPerson("Wilma", NewCommand(nil, nil))
	betty := NewPerson("Betty", NewCommand(&wilma, wilma.Listen))
	barney := NewPerson("Barney", NewCommand(&betty, betty.Gossip))
	fred := NewPerson("Fred", NewCommand(&barney, barney.PassOn))
	fred.Talk()

	assert.Equal(t, "Fred is talking.\n"+
		"Barney is passing on.\n"+
		"Betty is gossiping.\n"+
		"Wilma is listening.\n", outputWriter.(*bytes.Buffer).String())
}