package testing

import "bytes"

func formatTestmain(t *testFuncs) ([]byte, error) {
	var buf bytes.Buffer
	if err := testmainTmpl.Execute(&buf, t); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
