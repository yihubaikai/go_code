package main

import (
        "fmt"

        "github.com/wangbin/jiebago"
)

var seg jiebago.Segmenter

func init() {
        seg.LoadDictionary("dict.txt")
}

func print(ch <-chan string) {
        for word := range ch {
                fmt.Printf(" %s /", word)
        }
        fmt.Println()
}

//https://github.com/wangbin/jiebago
func main() {
        text := "君生我未生,我生君已老,恨不生同时,日日与君好"
        fmt.Print("【完全分析模式】：")
        print(seg.CutAll(text))

        fmt.Print("【精确分析模式】：")
        print(seg.Cut(text, false))

        fmt.Print("【新词识别模式】：")
        print(seg.Cut(text, true))

        fmt.Print("【搜索引擎模式】：")
        print(seg.CutForSearch(text, true))
}
