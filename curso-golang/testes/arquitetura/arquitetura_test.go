package arquitetura

func TestDependente(t *testing.T) {
	t.Parallel()
	if runtime.GOARCH == "amd64" {
		t.skip("Não funciona em arquitetura amd64")
	}
	// ...
	t.Fail()
}
