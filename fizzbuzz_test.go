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


func TestOne(t *testing.T) {
	t.Log("testing one, like it matters")
	derp := fbdetection(1)
	if derp != "1" {
		t.Errorf("expected 1, did not get 1, I got this instead: ",derp)
	}
	t.Log("This is what the output is: ", derp)
}

func TestThree(t *testing.T) {
	t.Log("testing 3, expecting Fizz")
	derp := fbdetection(3)
	if derp != "Fizz" {
		t.Errorf("expected Fizz, did not get Fizz, i got this instead: ",derp)
	}	
	t.Log("this is what the output is: ", derp)
}


func TestTwenty(t *testing.T) {
	// what? 
	t.Log("testing 20, expecting Buzz")
	derp := fbdetection(20)
	if derp != "Buzz" {
		t.Errorf("expected Buzz, did not get Buzz, i got this instead: ",derp)
	}	
	t.Log("this is what the output is: ", derp)
}


func TestThirty(t *testing.T) {
	t.Log("testing 30, expecting FizzBuzz")
	derp := fbdetection(30)
	if derp != "FizzBuzz" {
		t.Errorf("expected FizzBuzz, did not get FizzBuzz, i got this instead: ",derp)
	}
	t.Log("this is what the output is: ", derp)

}

func TestRun(t *testing.T) {
	t.Log("Full Run! 1 to 100")
	ass := fullrun(100)
	tits := []string {"FizzBuzz","1","2","Fizz","4","Buzz","Fizz","7","8","Fizz","Buzz","11","Fizz","13","14","FizzBuzz","16","17","Fizz","19","Buzz","Fizz","22","23","Fizz","Buzz","26","Fizz","28","29","FizzBuzz","31","32","Fizz","34","Buzz","Fizz","37","38","Fizz","Buzz","41","Fizz","43","44","FizzBuzz","46","47","Fizz","49","Buzz","Fizz","52","53","Fizz","Buzz","56","Fizz","58","59","FizzBuzz","61","62","Fizz","64","Buzz","Fizz","67","68","Fizz","Buzz","71","Fizz","73","74","FizzBuzz","76","77","Fizz","79","Buzz","Fizz","82","83","Fizz","Buzz","86","Fizz","88","89","FizzBuzz","91","92","Fizz","94","Buzz","Fizz","97","98","Fizz"}

 	potato := reflect.DeepEqual(ass,tits)
	if potato != true {
		t.Errorf("Expected Bigass Slice, got this instead", ass)	
	}
		
	t.Log("This is what the output is: ", ass)
}
