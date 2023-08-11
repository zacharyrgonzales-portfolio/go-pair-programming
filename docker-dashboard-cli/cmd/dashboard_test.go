package cmd

import (
	"io"
	"os"
	"strings"
	"testing"
)

func TestDashboard(t *testing.T) {
	// Redirect standard output to capture printed text
	oldStout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	// Call the dashboard function
	Dashboard()

	// Restore standard output and read captured text
	w.Close()
	os.Stdout = oldStout
	capturedOutput, _ := io.ReadAll(r)

	// Define the expected output
	expectedOutput := "Displaying the dashboard..."

	// Check if the captured output contains the expected text
	if !strings.Contains(string(capturedOutput), expectedOutput) {
		t.Errorf("Expected to find %q in output, but did not. Got: %q", expectedOutput, string(capturedOutput))
	}

}