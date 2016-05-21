package main

// Arguments to format are:
//	[1]: type name
const stringValueToNameMap = `func %[1]s_Parse(s string) (%[1]s, error) {
	if val, ok := _%[1]sNameToValue_map[s]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("%%s does not belong to %[1]s values", s)
}
`

func (g *Generator) buildValueToNameMap(runs [][]Value, typeName string, runsThreshold int) {
	// At this moment, either "g.declareIndexAndNameVars()" or "g.declareNameVars()" has been called
	g.Printf("\nvar _%sNameToValue_map = map[string]%s{\n", typeName, typeName)
	for _, values := range runs {
		for _, value := range values {
			g.Printf("\t%q: %s,\n", value.display, value.name)
		}
	}
	g.Printf("}\n\n")
	g.Printf(stringValueToNameMap, typeName)
}

// Arguments to format are:
//	[1]: type name
const jsonMethods = `
func (i %[1]s) MarshalText() ([]byte, error) {
	return []byte(i.String()), nil
}

func (i *%[1]s) UnmarshalText(data []byte) error {
	val, err := %[1]s_Parse(string(data))
	if err != nil {
		return err
	}
	*i = val
	return nil
}
`

func (g *Generator) buildJSONMethods(runs [][]Value, typeName string, runsThreshold int) {
	g.Printf(jsonMethods, typeName)
}
