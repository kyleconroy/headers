package headers

import (
	"fmt"
	"strings"
)

type parser struct {
	index  int
	chars  []rune
	r      rune
	err    error
	k, v   strings.Builder
	output map[string]string
	eof    bool
}

// XXX: Add a specific error message type
func ParseDirectives(input string) (map[string]string, error) {
	if input == "" {
		return map[string]string{}, nil
	}
	return newParser(input).parse()
}

func newParser(input string) *parser {
	var chars []rune
	for _, c := range input {
		chars = append(chars, c)
	}
	return &parser{chars: chars, r: chars[0], output: map[string]string{}}
}

func (p *parser) accept(r rune) bool {
	if p.r == r {
		p.next()
		return true
	}
	return false
}

func (p *parser) next() {
	p.index++
	if p.index == len(p.chars) {
		p.eof = true
		return
	}
	p.r = p.chars[p.index]
}

func (p *parser) done() bool {
	return p.eof || p.err != nil
}

func (p *parser) stop(msg string) {
	p.err = fmt.Errorf(msg)
}

func (p *parser) expect(s rune) bool {
	if p.accept(s) {
		return true
	}
	p.stop("expect: unexpected symbol")
	return false
}

func (p *parser) parse() (map[string]string, error) {
	if p.r == 0 {
		return p.output, nil
	}
	for {
		if p.directive(); p.done() {
			return p.output, p.err
		}
		if !p.expect(';') {
			return p.output, p.err
		}
	}
}

func (p *parser) directive() {
	p.lws()
	p.name()
	if p.accept('=') {
		p.value()
	}
	if p.k.Len() > 0 {
		p.output[p.k.String()] = p.v.String()
		p.k.Reset()
		p.v.Reset()
	}
	p.lws()
}

func (p *parser) lws() {
	for {
		if p.r == ' ' || p.r == rune(9) {
			p.next()
			continue
		}
		return
	}
}

func (p *parser) name() {
	p.k.Reset()
	for {
		if p.eof {
			return
		}
		switch p.r {
		case ';', ' ', rune(9), '=':
			return
		case '(', ')', '<', '>', '@', ',', ':', '\\', '/', '[', ']', '?', '{', '}', '"':
			p.stop("illegal char in name")
			return
		default:
			if _, err := p.k.WriteRune(p.r); err != nil {
				panic(err)
			}
			p.next()
		}
	}
}

func (p *parser) value() {
	p.v.Reset()
	quoted := p.accept('"')
	for {
		if p.done() {
			if quoted {
				p.stop("imbalanced quote")
			}
			return
		}
		if quoted {
			if p.accept('"') {
				return
			}
			p.accept('\\')
		} else {
			if p.eof {
				return
			}
			switch p.r {
			case ';', ' ', rune(9):
				return
			case '(', ')', '<', '>', '@', ',', ':', '\\', '/', '[', ']', '?', '{', '}', '=', '"':
				p.stop("illegal char in name")
				return
			}
		}
		if !p.done() {
			if _, err := p.v.WriteRune(p.r); err != nil {
				panic(err)
			}
			p.next()
		}
	}
}
