package client

import (
	"partIGoLanguageFoundation/chapterVIModularityAndDependencyManagement/series"
	"testing"
)

func TestPackage(t *testing.T) {
	t.Log(series.GetFibonacciSerie(5))
	t.Log(series.Square(5))
}
