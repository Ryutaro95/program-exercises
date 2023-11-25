package main

import (
	"fmt"
	"sort"
)

func mostActive(customers []string) []string {
	n := len(customers)
	m := make(map[string]int)

	for _, customer := range customers {
		if _, ok := m[customer]; !ok {
			m[customer] = 1
		} else {
			m[customer]++
		}
	}

	activeCustomers := []string{}

	for key, val := range m {
		if float64(val)/float64(n) >= 0.05 {
			activeCustomers = append(activeCustomers, key)
		}
	}

	fmt.Println(activeCustomers)
	sort.Slice(activeCustomers, func(i, j int) bool { return activeCustomers[i] < activeCustomers[j] })
	fmt.Println(activeCustomers)

	return []string{}
}

func main() {
}
