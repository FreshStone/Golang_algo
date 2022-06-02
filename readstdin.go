/*
package main

import (
        "fmt"
        "bufio"
        "strings"
        "strconv"
        "os"
)

func main(){
        var n int
        testarr := make([][]string, n)
        fmt.Scanf("%d", &n)
        reader := bufio.NewReader(os.Stdin)
        for i := 0; i < n; i++{
                testarr = append(testarr, strings.Split(readline(reader), " "))
        }

        for i := 0; i < n; i++{
                l, _ = strings.ParseInt(testarr[i][0], 10, 64)
                r, _ = strings.ParseInt(testarr[i][1], 10, 64)
                fmt.Println(f(l, r))
        }
}

func readline(r *bufio.Reader)string{
        s, _, _ := r.ReadLine()
        return strings.TrimRight(string(s), "\r\n")
}
*/
package main

import (
    "bufio"
    "fmt"
    "os"
    "io"
    "strings"
    "strconv"
)

func main() {
    reader := bufio.NewReader(os.Stdin)
    /*text, _ := reader.ReadString('\n')
    fmt.Println(text)
    */
    //read unknown lines of stdin 
   for {
	   text, _, _ := reader.ReadLine()
	   t := string(text)
	if t == "" {
		break
	}
	   fmt.Printf("%s\n", t)
   }
/*
    fmt.Println("Enter text: ")
    var text2, text3 string
    fmt.Scanf("%s %s", &text2, &text3)
    fmt.Println(text2)
   fmt.Println(text3)
*/

}

/*
func main() {
    reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

    stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer stdout.Close()

    writer := bufio.NewWriterSize(stdout, 1024 * 1024)
    //val, _ := strconv.Atoi(readLine(reader))
    tTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
    checkError(err)
    t := int32(tTemp)

    for tItr := 0; tItr < int(t); tItr++ {
        nm := strings.Split(readLine(reader), " ")
        //val, _ := strconv.Atoi(nm[0])
        nTemp, err := strconv.ParseInt(nm[0], 10, 64)
        checkError(err)
        n := int32(nTemp)

        mTemp, err := strconv.ParseInt(nm[1], 10, 64)
        checkError(err)
        m := int32(mTemp)

        result := solve(n, m)

        fmt.Fprintf(writer, "%d\n", result)
    }

    writer.Flush()
}

func readLine(reader *bufio.Reader) string {
    str, _, err := reader.ReadLine()
    if err == io.EOF {
        return ""
    }

    return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
    if err != nil {
        panic(err)
    }
}
*/
