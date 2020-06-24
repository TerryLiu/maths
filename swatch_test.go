package maths

import "testing"

func TestMean(t *testing.T) {
	s := NewSwatch()
	s.Push(1)
	s.Push(1)
	s.Push(1)
	s.Push(0)
	s.Push(0)
	s.Push(0)

	if actual, expected := s.Ma(), 0.5; expected != actual {
		t.Errorf("Expected %f, got %f", expected, actual)
	}
}

func TestStandardDeviation(t *testing.T) {
	s := NewSwatch()

	if actual, expected := s.Std(), 0.0; expected != actual {
		t.Errorf("Expected %f, got %f", expected, actual)
	}

	s.Push(1)
	s.Push(1)
	s.Push(1)

	if actual, expected := s.Std(), 0.0; expected != actual {
		t.Errorf("Expected %f, got %f", expected, actual)
	}
}

func TestNumDataValues(t *testing.T) {
	s := NewSwatch()
	s.Push(1)

	if actual, expected := s.NumOfData(), uint(1); expected != actual {
		t.Errorf("Expected %f, got %f", expected, actual)
	}
}
