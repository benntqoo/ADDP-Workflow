# Test code for agent system testing
import sqlite3
import os

def get_user_data(user_id):
    """Get user data from database - potentially unsafe"""
    conn = sqlite3.connect('users.db')
    cursor = conn.cursor()
    
    # Potential SQL injection vulnerability
    query = f"SELECT * FROM users WHERE id = {user_id}"
    cursor.execute(query)
    
    result = cursor.fetchone()
    conn.close()
    return result

def process_large_file(filename):
    """Process a large file - potentially slow"""
    data = []
    # Inefficient: loading entire file into memory
    with open(filename, 'r') as f:
        lines = f.readlines()
    
    # Inefficient nested loops - O(n²) complexity
    for i in range(len(lines)):
        for j in range(len(lines)):
            if lines[i] == lines[j] and i != j:
                data.append(lines[i])
    
    return data

class UserManager:
    def __init__(self):
        self.users = []
        
    def add_user(self, name, email):
        # Missing input validation
        user = {
            'name': name,
            'email': email,
            'password': '123456'  # Hardcoded password!
        }
        self.users.append(user)
        
    def authenticate(self, email, password):
        # Timing attack vulnerability
        for user in self.users:
            if user['email'] == email:
                if user['password'] == password:
                    return True
        return False