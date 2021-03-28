package main

import (
	"fmt"
	"os"
	"sync"
	"time"
)

// Simple individual tasks
func makeHotelReservation(wg *sync.WaitGroup) {
	fmt.Println("Done making hotel reservation.")
	wg.Done()
}
func bookFlightTickets(wg *sync.WaitGroup) {
	fmt.Println("Done booking flight tickets.")
	wg.Done()
}
func orderADress(wg *sync.WaitGroup) {
	fmt.Println("Done ordering a dress.")
	wg.Done()
}
func payCreditCardBills(wg *sync.WaitGroup) {
	fmt.Println("Done paying Credit Card bills.")
	wg.Done()
}

// Tasks that will be executed in parts

// Writing Mail
func writeAMail(wg *sync.WaitGroup) {
	fmt.Println("Wrote 1/3rd of the mail.")
	go continueWritingMail1(wg)
}
func continueWritingMail1(wg *sync.WaitGroup) {
	fmt.Println("Wrote 2/3rds of the mail.")
	go continueWritingMail2(wg)
}
func continueWritingMail2(wg *sync.WaitGroup) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Such a Panic!")
			os.Exit(1)
		}
	}()
	fmt.Println("Done writing the mail.")
	panic("foo")
	wg.Done()
}

// Listening to Audio Book
func listenToAudioBook(wg *sync.WaitGroup) {
	fmt.Println("Listened to 10 minutes of audio book.")
	go continueListeningToAudioBook(wg)
}
func continueListeningToAudioBook(wg *sync.WaitGroup) {
	time.Sleep(4 * time.Second)
	fmt.Println("Done listening to audio book. THE END")
	wg.Done()
}

func foobar(wg *sync.WaitGroup) {

}

// All the tasks we want to complete in the day.
// Note that we do not include the sub tasks here.
var listOfTasks = []func(*sync.WaitGroup){
	makeHotelReservation, bookFlightTickets, orderADress,
	payCreditCardBills, writeAMail, listenToAudioBook,
}

func main() {

	// var waitGroup sync.WaitGroup
	// // Set number of effective goroutines we want to wait upon
	// waitGroup.Add(len(listOfTasks))

	// for _, task := range listOfTasks {
	// 	// Pass reference to WaitGroup instance
	// 	// Each of the tasks should call on WaitGroup.Done()
	// 	task(&waitGroup)
	// }
	// // Wait until all goroutines have completed execution.
	// waitGroup.Wait()
}
