package mdparse



var bold_ansi string = "\033[1m"
var underline_ansi string = "\033[4m"
var reset_ansi string = "\033[m"
var block string = "█"
var working int 
var skip_next int
var new_string string = ""

func Parse(content string) (string){
	new_string = ""
	working = 0
	skip_next = 0

	for i, part := range content {
		if i+1 >= len(content) {
			new_string += string(content[i])
			continue
		}
		if part == '*' && rune(content[i+1]) == part {		
			if skip_next == 1 {
				continue
			}
			Bold()
		} else if part == '*' && rune(content[i+1]) != part {
			if skip_next == 1 {
				continue
			}
			Underline()
		} else if part == '_' && rune(content[i+1]) == part {
			if skip_next == 1 {
				continue
			}
			Bold()
		} else if part == '_' && rune(content[i+1]) != part {
			if skip_next == 1 {
				continue
			}
			Underline()
		} else if part == '>' && rune(content[i+1]) == ' ' {
			new_string += block
		} else if part == '>' && rune(content[i+1]) == part {
			new_string += block + block
		} else {
			new_string += string(part)
		}
	}
	return new_string+reset_ansi
}

func Bold(){
	if working == 0 {
		working = 1
		skip_next = 1
		new_string += bold_ansi
	} else if working == 1 {
		working = 0
		new_string += reset_ansi
		skip_next = 1
	}
}

func Underline() {
	if working == 0 {
		working = 1
		skip_next = 1
		new_string += underline_ansi
	} else if working == 1 {
		working = 0
		new_string += reset_ansi
		skip_next = 1
	}
}
