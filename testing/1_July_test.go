package main  
import (
  "reflect"
  "testing"
)

func toSlice(head *Node) []int {
    var out []int
    for curr := head; curr != nil; curr = curr.next {
        out = append(out, curr.value)
    }
    return out
}

func TestAppendValue(t *testing.T) {
  test := []struct {
    name string
    input []int 
    operation string
    target int  
    want interface{}
  }{
    {
      name: "append to empty list",
      input: []int{}, operation: "append", target: 0,
      want: []int{0},     
    },
    {
      name: "append multiple value",
      input: []int{10, 20}, operation: "append", target: 30,
      want: []int{10, 20, 30},     
    },
    {
      name: "find existing value",
      input: []int{10, 20, 30}, operation: "find", target: 20,
      want: 20, 
    },
    {
      name: "find missing value",
      input: []int{10, 20}, operation: "find", target: 99,
      want: nil,     
    },
    {
      name: "delete head",
      input: []int{10, 20, 30}, operation: "delete", target: 10,
      want: []int{20, 30},
    },
    {
      name: "delete tail",
      input: []int{10, 20, 30}, operation: "delete", target: 30,
      want: []int{10, 20},
    },
    {
      name: "delete non-existent",
      input: []int{10, 20}, operation: "delete", target: 99,
      want: []int{10, 20}, 
    },
  }

  for _, tt := range test {
    t.Run(tt.name, func (t *testing.T)  {
      switch tt.operation {
      case "append":
        values := append(tt.input, tt.target)
        result := toSlice(append_value(nil, values))
        if !reflect.DeepEqual(result, tt.want){
          t.Errorf("AppendValue() = %v, want %v", result, tt.want)
        }
      case "find":
        list := append_value(nil, tt.input)
        found := find_value(list, tt.target)
        result := -1
        if found != nil {
          result = found.value
        }
       if tt.want == nil {
          if result != -1 {
            t.Errorf("FindValue() = %v, want nil", result)
          }
        } else {
          expected := tt.want.(int)
          if result != expected {
            t.Errorf("FindValue() = %v, want %v", result, expected)
          }
        }
      case "delete":
        list := append_value(nil, tt.input)
        result := toSlice(delete_node(list, tt.target))
        if !reflect.DeepEqual(result, tt.want){
          t.Errorf("DeleteNode() = %v, want %v", result, tt.want)
        }
      }
    })
  }
}
