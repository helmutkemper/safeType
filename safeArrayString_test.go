package safeType

import (
  "fmt"
  "strconv"
  "sync"
)

func ExampleSafeArrayString_Append() {
  var t safeArrayString
  var wg sync.WaitGroup

  for i := 0; i != 1000000; i += 1 {
    wg.Add(1)
    go func(i int){

      t.Append( strconv.Itoa( i ) )

      wg.Done()
    }( i )
  }
  wg.Wait()

  for i := 0; i != 1000000; i += 1 {
    wg.Add(1)
    go func(i int){

      t.GetKey( i )

      wg.Done()
    }( i )
  }

  wg.Wait()
  fmt.Printf( "total: %v",t.Len()  )

  // Output:
  // total: 1000000
}
