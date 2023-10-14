// Code generated by templ@(devel) DO NOT EDIT.

package templ

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func Body(cards templ.Component) templ.Component {
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
		_, err = templBuffer.WriteString("<html lang=\"en\"><head><meta charset=\"utf-8\"><title>")
		if err != nil {
			return err
		}
		var_2 := `Radwan's Awesome Blog`
		_, err = templBuffer.WriteString(var_2)
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</title><link rel=\"stylesheet\" href=\"style.css\"></head><body class=\"flex flex-col bg-zinc-900\"><div class=\"h-24\"></div><h1 class=\"text-center font-serif text-6xl text-neutral-200\"><a href=\"\" class=\"bg-gradient-to-r from-emerald-400 to-indigo-500 bg-clip-text font-serif font-bold text-neutral-200 transition duration-200 hover:text-transparent hover:duration-200\">")
		if err != nil {
			return err
		}
		var_3 := `Radwan`
		_, err = templBuffer.WriteString(var_3)
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</a><span class=\"text-neutral-400\">")
		if err != nil {
			return err
		}
		var_4 := `.`
		_, err = templBuffer.WriteString(var_4)
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</span><span class=\"text-indigo-500\">")
		if err != nil {
			return err
		}
		var_5 := `Is`
		_, err = templBuffer.WriteString(var_5)
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</span><span class=\"text-neutral-400\">")
		if err != nil {
			return err
		}
		var_6 := `(`
		_, err = templBuffer.WriteString(var_6)
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</span><em class=\"text-emerald-400\">")
		if err != nil {
			return err
		}
		var_7 := `"writing about computer"`
		_, err = templBuffer.WriteString(var_7)
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</em><span class=\"text-neutral-400\">")
		if err != nil {
			return err
		}
		var_8 := `)`
		_, err = templBuffer.WriteString(var_8)
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</span></h1><div class=\"h-24\"></div><div class=\"ml-24 mr-24 flex flex-wrap justify-center\">")
		if err != nil {
			return err
		}
		err = cards.Render(ctx, templBuffer)
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</div><div class=\"h-32\"></div><footer class=\"bottom-0 w-full text-center\"><p class=\"text-sm text-neutral-400\">")
		if err != nil {
			return err
		}
		var_9 := `author: literally the header`
		_, err = templBuffer.WriteString(var_9)
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</p><p class=\"text-sm text-neutral-400\"><a href=\"mailto:rjabraouti@outlook.com\">")
		if err != nil {
			return err
		}
		var_10 := `rjabraouti@outlook.com`
		_, err = templBuffer.WriteString(var_10)
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</a></p><p class=\"text-sm text-neutral-400\">")
		if err != nil {
			return err
		}
		var_11 := `built with: love, friendship, and cupcakes `
		_, err = templBuffer.WriteString(var_11)
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</p></footer></body></html>")
		if err != nil {
			return err
		}
		if !templIsBuffer {
			_, err = templBuffer.WriteTo(w)
		}
		return err
	})
}
