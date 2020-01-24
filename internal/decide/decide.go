package decide

import (
	"github.com/antonjah/gleif/internal/utils"
)

// Get returns a random answer to a scenario
func Get() string {
	return utils.GetRandom([]string{"Yes.", "No.", "Definitely.", "Definitely not."})
}
