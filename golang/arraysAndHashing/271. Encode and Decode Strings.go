package arraysAndHashing

import "strings"

type Codec struct {
}

// Encodes a list of strings to a single string.
func (c *Codec) Encode(strs []string) string {
	joined := strings.Join(strs, "")
	return joined
}

// Decodes a single string to a list of strings.
func (c *Codec) Decode(strs string) []string {
	split := strings.Split(strs, "")
	return split
}

// Your Codec object will be instantiated and called as such:
// var codec Codec
// codec.Decode(codec.Encode(strs));
