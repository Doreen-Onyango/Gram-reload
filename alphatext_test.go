package reload

import (
	"reflect"
	"testing"
)

func TestCapital(t *testing.T) {
	input := []string{"its", "raining", "today", "and", "i", "(cap)", "am", "sick"}

	expected := []string{"its", "raining", "today", "and", "I", "am", "sick"}

	result := Capital(input)
	if !reflect.DeepEqual(expected, result) {
		t.Errorf("Expected %v Got %v", expected, result)
	}
}

func TestUpperCase(t *testing.T) {
	input := []string{"its", "raining", "today", "and", "i", "(up)", "am", "sick"}

	expected := []string{"its", "raining", "today", "and", "I", "am", "sick"}

	result := UpperCase(input)
	if !reflect.DeepEqual(expected, result) {
		t.Errorf("Expected %v Got %v", expected, result)
	}
}

func TestLowerCase(t *testing.T) {
	input := []string{"its", "raining", "today", "and", "I", "(low)", "am", "sick"}

	expected := []string{"its", "raining", "today", "and", "i", "am", "sick"}

	result := LowerCase(input)
	if !reflect.DeepEqual(expected, result) {
		t.Errorf("Expected %v Got %v", expected, result)
	}
}

func TestHexBinDec(t *testing.T) {
	input := []string{"its", "raining", "1E", "(hex)", "today", "and", "10", "(bin)", "am", "sick"}

	expected := []string{"its", "raining", "30", "today", "and", "2", "am", "sick"}

	result := HexBinDec(input)
	if !reflect.DeepEqual(expected, result) {
		t.Errorf("Expected %v Got %v", expected, result)
	}
}

func TestGrammerVow(t *testing.T) {
	input := []string{"its", "raining", "today", "and", "i", "am", "a", "honest", "sickler"}

	expected := []string{"its", "raining", "today", "and", "i", "am", "an", "honest", "sickler"}

	result := GrammerVow(input)
	if !reflect.DeepEqual(expected, result) {
		t.Errorf("Expected %v Got %v", expected, result)
	}
}

func TestPunctuation(t *testing.T) {
	input := []string{"its", "raining", "today", "and", "i", "...", "am", "sick"}

	expected := []string{"its", "raining", "today", "and", "i...", "am", "sick"}

	result := Punctuation(input)
	if !reflect.DeepEqual(expected, result) {
		t.Errorf("Expected %v Got %v", expected, result)
	}
}

func TestPunctute(t *testing.T) {
	input := []string{"its", "raining", "today", "and", "i", "know", ",am", "sick", "!"}

	expected := []string{"its", "raining", "today", "and", "i", "know,", "am", "sick!"}

	result := Punctuate(input)
	if !reflect.DeepEqual(expected, result) {
		t.Errorf("Expected %v Got %v", expected, result)
	}
}

func TestQuotes(t *testing.T) {
	input := []string{"its", "'", "raining", "today", "and", "i", "'", "am", "sick"}

	expected := []string{"its", "'raining", "today", "and", "i'", "am", "sick"}

	result := Quotes(input)
	if !reflect.DeepEqual(expected, result) {
		t.Errorf("Expected %v Got %v", expected, result)
	}
}
