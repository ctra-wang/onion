package cmdtools

import (
	"fmt"
	"github.com/ctra-wang/onion/internal/logic/commontools"
	"github.com/ctra-wang/onion/internal/logic/model"
	"os/exec"
)

// RpcGenerator rpc 生成器
func RpcGenerator(rpcParams model.RpcParams) {
	rpcCmd := fmt.Sprintf("goctls rpc new %s -e -m %s -p %d -d", rpcParams.RpcName, rpcParams.ModuleName, rpcParams.Port)
	// 要执行的命令和参数
	cmd := exec.Command(rpcCmd)

	// 获取命令输出
	output, err := cmd.Output()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// 打印输出
	fmt.Println(string(output))
}

// EntGenerator ent 生成器
func EntGenerator(databaseConf model.DatabaseConf) {
	entStr := fmt.Sprintf("goctls extra ent import -d \"mysql://%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local\" -t %s", databaseConf.Username, databaseConf.Password, databaseConf.Host, databaseConf.Port, databaseConf.DBName, databaseConf.TableName)
	// 要执行的命令和参数
	cmd := exec.Command(entStr)

	// 获取命令输出
	output, err := cmd.Output()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// 打印输出
	fmt.Println(string(output))
}

// ProtoGenerator proto 生成器
func ProtoGenerator(modelName string) {
	protoStr := fmt.Sprintf("make gen-rpc-ent-logic model=%s group=%s", modelName, commontools.ToLowerFirstChar(modelName))
	// 要执行的命令和参数
	cmd := exec.Command(protoStr)

	// 获取命令输出
	output, err := cmd.Output()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// 打印输出
	fmt.Println(string(output))
}
