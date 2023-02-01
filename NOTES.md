Convert input (type string) to integer
``` 
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

Fetch local IP address :
```
package main  
  
import (  
 "fmt"  
 "net"  
)  
  
func GetOutboundIP() net.IP {
    conn, err := net.Dial("udp", "8.8.8.8:80")
    if err != nil {
        log.Fatal(err)
    }
    defer conn.Close()

    localAddr := conn.LocalAddr().(*net.UDPAddr)

    return localAddr.IP
}
```
Find hostname of local machine:

`hostname = os.Hostname`

