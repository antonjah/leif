package flip

import "github.com/antonjah/gleif/internal/utils"

// Get returns an answer for heads or tails
func Get() string {
	return utils.GetRandom([]string{"Heads.", "Tails."})
}
