package translate

import (
	"testing"
)

func TestTranslate(t *testing.T) {
	config := Config{
		Text: "Hello",
		From: "en",
		To: "es",
	}
	translated, err := Translate(config)
	if err != nil {
		t.Error(err)
	}

	if translated != "Hola" {
		t.Errorf("Expected 'Hola', got '%s'", translated)
	}
}

var articles = []string{
	"Fairfax Harrison (March 13, 1869 – February 2, 1938) was an American lawyer and businessman. He became a lawyer for the Southern Railway Company in 1896, and by 1906 he was the company's vice-president of finance. In 1913 he was elected president of Southern; under his leadership, the company expanded to an 8,000-mile (13,000 km) network across 13 states. Following the United States's entry into World War I, the federal government took control of the railroads, running them through the United States Railroad Administration, on which Harrison served. After the war, Harrison worked to improve the railroad's public relations, upgrade the locomotive stock by introducing more powerful engines, increase the company's amount of railroad track and extend the area serviced by the railway. Harrison struggled to keep the railroad afloat during the Great Depression, but by 1936 Southern was once again profitable. Harrison retired in 1937 and died three months later.",
}

func TestTranslateArticles(t *testing.T) {
	for _, article := range articles {
		config := Config{
			Text: article,
			From: "en",
			To: "es",
			Email: "test@gmail.com",
		}
		translated, err := Translate(config)
		if err != nil {
			t.Error(err)
		}

		t.Logf("Translated: %s", translated)
		if translated == "" {
			t.Errorf("Expected a translation, got an empty string")
		}
	}
}
