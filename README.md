# Go Reloaded

A text completion, formatting, and transformation CLI tool written in Go.

`go-reloaded` reads a text file, applies grammatical, casing, base-conversion, and punctuation transformations specified by embedded flags or syntax rules, and writes the processed text to an output file.

---

## Features

### 1. Case Transformations
- `(up)` / `(up, n)`: Converts the preceding word (or preceding `n` words) to **UPPERCASE**.
- `(low)` / `(low, n)`: Converts the preceding word (or preceding `n` words) to **lowercase**.
- `(cap)` / `(cap, n)`: Capitalizes the first letter of the preceding word (or preceding `n` words).

### 2. Base Conversions
- `(hex)`: Converts the preceding hexadecimal number (e.g., `1A`) to decimal (`26`).
- `(bin)`: Converts the preceding binary number (e.g., `101`) to decimal (`5`).

### 3. Punctuation Formatting
- Automatically formats punctuation marks (`,`, `.`, `!`, `?`, `:`, `;`).
- Removes extraneous spaces before punctuation.
- Ensures a single space follows punctuation marks (unless followed by another punctuation mark).
- Groups contiguous punctuation (e.g., `...`, `!?`) correctly without unwanted spaces.

### 4. Quote Formatting
- Ensures single quotes (`'`) wrap around text tightly without inner leading or trailing whitespace.
  - *Example:* `' amazing '` → `'amazing'`

### 5. Article Grammar Correction
- Automatically changes `a` to `an` (or `A` to `An`) when the following word begins with a vowel (`a`, `e`, `i`, `o`, `u`) or `h`.
  - *Example:* `a apple` → `an apple`, `a hour` → `an hour`

---

## Getting Started

### Prerequisites
- [Go](https://go.dev/dl/) 1.25 or later.

### Installation & Execution

Clone or download the repository, then navigate to the project directory:

```bash
cd go-reloaded
```

To run the program with an input file and output file:

```bash
go run . sample.txt result.txt
```

---

## Usage Example

### Input (`sample.txt`)
```text
it (cap) was the best of times, it was the worst of times (up) , it was the age of wisdom, it was the age of foolishness (cap, 6) , it was the epoch of belief, it was the epoch of incredulity, it was the season of Light, it was the season of darkness, it was the spring of hope, IT WAS THE (low, 3) winter of despair.
```

### Command
```bash
go run . sample.txt result.txt
```

### Output (`result.txt`)
```text
It was the best of times, it was the worst of TIMES, it was the age of wisdom, It Was The Age Of Foolishness, it was the epoch of belief, it was the epoch of incredulity, it was the season of Light, it was the season of darkness, it was the spring of hope, it was the winter of despair.
```

---

## Project Structure

```text
go-reloaded/
├── main.go            # Entry point: handles CLI args and file I/O
├── transformation.go  # Core text manipulation & flag parsing logic
├── utils.go           # Helper functions for string parsing & conversions
├── main_test.go       # Unit tests for functions and pipeline
├── go.mod             # Go module definition
├── sampel.txt         # Sample input file
└── result.txt         # Output result file
```

---

## Running Tests

To run the full unit test suite:

```bash
go test -v .
```

All tests verify text casing, base conversion, punctuation, quote formatting, article fixes, and full pipeline integration.
