package headers

import "fmt"

type parser struct {
	index  int
	input  string
	r      int
	err    error
	cname  string
	cvalue string
	output map[string]string
}

const eof = -1

// XXX: Add a specific error message type
func ParseDirectives(input string) (map[string]string, error) {
	p := parser{input: input, r: int(input[0]), output: map[string]string{}}
	return p.parse()
}

func newParser(input string) *parser {
	return &parser{input: input, r: int(input[0]), output: map[string]string{}}
}

func (p *parser) accept(r int) bool {
	if p.r == r {
		p.next()
		return true
	}
	return false
}

func (p *parser) next() {
	p.index++
	if p.index == len(p.input) {
		p.r = eof
		return
	}
	p.r = int(p.input[p.index])
}

func (p *parser) done() bool {
	return p.r == eof || p.err != nil
}

func (p *parser) stop(msg string) {
	p.err = fmt.Errorf(msg)
}

func (p *parser) expect(s int) bool {
	if p.accept(s) {
		return true
	}
	p.stop("expect: unexpected symbol")
	return false
}

func (p *parser) parse() (map[string]string, error) {
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
	if p.cname != "" {
		p.output[p.cname] = p.cvalue
		p.cname = ""
		p.cvalue = ""
	}
	p.lws()
}

func (p *parser) lws() {
	for {
		if p.r == ' ' || p.r == 9 {
			p.next()
			continue
		}
		return
	}
}

func (p *parser) name() {
	p.cname = ""
	for {
		switch p.r {
		case eof, ';', ' ', 9, '=':
			return
		case '(', ')', '<', '>', '@', ',', ':', '\\', '/', '[', ']', '?', '{', '}', '"':
			p.stop("illegal char in name")
			return
		default:
			p.cname += string(p.r)
			p.next()
		}
	}
}

func (p *parser) value() {
	p.cvalue = ""
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
			switch p.r {
			case eof, ';', ' ', 9:
				return
			case '(', ')', '<', '>', '@', ',', ':', '\\', '/', '[', ']', '?', '{', '}', '=', '"':
				p.stop("illegal char in name")
				return
			}
		}
		if !p.done() {
			p.cvalue += string(p.r)
			p.next()
		}
	}
}
