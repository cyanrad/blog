// Code generated by templ@(devel) DO NOT EDIT.

package template

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

import "strconv"

func Post(body templ.Component, title string, hook string, tag string, date string, readingTime int) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) (err error) {
		templBuffer, templIsBuffer := w.(*bytes.Buffer)
		if !templIsBuffer {
			templBuffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templBuffer)
		}
		ctx = templ.InitializeContext(ctx)
		var_1 := templ.GetChildren(ctx)
		if var_1 == nil {
			var_1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, err = templBuffer.WriteString("<head><link rel=\"stylesheet\" href=\"/static/style.css\"><link rel=\"stylesheet\" href=\"/static/atom-one-dark-reasonable.min.css\"><script src=\"/static/highlight.min.js\">")
		if err != nil {
			return err
		}
		var_2 := ``
		_, err = templBuffer.WriteString(var_2)
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</script><meta name=\"viewport\" content=\"width=device-width, initial-scale=1.0\"></head><body class=\"mt-24 flex justify-center bg-zinc-900 font-serif\"><script>")
		if err != nil {
			return err
		}
		var_3 := `
            hljs.highlightAll();
        `
		_, err = templBuffer.WriteString(var_3)
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</script><a class=\"absolute left-8 top-8 w-16\" href=\"/\"><img src=\"/static/r-back-gradient.png\" alt=\"go back\" class=\"absolute\"><img src=\"/static/r-back.png\" alt=\"go back gradient\" class=\"absolute duration-300 hover:opacity-0\"></a><div class=\"mt-12 w-[64rem] pl-8 pr-8 sm:pl-16 sm:pr-16 sm:text-justify md:mt-0 md:pl-32 md:pr-32\"><h1 class=\"mb-6 text-5xl font-semibold text-neutral-200 sm:text-6xl\">")
		if err != nil {
			return err
		}
		var var_4 string = title
		_, err = templBuffer.WriteString(templ.EscapeString(var_4))
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</h1><p class=\"text-md text-neutral-400 sm:text-justify sm:text-lg\">")
		if err != nil {
			return err
		}
		var var_5 string = tag
		_, err = templBuffer.WriteString(templ.EscapeString(var_5))
		if err != nil {
			return err
		}
		var_6 := `&#x2022; `
		_, err = templBuffer.WriteString(var_6)
		if err != nil {
			return err
		}
		var var_7 string = date
		_, err = templBuffer.WriteString(templ.EscapeString(var_7))
		if err != nil {
			return err
		}
		var_8 := `&#x2022; `
		_, err = templBuffer.WriteString(var_8)
		if err != nil {
			return err
		}
		var var_9 string = strconv.Itoa(readingTime) + "  Minutes"
		_, err = templBuffer.WriteString(templ.EscapeString(var_9))
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</p><br><em class=\"text-justify text-xl text-neutral-400 sm:text-2xl\">")
		if err != nil {
			return err
		}
		var var_10 string = hook
		_, err = templBuffer.WriteString(templ.EscapeString(var_10))
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</em><hr class=\"mb-12 opacity-50\">")
		if err != nil {
			return err
		}
		err = body.Render(ctx, templBuffer)
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</div></body>")
		if err != nil {
			return err
		}
		if !templIsBuffer {
			_, err = templBuffer.WriteTo(w)
		}
		return err
	})
}
