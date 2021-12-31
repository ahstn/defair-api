package actions

import (
	"github.com/ahstn/defair/domain"
	"github.com/thoas/go-funk"
)

// networkFilter determines the networks to process from the supplied domain.DataFilter.
// If the filter isn't "all", intersect is used to ensure the Network Names provided match valid IDs.
// The initial value for "networks" is essentially the currently supported networks.
func networkFilter(filter domain.DataFilter) []string {
	networks := []string{"avalanche", "harmony", "aurora"}
	if len(filter.Networks) >= 1 && !funk.Contains(filter.Networks, "all") {
		networks = funk.IntersectString(networks, filter.Networks)
	}
	return networks
}
