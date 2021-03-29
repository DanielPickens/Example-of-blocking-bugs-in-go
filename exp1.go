func finishReq(timeout time.Duration) r ob {
  ch := make(chan ob)
  ch := make(chan ob, 1)
 go func() {
  result := fn()
  ch <- result // block
  } ()
  select {
 case result = <- ch:
return result
case <- time.After(timeout):
 return nil
 }
 }
