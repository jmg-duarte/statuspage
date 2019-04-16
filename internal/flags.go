package internal

import "strings"

type set map[string]struct{}

func ValidateFilterFlags(only, exclude string, services Services) Services {
	filteredServices := make(Services)
	if only != "" {
		for _, str := range ParseFlag(only) {
			filteredServices[str] = services[str]
		}
	} else {
		if exclude != "" {
			// Create a set containing all excluded services
			servSet := make(set)
			for _, str := range ParseFlag(exclude) {
				servSet[str] = struct{}{}
			}

			// Query all available services and only add those not in the set
			for sID, serv := range services {
				if _, ok := servSet[sID]; !ok {
					filteredServices[sID] = serv
				}
			}
		} else {
			// If both only and exclude are empty use all available services
			filteredServices = services
		}
	}
	return filteredServices
}

// ParseFlag parses multiple flag arguments into a list of services
func ParseFlag(flag string) []string {
	flags := strings.Split(flag, ",")
	res := make([]string, len(flags))
	for i, f := range flags {
		res[i] = string(f)
	}
	return res
}
