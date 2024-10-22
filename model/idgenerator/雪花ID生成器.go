// Package idgenerator 提供基于雪花算法的唯一ID生成功能
package idgenerator

import (
	"github.com/bwmarrin/snowflake"
)

// EID编码格式 定义了ID的具体文本编码格式
type EID编码格式 int

const (
	E文本   EID编码格式 = 0  // 直接转换到文本
	Base2  EID编码格式 = 2  // Base2文本编码格式
	Base32 EID编码格式 = 32 // Base32文本编码格式
	Base36 EID编码格式 = 36 // Base36文本编码格式
	Base58 EID编码格式 = 58 // Base58文本编码格式
	Base64 EID编码格式 = 64 // Base64文本编码格式
)

// E雪花算法类 提供基于雪花算法的ID生成功能
// 雪花算法(Snowflake)是twitter公司内部分布式项目采用的ID生成算法
// Snowflake算法产生是为了满足Twitter每秒上万条消息的请求,每条消息都必须分配一条唯一的id,
// 这些id还需要一些大致的顺序(方便客户端排序),并且在分布式系统中不同机器产生的 ID 必须不同
// 本类支持在多线程环境中安全使用
type E雪花算法类 struct {
	node *snowflake.Node
}

// E创建 创建并返回一个本类对象
// 参数:
//   - 结点ID: 提供本对象生成的所有ID值所处结点的ID,此值由用户自行定义,一般由"机房+服务器编号"组合构成,
//     用作确保在分布式系统中不同服务器上所生成的ID值不同
//     本参数值必须大于等于 0 且小于等于 1023 ,不然将失败
//
// 返回值:
//   - *E雪花算法类: 成功返回所创建的本类对象,失败返回nil
//   - error: 如果创建过程中发生错误,返回相应的错误信息
func E创建(结点ID int64) (*E雪花算法类, error) {
	node, err := snowflake.NewNode(结点ID)
	if err != nil {
		return nil, err
	}
	return &E雪花算法类{node: node}, nil
}

// E取ID 调用本方法生成并返回一个不重复的ID值
// 返回值:
//   - int64: 生成的唯一ID
func (e *E雪花算法类) E取ID() int64 {
	return e.node.Generate().Int64()
}

// EID到文本 将所生成的ID转换到文本
// 参数:
//   - ID值: 提供通过"E取ID"方法所生成的ID值
//   - 编码格式: 指定ID转换时所使用的编码格式
//
// 返回值:
//   - string: 转换后的文本形式ID
func EID到文本(ID值 int64, 编码格式 EID编码格式) string {
	id := snowflake.ID(ID值)
	switch 编码格式 {
	case Base2:
		return id.Base2()
	case Base32:
		return id.Base32()
	case Base36:
		return id.Base36()
	case Base58:
		return id.Base58()
	case Base64:
		return id.Base64()
	default:
		return id.String()
	}
}

// E文本到ID 将"EID到文本"方法返回的文本ID值转换回ID
// 参数:
//   - ID文本: 提供所欲转换的ID文本,必须为"EID到文本"方法的返回值
//   - 编码格式: 提供使用"EID到文本"方法转换到"ID文本"参数值时所使用的编码格式
//
// 返回值:
//   - int64: 转换后的ID值
//   - error: 如果转换过程中发生错误,返回相应的错误信息
func E文本到ID(ID文本 string, 编码格式 EID编码格式) (int64, error) {
	var id snowflake.ID
	var err error
	switch 编码格式 {
	case Base2:
		id, err = snowflake.ParseBase2(ID文本)
	case Base32:
		id, err = snowflake.ParseBase32([]byte(ID文本))
	case Base36:
		id, err = snowflake.ParseBase36(ID文本)
	case Base58:
		id, err = snowflake.ParseBase58([]byte(ID文本))
	case Base64:
		id, err = snowflake.ParseBase64(ID文本)
	default:
		id, err = snowflake.ParseString(ID文本)
	}
	if err != nil {
		return 0, err
	}
	return id.Int64(), nil
}
