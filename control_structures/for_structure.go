package control_structures

func collatzChainLength(n int) int {
		// even /2
		//odd 3n+1
		count :=0;

		for {
					if n % 2==0{
						n=n/2
						count = count +1
					}else{
						n=(n*3)+1
						count = count +1
					}
					if(n==1){
						break
					}
			}
			return count

}
