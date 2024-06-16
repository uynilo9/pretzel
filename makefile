INPUT = ./cmd/pretzel/pretzel.go
OUTDIR = ./bin
OUTPUT = $(OUTDIR)/pretzel

build:
	go build -o $(OUTPUT) $(INPUT)

clean:
	rm -rf $(OUTDIR)