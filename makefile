targets := $(wildcard *.go)

compile: $(targets)
	go install github.com/michaellihs/golab

gendoc: compile
	golab gendoc -p doc
	golab zsh-completion --path zsh/_golab

test: compile
	ginkgo -r -v
