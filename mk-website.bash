#!/bin/bash

TITLE="cli: tools for a more consistant command line interface"

START=$(pwd)
cd "$(dirname "$0")"

function checkApp() {
	APP_NAME="$(which "$1")"
	if [ "$APP_NAME" = "" ] && [ ! -f "./bin/$1" ]; then
		echo "Missing $APP_NAME"
		exit 1
	fi
}

function softwareCheck() {
	for APP_NAME in "$@"; do
		checkApp "$APP_NAME"
	done
}

function MakePage() {
	nav="$1"
	content="$2"
	html="$3"
	echo "Rendering $html"
	mkpage \
		"title=text:${TITLE}" \
		"nav=$nav" \
		"content=$content" \
		"sitebuilt=text:Updated $(date)" \
		"copyright=copyright.md" \
		page.tmpl >"$html"
}

echo "Checking necessary software is installed"
softwareCheck mkpage
echo "Generating website index.html"
MakePage nav.md README.md index.html
echo "Generating install.html"
MakePage nav.md INSTALL.md install.html
echo "Generating license.html"
MakePage nav.md "markdown:$(cat LICENSE)" license.html
echo "Generating docs/index.html"
MakePage docs/nav.md docs/index.md docs/index.html
echo "Generating docs/pkgassets.html"
MakePage docs/nav.md docs/pkgassets.md docs/pkgassets.html
echo "Generating docs/cligenerate.html"
MakePage docs/nav.md docs/cligenerate.md docs/cligenerate.html
echo "Generating examples/help.html"
MakePage nav.md "examples/help.md" "examples/help.html"
echo "Generating examples/htdocs.html"
MakePage nav.md "examples/htdocs.md" "examples/htdocs.html"

cd "$START"
