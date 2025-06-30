/*
* Big O       | Name          | Time for n=10 | 100     | 1000    | 10000           |
* O(1)        | Constant      |       1       |  1      |   1     |   1             |
* O(log n)    | Logarithmic   |      ~3       | ~7      |  ~10    |  ~14            |  
* O(n)        | Linear        |       10      | 100     | 1000    | 10000           |
* O(n log n)  | Linearithmic  |      ~33      | ~700    | ~10000  | ~140000         |
* O(n^2)      | Quadratic     |      100      | 10000   |   1M    |   100M          |
* O(2^n)      | Exponential   |      1024     | ~10^30  | ~10^300 | 💀 (infeasible) |
*
* O(1) - Getting the first item from a list 
* O(log n) - Binary search - like guessing a number between 1 - 100
* O(n) - Going through every item in a shopping list
* O(n log n) - Merge sort or quicksort
* O(n^2) - Comparign every student with every other student
* O(2^n) - Solving a puzzle where each step foubles choices
*
* Operation               |    Array      |      Slice       |            Map               |
* Access by index         |     O(1)      |       O(1)       |     ❌ Not by index          |
* Search by value         |     O(n)      |       O(n)       |  O(1)(average), O(n)(worst)  |
* Insert at end           | ❌ Fixed size |   O(1)amortized  |      O(1)(average)           |
* Insert in middle/start  |     ❌        |       O(n)       |          O(1)                |
* Delete at end           |     ❌        |       O(1)       |          O(1)                |
* Delete by value/index   |     ❌        |       O(n)       | O(1)(average), O(n)(worst)   |
* Memory allocation       | Fixed(static) |  Dynamic(grows)  |  Hash table(extra space)     |
* Iteration               |     O(n)      |       O(n)       |          O(n)                |
*
* Feature           |       Bubble sort        |          Merge sort                   |      
* Time complexity   |      O(n^2)always        |        O(n log n)always               |
* Space complexity  |           O(1)           |      O(n)(needs extra space)          |
* Stable?           |           Yes            |              Yes                      |
* Recursive?        |           No             |              Yes                      |
* In-place?         |           Yes            |      No(uses extra slices)            |
* Best Use          |  Education, small lists  |   Real-world sorting, large lists     |
* Worst Use         |       Large datasets     | Memory-limites systems (due to space) |
*
* 
* TIME COMPLEXITY                 SPACE COMPLEXITY  
* Case    | Complexity  |         Metric  | Value |
* Best    |   O(n)      |         Space   | O(1)  |
* Average |   O(n^2)    |
* Worst   |   O(n^2)    |
*/
