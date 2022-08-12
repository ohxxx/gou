package simple_factory

// 【创建型】
// 简单工厂：定义一个用于创建对象的接口，让子类决定实例化哪一个类。
// Go 本身没有构造函数，一般采用 NewName 方法来创建 对象/接口，当返回值是接口的时候其实就是简单工厂模式

type IRuleConfigParser interface {
	Parse(data []byte)
}

type JsonRuleConfigParser struct{}

type YamlRuleConfigParser struct{}

func (J JsonRuleConfigParser) Parse(data []byte) {
	panic("implement xxx")
}

func (Y YamlRuleConfigParser) Parse(data []byte) {
	panic("implement xxx")
}

func RuleConfigParser(name string) IRuleConfigParser {
	switch name {
	case "json":
		return JsonRuleConfigParser{}
	case "yaml":
		return YamlRuleConfigParser{}
	default:
		return nil
	}
}
