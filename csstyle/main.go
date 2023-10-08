package main

import (
	"bytes"
	"flag"
	"io"
	"log"
	"os"
)

func initArgs() (string, string) {
	inputFiles := ""
	flag.StringVar(&inputFiles, "i", "", "input HTML file")

	outputFile := ""
	flag.StringVar(&outputFile, "o", "", "output HTML file")

	flag.Parse()

	return inputFiles, outputFile
}

func main() {
	i, o := initArgs()

	inputHTMLFile, err := os.Open(i)
	if err != nil {
		log.Fatal(err)
	}

	inputHTML, err := io.ReadAll(inputHTMLFile)
	if err != nil {
		log.Fatal(err)
	}

	// This is exteremely terrible. There is absolutely a better way to do this. You're welcome
	inputHTML = bytes.ReplaceAll(inputHTML, []byte(`<p>`), []byte(`<p class="mb-4 text-justify text-xl text-neutral-200">`))
	inputHTML = bytes.ReplaceAll(inputHTML, []byte(`<strong>`), []byte(`<strong class="text-rose-400">`))
	inputHTML = bytes.ReplaceAll(inputHTML, []byte(`<em>`), []byte(`<em class="text-rose-400">`))
	inputHTML = bytes.ReplaceAll(inputHTML, []byte(`<ul>`), []byte(`<ul class="mb-4 list-disc pl-4 text-xl text-neutral-200">`))
	inputHTML = bytes.ReplaceAll(inputHTML, []byte(`<ol`), []byte(`<ol class="mb-4 list-decimal pl-4 text-xl text-neutral-200" `))
	inputHTML = bytes.ReplaceAll(inputHTML, []byte(`<code>`), []byte(`<code class="rounded-md bg-zinc-700 p-1 font-mono text-lg text-emerald-300">`))
	inputHTML = bytes.ReplaceAll(inputHTML, []byte(`<h1`), []byte(`<h1 id="a-testing-file" class="mb-4 text-center text-4xl text-emerald-400" `))
	inputHTML = bytes.ReplaceAll(inputHTML, []byte(`<h2`), []byte(`<h2 id="a-testing-file" class="mb-4 text-center text-4xl text-emerald-400" `))
	inputHTML = bytes.ReplaceAll(inputHTML, []byte(`<h3`), []byte(`<h3 id="a-testing-file" class="mb-4 text-2xl text-indigo-300" `))
	inputHTML = bytes.ReplaceAll(inputHTML, []byte(`<a`), []byte(`<a class="hover:grad bg-gradient-to-tr from-emerald-400 to-indigo-400 bg-clip-text font-bold text-transparent hover:from-rose-400 hover:to-indigo-400" `))
	inputHTML = bytes.ReplaceAll(inputHTML, []byte(`<img`), []byte(`<img class="rounded-md" `))
	inputHTML = bytes.ReplaceAll(inputHTML, []byte(`<hr`), []byte(`<hr class="mb-12 opacity-50" `))

	outputHTMLFile, err := os.OpenFile(o, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0222)
	if err != nil {
		log.Fatal(err)
	}

	io.Copy(outputHTMLFile, bytes.NewReader(inputHTML))
}
