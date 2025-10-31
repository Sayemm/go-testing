package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello()
	want := "Hello World"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

/*
Writing a test is just like writing a function, with a few rules
- It needs to be in a file with a name like xxx_test.go
- The test function must start with the word Test
- The test function takes one argument only t *testing.T
	- t of type *testing.T is our "hook" into the testing framework
	- so we can do things like t.Fail() when we want to fail.

- Errorf method on t, which will print out a message and fail the test.
- The f stands for format, which allows us to build a string with values inserted into the placeholder values %q.
- When we make the test fail, it should be clear how it works.
- For tests, %q is very useful as it wraps your values in double quotes.


- The name of the module, usually refers to a URL where the module can be found and downloaded.
*/
