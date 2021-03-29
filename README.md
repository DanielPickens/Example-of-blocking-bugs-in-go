To demonstrate errors in message passing, I am going to use a blocking bug from Kubernetes.

The finishReq function creates a child goroutine using an anonymous function at line 4 to handle a request—a common practice in
  Go server programs. The child goroutine executes fn() and
 sends result back to the parent goroutine through channel
ch at line 6. The child will block at line 6 until the parent
pulls result from ch at line 9. Meanwhile, the parent will
block at select until either when the child sends result to ch
(line 9) or when a timeout happens (line 11). If timeout happens earlier or if Go runtime (non-deterministically) chooses
the case at line 11 when both cases are valid, the parent will
return from requestReq() at line 12, and no one else can
pull result from ch any more, resulting in the child being
blocked forever. The fix is to change ch from an unbuffered
channel to a buffered one, so that the child goroutine can
always send the result even when the parent has exit.
This bug demonstrates the complexity of using new features in Go and the difficulty in writing correct Go programs
like this. Programmers have to have a clear understanding
of goroutine creation with anonymous function, a feature
Go proposes to ease the creation of goroutines, the usage
of buffered vs. unbuffered channels, the non-determinism
of waiting for multiple channel operations using select,
