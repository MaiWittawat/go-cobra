package cmd

import (
	"crypto/sha256"
	"crypto/sha512"
	"encoding/base64"
	"fmt"
	"log"
	"math/rand"
	"strconv"

	"github.com/spf13/cobra"
	"golang.org/x/crypto/argon2"
	"golang.org/x/crypto/bcrypt"
)

func init() {
	rootCmd.AddCommand(sha256HashCmd, sha512HashCmd, bcryptHashCmd, argon2HashCmd)
}

// ---------------- SHA-512 ---------------- //
var sha512HashCmd = &cobra.Command{
	Use:   "sha512 [string]",
	Short: "Hash input string to SHA-512",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		input := args[0]
		hash := sha512.Sum512([]byte(input))
		fmt.Println(hash)
	},
}

// ---------------- SHA-256 ---------------- //
var sha256HashCmd = &cobra.Command{
	Use:   "sha256 [string]",
	Short: "Hash input string to SHA-256",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		input := args[0]
		hash := sha256.Sum256([]byte(input))
		fmt.Println(hash)
	},
}

// ---------------- Bcrypt ---------------- //
var bcryptHashCmd = &cobra.Command{
	Use:   "bcrypt [string] [int]",
	Short: "Hash input string to Bcrypt",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		password := args[0]
		cost, err := strconv.Atoi(args[1])
		if err != nil {
			fmt.Println("error: %w", err)
		}

		hash, err := bcrypt.GenerateFromPassword([]byte(password), cost)
		if err != nil {
			fmt.Println("error: %w", err)
		}

		fmt.Println(string(hash))
	},
}

// ---------------- Argon ---------------- //
var argon2HashCmd = &cobra.Command{
	Use:   "argon2 [string] [int]",
	Short: "Hash input string to argon2",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		const time = 1           // iterations
		const memory = 64 * 1024 // 64 MB
		const threads = 4
		const keyLen = 32
		password := args[0]
		cost, err := strconv.Atoi(args[1])
		if err != nil {
			fmt.Println("error: %w", err)
		}

		salt := generateSalt(cost)
		hash := argon2.IDKey([]byte(password), salt, time, memory, threads, keyLen)
		fmt.Println(base64.RawStdEncoding.EncodeToString(hash))
	},
}

func generateSalt(size int) []byte {
	salt := make([]byte, size)
	if _, err := rand.Read(salt); err != nil {
		log.Fatal(err)
	}
	return salt
}
