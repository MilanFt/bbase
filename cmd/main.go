package main

import (
	"bbase/internal/logger"
	"bbase/internal/profile"
	"bbase/internal/proxy"
	"bbase/internal/reader"
	"bbase/internal/site"
	"bbase/internal/site/example"
	"bbase/internal/task"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// Example flow to showcase the most important features
func main() {
	// Initialize checkout sounds
	// only needs to be done once
	// err := sound.InitSound("pathtofile.mp3")

	// Play the checkout sound
	// sound.PlayCheckoutSound()

	// Enable task status logging, this is optional
	logger.EnableLogging()

	/*
		Profile group + profile
	*/
	profileGroup := profile.NewProfileGroup()

	address := profile.NewAddress(
		"USA",
		"Michael",
		"Robert",
		"Test Street 2",
		"Test City",
		"Test State",
		"12345",
		"1234567890",
	)

	paymentMethod := profile.NewPaymentMethod(
		"Michael Robert",
		"4242424242424242",
		"123",
		"12",
		"2020",
	)

	// Create profile with address and payment method
	// and set it to the profile group
	profile.NewProfile(
		profileGroup,
		"test@test.com",
		address,
		address,
		paymentMethod,
		nil, // Optional web3 wallet information
	)

	/*
		Proxy group + proxies
	*/
	proxyGroup := proxy.NewProxyGroup()

	proxies := reader.ReadTXT("data/test.txt")
	for _, p := range proxies {
		proxy.NewProxy(proxyGroup, p.IP)
	}

	/*
		Task group + tasks
	*/
	taskGroup, err := task.NewTaskGroup(
		time.Duration(500*time.Millisecond),
		proxyGroup,
		profileGroup,
	)
	if err != nil {
		panic(err)
	}

	tasks := reader.ReadCSV("data/test.csv")
	// We currently do not use any information
	// from the CSV file - only the number of tasks,
	// but this process is easily modifiable
	for range tasks {
		task.NewTask(taskGroup)
	}

	/*
		Initializing a site and running the tasks
	*/
	exampleSite := example.NewExample()

	// Run all tasks in the task group
	site.RunTasks(exampleSite, taskGroup)

	// Prevent the program from exiting after
	// all tasks have been completed
	// The user can stop the program by pressing CTRL+C
	quitChannel := make(chan os.Signal, 1)
	signal.Notify(quitChannel, syscall.SIGINT, syscall.SIGTERM)
	<-quitChannel

	fmt.Println("Exiting...")
}
