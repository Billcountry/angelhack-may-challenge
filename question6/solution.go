package question6

import (
	"fmt"
	"regexp"
	"strings"
)

/*
A secure communication network has been compromised and the cyber security team must
restore its security. The network, series, is represented as a series of nodes identified using
lowercase English letters. The nodes must be disconnected in order to remove the threat. In
a single operation, any number of adjacent nodes identified by the same character can be
disconnected. Find the minimum number of operations required to disconnect all the nodes
and secure the network.

** Reasoning:
- It doesn't matter how long a connection is as long as it's the same character
	- i.e. 'aaaaacaaaa' will have the same result as 'aca'
- Items with the least appearances need to be removed first. In the case above c is removed before
- Repeat the process while counting the removals
*/

func simplifyNodes(nodes string) string {
	for r := 'a'; r <= 'z'; r++ {
		nodes = regexp.MustCompile(fmt.Sprintf("%s+", string(r))).ReplaceAllString(nodes, string(r))
	}
	return nodes
}

func processNodes(nodes string) (int, string) {
	simplified := simplifyNodes(nodes)
	lowest_count := 0
	character := ""
	for r := 'a'; r <= 'z'; r++ {
		count := strings.Count(simplified, string(r))
		if count > 0 && (lowest_count == 0 || count < lowest_count) {
			lowest_count = count
			character = string(r)
		}
	}
	return lowest_count, strings.ReplaceAll(simplified, character, "")
}

func DisconnectNodes(nodes string) int {
	totalSteps := 0
	for {
		var steps int
		steps, nodes = processNodes(nodes)

		totalSteps += steps
		if len(nodes) == 0 {
			break
		}
	}
	return totalSteps
}
