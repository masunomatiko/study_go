package bank

var deposits = make(chan int)
var balances = make(chan int)
var withdraw = make(chan int)
var result = make(chan bool)

func Deposit(amount int) { deposits <- amount }
func Balance() int       { return <-balances }
func WithDraw(amount int) bool {
	withdraw <- amount
	return <-result
}

func teller() {
	var balance int
	for {
		select {
		case amount := <-deposits:
			balance += amount
		case balances <- balance:
		case amount := <-withdraw:
			if amount <= balance {
				balance -= amount
				result <- true
			} else {
				result <- false
			}
		}
	}
}

func init() {
	go teller()
}
