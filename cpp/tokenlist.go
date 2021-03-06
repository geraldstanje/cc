package cpp

import "container/list"

//list of tokens

type tokenList struct {
	l *list.List
}

func newTokenList() *tokenList {
	return &tokenList{list.New()}
}

func (tl *tokenList) copy() *tokenList {
	ret := newTokenList()
	ret.appendList(tl)
	return ret
}

func (tl *tokenList) isEmpty() bool {
	return tl.l.Len() == 0
}

func (tl *tokenList) popFront() *Token {
	if tl.isEmpty() {
		panic("internal error")
	}
	fronte := tl.l.Front()
	ret := fronte.Value.(*Token)
	tl.l.Remove(fronte)
	return ret
}

//Makes a copy of all tokens.
func (tl *tokenList) appendList(toAdd *tokenList) {
	l := toAdd.l
	for e := l.Front(); e != nil; e = e.Next() {
		tl.l.PushBack(e.Value.(*Token).copy())
	}
}

func (tl *tokenList) append(toAdd *Token) {
	tl.l.PushBack(toAdd.copy())
}

func (tl *tokenList) prepend(toAdd *Token) {
	tl.l.PushFront(toAdd.copy())
}

//Makes a copy of all tokens.
func (tl *tokenList) prependList(toAdd *tokenList) {
	l := toAdd.l
	for e := l.Back(); e != nil; e = e.Prev() {
		tl.l.PushFront(e.Value.(*Token).copy())
	}
}

func (tl *tokenList) front() *list.Element {
	return tl.l.Front()
}

func (tl *tokenList) addToHideSets(tok *Token) {
	for e := tl.front(); e != nil; e = e.Next() {
		e.Value.(*Token).hs = e.Value.(*Token).hs.add(tok.Val)
	}
}

func (tl *tokenList) setHideSets(hs *hideset) {
	for e := tl.front(); e != nil; e = e.Next() {
		e.Value.(*Token).hs = hs
	}
}

func (tl *tokenList) setPositions(pos FilePos) {
	for e := tl.front(); e != nil; e = e.Next() {
		e.Value.(*Token).Pos = pos
	}
}

func (tl *tokenList) String() string {
	ret := ""
	for e := tl.front(); e != nil; e = e.Next() {
		t := e.Value.(*Token)
		ret += " " + t.Val
	}
	return ret[1:]
}
