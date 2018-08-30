package safeType

import "sync"

type safeArrayString struct {
  l sync.Mutex
  a []string
  init bool
}
func(el *safeArrayString)Clear(){
  defer el.l.Unlock()

  el.l.Lock()
  el.init = true
  el.a = make( []string, 0 )
}
func(el *safeArrayString)Append( v string ){
  defer el.l.Unlock()

  el.l.Lock()
  if el.init == false {
    el.init = true
    el.a = make( []string, 0 )
  }

  el.a = append( el.a, v )
}
func(el *safeArrayString)Len() int {
  defer el.l.Unlock()

  el.l.Lock()
  return len( el.a )
}
func(el *safeArrayString)GetKey( i int ) string {
  defer el.l.Unlock()

  el.l.Lock()
  return el.a[ i ]
}
func(el *safeArrayString)Get() []string {
  defer el.l.Unlock()

  el.l.Lock()
  return el.a
}
