package cmd

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/spf13/cobra"
)

const (
	letterBytes  = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	specialBytes = "!@#$%^&*()_+-=[]{}\\|;':\",.<>/?`~"
	numBytes     = "0123456789"
)

func init() {
	rand.NewSource(time.Now().UnixNano())
	rootCmd.AddCommand(passwordGenCmd)
}

func generatePassword(length int, useLetters bool, useSpecial bool, useNum bool) string {
	b := make([]byte, length)
	for i := range b {
		if useLetters {
			b[i] = letterBytes[rand.Intn(len(letterBytes))]
		} else if useSpecial {
			b[i] = specialBytes[rand.Intn(len(specialBytes))]
		} else if useNum {
			b[i] = numBytes[rand.Intn(len(numBytes))]
		}
	}
	return string(b)
}

var passwordGenCmd = &cobra.Command{
	Use:   "passgen",
	Short: "Generated password random",
	Run: func(cmd *cobra.Command, args []string) {
		password := generatePassword(12, true, true, true)
		fmt.Println(password)
	},
}
