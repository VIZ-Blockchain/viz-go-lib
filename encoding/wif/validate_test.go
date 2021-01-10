package wif

import "testing"

type testWifData struct {
	WIF       string
	PublicKey string
	Valid     bool
}

var wifs = []testWifData{
	{
		WIF:       "5JjAbnrwMn2LB8vkbVov8v9wLpDfA8UGXMfkMcH8z4yKwD1opE1",
		PublicKey: "VIZ5yWvkQU2KMtGkwztxwmjytw85ViZgfjfQY8zKokEv1tUd3ixvy",
		Valid:     true,
	},
	{
		WIF:       "5J4iGcTtH9EhhKWYaj7wnMJGUXKpo96uhY38B25Fq3B9VdSBzcE",
		PublicKey: "VIZ6PA9pdW6XxRLRgXhgiUCVDCgae8CFQb247hiZ5nrSYbLT5efee",
		Valid:     true,
	},
	{
		WIF:       "5Hq635ErPaxyd58Vc9A445BdcT8fVBETQREE9X4f799dgsWr6BJ",
		PublicKey: "VIZ666Wj5QV7e3fRcQry2HeFAqLKxgfWc95fj1bCD6txTFTnDypbo",
		Valid:     true,
	},
	{
		WIF:       "5JPtM8smAiTqPoLM9fQcA5hzHFM9QAmbNmTPpqgYyf8apUEzYQg",
		PublicKey: "VIZ8SMg2tZiNdLgwkAGJygghCdoq2NYN7HKKJx7WcHc7tK3BVF3o3",
		Valid:     true,
	},
	{
		WIF:       "5KDXzqjMtTvG47PweJE6a8xdfBJj5nxosSVYmjZWCVY15bvRBh1",
		PublicKey: "VIZ74wwCbN2X2sei6WnzX2azL8HLxxUDKzUCAZWCuQQXYevVQhmLj",
		Valid:     false,
	},
	{
		WIF:       "5KHDtj69h69JL4K6bZaex87hpQiVcGLaLa1VgozWFLWWv3vPM7Y",
		PublicKey: "VIZ6abfkaSNRFnLqwJEqHb7hCmLSP9iz4dJRrXGpKKh42bbgrb6aY",
		Valid:     false,
	},
}

func TestWifIsValid(t *testing.T) {
	for _, d := range wifs {
		b, err := WifIsValid(d.WIF, d.PublicKey)
		if err != nil {
			t.Error(err)
		}
		if b != d.Valid {
			t.Errorf("expected %v for wif %v", d.Valid, d.WIF)
		}
	}
}
