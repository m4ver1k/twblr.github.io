package control_structures

func collatzChainLength(n int) int {
		// even /2
		//odd 3n+1
		count :=0;

		for n!=1{
					if n % 2==0{
						n=n/2
						count = count +1
					}else{
						n=(n*3)+1
						count = count +1
					}
			}
			return count

}
