package main

import "errors"

func min(vals ...int) int {
	n := vals[0]
	for _, val := range vals {
		if val < n {
			n = val
		}
	}
	return n
}

func max(vals ...int) int {
	n := vals[0]
	for _, val := range vals {
		if val > n {
			n = val
		}
	}
	return n
}

func minWithValidation(vals ...int) (n int, err error) {
	if len(vals) < 1 {
		err = errors.New("At least 1 arg is needed.")
		return
	}
	n = vals[0]
	for _, val := range vals {
		if val < n {
			n = val
		}
	}
	return n, nil
}

func maxWithValidation(vals ...int) (n int, err error) {
	if len(vals) < 1 {
		err = errors.New("At least 1 arg is needed.")
		return
	}
	n = vals[0]
	for _, val := range vals {
		if val > n {
			n = val
		}
	}
	return n, nil
}
