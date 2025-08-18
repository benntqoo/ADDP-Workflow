# Code with an error
def calculate_average(numbers):
    total = sum(numbers)
    return total / len(numbers)

# This will crash
result = calculate_average([])
print(f"Average: {result}")