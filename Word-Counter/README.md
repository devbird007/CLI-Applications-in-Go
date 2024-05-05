# WORD COUNTER
This is a word counter command-line tool. It receives a body of text and counts the words, lines or bytes present in it.

## Options

- `-l` to count the number of lines

- `-b` to count the number of bytes


## Usage
### Build the tool
```
git clone git@github.com:devbird007/CLI-Applications-in-Go.git

cd Word-Counter

go build
```


### Examples
#### 1. Count Words
```bash
echo "This sentence is 5 words." | ./wc 
```

Output:
```
5
```

#### 2. Count Lines
```bash
echo -e "This sentence\n is 2 lines." | ./wc -l
```
Output:
```
2
```

#### 3. Count Bytes
```bash
echo "This sentence is 27 bytes." | ./wc -b
```
Output:
```
27
```

### Counting from a file
You can also count from a file. Below is a demonstration:
```
cat myfile.txt | ./wc
```
<hr style="border:2px solid black">