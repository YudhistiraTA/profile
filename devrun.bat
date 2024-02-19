@echo off
start cmd.exe /K "npx tailwindcss -i ./view/input.css -o ./service/static/tailwind.css --watch --minify"
air