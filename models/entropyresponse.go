package models

type EntropyResponse struct {
	EntropyDetail []float64      `json:"entropyDetail"`
	Summary       map[string]int `json:"summary"`
}
