run:
	make tailwind
	make posts
	templ generate -path "template"
	go run .

posts:
	nu genposts.nu
	
tailwind:
	tailwind -i tailwind.css -o ./content/static/style.css
