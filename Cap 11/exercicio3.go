package main
import "fmt"

type veiculo struct {
	portas int
	cor string
}
type caminhonete struct{
	veiculo
	x4 bool
}
type sedan struct{
	veiculo
	modeloLuxo bool
}

func main(){
	generico := veiculo {4,"Preto"}
	Frontier := caminhonete{generico, true}
	Civic := sedan{generico, true}
	S10 := caminhonete{veiculo{4, "Branco"},true}

	fmt.Println(Frontier.cor)
	fmt.Println(Frontier)
	fmt.Println(Civic.cor)
	fmt.Println(Civic)
	fmt.Println(S10.cor)
	fmt.Println(S10)



}