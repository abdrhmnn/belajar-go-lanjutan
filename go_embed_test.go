package belajargolanjutan

import (
	"embed"
	_ "embed"
	"fmt"
	"io/fs"
	"os"
	"testing"
)

//go:embed file_embed.txt
var fileEmbed string

//go:embed test_img.png
var gambarEmbed []byte

func TestEmbed(t *testing.T) {
	fmt.Println(fileEmbed)
}

func TestEmbedGambar(t *testing.T) {
	err := os.WriteFile("test_img_new.png", gambarEmbed, fs.ModePerm)
	if err != nil {
		panic(err)
	}
	fmt.Println("Berhasil embed gambar!")
}

// embed multiple files

//go:embed files/a.txt
//go:embed files/b.txt
var files embed.FS

func TestMultipleEmbed(t *testing.T) {
	a, _ := files.ReadFile("files/a.txt")
	fmt.Println(string(a))
	b, _ := files.ReadFile("files/b.txt")
	fmt.Println(string(b))
}

// path matcher

//go:embed files/*.txt
var path embed.FS

func TestPathMatcher(t *testing.T) {
	dirEntries, _ := path.ReadDir("files")
	for _, entry := range dirEntries {
		if !entry.IsDir() {
			fmt.Println(entry.Name())
			content, _ := path.ReadFile("files/" + entry.Name())
			fmt.Println(string(content))
		}
	}
}
