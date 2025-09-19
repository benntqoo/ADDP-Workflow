package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"

	"ai-launcher/internal/project"
	"ai-launcher/internal/terminal"
)

// WebGUI Web版本的GUI
type WebGUI struct {
	configManager   *project.ConfigManager
	terminalManager *terminal.TerminalManager
	port            string
}

// NewWebGUI 创建Web GUI
func NewWebGUI() *WebGUI {
	return &WebGUI{
		configManager:   project.NewConfigManager(),
		terminalManager: terminal.NewTerminalManager(),
		port:            "8080",
	}
}

// Run 启动Web GUI
func (w *WebGUI) Run() {
	if err := w.configManager.LoadProjects(); err != nil {
		log.Printf("加载配置失败: %v", err)
	}

	// 静态文件和API路由
	http.HandleFunc("/", w.handleHome)
	http.HandleFunc("/api/projects", w.handleProjects)
	http.HandleFunc("/api/models", w.handleModels)
	http.HandleFunc("/api/launch", w.handleLaunch)
	http.HandleFunc("/api/save", w.handleSave)

	url := fmt.Sprintf("http://localhost:%s", w.port)
	fmt.Printf("🚀 AI启动器 Web版本启动成功！\n")
	fmt.Printf("📱 请在浏览器中打开: %s\n", url)
	fmt.Printf("💡 按 Ctrl+C 停止服务器\n\n")

	// 自动打开浏览器
	go w.openBrowser(url)

	log.Fatal(http.ListenAndServe(":"+w.port, nil))
}

// handleHome 主页面
func (w *WebGUI) handleHome(rw http.ResponseWriter, req *http.Request) {
	html := `
<!DOCTYPE html>
<html lang="zh-CN">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>🚀 AI启动器</title>
    <style>
        * { margin: 0; padding: 0; box-sizing: border-box; }
        body {
            font-family: 'Segoe UI', Tahoma, Geneva, Verdana, sans-serif;
            background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
            min-height: 100vh;
            padding: 20px;
        }
        .container {
            max-width: 1200px;
            margin: 0 auto;
            background: white;
            border-radius: 15px;
            box-shadow: 0 20px 40px rgba(0,0,0,0.1);
            overflow: hidden;
        }
        .header {
            background: linear-gradient(45deg, #4CAF50, #45a049);
            color: white;
            padding: 20px;
            text-align: center;
        }
        .content {
            display: grid;
            grid-template-columns: 300px 1fr;
            min-height: 500px;
        }
        .sidebar {
            background: #f8f9fa;
            padding: 20px;
            border-right: 1px solid #dee2e6;
        }
        .main {
            padding: 20px;
        }
        .project-card {
            background: white;
            border: 1px solid #dee2e6;
            border-radius: 8px;
            padding: 15px;
            margin-bottom: 10px;
            cursor: pointer;
            transition: all 0.3s;
        }
        .project-card:hover {
            transform: translateY(-2px);
            box-shadow: 0 4px 8px rgba(0,0,0,0.1);
        }
        .form-group {
            margin-bottom: 20px;
        }
        .form-group label {
            display: block;
            margin-bottom: 5px;
            font-weight: bold;
        }
        .form-control {
            width: 100%;
            padding: 10px;
            border: 1px solid #ddd;
            border-radius: 5px;
            font-size: 14px;
        }
        .btn {
            padding: 12px 24px;
            border: none;
            border-radius: 5px;
            cursor: pointer;
            font-size: 16px;
            margin-right: 10px;
            transition: all 0.3s;
        }
        .btn-primary {
            background: #007bff;
            color: white;
        }
        .btn-primary:hover {
            background: #0056b3;
        }
        .btn-success {
            background: #28a745;
            color: white;
        }
        .btn-success:hover {
            background: #1e7e34;
        }
        .status {
            margin-top: 20px;
            padding: 10px;
            border-radius: 5px;
            display: none;
        }
        .status.success {
            background: #d4edda;
            color: #155724;
            border: 1px solid #c3e6cb;
        }
        .status.error {
            background: #f8d7da;
            color: #721c24;
            border: 1px solid #f5c6cb;
        }
        .checkbox-group {
            display: flex;
            align-items: center;
            gap: 10px;
        }
        .mode-description {
            background: #f8f9fa;
            padding: 15px;
            border-radius: 5px;
            margin-top: 10px;
            font-size: 14px;
        }
    </style>
</head>
<body>
    <div class="container">
        <div class="header">
            <h1>🚀 AI启动器 - 智能多AI工具启动器</h1>
            <p>Web版本 - 一键启动各种AI编程助手</p>
        </div>
        <div class="content">
            <div class="sidebar">
                <h3>📁 最近项目</h3>
                <div id="recent-projects"></div>
                <button class="btn btn-primary" onclick="clearForm()">➕ 新建项目</button>
            </div>
            <div class="main">
                <h3>⚙️ 项目配置</h3>
                <form id="project-form">
                    <div class="form-group">
                        <label>📁 项目路径</label>
                        <input type="text" id="project-path" class="form-control" placeholder="请输入项目目录路径">
                    </div>
                    <div class="form-group">
                        <label>📝 项目名称</label>
                        <input type="text" id="project-name" class="form-control" placeholder="项目名称">
                    </div>
                    <div class="form-group">
                        <label>🤖 AI模型</label>
                        <select id="ai-model" class="form-control">
                            <option value="claude_code">🤖 Claude Code</option>
                            <option value="gemini_cli">💎 Gemini CLI</option>
                            <option value="codex">🔧 Codex</option>
                        </select>
                    </div>
                    <div class="form-group">
                        <label>⚡ 运行模式</label>
                        <div class="checkbox-group">
                            <input type="checkbox" id="yolo-mode">
                            <label for="yolo-mode">启用YOLO模式 (跳过安全确认)</label>
                        </div>
                        <div class="mode-description">
                            <strong>🛡️ 普通模式：</strong> 需要用户确认重要操作，适合生产环境<br>
                            <strong>🚀 YOLO模式：</strong> 跳过安全检查，适合实验和快速原型
                        </div>
                    </div>
                    <div class="form-group">
                        <button type="button" class="btn btn-success" onclick="launchAI()">🚀 启动AI工具</button>
                        <button type="button" class="btn btn-primary" onclick="saveProject()">💾 保存配置</button>
                    </div>
                </form>
                <div id="status" class="status"></div>
            </div>
        </div>
    </div>

    <script>
        // 加载最近项目
        async function loadRecentProjects() {
            try {
                const response = await fetch('/api/projects');
                const projects = await response.json();
                const container = document.getElementById('recent-projects');

                container.innerHTML = '';
                projects.forEach(project => {
                    const card = document.createElement('div');
                    card.className = 'project-card';
                    card.onclick = () => loadProject(project);
                    card.innerHTML = ` + "`" + `
                        <div><strong>${project.name}</strong></div>
                        <div>📍 ${project.path}</div>
                        <div>🤖 ${getModelIcon(project.ai_model)} ${getModelName(project.ai_model)}</div>
                        <div>⚡ ${project.yolo_mode ? '🚀 YOLO' : '🛡️ 普通'}</div>
                    ` + "`" + `;
                });
            } catch (error) {
                console.error('加载项目失败:', error);
            }
        }

        // 加载项目到表单
        function loadProject(project) {
            document.getElementById('project-path').value = project.path;
            document.getElementById('project-name').value = project.name;
            document.getElementById('ai-model').value = project.ai_model;
            document.getElementById('yolo-mode').checked = project.yolo_mode;
            showStatus('已加载项目: ' + project.name, 'success');
        }

        // 清空表单
        function clearForm() {
            document.getElementById('project-form').reset();
            showStatus('准备创建新项目', 'success');
        }

        // 启动AI工具
        async function launchAI() {
            const config = getFormData();
            if (!validateForm(config)) return;

            try {
                const response = await fetch('/api/launch', {
                    method: 'POST',
                    headers: {'Content-Type': 'application/json'},
                    body: JSON.stringify(config)
                });

                const result = await response.json();
                if (result.success) {
                    showStatus('✅ AI工具启动成功！', 'success');
                    loadRecentProjects();
                } else {
                    showStatus('❌ 启动失败: ' + result.error, 'error');
                }
            } catch (error) {
                showStatus('❌ 启动失败: ' + error.message, 'error');
            }
        }

        // 保存项目
        async function saveProject() {
            const config = getFormData();
            if (!validateForm(config)) return;

            try {
                const response = await fetch('/api/save', {
                    method: 'POST',
                    headers: {'Content-Type': 'application/json'},
                    body: JSON.stringify(config)
                });

                const result = await response.json();
                if (result.success) {
                    showStatus('✅ 配置已保存', 'success');
                    loadRecentProjects();
                } else {
                    showStatus('❌ 保存失败: ' + result.error, 'error');
                }
            } catch (error) {
                showStatus('❌ 保存失败: ' + error.message, 'error');
            }
        }

        // 获取表单数据
        function getFormData() {
            return {
                path: document.getElementById('project-path').value,
                name: document.getElementById('project-name').value,
                ai_model: document.getElementById('ai-model').value,
                yolo_mode: document.getElementById('yolo-mode').checked
            };
        }

        // 验证表单
        function validateForm(config) {
            if (!config.path) {
                showStatus('❌ 请输入项目路径', 'error');
                return false;
            }
            if (!config.name) {
                showStatus('❌ 请输入项目名称', 'error');
                return false;
            }
            return true;
        }

        // 显示状态
        function showStatus(message, type) {
            const status = document.getElementById('status');
            status.textContent = message;
            status.className = 'status ' + type;
            status.style.display = 'block';
            setTimeout(() => {
                status.style.display = 'none';
            }, 5000);
        }

        // 获取模型图标
        function getModelIcon(model) {
            switch(model) {
                case 'claude_code': return '🤖';
                case 'gemini_cli': return '💎';
                case 'codex': return '🔧';
                default: return '❓';
            }
        }

        // 获取模型名称
        function getModelName(model) {
            switch(model) {
                case 'claude_code': return 'Claude Code';
                case 'gemini_cli': return 'Gemini CLI';
                case 'codex': return 'Codex';
                default: return 'Unknown';
            }
        }

        // 页面加载时初始化
        document.addEventListener('DOMContentLoaded', function() {
            loadRecentProjects();
            showStatus('🚀 AI启动器已就绪', 'success');
        });
    </script>
</body>
</html>
`
	rw.Header().Set("Content-Type", "text/html; charset=utf-8")
	rw.Write([]byte(html))
}

// handleProjects 处理项目API
func (w *WebGUI) handleProjects(rw http.ResponseWriter, req *http.Request) {
	projects := w.configManager.GetRecentProjects(10)
	rw.Header().Set("Content-Type", "application/json")
	json.NewEncoder(rw).Encode(projects)
}

// handleModels 处理模型API
func (w *WebGUI) handleModels(rw http.ResponseWriter, req *http.Request) {
	models := w.configManager.GetAvailableModels()
	rw.Header().Set("Content-Type", "application/json")
	json.NewEncoder(rw).Encode(models)
}

// handleLaunch 处理启动API
func (w *WebGUI) handleLaunch(rw http.ResponseWriter, req *http.Request) {
	if req.Method != "POST" {
		http.Error(rw, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var config project.ProjectConfig
	if err := json.NewDecoder(req.Body).Decode(&config); err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return
	}

	// 验证路径
	if err := w.configManager.ValidateProjectPath(config.Path); err != nil {
		response := map[string]interface{}{
			"success": false,
			"error":   err.Error(),
		}
		json.NewEncoder(rw).Encode(response)
		return
	}

	// 保存配置
	if err := w.configManager.AddProject(config); err != nil {
		response := map[string]interface{}{
			"success": false,
			"error":   err.Error(),
		}
		json.NewEncoder(rw).Encode(response)
		return
	}

	// 启动AI工具
	terminalConfig := terminal.TerminalConfig{
		Type:       w.getTerminalType(config.AIModel),
		Name:       fmt.Sprintf("%s-%s", config.AIModel, config.Name),
		WorkingDir: config.Path,
		Command:    config.AIModel.GetCommand(config.YoloMode),
		YoloMode:   config.YoloMode,
	}

	err := w.terminalManager.StartTerminal(terminalConfig)
	response := map[string]interface{}{
		"success": err == nil,
	}
	if err != nil {
		response["error"] = err.Error()
	}

	rw.Header().Set("Content-Type", "application/json")
	json.NewEncoder(rw).Encode(response)
}

// handleSave 处理保存API
func (w *WebGUI) handleSave(rw http.ResponseWriter, req *http.Request) {
	if req.Method != "POST" {
		http.Error(rw, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var config project.ProjectConfig
	if err := json.NewDecoder(req.Body).Decode(&config); err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return
	}

	// 验证路径
	if err := w.configManager.ValidateProjectPath(config.Path); err != nil {
		response := map[string]interface{}{
			"success": false,
			"error":   err.Error(),
		}
		json.NewEncoder(rw).Encode(response)
		return
	}

	// 保存配置
	err := w.configManager.AddProject(config)
	response := map[string]interface{}{
		"success": err == nil,
	}
	if err != nil {
		response["error"] = err.Error()
	}

	rw.Header().Set("Content-Type", "application/json")
	json.NewEncoder(rw).Encode(response)
}

// getTerminalType 获取终端类型
func (w *WebGUI) getTerminalType(model project.AIModelType) terminal.TerminalType {
	switch model {
	case project.ModelClaudeCode:
		return terminal.TypeClaudeCode
	case project.ModelGeminiCLI:
		return terminal.TypeGeminiCLI
	case project.ModelCodex:
		return terminal.TypeCustom
	default:
		return terminal.TypeCustom
	}
}

// openBrowser 打开浏览器
func (w *WebGUI) openBrowser(url string) {
	var cmd string
	var args []string

	switch runtime.GOOS {
	case "windows":
		cmd = "cmd"
		args = []string{"/c", "start", url}
	case "darwin":
		cmd = "open"
		args = []string{url}
	default: // "linux", "freebsd", "openbsd", "netbsd"
		cmd = "xdg-open"
		args = []string{url}
	}

	exec.Command(cmd, args...).Start()
}