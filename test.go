package main


func main() {
  t:=make(chan int)
  go func() {
	t<-1
	}()
   <-t
}
