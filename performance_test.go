package main

import (
	"regexp"
	"strings"
	"testing"
)

// read the write-up: https://abanoubhanna.com/posts/regexp-vs-string-manipulation/

// Original implementation using regexp
func generateCommandRegexp(template, pkgName string) string {
	re := regexp.MustCompile(`\bx\b`)
	return re.ReplaceAllStringFunc(template, func(s string) string {
		return pkgName
	})
}

// Pre-compiled regexp implementation (fairer comparison)
var cmdRe = regexp.MustCompile(`\bx\b`)

func generateCommandRegexpCompiled(template, pkgName string) string {
	return cmdRe.ReplaceAllStringFunc(template, func(s string) string {
		return pkgName
	})
}

// using string manipulation
func generateCommandString(template, pkgName string) string {
	cmdStr := template
	// if template ends with ".x" or " x" remove x and add pkgName
	if strings.HasSuffix(template, ".x") || strings.HasSuffix(template, " x") {
		cmdStr = strings.TrimSuffix(template, "x") + pkgName
	}
	return cmdStr
}

func BenchmarkRegexpReplacement(b *testing.B) {
	template := "apt install -y x"
	pkgName := "vim"
	for b.Loop() {
		generateCommandRegexp(template, pkgName)
	}
}

func BenchmarkRegexpPrecompiledReplacement(b *testing.B) {
	template := "apt install -y x"
	pkgName := "vim"
	for b.Loop() {
		generateCommandRegexpCompiled(template, pkgName)
	}
}

func BenchmarkStringReplacement(b *testing.B) {
	template := "apt install -y x"
	pkgName := "vim"
	for b.Loop() {
		generateCommandString(template, pkgName)
	}
}
