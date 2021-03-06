/*

https://old.reddit.com/r/dailyprogrammer/comments/ffxabb/20200309_challenge_383_easy_necklace_matching/

Imagine a necklace with lettered beads that can slide along the string. Here's
an example image.

In this example, you could take the N off NICOLE and slide it around to the
other end to make ICOLEN. Do it again to get COLENI, and so on. For the purpose
of today's challenge, we'll say that the strings "nicole", "icolen", and
"coleni" describe the same necklace.

Generally, two strings describe the same necklace if you can remove some number
of letters from the beginning of one, attach them to the end in their original
ordering, and get the other string. Reordering the letters in some other way
does not, in general, produce a string that describes the same necklace.

Write a function that returns whether two strings describe the same necklace.
Examples

same_necklace("nicole", "icolen") => true
same_necklace("nicole", "lenico") => true
same_necklace("nicole", "coneli") => false
same_necklace("aabaaaaabaab", "aabaabaabaaa") => true
same_necklace("abc", "cba") => false
same_necklace("xxyyy", "xxxyy") => false
same_necklace("xyxxz", "xxyxz") => false
same_necklace("x", "x") => true
same_necklace("x", "xx") => false
same_necklace("x", "") => false
same_necklace("", "") => true

Optional Bonus 1

If you have a string of N letters and you move each letter one at a time from
the start to the end, you'll eventually get back to the string you started with,
after N steps. Sometimes, you'll see the same string you started with before N
steps. For instance, if you start with "abcabcabc", you'll see the same string
("abcabcabc") 3 times over the course of moving a letter 9 times.

Write a function that returns the number of times you encounter the same
starting string if you move each letter in the string from the start to the
end, one at a time.

repeats("abc") => 1
repeats("abcabcabc") => 3
repeats("abcabcabcx") => 1
repeats("aaaaaa") => 6
repeats("a") => 1
repeats("") => 1

Optional Bonus 2

There is exactly one set of four words in the enable1 word list that all
describe the same necklace. Find the four words.

https://raw.githubusercontent.com/dolph/dictionary/master/enable1.txt

*/
package main

import "fmt"

func sameNecklace(a, b string) bool {

	// if the length of the two strings are different,
	// they can't be the same necklace
	if len(a) != len(b) {
		return false
	}

	// if the strings are the same, then it's the same necklace
	if a == b {
		return true
	}

	// shift a letter from front to back
	mutiA := shiftLetter(a)

	for i := 0; i < len(a); i++ {
		// compare the shifted letter to b
		if mutiA == b {
			return true
		}

		// shift another letter for the next comparison
		mutiA = shiftLetter(mutiA)
	}

	// if we've gotten here, nothing has matched so return false
	return false
}

func repeats(a string) int {
	count := 0

	// if we get a string length less than 2, return a single count
	if len(a) < 2 {
		return 1
	}

	// shift a letter and start our comparison loop
	mutiA := shiftLetter(a)
	for i := 0; i < len(a); i++ {
		if mutiA == a {
			count++
		}
		mutiA = shiftLetter(mutiA)
	}
	return count
}

func shiftLetter(a string) string {
	return a[1:] + a[:1]
}

// taken from a solution on reddit by skeeto
// to try to understand it
// https://old.reddit.com/r/dailyprogrammer/comments/ffxabb/20200309_challenge_383_easy_necklace_matching/fk1gzsp/
func canonicalize(s string) string {

	// this seems to produce a string that's easily sliced for a comparison
	// below by appending the string to itself, but leaving off the last letter
	// canonicalizing 'fish' by 'fish' + 'fis' => 'fishfis'
	c := s + s[:len(s)-1]
	//fmt.Printf("canonicalizing '%v' by '%v' + '%v' => '%v'\n", s, s, s[:len(s)-1], c)
	best := s
	//fmt.Printf("Best '%v'\n", s)

	// then we loop over the string, moving a letter to the end of the necklace
	// for each iteration. The c[i:i+len(s)] performs the rotation based on the
	// for loop iteration count, which starts at 1.
	for i := 1; i < len(s); i++ {
		//fmt.Printf("l%v checking if '%v' < best '%v'\n", i, c[i:i+len(s)], best)
		// this comparison seems to be an alphabetical comparison, effectively
		// giving a unique primary key to identify words that can produce
		// the same necklace given the rotation specified
		// From the spec: String values are comparable and ordered, lexically
		// byte-wise.
		if c[i:i+len(s)] < best {
			best = c[i : i+len(s)]
			//fmt.Printf("l%v New Best '%v'\n", i, best)
		}
	}
	return best
}

func main() {
	fmt.Println("Run the tests!")
}
