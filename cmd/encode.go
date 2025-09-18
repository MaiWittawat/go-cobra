package cmd

import (
	"encoding/base32"
	"encoding/base64"
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(base32EncodeCmd, base64EncodeCmd)
}

// ---------------- Base64 ---------------- //
var base64EncodeCmd = &cobra.Command{
	Use:   "base64 [string]",
	Short: "Encode input string to Base64",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		input := args[0] // รับค่าจาก args
		encoded := base64.StdEncoding.EncodeToString([]byte(input))
		fmt.Println(encoded)
	},
}

// ---------------- Base64 ---------------- //
var base32EncodeCmd = &cobra.Command{
	Use:   "base32 [string]",
	Short: "Encode input string to Base32",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		input := args[0]
		encoded := base32.StdEncoding.EncodeToString([]byte(input))
		fmt.Println(encoded)
	},
}
