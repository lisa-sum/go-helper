package string

import (
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strings"
)

// GetPath 获取路径
// dir: 目录
// filename: 文件名
// 返回值: 目录路径, 文件路径
func GetPath(dir, filename string) (string, string) {
	_, workdir, _, _ := runtime.Caller(0)
	root := path.Dir(path.Dir(workdir)) // 获取当前工作目录
	dirPath := path.Dir(root + dir)     // 获取配置文件的目录

	var filePath string
	if filename == "" {
		filePath := path.Join(dirPath) // 获取配置文件
		return dirPath, filePath
	}
	filePath = path.Join(dirPath, filename) // 获取配置文件
	return filePath, dirPath
}

// GetRunPath 获取执行目录作为默认目录
// 返回值: 执行目录
func GetRunPath() string {
	currentPath, err := os.Getwd()
	if err != nil {
		return ""
	}
	return currentPath
}

// GetFileDirectoryToCaller 根据运行堆栈信息获取文件目录，skip 默认1
// 返回值: 文件目录, 是否成功
func GetFileDirectoryToCaller(opts ...int) (directory string, ok bool) {
	var filename string
	directory = ""
	skip := 1
	if opts != nil {
		skip = opts[0]
	}
	if _, filename, _, ok = runtime.Caller(skip); ok {
		directory = path.Dir(filename)
	}
	return
}

// GetCurrentAbPathByExecutable 获取当前执行文件绝对路径
// 返回值: 当前执行文件绝对路径, 错误信息
func GetCurrentAbPathByExecutable() (string, error) {
	exePath, err := os.Executable()
	if err != nil {
		return "", err
	}
	res, _ := filepath.EvalSymlinks(exePath)
	return filepath.Dir(res), nil
}

// GetCurrentPath 获取当前执行文件路径，如果是临时目录则获取当前文件的的执行路径
// 返回值: 当前执行文件路径, 错误信息
func GetCurrentPath() (dir string, err error) {
	dir, err = GetCurrentAbPathByExecutable()
	if err != nil {
		return "", err
	}

	tmpDir, err := filepath.EvalSymlinks(os.TempDir())
	if err != nil {
		return "", err
	}

	if strings.Contains(dir, tmpDir) {
		var ok bool
		if dir, ok = GetFileDirectoryToCaller(2); !ok {
			return "", errors.New("failed to get path")
		}
	}
	return dir, nil
}

// GetDefaultPath 获取当前执行文件路径，如果是临时目录则获取运行命令的工作目录
// 返回值: 当前执行文件路径, 错误信息
func GetDefaultPath() (dir string, err error) {
	dir, err = GetCurrentAbPathByExecutable()
	if err != nil {
		return "", err
	}

	tmpDir, err := filepath.EvalSymlinks(os.TempDir())
	if err != nil {
		return "", err
	}

	if strings.Contains(dir, tmpDir) {
		return GetRunPath(), nil
	}
	return dir, nil
}
