def process_large_file_optimized(filename):
    """Optimized version - O(n) complexity with streaming"""
    from collections import Counter
    
    # Stream processing - don't load entire file
    line_counts = Counter()
    
    # First pass: count occurrences
    with open(filename, 'r') as f:
        for line in f:
            line_counts[line.rstrip()] += 1
    
    # Collect duplicates (lines that appear more than once)
    duplicates = []
    for line, count in line_counts.items():
        if count > 1:
            # Add the line 'count' times (matching original behavior)
            duplicates.extend([line] * (count * (count - 1)))
    
    return duplicates

# Alternative: If you just need unique duplicates
def find_duplicate_lines(filename):
    """Find lines that appear more than once - O(n) with minimal memory"""
    seen_once = set()
    duplicates = set()
    
    with open(filename, 'r') as f:
        for line in f:
            line = line.rstrip()
            if line in seen_once:
                duplicates.add(line)
            else:
                seen_once.add(line)
    
    return list(duplicates)