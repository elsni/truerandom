package truerandom

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
)

func fetch(min, max, num int) ([]int, error) {
	result := []int{}
	if max <= min {
		return result, errors.New("max must be greater than min")
	}
	if min < -(1e9) {
		return result, errors.New("min must be -1e9 or greater")
	}
	if max > 1e9 {
		return result, errors.New("max must be 1e9 or less")
	}
	if num < 1 || num > 10000 {
		return result, errors.New("num must be between 1 and 10000")
	}
	resp, err := http.Get(fmt.Sprintf("https://www.random.org/integers/?num=%d&min=%d&max=%d&col=1&base=10&format=plain&rnd=new", num, min, max))
	if err != nil {
		return result, err
	}
	resBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return result, err
	}
	parts := strings.Split(string(resBody), "\n")
	for _, numstring := range parts {
		num, err = strconv.Atoi(numstring)
		if err == nil {
			result = append(result, num)
		}
	}
	return result, nil
}
