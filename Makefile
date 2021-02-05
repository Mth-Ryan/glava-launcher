build:
	go build -o bin/glava-launcher

install:
	cp $(CURDIR)/bin/glava-launcher /usr/bin
	cp $(CURDIR)/assets/glava.desktop /usr/share/applications
	cp $(CURDIR)/assets/glava.png /usr/share/icons
