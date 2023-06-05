.PHONY: version

SHELL=/bin/bash

version:
	git branch --show-current > public/version.txt
	git for-each-ref --points-at `git branch --show-current` | head -1 | cut -d' ' -f 1 >> public/version.txt
