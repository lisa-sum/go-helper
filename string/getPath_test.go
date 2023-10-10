package string

import (
	"os"
	"path/filepath"
	"testing"
)

func TestGetPath(t *testing.T) {
	// 测试获取路径是否正确
	workPath, filePath := GetPath("/upload/video", "title.mp4")
	t.Log("workPath:", workPath)
	t.Log("filePath:", filePath)
	testPath, err := os.Getwd()
	if err != nil {
		t.Fatal("testPath", err)
	}

	// sysType := runtime.GOOS
	// // 如果是Linux环境, 则将路径分割符替换成/,并进行比对
	// if sysType == "linux" {
	// 	testPath = filepath.FromSlash(testPath)
	// 	if workPath != testPath {
	// 		t.Fatal("获取根目录失败:" + err.Error())
	// 	}
	// }
	testPath = filepath.ToSlash(testPath)

	t.Log("testPath:", testPath)
	t.Log("workPath:", workPath)
	if testPath != workPath {
		t.Logf("workPath != testPath: %s", err)
	}

	// 获取执行目录作为默认目录
	runPath := GetRunPath()
	t.Logf("获取执行目录作为默认目录:%s", runPath)

	// 获取执行目录作为默认目录
	FileDirectoryToCaller, ok := GetFileDirectoryToCaller()
	if !ok {
		t.Errorf("获取执行目录作为默认目录:%s", FileDirectoryToCaller)
	}
	t.Logf("获取执行目录作为默认目录:%s", FileDirectoryToCaller)

	// 获取当前执行文件绝对路径
	executable, err := GetCurrentAbPathByExecutable()
	if err != nil {
		t.Error(err)
	}
	t.Logf("获取当前执行文件绝对路径:%s", executable)

	// 获取当前执行文件路径，如果是临时目录则获取当前文件的的执行路径
	currentPath, err := GetCurrentPath()
	if err != nil {
		t.Error(err)
	}
	t.Logf("获取当前执行文件路径，如果是临时目录则获取当前文件的的执行路径:%s", currentPath)

	// 获取当前执行文件路径
	defaultPath, err := GetDefaultPath()
	if err != nil {
		t.Error(err)
	}
	t.Logf("获取当前执行文件路径:%s", defaultPath)
}
