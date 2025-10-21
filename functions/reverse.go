package lemino

func (bfs *Bfs)Reverse(){
	for _, value := range bfs.Chemin {
		for k := 0 ; k < len(value)/2 ; k++ {
            	j := len(value) - 1 - k
			value[k], value[j] = value[j], value[k]
         }
		}
}
