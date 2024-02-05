/*
Copyright © 2024 Devbird007
Copyrights apply to this source code.
Check LICENSE for details.
*/
package cmd

import (
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	"github.com/devbird007/pScan/scan"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// scanCmd represents the scan command
var scanCmd = &cobra.Command{
	Use:   "scan",
	Short: "Run a port scan on the hosts",

	RunE: func(cmd *cobra.Command, args []string) error {
		hostsFile := viper.GetString("hosts-file")

		portRange, err := cmd.Flags().GetString("port-range")
		if err != nil {
			return err
		}

		if portRange == "" {
			ports, err := cmd.Flags().GetIntSlice("ports")
			if err != nil {
				return err
			}
			return scanAction(os.Stdout, hostsFile, ports)

		} else {
			ports, err := parseString(portRange)
			if err != nil {
				return err
			}

			return scanAction(os.Stdout, hostsFile, ports)
		}
	},
}

func scanAction(out io.Writer, hostsFile string, ports []int) error {
	hl := &scan.HostsList{}

	if err := hl.Load(hostsFile); err != nil {
		return err
	}

	results := scan.Run(hl, ports)

	return printResults(out, results)
}

func printResults(out io.Writer, results []scan.Results) error {
	message := ""

	for _, r := range results {
		message += fmt.Sprintf("%s:", r.Host)

		if r.NotFound {
			message += " Host not found\n\n"
			continue
		}

		message += fmt.Sprintln()

		for _, p := range r.PortStates {
			message += fmt.Sprintf("\t%d: %s\n", p.Port, p.Open)
		}

		message += fmt.Sprintln()
	}

	_, err := fmt.Fprint(out, message)
	return err
}

func parseString(s string) ([]int, error) {
	parts := strings.Split(s, "-")

	if len(parts) != 2 {
		return nil, errors.New("invalid range format")
	}

	start, err := strconv.Atoi(parts[0])
	if err != nil {
		return nil, err
	}

	end, err := strconv.Atoi(parts[1])
	if err != nil {
		return nil, err
	}

	intSlice := make([]int, end-start+1)
	for i := range intSlice {
		intSlice[i] = start + i
	}

	return intSlice, nil
}

func init() {
	rootCmd.AddCommand(scanCmd)

	scanCmd.Flags().IntSliceP("ports", "p", []int{22, 88, 443}, "ports to scan")
	scanCmd.Flags().StringP("port-range", "r", "", "range of ports to scan")
}
