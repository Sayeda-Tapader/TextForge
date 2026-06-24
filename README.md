# Go-Reloaded

---

This is the first project in the 01-Founders' Full-stack course. Using functions created from the piscine repository, I used them to make a simple text completion/editing/auto-correction tools.


x x  x



# Usage Example

---

$ cat sample.txt

it (cap) was the best of times, it was the worst of times (up) , it was the age of wisdom, it was the age of foolishness (cap, 6) , it was the epoch of belief, it was the epoch of incredulity, it was the season of Light, it was the season of darkness, it was the spring of hope, IT WAS THE (low, 3) winter of despair.

$ go run . sample.txt result.txt

$ cat result.txt

It was the best of times, it was the worst of TIMES, it was the age of wisdom, It Was The Age Of Foolishness, it was the epoch of belief, it was the epoch of incredulity, it was the season of Light, it was the season of darkness, it was the spring of hope, it was the winter of despair.

$ cat sample.txt

Simply add 42 (hex) and 10 (bin) and you will see the result is 68.

$ go run . sample.txt result.txt

$ cat result.txt

Simply add 66 and 2 and you will see the result is 68.

$ cat sample.txt

There is no greater agony than bearing a untold story inside you.

$ go run . sample.txt result.txt

$ cat result.txt

There is no greater agony than bearing an untold story inside you.

$ cat sample.txt

Punctuation tests are ... kinda boring ,what do you think ?

$ go run . sample.txt result.txt

$ cat result.txt


# Meta

---

Sayeda Tapader - githubToAdd - Sayeda.Tapader@outlook.com

https://learn.01founders.co/git/stapader/Go-Reloaded.git

https://stackoverflow.com/questions/37471740/how-to-copy-commits-from-one-git-repo-to-another
-- how to copy commits from one repo to another


https://github.com/01-edu/public/tree/master/subjects/go-reloaded/audit
audit questions