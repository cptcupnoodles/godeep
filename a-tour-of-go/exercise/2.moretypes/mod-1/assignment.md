# Go practice: structs, arrays, pointers, and slices

Complete these exercises before [Creating a slice with make](https://go.dev/tour/moretypes/13).

Use what you have studied: basic types, functions, flow control, structs, pointers, arrays, slice literals, slice expressions, `len`, and `cap`.

Use fixed test values. Do not use `make`, `append`, `copy`, maps, methods, or packages other than `fmt`. Use index-based `for` loops; `range` is not required. You do not need keyboard input or automated tests.

Create a separate folder for each exercise: `easy`, `medium`, and `hard`. Put each solution in its own `main.go` with `package main`. Do not put all three solutions in one package.

Output layout is your choice. The values and state changes must match the requirements. Write your own solutions. No solution code is included below.

## 1. Easy: Update a book

### Task

Create a `Book` struct with these fields:

| Field | Type |
| --- | --- |
| Title | string |
| Pages | int |
| Read | bool |

Write these functions:

```go
func markRead(book *Book)
func addPages(book *Book, extra int) bool
```

### Requirements

1. Create a `[3]Book` array with the initial values below. Use struct literals with field names.
2. Create a slice named `selected` that refers to the first two books in the array.
3. In `markRead`, set the book's `Read` field to `true` through the pointer.
4. In `addPages`, reject `extra <= 0`: return `false` and do not change the book. Otherwise, increase `Pages` and return `true`.
5. Call both functions with a pointer to `selected[0]`. Use `20` for the page increase.
6. Print the first book through both the slice and the array. Both must show the updated fields.
7. Print the length and capacity of `selected`.
8. You can assume each pointer passed to these functions is non-nil.

### Initial values

| Title | Pages | Read |
| --- | --- | --- |
| Go Basics | 100 | false |
| Short Stories | 80 | false |
| Field Notes | 40 | true |

### Checks

Run each page-increase check from the initial values.

| Call | Returned value | First book's Pages |
| --- | --- | --- |
| `addPages(&selected[0], 20)` | true | 120 |
| `addPages(&selected[0], 0)` | false | 100 |
| `addPages(&selected[0], -5)` | false | 100 |

After `markRead`, the first book's `Read` field is `true`. The other books stay unchanged. For `selected`, length is `2` and capacity is `3`.

### Explain

Why does a change through a pointer to `selected[0]` also change the first array element?

## 2. Medium: Adjust a group of scores

### Task

Create a `Student` struct with `Name string` and `Score int` fields.

Write these functions:

```go
func addBonus(students []Student, bonus int) bool
func summarize(students []Student) (total int, passed int)
```

### Requirements

1. Create a `[5]Student` array with the initial values below.
2. Create `group` from array elements at indexes `1` through `3`, inclusive. Use a slice expression.
3. In `addBonus`, reject a negative bonus before any changes. Return `false` and leave every student unchanged.
4. For a non-negative bonus, use an index-based loop. Inside the loop, store a pointer to the current slice element in a variable. Update its score through that pointer.
5. Limit each updated score to `100`. Return `true`, including when the bonus is zero or the slice is empty.
6. In `summarize`, return the sum of scores and the number of students with a score of at least `60`. Do not change the students.
7. Call `addBonus(group, 5)`. Print the full array and the summary of `group`.
8. Print `len(group)` and `cap(group)`.
9. Also call both functions with a nil `[]Student` variable. Neither function must panic.
10. Assume initial scores are between `0` and `100`, and test bonuses are small integers.

### Initial values

| Index | Name | Score |
| --- | --- | --- |
| 0 | Ana | 55 |
| 1 | Ben | 58 |
| 2 | Cara | 76 |
| 3 | Dan | 98 |
| 4 | Eli | 40 |

### Checks

Start each check with a fresh array.

| Input to `addBonus` | Returned value | Group scores afterward | Group total | Group passed |
| --- | --- | --- | --- | --- |
| `group, 5` | true | 63, 81, 100 | 244 | 3 |
| `group, 0` | true | 58, 76, 98 | 232 | 2 |
| `group, -1` | false | 58, 76, 98 | 232 | 2 |
| nil slice, 5 | true | no elements | 0 | 0 |

Ana and Eli must stay unchanged. For `group`, length is `3` and capacity is `4`.

### Explain

Why can `addBonus` change array elements when its parameter is a slice? What changes if you update a local struct copy instead of using a pointer to the element?

## 3. Hard: Transfer stock between bins

### Task

Create a `Bin` struct with `Label string` and `Units int` fields.

Write these functions:

```go
func transfer(bins []Bin, from int, to int, amount int) bool
func totalUnits(bins []Bin) int
```

The indexes `from` and `to` refer to positions in the supplied slice.

### Requirements

1. Create a `[5]Bin` array with the initial values below.
2. Create two slices of that array: `left` contains indexes `0` through `2`; `right` contains indexes `2` through `4`. They share one element.
3. In `transfer`, reject any of these conditions: an index is negative; an index is outside the slice length; the indexes are equal; `amount <= 0`; the source has fewer units than `amount`.
4. Check index bounds before you access an element. A nil or empty slice must return `false` without a panic.
5. For every rejected transfer, return `false` and leave all bins unchanged.
6. For an accepted transfer, store pointers to the source and destination elements. Subtract units from the source and add the same amount to the destination through those pointers. Return `true`.
7. In `totalUnits`, use an index-based loop to calculate the total. Do not change the bins.
8. Perform the two successful transfers below in order. Print the full array after each transfer. Also print the full-array total before and after each transfer.
9. Print the length and capacity of both slices.
10. Create `empty := right[:0]`. Print its length, capacity, and whether it equals `nil`. Pass it to `transfer` and `totalUnits`.
11. Assume initial unit counts are non-negative and all counts fit in an `int`.

### Initial values

| Array index | Label | Units |
| --- | --- | --- |
| 0 | A | 10 |
| 1 | B | 20 |
| 2 | C | 30 |
| 3 | D | 40 |
| 4 | E | 50 |

### Successful sequence

| Step | Call | Returned value | Full array units afterward |
| --- | --- | --- | --- |
| 1 | `transfer(left, 0, 2, 4)` | true | 6, 20, 34, 40, 50 |
| 2 | `transfer(right, 0, 2, 10)` | true | 6, 20, 24, 40, 60 |

The full-array total stays `150`. After step 2, both `left[2]` and `right[0]` show `24` units.

### Rejected transfers

Start each check with the initial array. All calls below must return `false` and leave the array unchanged.

- `transfer(left, -1, 1, 1)`
- `transfer(left, 0, 3, 1)` — index 3 is outside the length, even though the capacity is larger.
- `transfer(left, 1, 1, 1)`
- `transfer(left, 0, 1, 0)`
- `transfer(left, 0, 1, -2)`
- `transfer(left, 0, 1, 11)`
- `transfer(empty, 0, 1, 1)`
- `transfer(nil, 0, 1, 1)`

Also check an exact-stock transfer: from a fresh array, `transfer(left, 0, 1, 10)` must return `true` and produce units `0, 30, 30, 40, 50`.

| Slice | Length | Capacity | Equals nil |
| --- | --- | --- | --- |
| left | 3 | 5 | false |
| right | 3 | 3 | false |
| empty | 0 | 3 | false |

`totalUnits(empty)` and `totalUnits(nil)` must both return `0`.

### Explain

1. Why does the second transfer also change the value seen through `left[2]`?
2. Why must index checks use the slice length instead of its capacity?
3. Why does `empty` have zero length but still differ from a nil slice?

## Review rubric

Each exercise is worth 100 points:

- **Correctness: 50 points.** Results, validation, and state changes match the checks.
- **Required concepts: 30 points.** The solution uses the required types, functions, pointers, and slice operations. Include your explanations as comments or separate notes.
- **Clarity: 20 points.** Names and control flow are easy to understand. Keep the code formatted.

The overall grade is the average of the three exercise grades. Exact output formatting is not graded. Do not add advanced features to earn points. When finished, request a review without changes to your Go files.

## Initial submission review — September 20, 2026

These grades record the submitted solutions reviewed on this date. They do not confirm later changes. The review was based on reading the source; the programs were not run.

| Exercise | Correctness /50 | Requirements /30 | Clarity /20 | Total /100 |
| --- | ---: | ---: | ---: | ---: |
| Easy: Update a book | 45 | 26 | 18 | 89 |
| Medium: Adjust a group of scores | 30 | 17 | 17 | 64 |
| Hard: Transfer stock between bins | 39 | 19 | 17 | 75 |

**Overall grade: 76/100.** Each exercise has equal weight: `(89 + 64 + 75) / 3 = 76` when rounded to the nearest whole number.

### Deductions and correction priorities

1. **Easy — page checks:** The struct, array, slice, pointer functions, successful `+20` call, and output through both the slice and array are correct. The submitted `main` does not demonstrate the `0` and negative `addPages` checks from fresh initial values (−5 correctness, −4 requirements). The explanation has the right idea, but say that the slice refers to the array's elements; the slice itself is not a pointer (−2 clarity).
2. **Medium — summary total:** `summarize` adds a score to `sum` only when that score is at least 60. The required total includes every student's score; only `passed` should depend on the 60-point condition (−15 correctness). The code calls `summarize(students[:])`; that is a valid call, but the requirement asks to print the summary of `group`, which this call does not show (−3 requirements). The required full-array print and calls using a nil slice are also missing (−6 requirements). The copy explanation reverses the behavior: a local struct copy would not update the original element (−4 requirements).
3. **Hard — required demonstrations and explanation:** Transfer validation, pointer updates, and the index-based total loop are implemented. The output does not show the full-array total before and after each successful transfer, and does not call `transfer`/`totalUnits` with `empty` and `nil` or print whether `empty == nil` (−7 correctness, −6 requirements). The explanation of `empty` is inaccurate: `right[:0]` has zero elements but still refers to the same backing array and is non-nil (−5 requirements). The output also omits the requested total before the first transfer (included in the deduction above).

The Easy `-5` check is a valid rejection test; its issue is only that the submitted `main` also needs to demonstrate the required successful `+20` case, which it does. No deduction was made for calling `summarize` on a slice other than `group` as an additional call; the submitted code simply does not include the requested group summary.
