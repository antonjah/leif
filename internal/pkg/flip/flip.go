package flip

import "github.com/antonjah/leif/internal/pkg/utils"

// Get returns an answer for heads or tails
func Get() string {
	return utils.GetRandom([]string{"Heads.", "Tails."})
}
