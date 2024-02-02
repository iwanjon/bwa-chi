name ="my name is iwan"
# print(name.title())

# print(name[::-1])

Upppername= name.title()
print(Upppername)
countBlank = 0

for i in Upppername:
    if i == " ":
        countBlank+=1
        
print("white space", count)

count = 0
for i in Upppername:
    if i.isupper():
        count+=1
print("upper", count)

count = 0
for i in Upppername:
    if i.islower():
        count+=1
print("lower", count)

