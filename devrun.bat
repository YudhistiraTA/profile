@echo off
start cmd.exe /K "npx tailwindcss -i ./lib/input.css -o ./static/tailwind.css --watch --minify"
air