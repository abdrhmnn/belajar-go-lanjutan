package belajargolanjutan

import (
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
