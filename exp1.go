func finishReq(timeout time.Duration) r ob {
2 - ch := make(chan ob)
3 + ch := make(chan ob, 1)
4 go func() {
5 result := fn()
6 ch <- result // block
7 } ()
8 select {
9 case result = <- ch:
10 return result
11 case <- time.After(timeout):
12 return nil
13 }
14 }
