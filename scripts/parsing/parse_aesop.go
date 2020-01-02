package main

import (
	"bufio"
	"encoding/xml"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"

	"golang.org/x/text/unicode/norm"

	"./types"
)

const (
	AESOP_PATH  = "src/First1KGreek-Aesop-XML"
	BOOK_FOLDER = "src/First1KGreek-Aesop-RAW"
	BOOK_FILE   = AESOP_PATH + "/tlg002" + "/tlg0096.tlg002.First1K-grc1.xml"
)

func main() {
	xmlFile, err := os.Open(BOOK_FILE)
	check(err)
	// defer the closing of our xmlFile so that we can parse it later on
	defer xmlFile.Close()

	// read our opened xmlFile as a byte array.
	byteValue, _ := ioutil.ReadAll(xmlFile)

	// we initialize our Users array
	var rawFormat types.Book

	// we unmarshal our byteArray which contains our
	// xmlFiles content into 'users' which we defined above
	xml.Unmarshal(byteValue, &rawFormat)
	bookName := rawFormat.TeiHeader.FileDesc.TitleStmt.Title.Text
	log.Println("Book: ", bookName)
	//log.Println("Book: ", bookName, book)
	text := rawFormat.Text.Body.Div

	targetName := "Aesop.txt"
	f, err := os.Create(targetName)
	check(err)

	w := bufio.NewWriter(f)
	wc := norm.NFC.Writer(w)
	defer wc.Close()
	defer f.Close()

	for chapter := 0; chapter < len(text.Div); chapter++ {
		chapterNumber := text.Div[chapter].N
		chapterHead := text.Div[chapter].Head
		sanitizedHead := strings.TrimSpace(strings.Replace(chapterHead, "\n", "", -1))
		textByte := []byte("Aesop" + "." + chapterNumber + ".HEAD " + sanitizedHead + "\n")
		_, errW := wc.Write(textByte)
		check(errW)
		w.Flush()

		for verse := 0; verse < len(text.Div[chapter].P); verse++ {
			verseText := text.Div[chapter].P[verse].Text
			verseNumber := strconv.Itoa(verse + 1)

			sanitizedText := strings.TrimSpace(strings.Replace(verseText, "\n", "", -1))
			textByte := []byte("Aesop" + "." + chapterNumber + "." + verseNumber + " " + sanitizedText + "\n")
			_, errW := wc.Write(textByte)
			check(errW)
			w.Flush()
		}
	}
	log.Printf("Done processing")
}

func check(e error) {
	if e != nil {
		log.Fatalln("ERROR: ", e)
	}
}
