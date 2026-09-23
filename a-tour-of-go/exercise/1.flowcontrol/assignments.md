# Go practice assignments

These three assignments increase in difficulty. Use them to practise the Go basics and flow control topics that you completed.

Use fixed test values. You do not need keyboard input, slices, maps, or structs. Write the solutions yourself.

Reference lessons:

- [A Tour of Go: Basics](https://go.dev/tour/basics/1)
- [A Tour of Go: Flow control](https://go.dev/tour/flowcontrol/1)

## Assignment 1: Temperature converter

Write a function with this signature:

```go
func convertTemperature(celsius float64) (float64, float64)
```

### Requirements

1. Return Fahrenheit first, then Kelvin.
2. Use these formulas:
   - Fahrenheit = Celsius × 9 / 5 + 32
   - Kelvin = Celsius + 273.15
3. Declare numeric constants for the offsets `32` and `273.15`.
4. Declare an integer temperature with `var` and an initial value.
5. Convert that integer to `float64` when you call the function.
6. Use `:=` to receive both results.
7. Store a location name in a `string` variable.
8. Store whether the Celsius temperature is below zero in a `bool` variable.
9. Declare a variable without an initial value. Print its value before you assign a new value.
10. Print the location, the Celsius, Fahrenheit, and Kelvin temperatures, and the freezing status.

### Test values and expected results

| Celsius | Fahrenheit | Kelvin | Below zero |
| --- | --- | --- | --- |
| 0 | 32 | 273.15 | false |
| 25 | 77 | 298.15 | false |
| -10 | 14 | 263.15 | true |

Exact decimal formatting is optional.

## Assignment 2: Score evaluator

Write a function with this signature:

```go
func evaluateScore(score int) (grade string, passed bool)
```

### Requirements

1. Use the named results `grade` and `passed`.
2. If the score is outside `0` through `100`, return `"Invalid"` and `false`.
3. Declare a passing-score constant with the value `60`.
4. Use `if` and `else` to set `passed`.
5. Use a `switch` with no condition to select the grade. Use lower-bound comparisons. Put the cases in the correct order.
6. For a valid score, assign the named results and use a bare `return`.
7. In the caller, use an `if` with a short statement to receive both results. Print a pass or fail message.
8. Write a second function that accepts a grade and returns a short message. Use a regular `switch` on the grade. Choose the function name and messages yourself.

### Grade rules

| Score range | Grade |
| --- | --- |
| 90–100 | A |
| 80–89 | B |
| 70–79 | C |
| 60–69 | D |
| 0–59 | F |

### Test values and expected results

| Score | Grade | Passed |
| --- | --- | --- |
| -1 | Invalid | false |
| 59 | F | false |
| 60 | D | true |
| 85 | B | true |
| 95 | A | true |
| 101 | Invalid | false |

### Explain your result

Explain why testing `>= 60` before `>= 90` gives the wrong grade.

## Assignment 3: Countdown report

Write a function with this signature:

```go
func countdown(start int) (sum int, valid bool)
```

### Requirements

1. Print `Starting countdown`.
2. Use `defer` to print `Countdown finished` whenever the function returns.
3. If `start < 1`, print `Invalid start` and return `0, false`.
4. Use a loop to count from `start` down to `1`.
5. In each iteration:
   - Immediately print `Now:` and the current number.
   - Add the current number to `sum`.
   - Defer printing `Saved:` and the current number.
6. After the loop, print `Sum:` and the sum.
7. Return the sum and `true`.

### Test values and expected results

| Start | Returned sum | Returned valid |
| --- | --- | --- |
| 3 | 6 | true |
| 0 | 0 | false |

For `countdown(3)`, print this exact output:

```text
Starting countdown
Now: 3
Now: 2
Now: 1
Sum: 6
Saved: 1
Saved: 2
Saved: 3
Countdown finished
```

Return `6, true`. The return values are separate from the printed output.

For `countdown(0)`, print this exact output:

```text
Starting countdown
Invalid start
Countdown finished
```

Return `0, false`. The return values are separate from the printed output.

### Extra challenge

Put two print statements inside each deferred function. Pass the current iteration number as an explicit argument to the deferred function. The extra print statements can change the output shown above.

## Next step

Start with Assignment 1. When you are ready, request a review of your solution. Ask for a review without changes to your files.

## Initial submission review — September 5, 2026

These grades record the first review. They do not confirm that later changes were tested or reviewed.

| Assignment | Correctness /50 | Requirements /30 | Clarity /20 | Total /100 |
| --- | --- | --- | --- | --- |
| Temperature converter | 50 | 22 | 20 | 92 |
| Score evaluator | 40 | 26 | 20 | 86 |
| Countdown report | 48 | 30 | 20 | 98 |

**Overall grade: 92/100.** Each assignment has equal weight: `(92 + 86 + 98) / 3 = 92`.

### Deductions and correction priorities

Line numbers refer to the initial submission.

1. **Temperature:** In `temperature/main.go`, line 5 uses `converter` instead of the required name `convertTemperature` (−3 requirements). Lines 14–21 omit the zero-value demonstration (−5 requirements). Use the required function name. Declare a variable without an initial value and print its value before you assign a new value.
2. **Score output:** In `score/main.go`, lines 56–59 print booleans without explicit pass or fail messages (−5 correctness). Line 59 also omits the comment from `message(g)` in the failure branch (−5 correctness). Add the required messages to the caller.
3. **Score explanation:** In `score/main.go`, line 64 gives an incorrect explanation (−4 requirements). The first true switch case runs. If the `>= 60` case and its D assignment come before the `>= 90` case, scores from 60 through 100 receive D. Correct the explanation.
4. **Countdown output:** In `countdown/main.go`, line 16 prints `Now` instead of `Now:` (−2 correctness). Add the colon to meet the exact output requirement.

### Tests from the initial review

Tests used temporary copies. No Go source files were changed.

- **Temperature:** Values `0`, `25`, `-10`, `12.5`, and `-0.5` gave correct Fahrenheit and Kelvin results.
- **Score:** Every integer from `-2` through `102` gave the correct grade and boolean result. Caller checks confirmed the missing output: `59` printed `F false`; `60` printed `60 true D Bad`; `-1` and `101` printed `Invalid false`.
- **Countdown:** Values `3`, `0`, `-1`, `1`, and `6` returned the correct sum and validity. Deferred prints ran in the correct order. The missing colon was the only difference in the function output.

No clarity points were deducted. The spelling `Fahrentheit` is optional cleanup. There was no deduction for the location name, temperature formatting, comment wording, inferred boolean type, or omitted extra challenge.
