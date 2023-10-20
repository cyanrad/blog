package post

import "bytes"

// This one of the most terrible functions I've ever wrote
func StyleHTML(inputHTML []byte) []byte {
	outputHTML := make([]byte, len(inputHTML))
	copy(outputHTML, inputHTML)
	outputHTML = bytes.ReplaceAll(outputHTML, []byte(`<p>`), []byte(`<p class="mb-4 text-justify text-xl text-neutral-200">`))
	outputHTML = bytes.ReplaceAll(outputHTML, []byte(`<strong>`), []byte(`<strong class="text-indigo-300">`))
	outputHTML = bytes.ReplaceAll(outputHTML, []byte(`<em>`), []byte(`<em class="text-indigo-300">`))
	outputHTML = bytes.ReplaceAll(outputHTML, []byte(`<ul>`), []byte(`<ul class="mb-4 list-disc pl-4 text-xl text-neutral-200">`))
	outputHTML = bytes.ReplaceAll(outputHTML, []byte(`<ol`), []byte(`<ol class="mb-4 list-decimal pl-4 text-xl text-neutral-200" `))
	outputHTML = bytes.ReplaceAll(outputHTML, []byte(`<code>`), []byte(`<code class="rounded-md bg-zinc-700 p-1 font-mono text-lg text-emerald-300">`))
	outputHTML = bytes.ReplaceAll(outputHTML, []byte(`<h1`), []byte(`<h1 id="a-testing-file" class="mb-4 text-center text-4xl text-emerald-400" `))
	outputHTML = bytes.ReplaceAll(outputHTML, []byte(`<h2`), []byte(`<h2 id="a-testing-file" class="mb-4 text-center text-4xl text-emerald-400" `))
	outputHTML = bytes.ReplaceAll(outputHTML, []byte(`<h3`), []byte(`<h3 id="a-testing-file" class="mb-4 text-2xl text-emerald-400" `))
	outputHTML = bytes.ReplaceAll(outputHTML, []byte(`<a`), []byte(`<a class="hover:grad bg-gradient-to-tr from-emerald-400 to-indigo-400 bg-clip-text font-bold text-transparent hover:from-rose-400 hover:to-indigo-400" `))
	outputHTML = bytes.ReplaceAll(outputHTML, []byte(`<img`), []byte(`<img class="rounded-md" `))
	outputHTML = bytes.ReplaceAll(outputHTML, []byte(`<hr`), []byte(`<hr class="mb-12 opacity-50" `))

	return outputHTML
}
