---
name: python-context-detector
model: haiku
description: "Intelligent Python context detection for ML/AI, Django, FastAPI, Flask, Data Science, and more"
trigger: "*.py, *.ipynb, requirements.txt, pyproject.toml"
tools: Read, Grep
---

# Python Context Detector - Python 上下文智能檢測器

I analyze Python code context to determine the exact type of Python project and route to the appropriate specialist.

## Detection Strategy

### 1. Priority-Based Detection

```yaml
detection_priority:
  1_machine_learning:
    confidence: HIGH
    indicators:
      - "import torch"
      - "import tensorflow"
      - "import sklearn"
      - "from transformers import"
      - "import keras"
      - "model.fit("
      - "model.train("
      - "*.ipynb files"
    agent: python-ml-specialist
    
  2_llm_development:
    confidence: HIGH
    indicators:
      - "import openai"
      - "from langchain"
      - "from llama_index"
      - "import anthropic"
      - "from transformers import AutoModel"
      - "ChatCompletion"
      - "embedding"
    agent: llm-development-expert
    
  3_data_science:
    confidence: HIGH
    indicators:
      - "import pandas as pd"
      - "import numpy as np"
      - "import matplotlib"
      - "import seaborn"
      - "import plotly"
      - "DataFrame"
      - "*.ipynb with data analysis"
    agent: data-scientist
    
  4_django_web:
    confidence: HIGH
    indicators:
      - "from django."
      - "django.db.models"
      - "settings.py"
      - "manage.py"
      - "urls.py"
      - "views.py"
      - "models.py"
      - "INSTALLED_APPS"
    agent: django-developer
    
  5_fastapi_backend:
    confidence: HIGH
    indicators:
      - "from fastapi import"
      - "FastAPI()"
      - "@app.get"
      - "@app.post"
      - "from pydantic import"
      - "BaseModel"
      - "uvicorn.run"
    agent: fastapi-developer
    
  6_flask_web:
    confidence: MEDIUM
    indicators:
      - "from flask import"
      - "Flask(__name__)"
      - "@app.route"
      - "render_template"
      - "request.form"
    agent: flask-developer
    
  7_streamlit_app:
    confidence: HIGH
    indicators:
      - "import streamlit as st"
      - "st.write"
      - "st.sidebar"
      - "st.columns"
      - "st.button"
    agent: streamlit-developer
    
  8_gradio_app:
    confidence: HIGH
    indicators:
      - "import gradio as gr"
      - "gr.Interface"
      - "gr.Blocks"
      - "demo.launch()"
    agent: gradio-developer
    
  9_pytest_testing:
    confidence: MEDIUM
    indicators:
      - "import pytest"
      - "def test_"
      - "@pytest.fixture"
      - "@pytest.mark"
      - "conftest.py"
    agent: python-test-engineer
    
  10_airflow_pipeline:
    confidence: MEDIUM
    indicators:
      - "from airflow import"
      - "DAG("
      - "@task"
      - "PythonOperator"
      - "BashOperator"
    agent: airflow-data-engineer
    
  11_scrapy_crawler:
    confidence: MEDIUM
    indicators:
      - "import scrapy"
      - "class.*Spider"
      - "def parse("
      - "scrapy.Request"
      - "yield {"
    agent: web-scraping-expert
    
  12_async_backend:
    confidence: MEDIUM
    indicators:
      - "import asyncio"
      - "async def"
      - "await "
      - "aiohttp"
      - "asyncpg"
    agent: async-python-developer
    
  13_computer_vision:
    confidence: MEDIUM
    indicators:
      - "import cv2"
      - "from PIL import"
      - "import torchvision"
      - "detectron2"
      - "image processing"
    agent: computer-vision-expert
    
  14_nlp_project:
    confidence: MEDIUM
    indicators:
      - "import spacy"
      - "import nltk"
      - "from transformers import pipeline"
      - "tokenizer"
      - "word2vec"
    agent: nlp-specialist
    
  15_game_development:
    confidence: LOW
    indicators:
      - "import pygame"
      - "import arcade"
      - "game loop"
      - "sprite"
      - "collision"
    agent: pygame-developer
    
  16_general_python:
    confidence: FALLBACK
    indicators:
      - "None of the above matched"
    agent: python-general-expert
```

### 2. Deep Context Analysis

```python
class PythonContextAnalyzer:
    def analyze_project(self, file_path: str) -> ProjectContext:
        detections = [
            self.detect_by_imports(file_path),
            self.detect_by_requirements(file_path),
            self.detect_by_project_structure(file_path),
            self.detect_by_file_patterns(file_path),
            self.detect_by_pyproject_toml(file_path),
            self.detect_by_notebooks(file_path)
        ]
        
        primary_context = max(detections, key=lambda d: d.confidence)
        
        return ProjectContext(
            primary_type=primary_context.type,
            confidence=primary_context.confidence,
            secondary_types=self.extract_secondary_types(detections),
            package_manager=self.detect_package_manager(),
            python_version=self.detect_python_version(),
            frameworks=self.detect_frameworks(detections)
        )
    
    def detect_by_imports(self, file_path: str) -> Detection:
        imports = self.extract_imports(file_path)
        
        # ML/AI Detection - Highest Priority
        ml_libraries = {
            'torch': ('pytorch', 0.95),
            'tensorflow': ('tensorflow', 0.95),
            'keras': ('keras', 0.9),
            'sklearn': ('scikit-learn', 0.9),
            'transformers': ('huggingface', 0.95),
            'jax': ('jax', 0.9),
            'lightning': ('pytorch-lightning', 0.9)
        }
        
        for lib, (framework, confidence) in ml_libraries.items():
            if any(lib in imp for imp in imports):
                return Detection(ProjectType.MACHINE_LEARNING, confidence)
        
        # LLM/RAG Detection
        llm_libraries = ['langchain', 'llama_index', 'openai', 
                        'anthropic', 'chromadb', 'pinecone']
        if any(lib in imp for lib in llm_libraries for imp in imports):
            return Detection(ProjectType.LLM_APPLICATION, 0.9)
        
        # Data Science Detection
        data_libraries = ['pandas', 'numpy', 'matplotlib', 
                         'seaborn', 'plotly', 'scipy']
        data_count = sum(1 for lib in data_libraries 
                         if any(lib in imp for imp in imports))
        if data_count >= 2:
            return Detection(ProjectType.DATA_SCIENCE, 0.85)
        
        # Web Framework Detection
        if any('django' in imp for imp in imports):
            return Detection(ProjectType.DJANGO, 0.95)
        
        if any('fastapi' in imp for imp in imports):
            return Detection(ProjectType.FASTAPI, 0.95)
        
        if any('flask' in imp for imp in imports):
            return Detection(ProjectType.FLASK, 0.9)
        
        # Specialized Detection
        if any('streamlit' in imp for imp in imports):
            return Detection(ProjectType.STREAMLIT, 0.9)
        
        if any('gradio' in imp for imp in imports):
            return Detection(ProjectType.GRADIO, 0.9)
        
        return Detection(ProjectType.GENERAL, 0.2)
    
    def detect_by_requirements(self, project_path: str) -> Detection:
        """Analyze requirements.txt or pyproject.toml"""
        req_file = os.path.join(project_path, 'requirements.txt')
        if not os.path.exists(req_file):
            return Detection(ProjectType.UNKNOWN, 0)
        
        with open(req_file, 'r') as f:
            requirements = f.read().lower()
        
        # ML/AI packages
        if any(pkg in requirements for pkg in 
               ['torch', 'tensorflow', 'transformers', 'scikit-learn']):
            return Detection(ProjectType.MACHINE_LEARNING, 0.9)
        
        # Web frameworks
        if 'django' in requirements:
            return Detection(ProjectType.DJANGO, 0.9)
        if 'fastapi' in requirements:
            return Detection(ProjectType.FASTAPI, 0.9)
        if 'flask' in requirements:
            return Detection(ProjectType.FLASK, 0.85)
        
        return Detection(ProjectType.GENERAL, 0.3)
```

### 3. Mixed Context Handling

```python
class MixedPythonContextHandler:
    def handle_mixed_context(self, contexts: List[ProjectType]) -> AgentSelection:
        # ML + Web API (MLOps pattern)
        if ProjectType.MACHINE_LEARNING in contexts and \
           ProjectType.FASTAPI in contexts:
            return AgentSelection(
                primary='mlops-engineer',
                support=['python-ml-specialist', 'fastapi-developer']
            )
        
        # Data Science + Streamlit (Data App)
        if ProjectType.DATA_SCIENCE in contexts and \
           ProjectType.STREAMLIT in contexts:
            return AgentSelection(
                primary='data-app-developer',
                support=['data-scientist', 'streamlit-developer']
            )
        
        # Django + React (Full-stack)
        if ProjectType.DJANGO in contexts and \
           'react' in self.detect_frontend():
            return AgentSelection(
                primary='fullstack-python-developer',
                support=['django-developer', 'react-developer']
            )
        
        # LLM + FastAPI (AI Service)
        if ProjectType.LLM_APPLICATION in contexts and \
           ProjectType.FASTAPI in contexts:
            return AgentSelection(
                primary='ai-service-architect',
                support=['llm-specialist', 'fastapi-developer']
            )
        
        # Jupyter + Multiple ML frameworks
        if self.is_notebook() and \
           len([c for c in contexts if 'ML' in str(c)]) > 1:
            return AgentSelection(
                primary='ml-researcher',
                support=['python-ml-specialist', 'data-scientist']
            )
        
        return self.select_best_match(contexts)
```

## Smart Detection Examples

### Example 1: PyTorch Deep Learning
```python
# File: model.py
import torch
import torch.nn as nn
from torch.utils.data import DataLoader
from transformers import AutoModel

class CustomModel(nn.Module):
    def __init__(self):
        super().__init__()
        # DETECTED: Machine Learning / Deep Learning Project
        # AGENT: python-ml-specialist
        # CONFIDENCE: 95%
```

### Example 2: FastAPI Backend
```python
# File: main.py
from fastapi import FastAPI, HTTPException
from pydantic import BaseModel
from typing import List

app = FastAPI()

@app.post("/users/")
async def create_user(user: UserCreate):
    # DETECTED: FastAPI Backend Service
    # AGENT: fastapi-developer
    # CONFIDENCE: 95%
```

### Example 3: Django Web Application
```python
# File: views.py
from django.shortcuts import render
from django.views.generic import ListView
from .models import Product

class ProductListView(ListView):
    model = Product
    template_name = 'products/list.html'
    # DETECTED: Django Web Application
    # AGENT: django-developer
    # CONFIDENCE: 95%
```

### Example 4: Data Science Notebook
```python
# File: analysis.ipynb
import pandas as pd
import numpy as np
import matplotlib.pyplot as plt
import seaborn as sns

df = pd.read_csv('data.csv')
df.describe()
# DETECTED: Data Science / Analysis Project
# AGENT: data-scientist
# CONFIDENCE: 90%
```

### Example 5: LangChain LLM Application
```python
# File: chatbot.py
from langchain import OpenAI, ConversationChain
from langchain.memory import ConversationBufferMemory
from langchain.embeddings import OpenAIEmbeddings
from langchain.vectorstores import Chroma

llm = OpenAI(temperature=0.7)
# DETECTED: LLM/RAG Application
# AGENT: llm-development-expert
# CONFIDENCE: 95%
```

## Framework-Specific Patterns

```python
def detect_framework_patterns(self, code: str) -> Dict[str, float]:
    patterns = {
        'django': {
            'indicators': ['class.*View', 'models.Model', 'forms.Form'],
            'confidence': 0.9
        },
        'fastapi': {
            'indicators': ['@app.', 'async def', 'Depends('],
            'confidence': 0.9
        },
        'flask': {
            'indicators': ['@app.route', 'render_template', 'Blueprint'],
            'confidence': 0.85
        },
        'pytest': {
            'indicators': ['def test_', '@pytest.', 'assert '],
            'confidence': 0.8
        },
        'airflow': {
            'indicators': ['DAG(', '@task', 'PythonOperator'],
            'confidence': 0.85
        }
    }
    
    detected = {}
    for framework, config in patterns.items():
        if any(pattern in code for pattern in config['indicators']):
            detected[framework] = config['confidence']
    
    return detected
```

## Notebook Detection

```python
def analyze_notebook(self, notebook_path: str) -> ProjectContext:
    """Special handling for Jupyter notebooks"""
    with open(notebook_path, 'r') as f:
        notebook = json.load(f)
    
    all_code = []
    for cell in notebook.get('cells', []):
        if cell['cell_type'] == 'code':
            all_code.extend(cell['source'])
    
    code_text = '\n'.join(all_code)
    
    # Check for ML/Data Science patterns
    if 'model.fit' in code_text or 'train(' in code_text:
        return ProjectContext(ProjectType.MACHINE_LEARNING, 0.9)
    
    if 'pd.DataFrame' in code_text or 'plt.show()' in code_text:
        return ProjectContext(ProjectType.DATA_SCIENCE, 0.85)
    
    return ProjectContext(ProjectType.GENERAL_NOTEBOOK, 0.5)
```

## Contextual Questions

```
I detected Python code but need more context:

1. Machine Learning/AI (PyTorch, TensorFlow, Scikit-learn)
2. LLM Development (LangChain, OpenAI, RAG systems)
3. Data Science/Analysis (Pandas, NumPy, Jupyter)
4. Django Web Application
5. FastAPI Backend Service
6. Flask Web Application
7. Streamlit Data App
8. Data Engineering (Airflow, ETL)
9. Web Scraping (Scrapy, BeautifulSoup)
10. Testing (Pytest, Unittest)
11. Desktop Application (Tkinter, PyQt)
12. Game Development (Pygame)
13. DevOps/Automation Scripts
14. Library/Package Development

Additional context:
- Are you using Jupyter notebooks?
- Is this a REST API or web application?
- Are you training models or deploying them?
```