package post

import "bytes"

// This one of the most terrible functions I've ever wrote
// I thought about making it a map then looping, but then I thought
// ... why, this is probably faster...yeah I know I'm itrating over string
// n times, but who cares
func StyleHTML(inputHTML []byte) []byte {
	outputHTML := make([]byte, len(inputHTML))
	copy(outputHTML, inputHTML)
	outputHTML = bytes.ReplaceAll(outputHTML, []byte(`<p>`), []byte(`<p class="mb-4 text-justify text-lg text-neutral-200 sm:text-left sm:text-xl">`))
	outputHTML = bytes.ReplaceAll(outputHTML, []byte(`<strong>`), []byte(`<strong class="text-indigo-300">`))
	outputHTML = bytes.ReplaceAll(outputHTML, []byte(`<em>`), []byte(`<em class="text-indigo-300">`))
	outputHTML = bytes.ReplaceAll(outputHTML, []byte(`<ul class="incremental">`), []byte(`<ul class="mb-4 list-disc pl-4 text-lg sm:text-xl text-neutral-200 marker:text-indigo-400">`))
	outputHTML = bytes.ReplaceAll(outputHTML, []byte(`<ol`), []byte(`<ol class="mb-4 list-decimal pl-4 text-lg sm:text-xl text-neutral-200 marker:text-indigo-300" `))
	outputHTML = bytes.ReplaceAll(outputHTML, []byte(`<code>`), []byte(`<code class="rounded bg-zinc-800 px-1.5 py-0.5 font-mono text-sm text-emerald-300">`))
	outputHTML = bytes.ReplaceAll(outputHTML, []byte(`<pre class="`), []byte(`<pre class="my-6 overflow-x-auto rounded-md border border-neutral-700 bg-neutral-800 p-4 text-sm leading-relaxed min-[320px]:max-w-[20rem] min-[420px]:max-w-[25rem] min-[500px]:max-w-[30rem] sm:max-w-[35rem] md:max-w-[40rem] lg:max-w-[100%] `))
	outputHTML = bytes.ReplaceAll(outputHTML, []byte(`<h1`), []byte(`<h1 class="mb-4 text-center text-3xl text-emerald-400 sm:text-4xl" `))
	outputHTML = bytes.ReplaceAll(outputHTML, []byte(`<h2`), []byte(`<h2 class="mb-4 text-center text-3xl text-emerald-400 sm:text-4xl" `))
	outputHTML = bytes.ReplaceAll(outputHTML, []byte(`<h3`), []byte(`<h3 class="mb-2 mt-12 text-2xl font-bold text-emerald-400" `))
	outputHTML = bytes.ReplaceAll(outputHTML, []byte(`<a`), []byte(`<a class="hover:grad bg-gradient-to-tr from-emerald-400 to-indigo-400 bg-clip-text font-bold text-transparent hover:from-rose-400 hover:to-indigo-400" `))
	outputHTML = bytes.ReplaceAll(outputHTML, []byte(`<img`), []byte(`<img class="ml-auto mr-auto mt-4 block rounded-md" `))
	outputHTML = bytes.ReplaceAll(outputHTML, []byte(`<hr`), []byte(`<hr class="mb-12 opacity-50" `))

	return outputHTML
}
