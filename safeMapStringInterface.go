package safeType

import "sync"

type safeMapStringInterface struct {
  l sync.Mutex
  m map[string]interface{}
  init bool
}
func(el *safeMapStringInterface)Clear(){
  el.init = true
  el.m = make( map[string]interface{} )
}
func(el *safeMapStringInterface)Set( k string, v interface{} ){
  defer el.l.Unlock()

  el.l.Lock()

  if el.init == false {
    el.init = true
    el.m = make( map[string]interface{} )
  }

  el.m[ k ] = v
}
func(el *safeMapStringInterface)Len() int {
  defer el.l.Unlock()

  el.l.Lock()
  return len( el.m )
}
func(el *safeMapStringInterface)GetKey( k string ) interface{} {
  defer el.l.Unlock()

  el.l.Lock()
  return el.m[ k ]
}
func(el *safeMapStringInterface)DeleteKey( k string ) {
  defer el.l.Unlock()

  el.l.Lock()
  delete( el.m, k )
}
func(el *safeMapStringInterface)Get() map[string]interface{} {
  defer el.l.Unlock()

  el.l.Lock()
  return el.m
}
