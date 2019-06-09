package test_httptest_test

import (
	"net/http"
	"testing"
)

const checkMarkc = "\u2713"
const ballotX = "\u2717"

// TestDownload vvalida se a função http Get é capaz de fazer download de conteúdo
func TestDownload(t *testing.T) {
	url := "http://www.goinggo.net/feeds/posts/defaults?alt=rss"
	statusCode := 200

	t.Log("Given the need to teste downloading content.")
	{
		t.Logf("\tWhen checking \"%s\" for status code \"%d\"", url, statusCode)
		{
			resp, err := http.Get(url)
			if err != nil {
				t.Fatal("\t\tShould be able to make the Get call.", ballotX, err)
			}
			t.Log("\t\tShould be able to make the Get call.", checkMarkc)

			defer resp.Body.Close()

			if resp.StatusCode == statusCode {
				t.Logf("\t\tShould receive a \"%d\" status. %v", statusCode, checkMarkc)
			} else {
				t.Errorf("\t\tShould receive a \"%d\" status. %v %v", statusCode, ballotX, resp.StatusCode)
			}
		}
	}
}
