package safeType

import "sync"

type safeMapStringPoint struct {
  l sync.Mutex
  m map[string][2]float64
  init bool
}
func(el *safeMapStringPoint)Clear(){
  el.init = true
  el.m = make( map[string][2]float64, 0 )
}
func(el *safeMapStringPoint)Set( k string, v [2]float64 ){
  defer el.l.Unlock()

  el.l.Lock()

  if el.init == false {
    el.init = true
    el.m = make( map[string][2]float64 )
  }

  el.m[ k ] = v
}
func(el *safeMapStringPoint)Len() int {
  defer el.l.Unlock()

  el.l.Lock()
  return len( el.m )
}
func(el *safeMapStringPoint)GetKey( k string ) [2]float64 {
  defer el.l.Unlock()

  el.l.Lock()
  return el.m[ k ]
}
func(el *safeMapStringPoint)DeleteKey( k string ) {
  defer el.l.Unlock()

  el.l.Lock()
  delete( el.m, k )
}
func(el *safeMapStringPoint)Get() map[string][2]float64 {
  defer el.l.Unlock()

  el.l.Lock()
  return el.m
}
