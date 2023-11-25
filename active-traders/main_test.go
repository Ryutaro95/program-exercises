package main

import (
	"reflect"
	"testing"
)

func TestMostActive(t *testing.T) {
	cases := []struct {
		desc      string
		customers []string
		expected  []string
	}{
		{
			desc:      "",
			customers: []string{"Omega", "Alpha", "Omega", "Alpha", "Omega", "Alpha", "Omega", "Alpha", "Omega", "Alpha", "Omega", "Alpha", "Omega", "Alpha", "Omega", "Alpha", "Omega", "Alpha", "Omega", "Beta"},
			expected:  []string{"Alpha", "Beta", "Omega"},
		},
		{
			desc:      "",
			customers: []string{"Omega", "Alpha", "Omega", "Alpha", "Omega", "Alpha", "Omega", "Alpha", "Omega", "Alpha", "Omega", "Alpha", "Omega", "Alpha", "Omega", "Alpha", "Omega", "Alpha", "Omega", "Beta"},
			expected:  []string{"hog"},
		},
	}

	for _, c := range cases {
		t.Run(c.desc, func(t *testing.T) {
			actual := mostActive(c.customers)
			if !reflect.DeepEqual(actual, c.expected) {
				t.Errorf("mostActive(%q) == %q, expected %q", c.customers, actual, c.expected)
			}
		})
	}
}
