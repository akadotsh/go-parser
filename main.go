package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

func main() {
	
	
	scanner:= bufio.NewScanner(os.Stdin);
	
	 scanner.Scan();
	 
	str:= scanner.Text();

	tokenizer(str)
}


type Token  struct {
	Type  string
	value string
}
 
func tokenizer(str string) {
 

current := 0;

fmt.Println("String",str)

tokens:= make([]Token,0);
fmt.Println(" len(str)", len(str))
regex:= regexp.MustCompile(`/[\d\w]/`);
whiteSpaceRegex:= regexp.MustCompile(`/\s/`)
for current < len(str) {	 
	curChar := rune(str[current])
	fmt.Printf("curChar: %c\n", curChar)
     
	if curChar == '{'{
		fmt.Println("BraceOpen")
		 	obj:= Token{Type:"BraceOpen",value:string(curChar)}
          	tokens= append(tokens, obj)
	}else if curChar == '}'{
		fmt.Println("BraceClose")
		obj:= Token{Type:"BraceClose",value:string(curChar)}
		tokens= append(tokens, obj)
	}else if curChar == '['{
	   obj:= Token{Type:"BracketOpen",value:string(curChar)};
	   tokens = append(tokens, obj)
	}else if curChar == ']'{
		obj:= Token{Type:"BracketClose",value:string(curChar)};
		tokens = append(tokens, obj)
	}else if curChar == ':' {
		obj:= Token{Type:"Colon",value:string(curChar)};
		tokens = append(tokens, obj)
	}else if curChar == ',' {
		obj:= Token{Type:"Comma",value:string(curChar)};
		tokens = append(tokens, obj)
	}else if curChar == '"' {
		var value = "";
		 current++;
		 curChar := rune(str[current]);

		 for curChar != '"' {
			value += string(curChar);
			current++;
			curChar = rune(str[current]);
		 }

		obj:= Token{Type:"String",value:value};
		tokens = append(tokens, obj)

	}else if regex.MatchString(string(curChar)) {
		
	}else if  whiteSpaceRegex.MatchString(string(curChar)) {
		fmt.Println("whiteSpace")
            current++
	}
	current++
	
}

fmt.Println("tokens",tokens)
 
 
//  parser(tokens)
}


// func parser(tokens []Token) {

// 	if len(tokens) == 0 {
// 		panic("Nothing to parse")
// 	}


// 	var  current=0

  

// 	// advance:= func () Token   {
// 	// 	current ++;
// 	// 	return tokens[current];
// 	// }


// 	parseValue:= func () Token   {
// 		var token = tokens[current];

// 		switch token.Type {
// 			case "BraceOpen":
// 				fmt.Println("BraceOpen",token.value)
// 				return Token{Type:"BraceOpen",value:token.value}
// 			case "BraceClose":
// 				fmt.Println("BraceClose",token.value)
// 				return Token{Type:"BraceClose",value:token.value}
// 			case "BracketOpen":
// 				fmt.Println("BracketOpen",token.value)
// 				return Token{Type:"BracketOpen",value:token.value}
// 			case "BracketClose":
// 				fmt.Println("BracketClose:value",token.value)
// 				return Token{Type:"BracketClose",value:token.value}
// 			case "Colon":
// 				fmt.Println("Colon:value",token.value)
// 				return Token{Type:"Colon",value:token.value}
// 			case "Comma":
// 				fmt.Println("Comma",token.value)
// 				return Token{Type:"Comma",value:token.value}
// 			case "String":
// 				fmt.Println("String",token.value)
// 				return Token{Type:"String",value:token.value}
// 			default:
// 				fmt.Println("default")
// 				return Token{Type:"",value:""}
// 		}
// 	}

// 	 AST:= parseValue()

// 	fmt.Println("AST",AST) 
// }


 

