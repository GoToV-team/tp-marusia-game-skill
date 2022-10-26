package lemonadescript

import (
	"github.com/evrone/go-clean-template/pkg/convertor/words2num"
	"strconv"
)

type NumberMatchers struct{}

func (nm NumberMatchers) Match(message string) (bool, string) {
	res, _ := words2num.Convert(message)
	return res != 0, strconv.FormatInt(res, 10)
}

var PositiveNumber = NumberMatchers{}
