package factory_method

type IRuleConfigParser interface {
	Parse(data []byte)
}

type jsonRuleConfigParser struct{}

type yamlRuleConfigParser struct{}

func (J jsonRuleConfigParser) Parse(data []byte) {
	panic("implement xxx")
}

func (Y yamlRuleConfigParser) Parse(data []byte) {
	panic("implement xxx")
}

// 【创建型】
// 工厂方法：定义一个用于创建对象的接口，让子类决定实例化哪一个类，工厂方法使一个类的实例化延迟到其子类。
// Go 中不能存在继承，所以使用匿名组合来实现

// IRuleConfigParserFactory 工厂方法接口
type IRuleConfigParserFactory interface {
	CreateParser() IRuleConfigParser
}

// yamlRuleConfigParserFactory 的工厂类
type yamlRuleConfigParserFactory struct{}

// jsonRuleConfigParserFactory 的工厂类
type jsonRuleConfigParserFactory struct{}

func (y yamlRuleConfigParserFactory) CreateParser() IRuleConfigParser {
	return yamlRuleConfigParser{}
}

func (j jsonRuleConfigParserFactory) CreateParser() IRuleConfigParser {
	return jsonRuleConfigParser{}
}

func NewRuleConfigParserFactory(t string) IRuleConfigParserFactory {
	switch t {
	case "json":
		return jsonRuleConfigParserFactory{}
	case "yaml":
		return yamlRuleConfigParserFactory{}
	default:
		return nil
	}
}
