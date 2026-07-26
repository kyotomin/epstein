package prompt

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"golang.org/x/term"
)

func ReadPassword(label string) (string, error) {
	fmt.Print(label)

	fd := int(os.Stdin.Fd())
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGTERM, os.Interrupt)
	defer signal.Stop(sigChan)

	type res struct {
		pass []byte
		err  error
	}
	resChan := make(chan res, 1)

	go func() {
		pass, err := term.ReadPassword(fd)
		resChan <- res{pass, err}
	}()

	select {
	case <-sigChan:
		return "", fmt.Errorf("input interrupt")
	case res := <-resChan:
		fmt.Println()
		if res.err != nil {
			return "", fmt.Errorf("error entering password")
		}

		return string(res.pass), nil
	}
}

func ReadAndConfirmPassword() (string, error) {
	for {
		pass, err := ReadPassword("Password: ")
		if err != nil {
			return "", err
		}

		confirm, err := ReadPassword("Confirm password: ")
		if err != nil {
			return "", err
		}

		if pass != confirm {
			fmt.Println("Passwords do not match.")
			continue
		}

		if len(pass) == 0 {
			fmt.Println("Password cannot be empty.")
			continue
		}
		return pass, nil
	}
}
