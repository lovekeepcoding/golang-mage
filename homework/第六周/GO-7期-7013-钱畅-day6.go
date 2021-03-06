package main

import (
	"fmt"
	"io"

	//"io/ioutil"
	//"log"
	"bufio"
	"compress/zlib"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

// var (
// 	loc *time.Location
// )

// func init() {
// 	loc, _ = time.LoadLocation("Asia/Shanghai")
// }

func printtime() {
	TIME_FMT := "2006-01-02 15:04:05"
	loc, _ := time.LoadLocation("Asia/Shanghai")
	thetime, _ := time.ParseInLocation(TIME_FMT, "1998-10-01 08:10:00", loc)
	//thetime := time.Now()
	//ts := thetime.Format(TIME_FMT)
	fmt.Println(thetime)
	fmt.Println(thetime.Day())
	timestring := strconv.Itoa(thetime.Year()) + strconv.Itoa(int(thetime.Month())) + strconv.Itoa(thetime.Day()) + strconv.Itoa(thetime.Hour()) + strconv.Itoa(thetime.Minute())
	//fmt.Println(int(thetime.Month()))
	fmt.Println(timestring)
}

func classday() {
	nowday := time.Now()
	//nowday := time.Now().Day()
	//fmt.Println(int(nowday.Weekday()))
	count := 0
	for count < 4 {
		if nowday.Weekday() == 6 {
			fmt.Println(nowday.Month(), nowday.Day())
			count += 1
		}
		nowday = nowday.Add(time.Hour * 24)
	}

}

// func compact(root, string) ([]string, error) {
// 	var Matches []string
// 	err := filepath.Walk(root,func(path string, info os.FileInfo, err error) error)
// 	files, err := ioutil.ReadDir(".")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	for _, f := range files {
// 		fmt.Println(f.Name())
// 	}
// }

var Matches []string
var Lines []string

func WalkMatch(root, pattern string) ([]string, error) {
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		if matched, err := filepath.Match(pattern, filepath.Base(path)); err != nil {
			return err
		} else if matched {
			Matches = append(Matches, path)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	fmt.Println(Matches)
	return Matches, nil
}

//usage:   files, err := WalkMatch("/root/", "*.md")

func read_file(files []string) {
	for _, f := range files {
		if fin, err := os.Open(f); err != nil {
			fmt.Printf("open file faied: %v\n", err) //?????????????????????
		} else {
			defer fin.Close() //???????????????????????????

			//????????????????????????bufio.Reader
			fin.Seek(0, 0) //?????????????????????
			reader := bufio.NewReader(fin)
			for { //????????????
				if line, err := reader.ReadString('\n'); err != nil { //???????????????
					if err == io.EOF {
						if len(line) > 0 { //??????????????????????????????????????????????????????????????????line???
							//fmt.Println(line)
						}
						break //?????????????????????
					} else {
						fmt.Printf("read file failed: %v\n", err)
					}
				} else {
					line = strings.TrimRight(line, "\n") //line??????????????????????????????????????????
					//fmt.Println(line)
					Lines = append(Lines, line)
				}
			}
		}
	}
}

func write_file(filename string) {
	//OpenFile()???Open()???????????????????????????os.O_WRONLY?????????????????????????????????os.O_TRUNC???????????????????????????????????????os.O_CREATE????????????????????????????????????0666???????????????????????????
	if fout, err := os.OpenFile(filename, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666); err != nil {
		fmt.Printf("open file faied: %v\n", err)
	} else {
		defer fout.Close() //???????????????????????????

		//???????????????????????????
		writer := bufio.NewWriter(fout)
		for _, l := range Lines {
			writer.WriteString(l)
			writer.WriteString("\n") //???????????????????????????
		}
		writer.Flush() //buffer????????????????????????????????????????????????????????????????????????Flush?????????????????????????????????????????????
	}
}

func create_file(outfile string) {
	//os.Remove(outfile) //????????????????????????Remove???????????????error
	if file, err := os.Create(outfile); err != nil {
		fmt.Printf("create file faied: %v\n", err)
	} else {
		file.Chmod(0666)                 //??????????????????
		fmt.Printf("fd=%d\n", file.Fd()) //?????????????????????file descriptor?????????????????????
		info, _ := file.Stat()
		fmt.Printf("is dir %t\n", info.IsDir())
		fmt.Printf("modify time %s\n", info.ModTime())
		fmt.Printf("mode %v\n", info.Mode()) //-rw-rw-rw-
		fmt.Printf("file name %s\n", info.Name())
		fmt.Printf("size %d\n", info.Size())
	}
}

func compress(filename string) {
	fin, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		return
	}
	stat, _ := fin.Stat()
	fmt.Printf("????????????????????? %dB\n", stat.Size())

	fout, err := os.OpenFile(filename+".zlib", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}

	bs := make([]byte, 1024)
	writer := zlib.NewWriter(fout) //????????????
	for {
		n, err := fin.Read(bs)
		if err != nil {
			if err == io.EOF {
				break
			} else {
				fmt.Println(err)
			}
		} else {
			writer.Write(bs[:n])
		}
	}
	writer.Close()
	fout.Close()
	fin.Close()

	fin, err = os.Open(filename + ".zlib")
	if err != nil {
		fmt.Println(err)
		return
	}
	stat, _ = fin.Stat()
	fmt.Printf("????????????????????? %dB\n", stat.Size())

	reader, err := zlib.NewReader(fin) //??????
	io.Copy(os.Stdout, reader)         //????????????????????????????????????
	reader.Close()
	fin.Close()
}

func main() {
	printtime()
	classday()
	//compact()
	WalkMatch("./", "*.txt")
	//fmt.Println(Matches)
	//compact(Matches)
	read_file(Matches)
	filename := strconv.FormatInt(time.Now().Unix(), 10) + ".txt"
	create_file(filename)
	write_file(filename)
	compress(filename)
}
