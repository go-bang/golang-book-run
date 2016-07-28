def fibonacci(n):
	if 0 <= n <= 1:
		return n
	else:
		return fibonacci(n-1) + fibonacci(n-2)

def fibolist(n):
	list = []
	for i in range(n):
		list.append(fibonacci(i))
	return list


print(fibolist(40))
