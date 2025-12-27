data = [0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15]

test_x= []
train_x = []

for i in range(2, 8):
    print(i)
    if i % 3 == 2:
        test_x.append(data[i])
    else:
        train_x.append(data[i])

print(train_x, test_x)
