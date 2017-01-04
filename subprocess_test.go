// +build !windows

package gexpect

import (
	"testing"
)

func TestSpawn(t *testing.T) {
	t.Logf("Testing Hello World... ")
	child, err := Spawn("echo \"Hello World\"")
	if err != nil {
		t.Fatal(err)
	}
	err = child.Expect("Hello World")
	if err != nil {
		t.Fatal(err)
	}
}

func TestCommandStart(t *testing.T) {
	t.Logf("Testing Command... ")

	// Doing this allows you to modify the cmd struct prior to execution, for example to add environment variables
	child, err := Command("echo 'Hello World'")
	if err != nil {
		t.Fatal(err)
	}
	child.Start()
	child.Expect("Hello World")
}
