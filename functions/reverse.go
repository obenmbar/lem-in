package lemino
// 🔄 Reverse inverse l'ordre des salles dans chaque chemin trouvé par BFS
// Elle permet de passer d'un chemin "fin → début" à un chemin "début → fin"
func (bfs *Bfs)Reverse(){
	for _, value := range bfs.Chemin {
		for k := 0 ; k < len(value)/2 ; k++ {
            	j := len(value) - 1 - k
			value[k], value[j] = value[j], value[k]
         }
		}
}
