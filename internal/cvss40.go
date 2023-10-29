package internal

import (
	"fmt"

	"github.com/gopherjs/gopherjs/js"
	gocvss40 "github.com/pandatix/go-cvss/40"
)

var (
	order = []string{
		// Base metrics (# = 11)
		"AV", "AC", "AT", "PR", "UI", "VC", "VI", "VA", "SC", "SI", "SA",
		// Threat (# = 1)
		"E",
		// Environmental (# = 14)
		"CR", "IR", "AR", "MAV", "MAC", "MAT", "MPR", "MUI", "MVC", "MVI", "MVA", "MSC", "MSI", "MSA",
		// Supplemental (# = 6)
		"S", "AU", "R", "V", "RE", "U",
	}
)

type CVSS40 struct {
	*js.Object
	vec *gocvss40.CVSS40
}

func New40() *js.Object {
	cvss40 := &CVSS40{
		vec: &gocvss40.CVSS40{}, // default value is fine
	}
	return js.MakeWrapper(cvss40)
}

func (cvss40 *CVSS40) Parse(vector string) error {
	vec, err := gocvss40.ParseVector(vector)
	if err != nil {
		return err
	}
	cvss40.vec = vec
	return nil
}

func (cvss40 *CVSS40) Vector() string {
	// cvss40.vec.Vector will be interpreted as an Object due to byte and unsafe
	// memory manipulation under the hood.
	// As there is no need of high performances on this implementation, let's
	// recode the Vector with allocs.
	vec := "CVSS:4.0"
	for i, metric := range order {
		v, _ := cvss40.vec.Get(metric)
		if i < 11 || v != "X" {
			vec += fmt.Sprintf("/%s:%s", metric, v)
		}
	}
	return vec
}

func (cvss40 *CVSS40) Get(metric string) (string, error) {
	return cvss40.vec.Get(metric)
}

func (cvss40 *CVSS40) Set(metric, value string) error {
	return cvss40.vec.Set(metric, value)
}

func (cvss40 *CVSS40) Score() float64 {
	return cvss40.vec.Score()
}

func (cvss40 *CVSS40) Nomenclature() string {
	return cvss40.vec.Nomenclature()
}
