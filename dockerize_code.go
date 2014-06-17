package main

import (
  //DockerClient "github.com/fsouza/go-dockerclient"
  //"time"
  "bytes"
  //"archive/tar"
  //"fmt"
  //"net/url"
  "net/http"
  "io/ioutil"
  "archive/zip"
)

//var client = DockerClient.NewCLient("unix:///var/run/docker.sock")

//download the zip file at url, unzip it and return the
//io.ReadCloser that has the data
func downloadZip(url string) (*zip.Reader, error) {
  resp, err := http.Get(url)
  if err != nil {
    return nil, err
  }
  zipped, err := ioutil.ReadAll(resp.Body)
  defer resp.Body.Close()
  if err != nil {
    return nil, err
  }

  zippedReader := bytes.NewReader(zipped)
  zippedLen := int64(len(zipped))
  zipReader, err := zip.NewReader(zippedReader, zippedLen)
  if err != nil {
    return nil, err
  }

  return zipReader, nil
}


/*
func dockerfileTxt(folderName string) string {
  return fmt.Sprintf(`
  FROM ubuntu:12.10\n
  ADD %s
  `, folderName)
}

//create an image from a folder, and return
func DockerizeFolder(name string, folderPath string) {
  t := time.Now()
  inputbuf, outputbuf := bytes.NewBuffer(nil), bytes.NewBuffer(nil)
  tr := tar.NewWriter(inputbuf)
  tr.WriteHeader(&tar.Header{Name: "Dockerfile", Size: 10, ModTime: t, AccessTime: t, ChangeTime: t})
  tr.Write([]byte(`
    FROM ubuntu:12.10\n
    ADD
  ))
  tr.Close()
  opts := docker.BuildImageOptions{
    Name: name,
    InputStream: inputbuf,
    OutputStream: outputbuf,
    RmTmpContainer: false,
    NoCache: false,
  }
  if err := client.BuildImage(opts); err != nil {
    log.Fatal(err)
  }
}
*/
