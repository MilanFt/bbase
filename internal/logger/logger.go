package logger

import (
	"fmt"
	"sync"
	"time"

	"github.com/jedib0t/go-pretty/table"
)

type TaskStatus struct {
	Status  string
	Site    string
	Product string
}

var (
	TaskStatuses map[string]*TaskStatus
	enabled      bool
	statusMut    sync.Mutex
)

func EnableLogging() {
	enabled = true
	go func() {
		for {
			printTable()
			// Refresh rate
			time.Sleep(250 * time.Millisecond)
		}
	}()
}

func SetTaskStatus(id string, status string, site string, product string) {
	if !enabled {
		return
	}

	statusMut.Lock()
	defer statusMut.Unlock()
	if TaskStatuses == nil {
		TaskStatuses = make(map[string]*TaskStatus)
	}
	TaskStatuses[id] = &TaskStatus{
		Status:  status,
		Site:    site,
		Product: product,
	}
}

func printTable() {
	statusMut.Lock()
	defer statusMut.Unlock()

	if TaskStatuses == nil {
		return
	}

	tw := table.NewWriter()
	tw.AppendHeader(table.Row{"Site", "Product", "Status", "#"})

	s := make(map[string]map[string]map[string]int)
	for _, v := range TaskStatuses {
		if s[v.Site] == nil {
			s[v.Site] = make(map[string]map[string]int)
		}
		if s[v.Site][v.Product] == nil {
			s[v.Site][v.Product] = make(map[string]int)
		}
		s[v.Site][v.Product][v.Status]++
	}

	for k, v := range s {
		for k2, v2 := range v {
			for k3, v3 := range v2 {
				tw.AppendRow(table.Row{k, k2, k3, v3})
			}
		}
	}

	tw.SetAutoIndex(true)
	tw.SetStyle(table.StyleColoredDark)

	fmt.Printf("%s\n\n", tw.Render())
}
