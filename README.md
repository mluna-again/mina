# Mina

have you ever wished for a slower and with 0.1% of the functionality, alternative to fzf? no? oh ok.

i have.

## Usage
```sh
$ mina -help
Usage of mina:
  -height int
    	height, if 0 or empty it takes the full screen
  -icon string
    	prompt icon (default "\uf002")
  -mode string
    	modes available: [prompt, fzf, confirm, menu] (default "fzf")
  -nth string
    	display specific columns. eg: -nth 1 displays only the second column, -nth 0,3 displays 1st, 2nd and 3rd column.
  -sep string
    	separator used with -nth (default " ")
  -title string
    	prompt title (default "Mina")
  -width int
    	width, if 0 or empty it takes the full screen
```


## FZF MODE
```sh
$ ls | mina -mode fzf -height 10 -title "FZF MODE"
```
<img width="1906" height="652" alt="fzf" src="https://github.com/user-attachments/assets/377cdd69-3fad-495c-80b3-01e33d54145b" />


## PROMPT MODE
```sh
$ mina -mode prompt -title "PROMPT MODE"
```
<img width="1906" height="652" alt="prompt" src="https://github.com/user-attachments/assets/ebee1e3d-0314-4f81-a3f3-45cd7f8e684a" />


## MENU MODE
```sh
$ echo -e "Run something@r\nTest something@t" | mina -mode menu -title "MENU MODE" -sep @
```
<img width="1906" height="652" alt="menu" src="https://github.com/user-attachments/assets/82df894f-7900-4c41-b31b-76a71bf9d046" />
