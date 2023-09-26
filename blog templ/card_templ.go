// Code generated by templ@(devel) DO NOT EDIT.

package main

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func Card(id string, title string, imgSrc string, tag string, date string, readingTime string) templ.Component {
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
		_, err = templBuffer.WriteString("<a href=\"")
		if err != nil {
			return err
		}
		var var_2 templ.SafeURL = templ.URL(id)
		_, err = templBuffer.WriteString(templ.EscapeString(string(var_2)))
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("\"><div class=\"group m-4 flex h-[30rem] w-96 flex-col justify-end rounded-md border-2 border-solid border-emerald-400 bg-neutral-800 p-2 shadow-md shadow-emerald-900 transition duration-300 hover:-translate-y-1 hover:scale-105 hover:border-indigo-500 hover:bg-neutral-900 hover:shadow-lg hover:shadow-indigo-900 hover:delay-100 hover:duration-300\"><div class=\"m-2 flex flex-col justify-end overflow-hidden rounded-md\"><img src=\"")
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString(templ.EscapeString(imgSrc))
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("\" class=\"mb-3 w-full rounded-md\"><em class=\"mb-3 font-serif text-2xl text-neutral-200 duration-200 hover:delay-100 group-hover:text-slate-300\">")
		if err != nil {
			return err
		}
		var var_3 string = title
		_, err = templBuffer.WriteString(templ.EscapeString(var_3))
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</em><p class=\"font-serif text-sm text-neutral-400 duration-200 hover:delay-100 group-hover:text-slate-400\">")
		if err != nil {
			return err
		}
		var var_4 string = tag
		_, err = templBuffer.WriteString(templ.EscapeString(var_4))
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString(" ")
		if err != nil {
			return err
		}
		var_5 := `&#x2022; `
		_, err = templBuffer.WriteString(var_5)
		if err != nil {
			return err
		}
		var var_6 string = date
		_, err = templBuffer.WriteString(templ.EscapeString(var_6))
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString(" ")
		if err != nil {
			return err
		}
		var_7 := `&#x2022; `
		_, err = templBuffer.WriteString(var_7)
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("<span class=\"text-emerald-400\">")
		if err != nil {
			return err
		}
		var var_8 string = readingTime
		_, err = templBuffer.WriteString(templ.EscapeString(var_8))
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</span></p></div></div></a>")
		if err != nil {
			return err
		}
		if !templIsBuffer {
			_, err = templBuffer.WriteTo(w)
		}
		return err
	})
}
