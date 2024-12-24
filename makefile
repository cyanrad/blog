####################################### WINDOWS/DEV #######################################
run:
	tailwind -i tailwind.css -o ./content/static/style.css
	nu genposts.nu
	templ generate -path "template"
	go run .

build-win:
	tailwind -i tailwind.css -o ./content/static/style.css
	nu genposts.nu
	templ generate -path "template"
	go build -o blog main.go

build-docker:
	-docker image rm blog:latest
	docker build -t blog . 

####################################### LINUX/PROD #######################################
build: 
	./deps/tailwind -i tailwind.css -o ./content/static/style.css
	sh genposts.sh
	./deps/templ generate -path "template"
	go build -o blog

dependencies:	
# init
	mkdir ./deps
	mkdir ./deps/installs

# pandoc
	curl -L https://github.com/jgm/pandoc/releases/download/3.6/pandoc-3.6-linux-amd64.tar.gz -o ./deps/installs/pandoc.tar
	tar -xzf ./deps/installs/pandoc.tar -C ./deps/installs 
	mv ./deps/installs/pandoc-3.6/bin/pandoc ./deps

# templ
	curl -L https://github.com/a-h/templ/releases/download/v0.2.793/templ_Linux_x86_64.tar.gz -o ./deps/installs/templ.tar
	mkdir ./deps/installs/templ
	tar -xzf ./deps/installs/templ.tar -C ./deps/installs/templ 
	mv ./deps/installs/templ/templ ./deps

# tailwind
	curl -L https://github.com/tailwindlabs/tailwindcss/releases/download/v3.4.17/tailwindcss-linux-x64 -o ./deps/tailwind
	chmod u+x ./deps/tailwind