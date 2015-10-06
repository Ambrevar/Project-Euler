/* Heuristic:
- Create sizeof(key) arrays of frequencies.
- For each array:
  - Find the character with max frequency.
  - XOR it with ' ' (most frequent character) to get the key.

We assume the input is a regular text using spaces as a word separator.
Otherwise we could have used the most frequent letter in the input language.

Source:
https://en.wikipedia.org/wiki/Letter_frequency#Relative_frequencies_of_letters_in_the_English_language
*/

package main

import (
	"fmt"
	"io/ioutil"
)

const (
	key_size       = 3
	most_freq_char = ' '
)

func main() {
	// Load file in array.
	input, err := ioutil.ReadFile("p059_cipher.txt")
	if err != nil {
		panic(err)
	}

	buf := make([]byte, len(input))
	count := 0
	for _, i := range input {
		if i >= '0' && i <= '9' {
			buf[count] = buf[count]*10 + (i - '0')
		} else {
			// We hit a letter separator.
			count++
		}
	}

	key := [key_size]byte{}

	// Find key.
	freq := [key_size][256]byte{}
	for k := 0; k < key_size; k++ {
		for i := k; i < count; i += key_size {
			freq[k][buf[i]]++
		}

		// Heuristic
		max := 0
		for i := 0; i < 256; i++ {
			if freq[k][i] > freq[k][max] {
				max = i
			}
		}
		key[k] = byte(max ^ most_freq_char)
	}

	result := 0
	round_limit := count - count%key_size
	for i := 0; i < round_limit; i += key_size {
		for j := 0; j < key_size; j++ {
			result += int(buf[i+j] ^ key[j])
		}
	}
	// Last part that does not fit the size of the key.
	for i := round_limit; i < count; i++ {
		result += int(buf[i] ^ key[i-count+count%key_size])
	}

	fmt.Println(result)
}
