package main

import (
	"fmt"
	"strconv"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

type Response struct {
	status  int
	message string
}

func sendMessageAsync(channel chan<- Response, message string) {
	time.Sleep(2 * time.Second)
	channel <- Response{status: 200, message: message}
}

func printMessage(channel <-chan Response) {
	response := <-channel
	fmt.Println("Response status code", response.status, "with message", response.message)
}

func TestCreateChannel(t *testing.T) {
	channel := make(chan Response, 5)
	defer close(channel)

	go sendMessageAsync(channel, "Hello")
	go sendMessageAsync(channel, "From Channel")
	go sendMessageAsync(channel, "Asd Asd")

	//  Kaya async await
	printMessage(channel)
	printMessage(channel)
	printMessage(channel)
}

func TestRangeChannel(t *testing.T) {
	channel := make(chan Response)

	go func() {
		for i := 0; i < 10; i++ {
			channel <- Response{status: 200, message: "Perulangan ke" + strconv.Itoa(i+1)}
		}

		close(channel)
	}()

	for data := range channel {
		fmt.Println(data.message)
	}
}

func TestSelectChannel(t *testing.T) {
	channel1 := make(chan Response)
	channel2 := make(chan Response)
	defer close(channel1)
	defer close(channel2)

	go sendMessageAsync(channel1, "Message from channel 1")
	go sendMessageAsync(channel2, "Message from channel 2")

	occurence := 0

	for {
		select {
		case data := <-channel1:
			fmt.Println(data.message)
			occurence++
		case data := <-channel2:
			fmt.Println(data.message)
			occurence++
		default:
			fmt.Println("loading")
		}

		if occurence == 2 {
			break
		}
	}
}

func TestSimulateRaceCondition(t *testing.T) {
	x := 0
	var mutex sync.Mutex
	for i := 1; i <= 1000; i++ {
		go func() {
			for j := 1; j <= 100; j++ {
				mutex.Lock()
				x = x + 1
				mutex.Unlock()
			}
		}()
	}

	time.Sleep(5 * time.Second)
	fmt.Println(x)
}

type BankAccount struct {
	Mutex   sync.RWMutex
	Balance int
}

func (account *BankAccount) AddBalance(amount int) {
	account.Mutex.Lock()
	account.Balance += amount
	account.Mutex.Unlock()
}

func (account *BankAccount) GetBalance() int {
	account.Mutex.RLock()
	currentBalance := account.Balance
	account.Mutex.RUnlock()

	return currentBalance
}

func TestSimulateRWRaceCondition(t *testing.T) {
	account := BankAccount{}

	for i := 0; i < 100; i++ {
		go func() {
			for j := 0; j < 100; j++ {
				account.AddBalance(1)
			}
		}()
	}

	time.Sleep(5 * time.Second)
	fmt.Println(account.GetBalance())
}

type UserBalance struct {
	sync.Mutex
	Name    string
	Balance int
}

func (user *UserBalance) Change(amount int) {
	user.Mutex.Lock()
	user.Balance = user.Balance + amount
	user.Mutex.Unlock()
}

func Transfer(user1 *UserBalance, user2 *UserBalance, amount int, waitGroup *sync.WaitGroup) {
	waitGroup.Add(1)

	go func() {
		defer waitGroup.Done()

		fmt.Println("lock user 1", user1.Name)
		user1.Change(-amount)
		time.Sleep(10 * time.Millisecond)

		fmt.Println("lock user 2", user2.Name)
		user2.Change(amount)
		time.Sleep(10 * time.Millisecond)
	}()
}

func TestDeadlock(t *testing.T) {
	group := &sync.WaitGroup{}
	user1 := UserBalance{Name: "John", Balance: 100000}
	user2 := UserBalance{Name: "Doe", Balance: 100000}

	for i := 0; i < 100000; i++ {
		go Transfer(&user1, &user2, 1, group)
	}

	group.Wait()
	fmt.Println(user1.Balance)
	fmt.Println(user2.Balance)
}

var counter = 0

func OnlyOnce() {
	time.Sleep(3 * time.Second)
	counter++
}

func TestOnce(t *testing.T) {
	once := new(sync.Once)
	group := new(sync.WaitGroup)

	for i := 0; i < 100; i++ {
		group.Add(1)

		go func() {
			once.Do(OnlyOnce)
			group.Done()
		}()
	}

	group.Wait()
	fmt.Println("Counter", counter)
}

func TestPool(t *testing.T) {
	group := new(sync.WaitGroup)
	pool := sync.Pool{
		New: func() any {
			return "Database Instance"
		},
	}

	for i := 0; i < 10; i++ {
		group.Add(1)
		go func() {
			defer group.Done()
			data := pool.Get()
			fmt.Println(data)

			time.Sleep(10 * time.Millisecond)
			pool.Put(data)
		}()
		group.Wait()
	}

	fmt.Println("Apps Killed")
}

type MapStore struct {
	instance *sync.Map
}

func (store *MapStore) addToMap(value int) {
	store.instance.Store(value, value)
}

func TestSyncMap(t *testing.T) {
	data := sync.Map{}
	group := new(sync.WaitGroup)
	store := MapStore{instance: &data}

	for i := 0; i < 100; i++ {
		group.Add(1)

		go func() {
			defer group.Done()
			store.addToMap(i)
		}()

		group.Wait()
	}

	group.Wait()
	data.Range(func(key, value any) bool {
		fmt.Println(key, ":", value)
		return true
	})
}

func TestSimulateRaceConditionAtomic(t *testing.T) {
	var x int32 = 0
	group := sync.WaitGroup{}
	for i := 1; i <= 1000; i++ {
		for j := 1; j <= 100; j++ {
			group.Add(1)

			go func() {
				defer group.Done()
				atomic.AddInt32(&x, 1)
			}()
		}
	}

	group.Wait()
	fmt.Println(x)
}
