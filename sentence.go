package sentence

import (
	"bufio"
	"bytes"
	_ "embed"
	"encoding/json"
	"errors"
	"math/rand"
	"sync"
)


//go:embed src/sentence1-10000.json
var sentenceData []byte

// lineOffsets stores the start and end indices of each line in sentenceData.
// Storing integers (2 ints per line) uses less memory than storing slice headers,
// satisfying strict efficiency requirements.
type lineOffset struct {
	start, end int
}

var (
	offsets   []lineOffset
	initOnce  sync.Once
)

// initIndex initializes the line offsets by scanning the byte slice once.
// This allows O(1) random access later without storing the full line content in separate slices.
func initIndex() {
	scanner := bufio.NewScanner(bytes.NewReader(sentenceData))
	// Set a reasonable buffer size if lines assume to be long, but default is usually efficient
	
	currentPos := 0
	offsets = make([]lineOffset, 0, 10000) // Pre-allocate for expected size

	for scanner.Scan() {
		lineLen := len(scanner.Bytes())
		if lineLen > 0 {
			offsets = append(offsets, lineOffset{
				start: currentPos,
				end:   currentPos + lineLen,
			})
		}
		// Calculate position for next line: +1 for the newline character consumed by Scan
		// Note: To be perfectly precise with Scanner, we need to handle potential \r\n vs \n.
		// However, scanner.Bytes() returns data without the newline.
		// Since we are indexing the *original* sentenceData, calculating exact position with Scanner is tricky
		// because Scanner doesn't report how many bytes it advanced (it could be \r\n or \n).
		//
		// BETTER STRATEGY: Use strict bytes.IndexByte for robust, high-perf offset calculation
		// that exactly matches standard byte slicing.
		currentPos += lineLen + 1 
		// Note: The simple addition above assumes \n. If strictly matching the user's reference `bufio` logic,
		// we should trust `bytes.Split` or manual byte scanning for 100% safety on offsets.
		//
		// Given the constraints, let's switch to the manual byte scan which is faster and safer for offsets.
	}
	
	// Re-do robustly with zero-copy byte scan manually to ensure offsets align exactly with sentenceData
	offsets = offsets[:0]
	start := 0
	for i, b := range sentenceData {
		if b == '\n' {
			if i > start {
				offsets = append(offsets, lineOffset{start, i})
			}
			start = i + 1
		}
	}
	// Handle last line if no newline at EOF
	if start < len(sentenceData) {
		offsets = append(offsets, lineOffset{start, len(sentenceData)})
	}
}


// Random returns a random sentence map.
// Unmarshalling is done on-demand to minimize startup memory usage while maintaining fast access.
func Random() (map[string]interface{}, error) {
	initOnce.Do(initIndex)

	if len(offsets) == 0 {
		return nil, errors.New("sentence data is empty")
	}

	// O(1) Selection
	idx := rand.Intn(len(offsets))
	loc := offsets[idx]

	// Zero-copy slice from original data
	line := sentenceData[loc.start:loc.end]
	
	// Trim just in case of carriage returns if manual scan didn't catch \r
	line = bytes.TrimSpace(line)

	var result map[string]interface{}
	if err := json.Unmarshal(line, &result); err != nil {
		return nil, err
	}

	return result, nil
}
