package apinine

func TestApiNine(t *testing.T){
	server := httptest.NewServer(http.HandlerFunc(ApiNine))
	resp, err := http.Get(server.URL)
	if err != nil {
		t.Error(err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("expect 200, but got %d", resp.StatusCode)
	}
	defer resp.Body.Close()

	expect := "9"
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Error(err)
	}

	if expect != string(b) {
		t.Errorf("expect %s, but got %s", expect, string(b))
	}
}