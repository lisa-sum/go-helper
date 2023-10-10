# go-helper

为了Golang不包含的常见的功能而写的一个工具, 目前只为自用, 所有方法均经过作者实践, 但是不保证没有bug, 请谨慎使用

## 意义
减少重复造轮子

## 使用

1. 安装
    
    ```shell
    go get github.com:lisa-sum/go-helper
    ```

2. 使用
示例: 判断一个字符串是否在一个字符串数组中
    ```go
    package main
    import "github.com/lisa-sum/go-helper"
    
    func main() {
        method := "POST"
        methods := []string{"POST", "PUT", "PATCH"}
    
        if method == "GET" {
            act = "read"
        } else if slice.Include(methods, method) {
            act = "write"
        }
    }
    ```

## 方法
快速找到对应的方法:
- 目录起名均对应Golang中某个数据类型, 例如`Slice`切片类型就对应`slice`目录
- 目录下的文件名均对应某个方法, 例如`Include`方法就对应`include.go`文件
- 每个目录都有自述文件, 可以通过自述文件查看该目录下的所有方法的说明
- 每个方法都有对应的测试用例文件, 可以通过测试文件查看该方法的使用示例
