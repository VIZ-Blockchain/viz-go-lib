package operations

import "sort"

// sortedStrings returns a sorted copy of v. fc::flat_set serializes its elements in
// ascending order, so any flat_set<string> field must be sorted before encoding.
func sortedStrings(v []string) []string {
	if len(v) == 0 {
		return v
	}
	out := make([]string, len(v))
	copy(out, v)
	sort.Strings(out)
	return out
}
