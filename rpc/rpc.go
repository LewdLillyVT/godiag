package rpc

import (
	"fmt"
	"time"

	"github.com/altfoxie/drpc"
)

var client *drpc.Client

// StartRPC initializes the Discord RPC connection and sets the activity with a GitHub link.
func StartRPC() error {
	var err error
	client, err = drpc.New("1316963964650782830") // Replace with your actual Discord application ID
	if err != nil {
		return fmt.Errorf("failed to initialize Discord RPC client: %w", err)
	}

	fmt.Println("Starting Discord RPC...")
	err = client.SetActivity(drpc.Activity{
		Details: "Generating diagnostics reports",
		State:   "Idle",
		Timestamps: &drpc.Timestamps{
			Start: time.Now(),
		},
		Assets: &drpc.Assets{
			LargeImage: "icon",
			LargeText:  "",
		},
		Buttons: []drpc.Button{
			{
				Label: "GitHub Repo",
				URL:   "https://github.com/LewdLillyVT/godiag",
			},
		},
	})
	if err != nil {
		return fmt.Errorf("failed to set Discord activity: %w", err)
	}

	fmt.Println("Discord RPC activity set successfully.")
	return nil
}

// StopRPC cleans up the Discord RPC client and closes the connection.
func StopRPC() {
	if client == nil {
		fmt.Println("No active Discord RPC client to stop.")
		return
	}

	fmt.Println("Stopping Discord RPC...")
	err := client.Close()
	if err != nil {
		fmt.Printf("Error closing Discord RPC client: %v\n", err)
	} else {
		fmt.Println("Discord RPC stopped successfully.")
	}
}
