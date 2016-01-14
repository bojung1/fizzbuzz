package fizzbuzz

import "testing"
import "reflect"

func fullrun (count int) []string {
	results := make([]string,count)
	for i := 0; i < count; i++ {
		results[i] = fbdetection(i)
	}
	return results
}

var outputis string = "This is what the output is: "

func TestOne(t *testing.T) {
	t.Log("testing one, like it matters")
	testresult := fbdetection(1)
	if testresult != "1" {
		t.Errorf("Expected 1, got this instead: ",testresult)
	}
	t.Log(outputis, testresult)
}

func TestThree(t *testing.T) {
	t.Log("testing 3, expecting Fizz")
	testresult := fbdetection(3)
	if testresult != "Fizz" {
		t.Errorf("Expected Fizz, got this instead: ",testresult)
	}	
	t.Log(outputis, testresult)
}

func TestTwenty(t *testing.T) {
	t.Log("testing 20, expecting Buzz")
	testresult := fbdetection(20)
	if testresult != "Buzz" {
		t.Errorf("Expected Buzz, got this instead: ",testresult)
	}	
	t.Log(outputis, testresult)
}

func TestThirty(t *testing.T) {
	t.Log("testing 30, expecting FizzBuzz")
	testresult := fbdetection(30)
	if testresult != "FizzBuzz" {
		t.Errorf("Expected FizzBuzz, got this instead: ",testresult)
	}
	t.Log(outputis, testresult)
}

func TestRun(t *testing.T) {
	t.Log("Full Run! 1 to 100")
	realresults := fullrun(100)
	saneresults := []string {"FizzBuzz","1","2","Fizz","4","Buzz","Fizz","7","8","Fizz","Buzz","11","Fizz","13","14","FizzBuzz","16","17","Fizz","19","Buzz","Fizz","22","23","Fizz","Buzz","26","Fizz","28","29","FizzBuzz","31","32","Fizz","34","Buzz","Fizz","37","38","Fizz","Buzz","41","Fizz","43","44","FizzBuzz","46","47","Fizz","49","Buzz","Fizz","52","53","Fizz","Buzz","56","Fizz","58","59","FizzBuzz","61","62","Fizz","64","Buzz","Fizz","67","68","Fizz","Buzz","71","Fizz","73","74","FizzBuzz","76","77","Fizz","79","Buzz","Fizz","82","83","Fizz","Buzz","86","Fizz","88","89","FizzBuzz","91","92","Fizz","94","Buzz","Fizz","97","98","Fizz"}

 	sanitycheck := reflect.DeepEqual(realresults,saneresults)
	if sanitycheck != true {
		t.Errorf("Expected Bigass Slice, got this instead", realresults)	
	}
		
	t.Log(outputis, realresults)
}
