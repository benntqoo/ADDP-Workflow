package project

import (
	"os"
	"path/filepath"
	"testing"
	"time"
)

func TestAIModelType_String(t *testing.T) {
	tests := []struct {
		model    AIModelType
		expected string
	}{
		{ModelClaudeCode, "Claude Code"},
		{ModelGeminiCLI, "Gemini CLI"},
		{ModelCodex, "Codex"},
		{ModelCustom, "Custom"},
		{AIModelType("unknown"), "Unknown"},
	}

	for _, test := range tests {
		result := test.model.String()
		if result != test.expected {
			t.Errorf("Expected %s, got %s", test.expected, result)
		}
	}
}

func TestAIModelType_GetCommand(t *testing.T) {
	tests := []struct {
		model    AIModelType
		yolo     bool
		expected []string
	}{
		{ModelClaudeCode, false, []string{"claude"}},
		{ModelClaudeCode, true, []string{"claude", "--dangerously-skip-permissions"}},
		{ModelGeminiCLI, false, []string{"gemini"}},
		{ModelGeminiCLI, true, []string{"gemini", "--yolo"}},
		{ModelCodex, false, []string{"codex"}},
		{ModelCodex, true, []string{"codex", "--dangerously-bypass-approvals-and-sandbox"}},
	}

	for _, test := range tests {
		result := test.model.GetCommand(test.yolo)
		if len(result) != len(test.expected) {
			t.Errorf("Expected %d args, got %d", len(test.expected), len(result))
			continue
		}

		for i, arg := range test.expected {
			if result[i] != arg {
				t.Errorf("Expected arg %d to be %s, got %s", i, arg, result[i])
			}
		}
	}
}

func TestAIModelType_GetIcon(t *testing.T) {
	tests := []struct {
		model    AIModelType
		expected string
	}{
		{ModelClaudeCode, "🤖"},
		{ModelGeminiCLI, "💎"},
		{ModelCodex, "🔧"},
		{ModelCustom, "⚙️"},
		{AIModelType("unknown"), "❓"},
	}

	for _, test := range tests {
		result := test.model.GetIcon()
		if result != test.expected {
			t.Errorf("Expected %s, got %s", test.expected, result)
		}
	}
}

func TestConfigManager_NewConfigManager(t *testing.T) {
	cm := NewConfigManager()

	if cm == nil {
		t.Fatal("Expected non-nil ConfigManager")
	}

	if cm.projects == nil {
		t.Error("Expected projects slice to be initialized")
	}

	if cm.configDir == "" {
		t.Error("Expected configDir to be set")
	}

	if cm.configFile == "" {
		t.Error("Expected configFile to be set")
	}
}

func TestConfigManager_LoadAndSaveProjects(t *testing.T) {
	// 创建临时目录进行测试
	tempDir := t.TempDir()

	cm := &ConfigManager{
		configDir:  tempDir,
		configFile: filepath.Join(tempDir, "test_projects.json"),
		projects:   make([]ProjectConfig, 0),
	}

	// 测试保存空项目列表
	err := cm.SaveProjects()
	if err != nil {
		t.Fatalf("Failed to save empty projects: %v", err)
	}

	// 测试加载空项目列表
	err = cm.LoadProjects()
	if err != nil {
		t.Fatalf("Failed to load projects: %v", err)
	}

	if len(cm.projects) != 0 {
		t.Errorf("Expected 0 projects, got %d", len(cm.projects))
	}

	// 添加测试项目
	testProject := ProjectConfig{
		Name:     "test-project",
		Path:     "/test/path",
		AIModel:  ModelClaudeCode,
		YoloMode: true,
		LastUsed: time.Now(),
	}

	err = cm.AddProject(testProject)
	if err != nil {
		t.Fatalf("Failed to add project: %v", err)
	}

	// 重新加载并验证
	cm2 := &ConfigManager{
		configDir:  tempDir,
		configFile: filepath.Join(tempDir, "test_projects.json"),
		projects:   make([]ProjectConfig, 0),
	}

	err = cm2.LoadProjects()
	if err != nil {
		t.Fatalf("Failed to load projects after adding: %v", err)
	}

	if len(cm2.projects) != 1 {
		t.Fatalf("Expected 1 project, got %d", len(cm2.projects))
	}

	loadedProject := cm2.projects[0]
	if loadedProject.Name != testProject.Name {
		t.Errorf("Expected name %s, got %s", testProject.Name, loadedProject.Name)
	}

	if loadedProject.Path != testProject.Path {
		t.Errorf("Expected path %s, got %s", testProject.Path, loadedProject.Path)
	}

	if loadedProject.AIModel != testProject.AIModel {
		t.Errorf("Expected model %s, got %s", testProject.AIModel, loadedProject.AIModel)
	}

	if loadedProject.YoloMode != testProject.YoloMode {
		t.Errorf("Expected yolo mode %v, got %v", testProject.YoloMode, loadedProject.YoloMode)
	}
}

func TestConfigManager_AddProject(t *testing.T) {
	tempDir := t.TempDir()

	cm := &ConfigManager{
		configDir:  tempDir,
		configFile: filepath.Join(tempDir, "test_projects.json"),
		projects:   make([]ProjectConfig, 0),
	}

	// 添加第一个项目
	project1 := ProjectConfig{
		Name:     "project1",
		Path:     "/path1",
		AIModel:  ModelClaudeCode,
		YoloMode: false,
	}

	err := cm.AddProject(project1)
	if err != nil {
		t.Fatalf("Failed to add project1: %v", err)
	}

	if len(cm.projects) != 1 {
		t.Errorf("Expected 1 project, got %d", len(cm.projects))
	}

	// 添加相同路径的项目（应该更新而不是新增）
	project1Updated := ProjectConfig{
		Name:     "project1-updated",
		Path:     "/path1", // 相同路径
		AIModel:  ModelGeminiCLI,
		YoloMode: true,
	}

	err = cm.AddProject(project1Updated)
	if err != nil {
		t.Fatalf("Failed to update project1: %v", err)
	}

	if len(cm.projects) != 1 {
		t.Errorf("Expected 1 project after update, got %d", len(cm.projects))
	}

	// 验证项目已更新
	updated := cm.projects[0]
	if updated.Name != "project1-updated" {
		t.Errorf("Expected updated name, got %s", updated.Name)
	}

	if updated.AIModel != ModelGeminiCLI {
		t.Errorf("Expected updated model, got %s", updated.AIModel)
	}

	// 添加不同路径的项目
	project2 := ProjectConfig{
		Name: "project2",
		Path: "/path2",
	}

	err = cm.AddProject(project2)
	if err != nil {
		t.Fatalf("Failed to add project2: %v", err)
	}

	if len(cm.projects) != 2 {
		t.Errorf("Expected 2 projects, got %d", len(cm.projects))
	}
}

func TestConfigManager_GetRecentProjects(t *testing.T) {
	cm := &ConfigManager{
		projects: []ProjectConfig{
			{Name: "old", Path: "/old", LastUsed: time.Now().Add(-2 * time.Hour)},
			{Name: "newest", Path: "/newest", LastUsed: time.Now()},
			{Name: "older", Path: "/older", LastUsed: time.Now().Add(-1 * time.Hour)},
		},
	}

	// 测试获取所有项目
	recent := cm.GetRecentProjects(0)
	if len(recent) != 3 {
		t.Errorf("Expected 3 projects, got %d", len(recent))
	}

	// 验证排序（最新的在前）
	if recent[0].Name != "newest" {
		t.Errorf("Expected newest project first, got %s", recent[0].Name)
	}

	if recent[1].Name != "older" {
		t.Errorf("Expected older project second, got %s", recent[1].Name)
	}

	if recent[2].Name != "old" {
		t.Errorf("Expected old project third, got %s", recent[2].Name)
	}

	// 测试限制数量
	recent = cm.GetRecentProjects(2)
	if len(recent) != 2 {
		t.Errorf("Expected 2 projects, got %d", len(recent))
	}

	if recent[0].Name != "newest" {
		t.Errorf("Expected newest project first, got %s", recent[0].Name)
	}
}

func TestConfigManager_ValidateProjectPath(t *testing.T) {
	cm := NewConfigManager()

	// 测试不存在的路径
	err := cm.ValidateProjectPath("/nonexistent/path")
	if err == nil {
		t.Error("Expected error for nonexistent path")
	}

	// 测试当前目录（应该存在）
	wd, _ := os.Getwd()
	err = cm.ValidateProjectPath(wd)
	if err != nil {
		t.Errorf("Expected no error for current directory, got: %v", err)
	}

	// 创建临时文件进行测试
	tempFile := filepath.Join(t.TempDir(), "testfile")
	file, _ := os.Create(tempFile)
	file.Close()

	// 测试文件路径（应该失败，因为需要目录）
	err = cm.ValidateProjectPath(tempFile)
	if err == nil {
		t.Error("Expected error for file path")
	}
}

func TestConfigManager_GetAvailableModels(t *testing.T) {
	cm := NewConfigManager()
	models := cm.GetAvailableModels()

	expected := []AIModelType{
		ModelClaudeCode,
		ModelGeminiCLI,
		ModelCodex,
	}

	if len(models) != len(expected) {
		t.Errorf("Expected %d models, got %d", len(expected), len(models))
	}

	for i, model := range expected {
		if models[i] != model {
			t.Errorf("Expected model %d to be %s, got %s", i, model, models[i])
		}
	}
}

func TestConfigManager_IsValidModel(t *testing.T) {
	cm := NewConfigManager()

	validModels := []AIModelType{
		ModelClaudeCode,
		ModelGeminiCLI,
		ModelCodex,
	}

	for _, model := range validModels {
		if !cm.IsValidModel(model) {
			t.Errorf("Expected %s to be valid", model)
		}
	}

	// 测试无效模型
	if cm.IsValidModel(ModelCustom) {
		t.Error("Expected ModelCustom to be invalid")
	}

	if cm.IsValidModel(AIModelType("unknown")) {
		t.Error("Expected unknown model to be invalid")
	}
}

func TestConfigManager_RemoveProject(t *testing.T) {
	tempDir := t.TempDir()
	cm := &ConfigManager{
		configDir:  tempDir,
		configFile: filepath.Join(tempDir, "projects.json"),
		projects: []ProjectConfig{
			{Name: "project1", Path: "/path1"},
			{Name: "project2", Path: "/path2"},
			{Name: "project3", Path: "/path3"},
		},
	}

	// 删除存在的项目
	err := cm.RemoveProject("/path2")
	if err != nil {
		t.Fatalf("Failed to remove project: %v", err)
	}

	if len(cm.projects) != 2 {
		t.Errorf("Expected 2 projects after removal, got %d", len(cm.projects))
	}

	// 验证正确的项目被删除
	for _, p := range cm.projects {
		if p.Path == "/path2" {
			t.Error("Project /path2 should have been removed")
		}
	}

	// 尝试删除不存在的项目
	err = cm.RemoveProject("/nonexistent")
	if err == nil {
		t.Error("Expected error when removing nonexistent project")
	}
}

func TestConfigManager_UpdateProjectUsage(t *testing.T) {
	tempDir := t.TempDir()
	oldTime := time.Now().Add(-1 * time.Hour)

	cm := &ConfigManager{
		configDir:  tempDir,
		configFile: filepath.Join(tempDir, "projects.json"),
		projects: []ProjectConfig{
			{Name: "project1", Path: "/path1", LastUsed: oldTime},
		},
	}

	// 更新存在的项目
	err := cm.UpdateProjectUsage("/path1")
	if err != nil {
		t.Fatalf("Failed to update project usage: %v", err)
	}

	// 验证时间已更新
	if !cm.projects[0].LastUsed.After(oldTime) {
		t.Error("Expected LastUsed to be updated")
	}

	// 尝试更新不存在的项目
	err = cm.UpdateProjectUsage("/nonexistent")
	if err == nil {
		t.Error("Expected error when updating nonexistent project")
	}
}

func TestConfigManager_GetProjectByPath(t *testing.T) {
	testProject := ProjectConfig{
		Name: "test",
		Path: "/test/path",
	}

	cm := &ConfigManager{
		projects: []ProjectConfig{testProject},
	}

	// 获取存在的项目
	project, err := cm.GetProjectByPath("/test/path")
	if err != nil {
		t.Fatalf("Failed to get project: %v", err)
	}

	if project.Name != "test" {
		t.Errorf("Expected project name 'test', got %s", project.Name)
	}

	// 尝试获取不存在的项目
	_, err = cm.GetProjectByPath("/nonexistent")
	if err == nil {
		t.Error("Expected error when getting nonexistent project")
	}
}