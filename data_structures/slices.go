package data_structures

type mapOperation func(int32) int32
type filterOperation func(int32) bool

func mapInts(op mapOperation, vals []int32) []int32 {

 result :=[]int32{}
	for _,i := range vals{
		result=append(result, op(i))
	}
return result

}

func filterInts(op filterOperation, vals []int32) []int32 {

	 result :=[]int32{}
		for _,i := range vals{
				if(op(i)){
								result=append(result, i)
				}
		}
	return result

}

func concatenate(dest []string, newValues ...string) []string {
  result :=make([]string, len(dest))
  copy(result,dest)
  result = append(result,newValues...)
  return result
}

func equals(list1 []string, list2 []string) bool {
  if len(list1)!=len(list2){
    return false
  }
  for index,item := range list1{
    if item != list2[index]{
      return false
    }
  }
	return true;
}

func partialReverse(src []int, from, to int) []int {
  result := []int{}
  if to > len(src){
    to = len(src)
  }
  for i :=to;i>=from;i-- {
    result=append(result,src[i])
  }
	return result
}
