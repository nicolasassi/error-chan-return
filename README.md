# error-chan-return

A simple pattern which I like to use when in need to use multiple **Goroutines** in a Loop and return its errors from **channels** without losing the capability to use **sync.WaitGroup**.

By running the `main.go` file everything will run as expected returning no errors.

By runnning the `test_main.go` there is a small chance that an error will return.

Any comments, issues, thoughts are warmly welcome :)
