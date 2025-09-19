package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"time"
)

// 项目配置
type ProjectConfig struct {
	Name     string    `json:"name"`
	Path     string    `json:"path"`
	AIModel  string    `json:"ai_model"`
	YoloMode bool      `json:"yolo_mode"`
	LastUsed time.Time `json:"last_used"`
}

// AI启动器
type AILauncher struct {
	configDir string
	projects  []ProjectConfig
}

// 创建新的启动器
func NewAILauncher() *AILauncher {
	homeDir, _ := os.UserHomeDir()
	configDir := filepath.Join(homeDir, ".ai-launcher")
	os.MkdirAll(configDir, 0755)

	launcher := &AILauncher{
		configDir: configDir,
		projects:  []ProjectConfig{},
	}
	launcher.loadProjects()
	return launcher
}

// 加载项目配置
func (a *AILauncher) loadProjects() {
	configFile := filepath.Join(a.configDir, "projects.json")
	data, err := os.ReadFile(configFile)
	if err != nil {
		return // 文件不存在是正常的
	}
	json.Unmarshal(data, &a.projects)
}

// 保存项目配置
func (a *AILauncher) saveProjects() error {
	configFile := filepath.Join(a.configDir, "projects.json")
	data, _ := json.MarshalIndent(a.projects, "", "  ")
	return os.WriteFile(configFile, data, 0644)
}

// 添加项目
func (a *AILauncher) addProject(config ProjectConfig) {
	config.LastUsed = time.Now()

	// 查找是否已存在
	for i, p := range a.projects {
		if p.Path == config.Path {
			a.projects[i] = config
			a.saveProjects()
			return
		}
	}

	// 添加新项目
	a.projects = append([]ProjectConfig{config}, a.projects...)
	if len(a.projects) > 10 {
		a.projects = a.projects[:10] // 只保留最近10个
	}
	a.saveProjects()
}

// 获取AI命令
func getAICommand(model string, yoloMode bool) []string {
	switch model {
	case "claude_code":
		if yoloMode {
			return []string{"claude", "--dangerously-skip-permissions"}
		}
		return []string{"claude"}
	case "gemini_cli":
		if yoloMode {
			return []string{"gemini", "--yolo"}
		}
		return []string{"gemini"}
	case "codex":
		if yoloMode {
			return []string{"codex", "--dangerously-bypass-approvals-and-sandbox"}
		}
		return []string{"codex"}
	default:
		return []string{"echo", "Unknown AI model"}
	}
}

// 启动AI工具
func (a *AILauncher) launchAI(config ProjectConfig) error {
	// 验证路径
	if _, err := os.Stat(config.Path); os.IsNotExist(err) {
		return fmt.Errorf("项目路径不存在: %s", config.Path)
	}

	// 获取命令
	cmdArgs := getAICommand(config.AIModel, config.YoloMode)
	if len(cmdArgs) == 0 {
		return fmt.Errorf("无效的AI模型: %s", config.AIModel)
	}

	// 保存配置
	a.addProject(config)

	// 启动命令
	cmd := exec.Command(cmdArgs[0], cmdArgs[1:]...)
	cmd.Dir = config.Path

	// 根据操作系统决定如何启动
	switch runtime.GOOS {
	case "windows":
		// Windows下启动新的命令行窗口
		windowsCmd := exec.Command("cmd", "/c", "start", "cmd", "/k")
		windowsCmd.Dir = config.Path

		// 创建批处理文件来执行命令
		batFile := filepath.Join(config.Path, "temp_ai_launch.bat")
		batContent := fmt.Sprintf("@echo off\ncd /d \"%s\"\n%s\npause\ndel \"%s\"",
			config.Path,
			fmt.Sprintf("%s %s", cmdArgs[0], joinArgs(cmdArgs[1:])),
			batFile)
		os.WriteFile(batFile, []byte(batContent), 0644)

		exec.Command("cmd", "/c", "start", batFile).Start()
	default:
		// Linux/macOS
		return cmd.Start()
	}

	return nil
}

// 辅助函数：连接参数
func joinArgs(args []string) string {
	result := ""
	for _, arg := range args {
		result += " " + arg
	}
	return result
}

// Web处理器
func (a *AILauncher) setupRoutes() {
	http.HandleFunc("/", a.handleHome)
	http.HandleFunc("/api/projects", a.handleProjects)
	http.HandleFunc("/api/launch", a.handleLaunch)
	http.HandleFunc("/api/save", a.handleSave)
}

// 主页面
func (a *AILauncher) handleHome(w http.ResponseWriter, r *http.Request) {
	html := `<!DOCTYPE html>
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
    </style>
</head>
<body>
    <div class="container">
        <div class="header">
            <h1>🚀 AI启动器 - 智能多AI工具启动器</h1>
            <p>Web版本 v2.0 - 无依赖简化版</p>
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
                        <input type="text" id="project-path" class="form-control" placeholder="例如: C:\\Users\\username\\my-project">
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
                        <div>🤖 ${getModelName(project.ai_model)}</div>
                        <div>⚡ ${project.yolo_mode ? '🚀 YOLO' : '🛡️ 普通'}</div>
                    ` + "`" + `;
                });
            } catch (error) {
                console.error('加载项目失败:', error);
            }
        }

        function loadProject(project) {
            document.getElementById('project-path').value = project.path;
            document.getElementById('project-name').value = project.name;
            document.getElementById('ai-model').value = project.ai_model;
            document.getElementById('yolo-mode').checked = project.yolo_mode;
            showStatus('已加载项目: ' + project.name, 'success');
        }

        function clearForm() {
            document.getElementById('project-form').reset();
            showStatus('准备创建新项目', 'success');
        }

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
                    showStatus('✅ AI工具启动成功！新窗口已打开', 'success');
                    loadRecentProjects();
                } else {
                    showStatus('❌ 启动失败: ' + result.error, 'error');
                }
            } catch (error) {
                showStatus('❌ 启动失败: ' + error.message, 'error');
            }
        }

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

        function getFormData() {
            return {
                path: document.getElementById('project-path').value,
                name: document.getElementById('project-name').value,
                ai_model: document.getElementById('ai-model').value,
                yolo_mode: document.getElementById('yolo-mode').checked
            };
        }

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

        function showStatus(message, type) {
            const status = document.getElementById('status');
            status.textContent = message;
            status.className = 'status ' + type;
            status.style.display = 'block';
            setTimeout(() => {
                status.style.display = 'none';
            }, 5000);
        }

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
            showStatus('🚀 AI启动器已就绪，Web版本运行中', 'success');
        });
    </script>
</body>
</html>`

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Write([]byte(html))
}

// 处理项目API
func (a *AILauncher) handleProjects(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(a.projects)
}

// 处理启动API
func (a *AILauncher) handleLaunch(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var config ProjectConfig
	if err := json.NewDecoder(r.Body).Decode(&config); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err := a.launchAI(config)
	response := map[string]interface{}{
		"success": err == nil,
	}
	if err != nil {
		response["error"] = err.Error()
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// 处理保存API
func (a *AILauncher) handleSave(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var config ProjectConfig
	if err := json.NewDecoder(r.Body).Decode(&config); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	a.addProject(config)
	response := map[string]interface{}{
		"success": true,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// 打开浏览器
func openBrowser(url string) {
	var cmd string
	var args []string

	switch runtime.GOOS {
	case "windows":
		cmd = "cmd"
		args = []string{"/c", "start", url}
	case "darwin":
		cmd = "open"
		args = []string{url}
	default:
		cmd = "xdg-open"
		args = []string{url}
	}

	exec.Command(cmd, args...).Start()
}

func main() {
	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "version":
			fmt.Println("AI启动器 v2.0.0 (Web版本)")
			fmt.Println("智能多AI工具启动器 - 无依赖简化版")
			return
		case "help":
			fmt.Println("AI启动器 - 智能多AI工具启动器")
			fmt.Println("")
			fmt.Println("使用方法:")
			fmt.Println("  ai-launcher        启动Web GUI界面")
			fmt.Println("  ai-launcher version 显示版本信息")
			fmt.Println("  ai-launcher help    显示帮助信息")
			fmt.Println("")
			fmt.Println("支持的AI模型:")
			fmt.Println("  🤖 Claude Code")
			fmt.Println("  💎 Gemini CLI")
			fmt.Println("  🔧 Codex")
			return
		}
	}

	launcher := NewAILauncher()
	launcher.setupRoutes()

	port := "8080"
	url := fmt.Sprintf("http://localhost:%s", port)

	fmt.Printf("🚀 AI启动器 Web版本启动成功！\n")
	fmt.Printf("📱 请在浏览器中打开: %s\n", url)
	fmt.Printf("💡 按 Ctrl+C 停止服务器\n\n")

	// 自动打开浏览器
	go func() {
		time.Sleep(1 * time.Second)
		openBrowser(url)
	}()

	log.Fatal(http.ListenAndServe(":"+port, nil))
}