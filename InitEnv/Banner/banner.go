/***
creat byteslzz time :2023-4-24
*/
package Banner

import (
	"fmt"
	"time"
)

//banner have logo and time（Startup time）
func PrintBanner() {
	fmt.Println("  ____                 _ _        ")
	fmt.Println(" / ___|_ __ __ ___   _(_) |_ __ _ ")
	fmt.Println("| |  _| '__/ _` \\ \\ / / | __/ _` |")
	fmt.Println("| |_| | | | (_| |\\ V /| | || (_| |")
	fmt.Println(" \\____|_|  \\__,_| \\_/ |_|\\__\\__,_|")
	fmt.Printf("  Started at %s\n", time.Now().Format("2006-01-02 15:04:05"))
}
