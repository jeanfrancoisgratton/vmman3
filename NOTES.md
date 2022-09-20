https://www.socketloop.com/tutorials/golang-how-to-check-if-input-from-os-args-is-integer

```
// convert input (type string) to integer
 first, err := strconv.ParseInt(os.Args[1], 10, 0)

 if err != nil {
             fmt.Println("First input parameter must be integer")
             os.Exit(1)
  }

 second, err := strconv.ParseInt(os.Args[2], 10, 0)

 if err != nil {
             fmt.Println("Second input parameter must be integer")
             os.Exit(1)
 }
```
