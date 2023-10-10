# String

- GetPath(dir, file string) string 获取文件路径
- GetFileDirectoryToCaller(opts ...int) 根据运行堆栈信息获取文件目录，skip 默认1
- GetCurrentAbPathByExecutable() (string, error) 获取当前执行文件绝对路径
- GetCurrentPath() (dir string, err error) 获取当前执行文件路径，如果是临时目录则获取当前文件的的执行路径
- GetDefaultPath() (dir string, err error) 获取当前执行文件路径，如果是临时目录则获取运行命令的工作目录
