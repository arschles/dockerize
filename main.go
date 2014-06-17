package main

import (
  "flag"
  "io/ioutil"
  "log"
)


const DefaultUrl = "https://github.com/arschles/ghdockerize/archive/master.zip"
var url = flag.String("url", DefaultUrl, "the URL of the zip file to dockerize")

func main() {
  flag.Parse()
  zipReader, err := downloadZip(*url)
  if err != nil {
    log.Fatal(err)
  }
  for _, f := range r.File {
    log.Printf("Contents of %s:\n", f.Name)
    rc, err := f.Open()
    if err != nil {
        log.Fatal(err)
    }
    _, err = io.CopyN(os.Stdout, rc, 68)
    if err != nil {
        log.Fatal(err)
    }
    rc.Close()
    fmt.Println()
}

  log.Printf(string(ioutil.ReadAll(*zipReader)))
}
