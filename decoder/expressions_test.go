package decoder

// import (
// 	"testing"

// 	"github.com/hashicorp/hcl/v2"
// 	"github.com/hashicorp/hcl/v2/hclsyntax"
// 	"github.com/kr/pretty"
// )

// func TestExpressions(t *testing.T) {
// 	testConfig := []byte(`
// providers = {
// 	source = "somethign"
// 	version = 3.5
// }
// `)

// 	f, diags := hclsyntax.ParseConfig(testConfig, "test.tf", hcl.InitialPos)
// 	if len(diags) > 0 {
// 		t.Fatal(diags)
// 	}
// 	t.Fatal(pretty.Sprint(f.Body))
// }
