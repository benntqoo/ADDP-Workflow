import unittest
import math

def calculate_average(numbers):
    """Calculate average with proper error handling"""
    if not numbers:
        return 0
    return sum(numbers) / len(numbers)


class TestCalculateAverage(unittest.TestCase):
    """Comprehensive unit tests for calculate_average function"""
    
    def test_empty_list(self):
        """Test with empty list returns 0"""
        self.assertEqual(calculate_average([]), 0)
    
    def test_single_element(self):
        """Test with single element returns that element"""
        self.assertEqual(calculate_average([5]), 5)
        self.assertEqual(calculate_average([-10]), -10)
        self.assertEqual(calculate_average([0]), 0)
    
    def test_positive_numbers(self):
        """Test with positive numbers"""
        self.assertEqual(calculate_average([1, 2, 3, 4, 5]), 3)
        self.assertEqual(calculate_average([10, 20, 30]), 20)
    
    def test_negative_numbers(self):
        """Test with negative numbers"""
        self.assertEqual(calculate_average([-1, -2, -3]), -2)
        self.assertEqual(calculate_average([-10, -20, -30]), -20)
    
    def test_mixed_numbers(self):
        """Test with mixed positive and negative numbers"""
        self.assertEqual(calculate_average([-1, 0, 1]), 0)
        self.assertEqual(calculate_average([-10, 10]), 0)
        self.assertEqual(calculate_average([1, -2, 3, -4]), -0.5)
    
    def test_floating_point_numbers(self):
        """Test with floating point numbers"""
        self.assertAlmostEqual(calculate_average([1.5, 2.5, 3.5]), 2.5)
        self.assertAlmostEqual(calculate_average([0.1, 0.2, 0.3]), 0.2)
    
    def test_large_numbers(self):
        """Test with large numbers"""
        large_nums = [10**6, 2*10**6, 3*10**6]
        self.assertEqual(calculate_average(large_nums), 2*10**6)
    
    def test_many_elements(self):
        """Test with many elements"""
        numbers = list(range(1, 101))  # 1 to 100
        self.assertEqual(calculate_average(numbers), 50.5)
    
    def test_precision(self):
        """Test floating point precision"""
        numbers = [1/3, 1/3, 1/3]
        result = calculate_average(numbers)
        self.assertAlmostEqual(result, 1/3, places=10)
    
    def test_special_values(self):
        """Test with special float values"""
        # Test with infinity
        self.assertEqual(calculate_average([float('inf')]), float('inf'))
        self.assertEqual(calculate_average([float('-inf')]), float('-inf'))
        
        # Test with NaN
        result = calculate_average([float('nan')])
        self.assertTrue(math.isnan(result))


if __name__ == '__main__':
    unittest.main()