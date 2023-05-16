package question5

import "strings"

/**
REASONING:
- While values can be repeated in different codes
- The codes are unique per character if you consider them as prefixes with the exception of y z \s
- By this assumption, it's best to start checking from the longest codes to the shortest such that
	- 11111111110001 as space is not confused with 1111111111 as y
	- This assumes that space and z supersedes y
	** revision **
	- As space is the only symbol, it shouldn't repeat twice
	- In a case where space is repeated, process it as the next character instead
- Therefore going through the encoded string from the begging
	- Check if the prefix matches the longest to the shortest code
	- If a prefix is found match it to the correct character
	- Repeat the process minus the prefix till the string is empty
*/

var characters = [27]string{" ", "z", "y", "x", "w", "v", "u", "t", "s", "r", "q", "p", "o", "n", "m", "l", "k", "j", "i", "h", "g", "f", "e", "d", "c", "b", "a"}
var codes = [27]string{"11111111110001", "11111111110000", "1111111111", "1111111110", "1111111101", "1111111100", "1111111011", "1111111010", "1111111001", "1111111000", "1111110111", "1111110110", "1111110101", "1111110100", "1111110011", "1111110010", "1111110001", "1111110000", "111110", "111101", "111100", "1110", "1101", "1100", "10", "01", "00"}

func get_possible_value(encoded, prev string) (string, string) {
	for i, code := range codes {
		// Find the prefix and remove it, get the rest of the string for further decoding
		rem, found := strings.CutPrefix(encoded, code)
		if found {
			out := characters[i]
			// Ignore continuous spaces and instead process as the next character
			if out == " " && prev == out {
				continue
			}
			return out, rem
		}
	}
	panic("Unable to find value to decode")
}

func Decode(encoded_value string) string {
	output := []string{}
	var decoded_char string = ""
	for {
		if len(encoded_value) == 0 {
			break
		}
		decoded_char, encoded_value = get_possible_value(encoded_value, decoded_char)
		output = append(output, decoded_char)
	}
	return strings.Join(output, "")
}
