package safeType

import "sync"

type safeMapStringString struct {
  l sync.Mutex
  m map[string]string
  init bool
}
func(el *safeMapStringString)Clear(){
  el.init = true
  el.m = make( map[string]string )
}
func(el *safeMapStringString)Set( k string, v string ){
  defer el.l.Unlock()

  el.l.Lock()

  if el.init == false {
    el.init = true
    el.m = make( map[string]string )
  }

  el.m[ k ] = v
}
func(el *safeMapStringString)Len() int {
  defer el.l.Unlock()

  el.l.Lock()
  return len( el.m )
}
func(el *safeMapStringString)GetKey( k string ) string {
  defer el.l.Unlock()

  el.l.Lock()
  return el.m[ k ]
}
func(el *safeMapStringString)DeleteKey( k string ) {
  defer el.l.Unlock()

  el.l.Lock()
  delete( el.m, k )
}
func(el *safeMapStringString)Get() map[string]string {
  defer el.l.Unlock()

  el.l.Lock()
  return el.m
}
