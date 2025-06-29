// TODO [Implement]
// - Parse file paths from arguments
// - Launch goroutines to read + count words from each file
// - Use channels or mutex to send/merge individual counts
// - Aggregate word counts across files
// - Sort and print most common words

// TODO [Refactor & Improve]
// - Add input validation and error handling
// - Extract logic into clean helper functions
// - Make word splitting/normalizing consistent

// TODO [Bonus]
// - Add flags (e.g., -top=10, -normalize=true)
// - Add test cases or benchmarks (optional)
// - Write a README.md with usage examples

// TODO [Git]
// - Stage and commit with a clear message: "feat: add word counter with concurrency"
// - Push to your GitHub repo or portfolio project folder

package main  
import (
  "os"
  "fmt"
  "log"
  "strings"
  "sync"
  "sort"
)

type WordCount struct {
  Word string
  Count int
}

var globalCounts = make(map[string]int)
var mu sync.Mutex

func main()  {
  var wg sync.WaitGroup

  if len(os.Args) < 4 {
    fmt.Println("Usage: go run main.go <file_name> <file_name> <file_name>")
  }

  files := os.Args[1:]
  if len(files) == 0 {
    log.Fatal("No files given")
  }

  for _, f := range files {
    wg.Add(1)
    go func (file string)  {
      defer wg.Done()
      content, err := os.ReadFile(file)
      if err != nil {
        panic(err)
      }

      words := strings.Fields(string(content))
      counts := map[string]int{}

      for _, w := range words {
        counts[w]++
      }

      mu.Lock()
      for w, c := range counts {
        globalCounts[w] += c 
      }
      mu.Unlock()
    }(f)
  }
  
  wg.Wait()

  var counts []WordCount
  for w, c := range globalCounts {
    counts = append(counts, WordCount{w, c})
    sort.Slice(counts, func (i, j int) bool {
      return counts[i].Count > counts[j].Count
    })
  }
  for i := 0; i < 10 && i < len(counts); i++ {
    fmt.Printf("%s: %d\n", counts[i].Word, counts[i].Count)
  }
}
