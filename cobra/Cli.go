package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/spf13/cobra"
)

var (
	persistRootFlag bool
	LocalRootFlag   bool
	time            int
	rootCmd         = &cobra.Command{
		Use:   " mt task",
		Short: "first CLI\n",
		Long:  `my first CLI attempt`,
		Run: func(cmd *cobra.Command, args []string) {
			// this runs first
			fmt.Println("rtt")
		},
	} // root command
	echoCmd = &cobra.Command{
		Use:   "echo [string to echo]",
		Short: "print given string to stdout",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Echo: " + strings.Join(args, " "))
		},
	} // sub command
	timesCmd = &cobra.Command{
		Use:   "times [Strings to echo times]",
		Short: "Prints given std command to stdout multiple times",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {

			for i := 0; i < time; i++ {
				fmt.Println("Echo! *  : " + strings.Join(args, " "))
			}
		}, // its a subcommand of echoCmd command <go run <name>  help echo times>

	}
)

func init() {
	// initilation
	rootCmd.PersistentFlags().BoolVarP(&persistRootFlag, "persistFlag", "p", false, "a presistent root flag") // only avalable global root < go run <name> help echo>
	rootCmd.Flags().BoolVarP(&LocalRootFlag, "localFlag", "l", false, "a local root flag")                    // only avalable in root < go run <name> --help>
	timesCmd.Flags().IntVarP(&time, "time", "n", 1, "number of times to echo to stdout")
	rootCmd.AddCommand(echoCmd)  // adding to root command
	echoCmd.AddCommand(timesCmd) // adding to echo command that makes it a subcmd of subcmd <go run echo times hello world -n 5>

}

func main() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}

}
