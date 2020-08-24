# unique-morse-code-words

## 題目解讀：

### 題目來源:
[unique-morse-code-words](https://leetcode.com/problems/unique-morse-code-words/)

### 原文:

International Morse Code defines a standard encoding where each letter is mapped to a series of dots and dashes, as follows: "a" maps to ".-", "b" maps to "-...", "c" maps to "-.-.", and so on.

For convenience, the full table for the 26 letters of the English alphabet is given below:

[".-","-...","-.-.","-..",".","..-.","--.","....","..",".---","-.-",".-..","--","-.","---",".--.","--.-",".-.","...","-","..-","...-",".--","-..-","-.--","--.."]
Now, given a list of words, each word can be written as a concatenation of the Morse code of each letter. For example, "cab" can be written as "-.-..--...", (which is the concatenation "-.-." + ".-" + "-..."). We'll call such a concatenation, the transformation of a word.

Return the number of different transformations among all words we have.
### 解讀：

給定一個字串陣列 把每個字串轉換成 摩斯密碼

然找所有不同的摩斯密碼個數 找出來

## 初步解法:
### 初步觀察:
建立一個 map 存放所有翻譯好的摩斯密碼

最後map的長度就是結果

### 初步設計:
Given a string array words,

Step 0: let valueMap = make([string]int), idx = 0

Step 1: if idx >= len(words) go to step 4

Step 2: first translate words[idx] into morse code string morseStr , set valueMap[morseStr] += 1
Step 3: idx = idx + 1 go to step 1

Step 4: return len(valueMap)

## 遇到的困難
### 題目上理解的問題
因為英文不是筆者母語

所以在題意解讀上 容易被英文用詞解讀給搞模糊

### pseudo code撰寫

一開始不習慣把pseudo code寫下來

因此 不太容易把自己的code做解析

### golang table driven test不熟
對於table driven test還不太熟析

所以對於寫test還是耗費不少時間
## 我的github source code
```golang
package unique_morse_representation

var MORSE_CODE = [...]string{".-", "-...", "-.-.", "-..", ".", "..-.", "--.", "....", "..", ".---", "-.-", ".-..", "--", "-.", "---", ".--.", "--.-", ".-.", "...", "-", "..-", "...-", ".--", "-..-", "-.--", "--.."}

func uniqueMorseRepresentations(words []string) int {
	morseMap := make(map[string]int)
	for _, word := range words {
		morseMap[translateStrToMorse(word)] += 1
	}
	return len(morseMap)
}

func translateStrToMorse(word string) string {
	result := ""
	for _, r := range word {
		result += MORSE_CODE[r-'a']
	}
	return result
}

```
## 測資的撰寫

```golang
package unique_morse_representation

import "testing"

func Test_uniqueMorseRepresentations(t *testing.T) {
	type args struct {
		words []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example1",
			args: args{
				words: []string{"gin", "zen", "gig", "msg"},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := uniqueMorseRepresentations(tt.args.words); got != tt.want {
				t.Errorf("uniqueMorseRepresentations() = %v, want %v", got, tt.want)
			}
		})
	}
}

```
## my self record
[golang leetcode 30 day 26th day](https://hackmd.io/79v0LJC1SpqqvNxeYOM_0Q?view)
## 參考文章

[golang test](https://ithelp.ithome.com.tw/articles/10204692)