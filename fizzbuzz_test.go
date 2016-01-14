package fizzbuzz

import "testing"

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
	potato := fullrun(100)
	t.Log(potato)
}
