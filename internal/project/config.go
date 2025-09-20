package project

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"time"
)

// ProjectConfig 项目配置
type ProjectConfig struct {
	Name        string            `json:"name"`
	Path        string            `json:"path"`
	AIModel     AIModelType       `json:"ai_model"`
	YoloMode    bool              `json:"yolo_mode"`
	LastUsed    time.Time         `json:"last_used"`
	Preferences map[string]string `json:"preferences"`
}

// AIModelType AI模型类型
type AIModelType string

const (
	ModelClaudeCode AIModelType = "claude_code"
	ModelGeminiCLI  AIModelType = "gemini_cli"
	ModelCodex      AIModelType = "codex"
	ModelAider      AIModelType = "aider"
	ModelCustom     AIModelType = "custom"
)

// String 返回模型类型的字符串表示
func (a AIModelType) String() string {
	switch a {
	case ModelClaudeCode:
		return "Claude Code"
	case ModelGeminiCLI:
		return "Gemini CLI"
	case ModelCodex:
		return "Codex"
	case ModelAider:
		return "Aider"
	case ModelCustom:
		return "Custom"
	default:
		return "Unknown"
	}
}

// GetCommand 获取模型对应的启动命令
func (a AIModelType) GetCommand(yoloMode bool) []string {
	switch a {
	case ModelClaudeCode:
		if yoloMode {
			return []string{"claude", "--dangerously-skip-permissions"}
		}
		return []string{"claude"}
	case ModelGeminiCLI:
		if yoloMode {
			return []string{"gemini", "--yolo"}
		}
		return []string{"gemini"}
	case ModelCodex:
		if yoloMode {
			return []string{"codex", "--dangerously-bypass-approvals-and-sandbox"}
		}
		return []string{"codex"}
	case ModelAider:
		if yoloMode {
			return []string{"aider", "--yes"}
		}
		return []string{"aider"}
	default:
		return []string{"echo", "Unknown model"}
	}
}

// GetIcon 获取模型图标
func (a AIModelType) GetIcon() string {
	switch a {
	case ModelClaudeCode:
		return "🤖"
	case ModelGeminiCLI:
		return "💎"
	case ModelCodex:
		return "🔧"
	case ModelAider:
		return "🔬"
	case ModelCustom:
		return "⚙️"
	default:
		return "❓"
	}
}

// ConfigManager 配置管理器
type ConfigManager struct {
	configDir  string
	configFile string
	projects   []ProjectConfig
}

// NewConfigManager 创建新的配置管理器
func NewConfigManager() *ConfigManager {
	homeDir, _ := os.UserHomeDir()
	configDir := filepath.Join(homeDir, ".ai-launcher")
	configFile := filepath.Join(configDir, "projects.json")

	return &ConfigManager{
		configDir:  configDir,
		configFile: configFile,
		projects:   make([]ProjectConfig, 0),
	}
}

// LoadProjects 加载项目配置
func (cm *ConfigManager) LoadProjects() error {
	// 确保配置目录存在
	if err := os.MkdirAll(cm.configDir, 0755); err != nil {
		return fmt.Errorf("创建配置目录失败: %v", err)
	}

	// 检查配置文件是否存在
	if _, err := os.Stat(cm.configFile); os.IsNotExist(err) {
		// 文件不存在，创建空配置
		return cm.SaveProjects()
	}

	// 读取配置文件
	data, err := os.ReadFile(cm.configFile)
	if err != nil {
		return fmt.Errorf("读取配置文件失败: %v", err)
	}

	// 解析JSON
	if err := json.Unmarshal(data, &cm.projects); err != nil {
		return fmt.Errorf("解析配置文件失败: %v", err)
	}

	return nil
}

// SaveProjects 保存项目配置
func (cm *ConfigManager) SaveProjects() error {
	// 确保配置目录存在
	if err := os.MkdirAll(cm.configDir, 0755); err != nil {
		return fmt.Errorf("创建配置目录失败: %v", err)
	}

	// 序列化为JSON
	data, err := json.MarshalIndent(cm.projects, "", "  ")
	if err != nil {
		return fmt.Errorf("序列化配置失败: %v", err)
	}

	// 写入文件
	if err := os.WriteFile(cm.configFile, data, 0644); err != nil {
		return fmt.Errorf("写入配置文件失败: %v", err)
	}

	return nil
}

// AddProject 添加项目配置
func (cm *ConfigManager) AddProject(project ProjectConfig) error {
	// 检查项目是否已存在
	for i, p := range cm.projects {
		if p.Path == project.Path {
			// 更新现有项目
			cm.projects[i] = project
			return cm.SaveProjects()
		}
	}

	// 添加新项目
	project.LastUsed = time.Now()
	cm.projects = append(cm.projects, project)
	return cm.SaveProjects()
}

// GetProjects 获取所有项目
func (cm *ConfigManager) GetProjects() []ProjectConfig {
	return cm.projects
}

// GetRecentProjects 获取最近使用的项目
func (cm *ConfigManager) GetRecentProjects(limit int) []ProjectConfig {
	// 按最后使用时间排序
	projects := make([]ProjectConfig, len(cm.projects))
	copy(projects, cm.projects)

	// 简单的排序，实际可以使用sort包
	for i := 0; i < len(projects)-1; i++ {
		for j := i + 1; j < len(projects); j++ {
			if projects[i].LastUsed.Before(projects[j].LastUsed) {
				projects[i], projects[j] = projects[j], projects[i]
			}
		}
	}

	// 限制返回数量
	if limit > 0 && limit < len(projects) {
		return projects[:limit]
	}

	return projects
}

// RemoveProject 删除项目配置
func (cm *ConfigManager) RemoveProject(path string) error {
	for i, p := range cm.projects {
		if p.Path == path {
			// 删除项目
			cm.projects = append(cm.projects[:i], cm.projects[i+1:]...)
			return cm.SaveProjects()
		}
	}
	return fmt.Errorf("项目不存在: %s", path)
}

// UpdateProjectUsage 更新项目使用时间
func (cm *ConfigManager) UpdateProjectUsage(path string) error {
	for i, p := range cm.projects {
		if p.Path == path {
			cm.projects[i].LastUsed = time.Now()
			return cm.SaveProjects()
		}
	}
	return fmt.Errorf("项目不存在: %s", path)
}

// GetProjectByPath 根据路径获取项目
func (cm *ConfigManager) GetProjectByPath(path string) (*ProjectConfig, error) {
	for _, p := range cm.projects {
		if p.Path == path {
			return &p, nil
		}
	}
	return nil, fmt.Errorf("项目不存在: %s", path)
}

// ValidateProjectPath 验证项目路径
func (cm *ConfigManager) ValidateProjectPath(path string) error {
	// 检查路径是否存在
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return fmt.Errorf("路径不存在: %s", path)
	}

	// 检查是否为目录
	info, err := os.Stat(path)
	if err != nil {
		return fmt.Errorf("无法访问路径: %v", err)
	}

	if !info.IsDir() {
		return fmt.Errorf("路径必须是目录: %s", path)
	}

	return nil
}

// GetAvailableModels 获取可用的AI模型
func (cm *ConfigManager) GetAvailableModels() []AIModelType {
	return []AIModelType{
		ModelClaudeCode,
		ModelGeminiCLI,
		ModelCodex,
		ModelAider,
	}
}

// IsValidModel 检查模型是否有效
func (cm *ConfigManager) IsValidModel(model AIModelType) bool {
	for _, m := range cm.GetAvailableModels() {
		if m == model {
			return true
		}
	}
	return false
}