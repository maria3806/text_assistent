# TEXT-ASSISTANT

 ## A program for working with text files. It allows you to analyze text, shorten it, rewrite it in different styles, find keywords, and manage files. ##

### === Features === ###

Rename files – change the name of a selected text file.

Check file size – display the size of the file in bytes.

Main idea of the text – identify the main point or purpose of the text.

Compress text – shorten text to a specified number of words.

Rewrite text – change the style of the text (formal, friendly, short).

Keywords – find the most important words in the text.

 ### === Structure === ###

text_assistant/
├─ main.go                 # Main program file
├─ go.mod                  # Go module
├─ texts/                  # Folder with text files for processing
└─ processor/              # Helper functions
   ├─ file_manager.go      # Rename, file size, list files
   ├─ compress.go          # Compress text
   ├─ creative.go          # Rewrite text in different styles
   └─ summary.go           # Keywords, word count, short summary



### Installation & Running 

Clone or download the project to any folder.

Create a texts folder and put your text files there.

## Install Go (version ≥ 1.25). ##

Initialize the module:

go mod tidy

## Run the program:

bash
Копировать код
go run main.go

## Follow the console instructions: select a file and a task.


# Requirements 

 Go ≥ 1.25

# texts folder with text files

# Usage

 The program lists all files in the texts folder.

The user inputs the number of the file to process.

The program prompts to select a task (rename, analyze, compress, etc.).

The result is displayed in the console.
