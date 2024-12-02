import sys

D = open("input.txt").read().strip()
ans = 0
for line in D.split('\n'):
    digits = []
    for i, c in enumerate(line):
        if c.isdigit():
            digits.append(c)
        for d, val in enumerate(['one', 'two', 'three', 'four', 'five', 'six', 'seven', 'eight', 'nine']):
            if line[i:].startswith(val):
                digits.append(str(d + 1))
    score = int(digits[0]+digits[-1])
    print(score)
    ans += score
print(ans)
