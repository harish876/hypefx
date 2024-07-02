// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.707
package layout

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func Base() templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<!doctype html><html lang=\"en\" data-theme=\"\"><head><meta charset=\"UTF-8\"><meta name=\"viewport\" content=\"width=device-width, initial-scale=1.0\"><meta name=\"description\" content=\"Hypefx D\"><meta name=\"google\" content=\"notranslate\"><link rel=\"stylesheet\" href=\"../assets/css/styles.css\"><script src=\"https://unpkg.com/htmx.org@1.9.9\" integrity=\"sha384-QFjmbokDn2DjBjq+fM+8LUIVrAgqcNW2s0PjAxHETgRn9l4fvX31ZxDxvwQnyMOX\" crossorigin=\"anonymous\"></script><script src=\"https://cdn.jsdelivr.net/npm/sweetalert2@11\"></script><script src=\"https://unpkg.com/hyperscript.org@0.9.12\"></script></head><body hx-boost=\"true\"><div><div class=\"drawer bg-base-100 lg:drawer-open\"><input id=\"drawer\" type=\"checkbox\" class=\"drawer-toggle\"><div class=\"drawer-content\"><div class=\"sticky top-0 z-30 flex h-16 w-full justify-center bg-base-100 bg-opacity-90 text-base-content backdrop-blur transition-shadow duration-100 [transform:translate3d(0,0,0)]\"><nav class=\"navbar w-full\"><div class=\"flex flex-1 md:gap-1 lg:gap-2\"><span class=\"tooltip tooltip-bottom before:text-xs before:content-[attr(data-tip)]\" data-tip=\"Menu\"><label aria-label=\"Open menu\" for=\"drawer\" class=\"btn btn-square btn-ghost drawer-button lg:hidden\"><svg width=\"20\" height=\"20\" xmlns=\"http://www.w3.org/2000/svg\" fill=\"none\" viewBox=\"0 0 24 24\" class=\"inline-block h-5 w-5 stroke-current md:h-6 md:w-6\"><path stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" d=\"M4 6h16M4 12h16M4 18h16\"></path></svg></label></span><div class=\"flex items-center gap-2 lg:hidden\"><a data-sveltekit-preload-data href=\"/\" aria-current=\"page\" aria-label=\"daisyUI\" class=\"flex-0 btn btn-ghost gap-1 px-2 md:gap-2\"><svg class=\"h-6 w-6 md:h-8 md:w-8\" width=\"32\" height=\"32\" viewBox=\"0 0 415 415\" xmlns=\"http://www.w3.org/2000/svg\"><rect x=\"82.5\" y=\"290\" width=\"250\" height=\"125\" rx=\"62.5\" fill=\"#1AD1A5\"></rect> <circle cx=\"207.5\" cy=\"135\" r=\"130\" fill=\"black\" fill-opacity=\".3\"></circle> <circle cx=\"207.5\" cy=\"135\" r=\"125\" fill=\"white\"></circle> <circle cx=\"207.5\" cy=\"135\" r=\"56\" fill=\"#FF9903\"></circle></svg> <span class=\"font-title text-lg text-base-content md:text-2xl\">daisyUI</span></a></div><div class=\"hidden w-full max-w-sm lg:flex\"></div></div><div class=\"flex-0\"><div class=\"hidden flex-none items-center lg:block\"><a data-sveltekit-preload-data href=\"/components/\" class=\"btn btn-ghost drawer-button font-normal\">Components</a></div><div class=\"hidden flex-none items-center lg:block\"><a data-sveltekit-preload-data href=\"/store/\" class=\"btn btn-ghost drawer-button font-normal\">Store</a></div><div class=\"btn btn-ghost cursor-wait\"><svg width=\"20\" height=\"20\" xmlns=\"http://www.w3.org/2000/svg\" fill=\"none\" viewBox=\"0 0 24 24\" class=\"h-5 w-5 stroke-current md:hidden\"><path stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" d=\"M7 21a4 4 0 01-4-4V5a2 2 0 012-2h4a2 2 0 012 2v12a4 4 0 01-4 4zm0 0h12a2 2 0 002-2v-4a2 2 0 00-2-2h-2.343M11 7.343l1.657-1.657a2 2 0 012.828 0l2.829 2.829a2 2 0 010 2.828l-8.486 8.485M7 17h.01\"></path></svg> <span class=\"hidden font-normal md:inline\">Theme</span> <svg width=\"12px\" height=\"12px\" class=\"hidden h-2 w-2 fill-current opacity-60 sm:inline-block\" xmlns=\"http://www.w3.org/2000/svg\" viewBox=\"0 0 2048 2048\"><path d=\"M1799 349l242 241-1017 1017L7 590l242-241 775 775 775-775z\"></path></svg></div><div class=\"btn btn-ghost cursor-wait\"><svg xmlns=\"http://www.w3.org/2000/svg\" viewBox=\"0 0 16 16\" fill=\"currentColor\" class=\"h-4 w-4\"><path fill-rule=\"evenodd\" d=\"M11 5a.75.75 0 0 1 .688.452l3.25 7.5a.75.75 0 1 1-1.376.596L12.89 12H9.109l-.67 1.548a.75.75 0 1 1-1.377-.596l3.25-7.5A.75.75 0 0 1 11 5Zm-1.24 5.5h2.48L11 7.636 9.76 10.5ZM5 1a.75.75 0 0 1 .75.75v1.261a25.27 25.27 0 0 1 2.598.211.75.75 0 1 1-.2 1.487c-.22-.03-.44-.056-.662-.08A12.939 12.939 0 0 1 5.92 8.058c.237.304.488.595.752.873a.75.75 0 0 1-1.086 1.035A13.075 13.075 0 0 1 5 9.307a13.068 13.068 0 0 1-2.841 2.546.75.75 0 0 1-.827-1.252A11.566 11.566 0 0 0 4.08 8.057a12.991 12.991 0 0 1-.554-.938.75.75 0 1 1 1.323-.707c.049.09.099.181.15.271.388-.68.708-1.405.952-2.164a23.941 23.941 0 0 0-4.1.19.75.75 0 0 1-.2-1.487c.853-.114 1.72-.185 2.598-.211V1.75A.75.75 0 0 1 5 1Z\" clip-rule=\"evenodd\"></path></svg> <svg width=\"12px\" height=\"12px\" class=\"hidden h-2 w-2 fill-current opacity-60 sm:inline-block\" xmlns=\"http://www.w3.org/2000/svg\" viewBox=\"0 0 2048 2048\"><path d=\"M1799 349l242 241-1017 1017L7 590l242-241 775 775 775-775z\"></path></svg></div></div></nav></div><div class=\"max-w-[100vw] px-6 pb-16\"><div id=\"content\" class=\"flex flex-col-reverse justify-between gap-6 xl:flex-row\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templ_7745c5c3_Var1.Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div></div><div class=\"toast toast-center z-10 [@supports(color:oklch(0%_0_0))]:hidden\"><div class=\"alert alert-warning grid-cols-[auto] py-2 text-xs\"><span><a class=\"link\" rel=\"nofollow, noreferrer\" target=\"_blank\" href=\"https://www.wikihow.com/Update-Your-Browser\">Please upgrade your browser</a></span></div></div></div><div class=\"drawer-side z-40\" style=\"scroll-behavior: smooth; scroll-padding-top: 5rem;\"><label for=\"drawer\" class=\"drawer-overlay\" aria-label=\"Close menu\"></label><aside class=\"min-h-screen w-80 bg-base-100\"><div data-sveltekit-preload-data class=\"sticky top-0 z-20 hidden items-center gap-2 bg-base-100 bg-opacity-90 px-4 py-2 backdrop-blur lg:flex\"><a href=\"/\" aria-current=\"page\" aria-label=\"Homepage\" class=\"flex-0 btn btn-ghost px-2\"><svg width=\"32\" height=\"32\" viewBox=\"0 0 415 415\" xmlns=\"http://www.w3.org/2000/svg\"><rect x=\"82.5\" y=\"290\" width=\"250\" height=\"125\" rx=\"62.5\" fill=\"#1AD1A5\"></rect> <circle cx=\"207.5\" cy=\"135\" r=\"130\" fill=\"black\" fill-opacity=\".3\"></circle> <circle cx=\"207.5\" cy=\"135\" r=\"125\" fill=\"white\"></circle> <circle cx=\"207.5\" cy=\"135\" r=\"56\" fill=\"#FF9903\"></circle></svg><div class=\"font-title inline-flex text-lg md:text-2xl\">Hypefx</div></a></div><ul class=\"menu px-4 py-0\"><li><details id=\"disclosure-docs\"><summary class=\"group\"><span><svg width=\"18\" height=\"18\" viewBox=\"0 0 48 48\" class=\"h-5 w-5 text-orange-400\" fill=\"none\" xmlns=\"http://www.w3.org/2000/svg\"><path d=\"M5 7H16C20.4183 7 24 10.5817 24 15V42C24 38.6863 21.3137 36 18 36H5V7Z\" fill=\"none\" stroke=\"currentColor\" stroke-width=\"4\" stroke-linejoin=\"bevel\"></path> <path d=\"M43 7H32C27.5817 7 24 10.5817 24 15V42C24 38.6863 26.6863 36 30 36H43V7Z\" fill=\"none\" stroke=\"currentColor\" stroke-width=\"4\" stroke-linejoin=\"bevel\"></path></svg></span> Docs</summary><ul><li hx-get=\"installation.html\" hx-target=\"#content\" class=\"group\"><span>Install</span></li><li><a href=\"/docs/use/\" class=\"group\"><span>Use</span></a></li><li><a href=\"/docs/customize/\" class=\"group\"><span>Customize components</span></a></li><li><a href=\"/docs/config/\" class=\"group\"><span>Config</span></a></li><li><a href=\"/docs/colors/\" class=\"group\"><span>Colors</span></a></li><li><a href=\"/docs/themes/\" class=\"group\"><span>Themes</span></a></li><li><a href=\"/docs/utilities/\" class=\"group\"><span>Utilities</span></a></li><li><a href=\"/docs/layout-and-typography/\" class=\"group\"><span>Layout & Typography</span></a></li></ul></details></li><li></li><li><a href=\"/store/\" class=\"group from-primary from-[-200%] to-primary/0 to-60% [background-image:linear-gradient(-35deg,var(--tw-gradient-stops))] hover:to-primary/10\"><span class=\"transition-colors group-hover:text-primary\"><svg class=\"h-5 w-5\" width=\"18\" height=\"18\" viewBox=\"0 0 48 48\" fill=\"none\" xmlns=\"http://www.w3.org/2000/svg\"><path d=\"M40.0391 22V42H8.03906V22\" stroke=\"currentColor\" stroke-width=\"4\" stroke-linecap=\"butt\" stroke-linejoin=\"bevel\"></path> <path d=\"M5.84231 13.7766C4.31276 17.7377 7.26307 22 11.5092 22C14.8229 22 17.5276 19.3137 17.5276 16C17.5276 19.3137 20.2139 22 23.5276 22H24.546C27.8597 22 30.546 19.3137 30.546 16C30.546 19.3137 33.2518 22 36.5655 22C40.8139 22 43.767 17.7352 42.2362 13.7723L39.2337 6H8.84523L5.84231 13.7766Z\" fill=\"none\" stroke=\"currentColor\" stroke-width=\"4\" stroke-linejoin=\"bevel\"></path></svg></span> <span>Store</span> <span class=\"badge badge-sm border-transparent bg-primary/10 font-mono font-sans font-bold uppercase text-[inherit] text-opacity-70\">new</span></a></li><li><a href=\"/blog/\" class=\"group\"><span><svg xmlns=\"http://www.w3.org/2000/svg\" viewBox=\"0 0 20 20\" fill=\"currentColor\" class=\"h-5 w-5\"><path d=\"M3.75 3a.75.75 0 00-.75.75v.5c0 .414.336.75.75.75H4c6.075 0 11 4.925 11 11v.25c0 .414.336.75.75.75h.5a.75.75 0 00.75-.75V16C17 8.82 11.18 3 4 3h-.25z\"></path> <path d=\"M3 8.75A.75.75 0 013.75 8H4a8 8 0 018 8v.25a.75.75 0 01-.75.75h-.5a.75.75 0 01-.75-.75V16a6 6 0 00-6-6h-.25A.75.75 0 013 9.25v-.5zM7 15a2 2 0 11-4 0 2 2 0 014 0z\"></path></svg></span> <span>Blog</span></a></li><li><a href=\"/resources/videos/\" class=\"group\"><span><svg width=\"18\" height=\"18\" viewBox=\"0 0 48 48\" class=\"h-5 w-5\" fill=\"none\" xmlns=\"http://www.w3.org/2000/svg\"><path d=\"M24 44C35.0457 44 44 35.0457 44 24C44 12.9543 35.0457 4 24 4C12.9543 4 4 12.9543 4 24C4 35.0457 12.9543 44 24 44Z\" fill=\"none\" stroke=\"currentColor\" stroke-width=\"4\" stroke-linejoin=\"bevel\"></path> <path d=\"M20 24V17.0718L26 20.5359L32 24L26 27.4641L20 30.9282V24Z\" fill=\"none\" stroke=\"currentColor\" stroke-width=\"4\" stroke-linejoin=\"bevel\"></path></svg></span> <span>Resources</span></a></li><li></li><li><a href=\"/tailwindplay/\" target=\"_blank\" rel=\"noopener noreferrer\" class=\"group\"><span><svg xmlns=\"http://www.w3.org/2000/svg\" class=\"h-5 w-5\" fill=\"none\" viewBox=\"0 0 54 33\"><g clip-path=\"url(#prefix__clip0)\"><path fill=\"currentColor\" fill-rule=\"evenodd\" d=\"M27 0c-7.2 0-11.7 3.6-13.5 10.8 2.7-3.6 5.85-4.95 9.45-4.05 2.054.513 3.522 2.004 5.147 3.653C30.744 13.09 33.808 16.2 40.5 16.2c7.2 0 11.7-3.6 13.5-10.8-2.7 3.6-5.85 4.95-9.45 4.05-2.054-.513-3.522-2.004-5.147-3.653C36.756 3.11 33.692 0 27 0zM13.5 16.2C6.3 16.2 1.8 19.8 0 27c2.7-3.6 5.85-4.95 9.45-4.05 2.054.514 3.522 2.004 5.147 3.653C17.244 29.29 20.308 32.4 27 32.4c7.2 0 11.7-3.6 13.5-10.8-2.7 3.6-5.85 4.95-9.45 4.05-2.054-.513-3.522-2.004-5.147-3.653C23.256 19.31 20.192 16.2 13.5 16.2z\" clip-rule=\"evenodd\"></path></g> <defs><clipPath id=\"prefix__clip0\"><path fill=\"#fff\" d=\"M0 0h54v32.4H0z\"></path></clipPath></defs></svg></span> <span>Playground</span> <svg width=\"12\" height=\"12\" class=\"opacity-0 transition-opacity duration-300 ease-out group-hover:opacity-100\" viewBox=\"0 0 48 48\" fill=\"none\" xmlns=\"http://www.w3.org/2000/svg\"><path d=\"M19 11H37V29\" stroke=\"currentColor\" stroke-width=\"4\" stroke-linecap=\"butt\" stroke-linejoin=\"bevel\"></path> <path d=\"M11.5439 36.4559L36.9997 11\" stroke=\"currentColor\" stroke-width=\"4\" stroke-linecap=\"butt\" stroke-linejoin=\"bevel\"></path></svg></a></li><li><a href=\"https://github.com/saadeghi/daisyui\" target=\"_blank\" rel=\"noopener noreferrer\" class=\"group\"><span><svg width=\"18\" height=\"18\" class=\"h-5 w-5\" viewBox=\"0 0 48 48\" fill=\"none\" xmlns=\"http://www.w3.org/2000/svg\"><path d=\"M29.3444 30.4765C31.7481 29.977 33.9292 29.1108 35.6247 27.8391C38.5202 25.6676 40 22.3136 40 18.9999C40 16.6752 39.1187 14.505 37.5929 12.6668C36.7427 11.6425 39.2295 3.99989 37.02 5.02919C34.8105 6.05848 31.5708 8.33679 29.8726 7.83398C28.0545 7.29565 26.0733 6.99989 24 6.99989C22.1992 6.99989 20.4679 7.22301 18.8526 7.6344C16.5046 8.23237 14.2591 5.99989 12 5.02919C9.74086 4.05848 10.9736 11.9632 10.3026 12.7944C8.84119 14.6051 8 16.7288 8 18.9999C8 22.3136 9.79086 25.6676 12.6863 27.8391C14.6151 29.2857 17.034 30.2076 19.7401 30.6619\" stroke=\"currentColor\" stroke-width=\"4\" stroke-linecap=\"butt\"></path> <path d=\"M19.7397 30.6619C18.5812 31.937 18.002 33.1478 18.002 34.2944C18.002 35.441 18.002 38.3464 18.002 43.0106\" stroke=\"currentColor\" stroke-width=\"4\" stroke-linecap=\"butt\"></path> <path d=\"M29.3446 30.4766C30.4423 31.9174 30.9912 33.211 30.9912 34.3576C30.9912 35.5042 30.9912 38.3885 30.9912 43.0107\" stroke=\"currentColor\" stroke-width=\"4\" stroke-linecap=\"butt\"></path> <path d=\"M6 31.2155C6.89887 31.3254 7.56554 31.7387 8 32.4554C8.65169 33.5303 11.0742 37.518 13.8251 37.518C15.6591 37.518 17.0515 37.518 18.0024 37.518\" stroke=\"currentColor\" stroke-width=\"4\" stroke-linecap=\"butt\"></path></svg></span> <span>GitHub</span> <svg width=\"12\" height=\"12\" class=\"opacity-0 transition-opacity duration-300 ease-out group-hover:opacity-100\" viewBox=\"0 0 48 48\" fill=\"none\" xmlns=\"http://www.w3.org/2000/svg\"><path d=\"M19 11H37V29\" stroke=\"currentColor\" stroke-width=\"4\" stroke-linecap=\"butt\" stroke-linejoin=\"bevel\"></path> <path d=\"M11.5439 36.4559L36.9997 11\" stroke=\"currentColor\" stroke-width=\"4\" stroke-linecap=\"butt\" stroke-linejoin=\"bevel\"></path></svg></a></li><li><a href=\"/discord/\" target=\"_blank\" rel=\"noopener noreferrer\" class=\"group\"><span><svg class=\"h-5 w-5\" x=\"0px\" y=\"0px\" width=\"16px\" height=\"16px\" viewBox=\"0 0 24 24\"><path fill=\"currentColor\" d=\"M 8.2363281 3.3886719 C 7.9184915 3.3507541 7.5922955 3.389759 7.2851562 3.5078125 C 6.3662291 3.8593556 5.168019 4.408137 4.1210938 5.2539062 L 4.09375 5.2753906 L 4.0683594 5.296875 C 3.6038679 5.7274657 3.3488069 6.2439566 3.0234375 6.9511719 C 2.6980681 7.6583872 2.3700247 8.5311318 2.0722656 9.5136719 C 1.4767474 11.478752 1 13.870256 1 16.164062 C 1 16.404094 1.06291 16.653616 1.1875 16.873047 C 1.9936894 18.296146 3.3875707 19.089172 4.6289062 19.564453 C 5.8601415 20.035867 6.9037775 20.195689 7.5 20.216797 C 7.51943 20.217478 7.5686445 20.224609 7.5703125 20.224609 C 7.9778976 20.224609 8.507273 20.074519 8.8046875 19.511719 L 9.7265625 17.767578 C 11.253027 18.034862 12.698188 18.031089 14.265625 17.751953 L 15.193359 19.507812 C 15.4917 20.076956 16.025986 20.222656 16.429688 20.222656 C 16.431188 20.222656 16.477491 20.217397 16.494141 20.216797 C 17.090411 20.196247 18.137319 20.035024 19.371094 19.5625 C 20.610912 19.087661 22.002426 18.296509 22.808594 16.876953 C 22.937245 16.653708 23 16.40296 23 16.164062 C 23 13.870256 22.522369 11.477589 21.923828 9.5097656 C 21.624558 8.5258538 21.295636 7.6497223 20.966797 6.9394531 C 20.637958 6.2291839 20.373325 5.7078286 19.908203 5.2773438 L 19.882812 5.2558594 L 19.857422 5.234375 C 18.821451 4.3968794 17.628961 3.8543294 16.712891 3.5058594 C 16.100244 3.2723752 15.409182 3.355343 14.867188 3.7285156 C 14.448635 4.0163534 14.306763 4.5143523 14.197266 5 L 9.8027344 5 C 9.6928169 4.5147837 9.5518548 4.0163226 9.1347656 3.7285156 C 8.8636418 3.5414302 8.5541647 3.4265897 8.2363281 3.3886719 z M 8 5.375 C 8 6.2610628 8.7389372 7 9.625 7 L 14.373047 7 C 15.257904 7 15.998928 6.2596472 16 5.375 L 16.001953 5.375 C 16.788319 5.6741308 17.79628 6.1579204 18.5625 6.7675781 C 18.58015 6.7895071 18.877477 7.1898224 19.150391 7.7792969 C 19.429239 8.3815902 19.733786 9.1844586 20.009766 10.091797 C 20.546111 11.855141 20.948842 14.026311 20.972656 15.992188 C 20.490724 16.743136 19.60263 17.332859 18.65625 17.695312 C 17.804404 18.021562 17.097701 18.106186 16.736328 18.146484 L 16.273438 17.269531 C 16.616654 17.168826 16.957602 17.069914 17.320312 16.947266 A 1.0001 1.0001 0 1 0 16.679688 15.052734 C 12.849602 16.34789 10.965085 16.349225 7.3339844 15.058594 A 1.0001 1.0001 0 0 0 6.9414062 14.996094 A 1.0001 1.0001 0 0 0 6.6660156 16.941406 C 7.0298926 17.070746 7.3750462 17.176206 7.7207031 17.28125 L 7.2636719 18.146484 C 6.9025364 18.106234 6.1955688 18.021456 5.34375 17.695312 C 4.3973996 17.332975 3.5097566 16.745657 3.0273438 15.994141 C 3.0509129 14.027751 3.4522238 11.856178 3.9863281 10.09375 C 4.261069 9.1871651 4.5637131 8.3873003 4.8398438 7.7871094 C 5.1096316 7.2007049 5.4070693 6.8013834 5.4160156 6.7871094 C 6.1966698 6.1664912 7.2108742 5.6768865 8 5.375 z M 9.0390625 9.9960938 C 8.2100625 9.9960938 7.5390625 10.889188 7.5390625 11.992188 C 7.5390625 13.095188 8.2100625 13.988281 9.0390625 13.988281 C 9.8680625 13.988281 10.539062 13.095188 10.539062 11.992188 C 10.521062 10.889187 9.8710625 9.8860938 9.0390625 9.9960938 z M 14.996094 10.011719 A 1.5 2 0 0 0 14.996094 14.011719 A 1.5 2 0 0 0 14.996094 10.011719 z\"></path></svg></span> <span>Discord</span> <svg width=\"12\" height=\"12\" class=\"opacity-0 transition-opacity duration-300 ease-out group-hover:opacity-100\" viewBox=\"0 0 48 48\" fill=\"none\" xmlns=\"http://www.w3.org/2000/svg\"><path d=\"M19 11H37V29\" stroke=\"currentColor\" stroke-width=\"4\" stroke-linecap=\"butt\" stroke-linejoin=\"bevel\"></path> <path d=\"M11.5439 36.4559L36.9997 11\" stroke=\"currentColor\" stroke-width=\"4\" stroke-linecap=\"butt\" stroke-linejoin=\"bevel\"></path></svg></a></li><li><a href=\"https://opencollective.com/daisyui\" target=\"_blank\" rel=\"noopener noreferrer\" class=\"group\"><span><svg width=\"18\" class=\"h-5 w-5\" height=\"18\" viewBox=\"0 0 48 48\" fill=\"none\" xmlns=\"http://www.w3.org/2000/svg\"><path d=\"M15 8C8.92487 8 4 12.9249 4 19C4 30 17 40 24 42.3262C31 40 44 30 44 19C44 12.9249 39.0751 8 33 8C29.2797 8 25.9907 9.8469 24 12.6738C22.0093 9.8469 18.7203 8 15 8Z\" fill=\"none\" stroke=\"currentColor\" stroke-width=\"4\" stroke-linecap=\"butt\" stroke-linejoin=\"bevel\"></path></svg></span> <span>Support daisyUI</span> <svg width=\"12\" height=\"12\" class=\"opacity-0 transition-opacity duration-300 ease-out group-hover:opacity-100\" viewBox=\"0 0 48 48\" fill=\"none\" xmlns=\"http://www.w3.org/2000/svg\"><path d=\"M19 11H37V29\" stroke=\"currentColor\" stroke-width=\"4\" stroke-linecap=\"butt\" stroke-linejoin=\"bevel\"></path> <path d=\"M11.5439 36.4559L36.9997 11\" stroke=\"currentColor\" stroke-width=\"4\" stroke-linecap=\"butt\" stroke-linejoin=\"bevel\"></path></svg></a></li></ul><div class=\"pointer-events-none sticky bottom-0 flex h-40 bg-base-100 [mask-image:linear-gradient(transparent,#000000)]\"></div></aside></div></div></div><script lang=\"js\">\n\t\t\t     document.body.addEventListener('htmx:beforeSwap', function(evt) {\n\t\t\t    if(evt.detail.xhr.status === 404){\n\t\t\t        // alert the user when a 404 occurs (maybe use a nicer mechanism than alert())\n\t\t\t        alert(\"Error: Could Not Find Resource\");\n\t\t\t    } else if(evt.detail.xhr.status === 422){\n\t\t\t        // allow 422 responses to swap as we are using this as a signal that\n\t\t\t        // a form was submitted with bad data and want to rerender with the\n\t\t\t        // errors\n\t\t\t        //\n\t\t\t        // set isError to false to avoid error logging in console\n                    alert(\"Bad Request Debug\")\n\t\t\t        evt.detail.shouldSwap = true;\n\t\t\t        evt.detail.isError = false;\n\t\t\t    } else if(evt.detail.xhr.status === 418){\n\t\t\t        // if the response code 418 (I'm a teapot) is returned, retarget the\n\t\t\t        // content of the response to the element with the id `teapot`s\n\t\t\t        evt.detail.shouldSwap = true;\n\t\t\t        evt.detail.target = htmx.find(\"#teapot\");\n\t\t\t    }\n\t\t\t});\n\t\t\t</script></body></html>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
