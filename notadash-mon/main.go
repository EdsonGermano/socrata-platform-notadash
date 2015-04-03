package main

import (
    "os"
    "fmt"
    "bytes"
    "github.com/codegangsta/cli"
    lib "github.com/boldfield/notadash/lib"
)


func main() {
    app := buildApp()
    app.Run(os.Args)
}


func validateContext(ctx *cli.Context, strings []string) (string, error) {
    var buffer bytes.Buffer
    var err error

    for _, v := range strings {
        if (ctx.String(v) == "") && (ctx.GlobalString(v) == "") {
            buffer.WriteString(v)
            buffer.WriteString(" ")
            err = lib.ErrParameterMissing
        }
    }
     return fmt.Sprintf(buffer.String()), err
}
