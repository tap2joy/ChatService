package utils

import "strings"

type PopularWords struct {
	WordList map[string]int
	TopCount int    // 当前出现最高的词的次数
	TopWord  string // 当前频率最高的词
}

func NewPopularWords() *PopularWords {
	return &PopularWords{
		WordList: make(map[string]int),
	}
}

// 重置
func (p *PopularWords) Reset() {
	p.WordList = make(map[string]int)
	p.TopCount = -1
	p.TopWord = ""
}

// 喂聊天日志
func (p *PopularWords) AddLog(log string) {
	words := strings.Trim(log, " ")
	splitedWords := strings.Split(words, " ")

	for _, v := range splitedWords {
		if _, ok := p.WordList[v]; ok {
			p.WordList[v]++
		} else {
			p.WordList[v] = 1
		}

		if p.WordList[v] > p.TopCount {
			p.TopCount = p.WordList[v]
			p.TopWord = v
		}
	}
}

func (p *PopularWords) GetPopularWord() string {
	return p.TopWord
}
