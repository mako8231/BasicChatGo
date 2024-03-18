/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"net"
	"time"

	"github.com/spf13/cobra"
)

// createServerCmd represents the createServer command
var createServerCmd = &cobra.Command{
	Use:   "createServer",
	Short: "Make a simple UDP server",
	Long:  `Make this session an active UDP server given the specified port`,
	Run: func(cmd *cobra.Command, args []string) {
		//fmt.Println("createServer called")
		listenServer()
	},
}

func init() {
	rootCmd.AddCommand(createServerCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createServerCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createServerCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// This should be open a UDP connection using specified port
func listenServer() {
	server, err := net.ListenPacket("udp", ":8088")
	if err != nil {
		log.Panic("Error while opening the connection", err)
	}

	defer server.Close()

	//set a infinite loop
	fmt.Printf("Server started\n")
	for {
		//get the buffer (might I need to make a custom data structure? IDK how UDP works)
		buf := make([]byte, 1024)
		_, addr, err := server.ReadFrom(buf)
		if err != nil {
			continue
		}
		//get the goroutine here:
		go response(server, addr, buf)

	}
}

func response(server net.PacketConn, addr net.Addr, buffer []byte) {

	time := time.Now().Format(time.ANSIC)
	respStr := fmt.Sprintf("Time: %v. Message: %v!\n", time, string(buffer))
	server.WriteTo([]byte(respStr), addr)
}
