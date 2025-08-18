# Fixed code with proper error handling
def calculate_average(numbers):
    """Calculate average with proper error handling"""
    if not numbers:  # Handle empty list
        return 0  # or raise ValueError("Cannot calculate average of empty list")
    
    total = sum(numbers)
    return total / len(numbers)

# Test cases
test_cases = [
    [],                    # Empty list
    [5],                   # Single number
    [1, 2, 3, 4, 5],      # Normal case
    [-1, 0, 1],           # With negative numbers
]

for numbers in test_cases:
    result = calculate_average(numbers)
    print(f"Average of {numbers}: {result}")