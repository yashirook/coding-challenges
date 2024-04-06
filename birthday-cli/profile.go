package main

import (
	"math"
	"time"
)

type Profile struct {
	Birthday time.Time
}

func NewProfile(birthday string) (p Profile, err error) {
	t, err := time.Parse("2006-01-02", birthday)
	if err != nil {
		return p, err
	}

	p = Profile{
		Birthday: t,
	}

	return p, nil
}

func (p Profile) Age() int {
	hourFromBirth := time.Since(p.Birthday).Hours()
	yearFromBirth := hourFromBirth / 24 / 365
	age := math.Trunc(yearFromBirth)

	return int(age)

}
