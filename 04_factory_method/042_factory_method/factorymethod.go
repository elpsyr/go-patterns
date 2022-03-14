package factory_method

type IRuleConfigParser interface {
	Parse(data []byte)
}

type jsonRuleConfigParser struct {
}

func (j jsonRuleConfigParser) Parse(data []byte) {
	panic("implement me")
}

type yamlRuleConfigParser struct {
}

func (y yamlRuleConfigParser) Parse(data []byte) {
	panic("implement me")
}

//工厂接口
type IRuleConfigParserFactory interface {
	CreateParser() IRuleConfigParser
}

//每种Parser都比较负责，交给他们自己的工厂去实现
type jsonRuleConfigParserFactory struct {
}

func (j jsonRuleConfigParserFactory) CreateParser() IRuleConfigParser {
	// 一些复杂的过程 ...
	panic("implement me")
}

type yamlRuleConfigParserFactory struct {
}

func (y yamlRuleConfigParserFactory) CreateParser() IRuleConfigParser {
	// 一些复杂的过程 ...
	panic("implement me")
}

func NewRuleConfigParser(t string) IRuleConfigParserFactory {
	switch t {
	case "json":
		return jsonRuleConfigParserFactory{}
	case "yaml":
		return yamlRuleConfigParserFactory{}
	}
	return nil

}
