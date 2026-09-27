# Go practice: maps

Complete this exercise after studying maps in the Tour of Go.

Use fixed test values. Use only the standard `fmt` package. Do not use structs, slices, sorting, or packages other than `fmt`. Write the solution yourself.

Create a separate `main.go` in this folder with `package main`.

## Exercise: Count and summarize words

Write these functions:

```go
func countWords(words []string) map[string]int
func mostCommon(counts map[string]int) (word string, count int)
func removeWord(counts map[string]int, word string) bool
```

### Requirements

1. In `main`, create a fixed `[]string` containing these words:

   ```text
   go, code, go, learn, code, go, map, learn
   ```

2. In `countWords`, create and return a map from each word to the number of times it appears.
3. If the input slice is nil or empty, return an empty non-nil map.
4. In `mostCommon`, return the word with the highest count and its count.
5. If the map is nil or empty, return `""` and `0`.
6. If two words have the same highest count, return the word that appears first when scanning the map. Do not assume map iteration order is stable; your `main` only needs to demonstrate the result for the supplied counts.
7. In `removeWord`, delete the requested key and return `true` if the key existed. Return `false` when the key was absent. Do not panic for a nil map.
8. Print the original words, the counts map, the most common word and count, and the result of removing `"code"`.
9. Print the map after removing `"code"`.
10. Also call `countWords(nil)`, `mostCommon(nil)`, and `removeWord(nil, "go")` to demonstrate that they do not panic.

### Expected counts

| Word | Count |
| --- | ---: |
| `go` | 3 |
| `code` | 2 |
| `learn` | 2 |
| `map` | 1 |

Before removal, the most common word is `go` with count `3`. Removing `code` returns `true`; afterward, looking up `counts["code"]` gives the zero value `0`, and the key is absent.

### Explain

Answer these questions in comments or a separate notes file:

1. Why can you read from a nil map but not assign a key into one?
2. Why does looking up a missing key return the element type's zero value?
3. Why is a map's iteration order not suitable for choosing a deterministic tie winner?

## Review rubric

This exercise is worth 100 points:

- **Correctness: 50 points.** Counting, lookup, deletion, empty inputs, and returned values match the requirements.
- **Required concepts: 30 points.** The solution uses map literals or `make`, key lookup, assignment, `delete`, the two-value lookup form, and nil-map behavior.
- **Clarity: 20 points.** Names and control flow are easy to understand. Keep the code formatted.

When finished, request a review without changes to your Go files.
