package user

import "testing"

type expectedCase struct {
	mock           []user
	expectedResult bool
}

func TestAtLeastTwice(t *testing.T) {
	cases := []expectedCase{
		{mocksA, true},
		{mocksB, true},
		{mocksC, true},
		{mocksD, false},
	}
	for _, v := range cases {
		if atLeastTwice(v.mock) != v.expectedResult {
			t.Errorf("expected %t", v.expectedResult)
		}
	}
}

func TestExactlyTwice(t *testing.T) {
	cases := []expectedCase{
		{mocksA, false},
		{mocksB, false},
		{mocksC, true},
		{mocksD, false},
	}
	for _, v := range cases {
		if exactlyTwice(v.mock) != v.expectedResult {
			t.Errorf("expected %t", v.expectedResult)
		}
	}
}

func TestConstrainedExactlyTwice(t *testing.T) {
	cases := []expectedCase{
		{mocksB, false},
		{mocksC, true},
		{mocksD, false},
	}
	for _, v := range cases {
		if constrainedExactlyTwice(v.mock) != v.expectedResult {
			t.Errorf("expected %t", v.expectedResult)
		}
	}
}
