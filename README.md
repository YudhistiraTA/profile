# Description

A dead simple Go web application that serves HTMX files. Currently under construction. I will be using it to document my coding journey. Not yet deployed.

# Tech Stacks

1. [Chi](https://go-chi.io) as router
    - I've chosen to use Chi because it is the closest thing to raw Go stdlib. This way, I can focus on learning Go, instead of learning about another framework.
2. [Templ](https://templ.guide/) as template engine
    - To use Go code directly inside the HTML files. A little bit more convenient than stdlib [html/template](https://pkg.go.dev/html/template) engine
3. [HTMX](https://htmx.org/) as "frontend" framework
    - I'm just tired of React and Typescript.
4. [Tailwind](https://tailwindcss.com/) as CSS framework
    - I just don't want to make .css files
5. [DaisyUI](https://daisyui.com/) as component library
    - I just don't want to write tailwind classes
6. [Air](https://github.com/cosmtrek/air)
    - It's the simplest Go live reload package I've found. Also very customizable.
7. [Goldmark](https://github.com/yuin/goldmark)
    - Most maintained markdown parser with plenty of extensions.
