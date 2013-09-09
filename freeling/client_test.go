/**

goling Go Freeling client

Version: 0.1
Author: Jonathan Hendler
Contact: jonathan dot hendler at gmail dot com
License: MIT
Repository: https://github.com/Hendler/goling

*/


package freeling

import (
    "fmt"
    "testing"
)

func Test() {
    conn := freeling.Init( "localhost" , "50005" )
    fmt.Printf("Connected\n\n")
    status := freeling.Send( conn , "Every good boy deserves fudge.")
    fmt.Printf("\nResults: %s\n", status)
    freeling.Close(conn)
}

