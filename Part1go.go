package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

func main() {

	start := time.Now()

	// CHILD MODE
	if len(os.Args) > 1 {

		// PRODUCER
		if os.Args[1] == "producer" {
			producer()
			return
		}

		// CONSUMER
		if os.Args[1] == "consumer" {
			consumer()
			return
		}
	}

	// =========================
	// PARENT PROCESS
	// =========================

	producerCmd := exec.Command(os.Args[0], "producer")
	consumerCmd := exec.Command(os.Args[0], "consumer")

	// Producer stdout -> Consumer stdin (numbers)
	numPipe, _ := producerCmd.StdoutPipe()
	consumerCmd.Stdin = numPipe

	// Consumer stdout -> Producer stdin (ACK)
	ackPipe, _ := consumerCmd.StdoutPipe()
	producerCmd.Stdin = ackPipe

	// Both print to terminal using stderr
	producerCmd.Stderr = os.Stdout
	consumerCmd.Stderr = os.Stdout

	// Start consumer first
	consumerCmd.Start()
	producerCmd.Start()

	producerCmd.Wait()
	consumerCmd.Wait()

	elapsed := time.Since(start)
	fmt.Println("Elapsed time:", elapsed)

}

// =========================
// PRODUCER PROCESS
// =========================

func producer() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	for i := 1; i <= 5; i++ {

		// Print to terminal
		fmt.Fprintf(os.Stderr, "Producer: %d\n", i)

		// Send number to consumer
		fmt.Fprintln(writer, i)
		writer.Flush()

		// Wait for ACK
		_, err := reader.ReadString('\n')
		if err != nil {
			return
		}
	}
}

// =========================
// CONSUMER PROCESS
// =========================

func consumer() {
	scanner := bufio.NewScanner(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	for count := 0; count < 5; count++ {

		if !scanner.Scan() {
			return
		}

		val := strings.TrimSpace(scanner.Text())
		n, _ := strconv.Atoi(val)

		// Print to terminal
		fmt.Fprintf(os.Stderr, "Consumer: %d\n", n)

		// Send ACK back to producer
		fmt.Fprintln(writer, "ok")
		writer.Flush()
	}
}
Part
