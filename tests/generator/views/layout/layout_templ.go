// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.663
package layout

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func Base(title string) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
		if !templ_7745c5c3_IsBuffer {
			templ_7745c5c3_Buffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<!doctype html><html lang=\"en\" data-theme=\"\"><head><meta charset=\"UTF-8\"><meta name=\"viewport\" content=\"width=device-width, initial-scale=1.0\"><meta name=\"description\" content=\"Go/Echo+Templ: Some Exploration for GO HTMX and Templ\"><meta name=\"google\" content=\"notranslate\"><link rel=\"shortcut icon\" href=\"/img/templ.png\" type=\"image/png\"><link rel=\"stylesheet\" href=\"/css/styles.css\"><title>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var2 string
		templ_7745c5c3_Var2, templ_7745c5c3_Err = templ.JoinStringErrs(title)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `tests/generator/views/layout/layout.templ`, Line: 16, Col: 17}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var2))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</title><script src=\"https://unpkg.com/htmx.org@1.9.9\" integrity=\"sha384-QFjmbokDn2DjBjq+fM+8LUIVrAgqcNW2s0PjAxHETgRn9l4fvX31ZxDxvwQnyMOX\" crossorigin=\"anonymous\"></script><script src=\"https://cdn.jsdelivr.net/npm/sweetalert2@11\"></script><script src=\"https://unpkg.com/hyperscript.org@0.9.12\"></script><!-- Alpine.js --><!-- Alpine Plugins --><script defer src=\"https://unpkg.com/@alpinejs/focus@3.x.x/dist/cdn.min.js\"></script><!-- Alpine Core --><script defer src=\"https://unpkg.com/alpinejs@3.x.x/dist/cdn.min.js\"></script><script>\n                /* Script to toggle light and  mode */\n            </script><script src=\"/js/confirm.js\" defer></script><link rel=\"stylesheet\" href=\"/css/static.css\"></head><body hx-boost=\"true\"><layout class=\"main-content flex flex-row flex-grow p-4\"><div class=\"flex justify-center mx-auto\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templ_7745c5c3_Var1.Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div></layout><script lang=\"js\">\n\t\t\t     document.body.addEventListener('htmx:beforeSwap', function(evt) {\n\t\t\t    if(evt.detail.xhr.status === 404){\n\t\t\t        // alert the user when a 404 occurs (maybe use a nicer mechanism than alert())\n\t\t\t        alert(\"Error: Could Not Find Resource\");\n\t\t\t    } else if(evt.detail.xhr.status === 422){\n\t\t\t        // allow 422 responses to swap as we are using this as a signal that\n\t\t\t        // a form was submitted with bad data and want to rerender with the\n\t\t\t        // errors\n\t\t\t        //\n\t\t\t        // set isError to false to avoid error logging in console\n                    alert(\"Bad Request Debug\")\n\t\t\t        evt.detail.shouldSwap = true;\n\t\t\t        evt.detail.isError = false;\n\t\t\t    } else if(evt.detail.xhr.status === 418){\n\t\t\t        // if the response code 418 (I'm a teapot) is returned, retarget the\n\t\t\t        // content of the response to the element with the id `teapot`s\n\t\t\t        evt.detail.shouldSwap = true;\n\t\t\t        evt.detail.target = htmx.find(\"#teapot\");\n\t\t\t    }\n\t\t\t});\n\t\t\t</script></body></html>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
