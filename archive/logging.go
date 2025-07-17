package main

import (
	"fmt"
	"os"
	"time"

	"codeberg.org/anaseto/gruid"
)

func (md *model) Print(message gruid.Msg) error {
    // Open the file in append mode, create it if it doesn't exist
    f, err := os.OpenFile(md.opt.LogFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    if err != nil {
        return fmt.Errorf("failed to open log file: %w", err)
    }
    defer f.Close()
    
    // Format the current time
    timestamp := time.Now().Format("2006-01-02 15:04:05")
    
    // Write the timestamped message
    logEntry := fmt.Sprintf("[%s] %s\n", timestamp, message)
    if _, err := f.WriteString(logEntry); err != nil {
        return fmt.Errorf("failed to write to log file: %w", err)
    }
    
    return nil
}
