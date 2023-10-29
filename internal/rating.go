package internal

import gocvss40 "github.com/pandatix/go-cvss/40"

func Rating(score float64) (string, error) {
	return gocvss40.Rating(score)
}
