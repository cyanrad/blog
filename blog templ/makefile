run:
	make tailwind
	templ generate -path "template"
	go run .

posts:
	nu genposts.nu
	
format:
	templ fmt .
	
tailwind:
	tailwind -i tailwind.css -o ./content/static/style.css

.PHONY: posts run
