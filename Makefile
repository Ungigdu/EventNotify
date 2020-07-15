BINDIR=bin

all: a

a:
	gomobile bind -v -o $(BINDIR)/dss.aar -target=android github.com/Ungigdu/EventNotify/android

clean:
	gomobile clean
	rm $(BINDIR)/*