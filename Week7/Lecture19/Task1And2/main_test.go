package main

import (
	"bytes"
	"io"
	"testing"
)

func TestRead(t *testing.T) {
	//Arrange
	testString := "test"
	expectedString := "tset"
	rsr := NewReverseStringReader(testString)
	var result1 bytes.Buffer
	var result2 bytes.Buffer

	//Act
	result1.ReadFrom(rsr)
	result2.ReadFrom(rsr)
	n, _ := rsr.Read([]byte(testString))

	//Assert
	t.Run("Test reader returns correct reversed string", func(t *testing.T) {
		if expectedString != result1.String() {
			t.Errorf("Failed test expected %s, but got %s", expectedString, result1.String())
		}
	})

	t.Run("Test reader returns an error EOF", func(t *testing.T) {
		if io.EOF.Error() == result2.String() {
			t.Errorf("Failed test expected %s, but got %s", expectedString, result1.String())
		}
	})

	t.Run("Test reader returns correct number of bytes", func(t *testing.T) {
		if n != 0 {
			t.Errorf("Failed test expected 0, but got %d", n)
		}
	})
}
