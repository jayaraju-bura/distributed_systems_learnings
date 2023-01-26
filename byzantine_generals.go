


package main
import "fmt"
import "strings"
import "os"

type Byzantine interface{
    call()
    om_algorithm()
    next_order()
    decision()
    
}

type General struct{
    id int
    other_generals []General
    orders []string
    is_traitor bool

}
func (g *General) next_order(is_traitor bool, order string, i int) string{
    if is_traitor {
        if i%2 == 0 {
            if order == "ATTACK" {
                return "RETREAT"
            }
            else {
                return "ATTACK"
            }
        } 
    }
    return order
    
}
func (g *General)om_algorithm(commander *General, m int, order string) {
    if m < 0 {
        g.orders = append(g.orders, order)
    }
    if m == 0 {
        for i, l := range(m.other_generals) {
            l.om_algorithm(g, m-1, g.next_order(g.is_traitor, order, i))
        }
        
    }
    else {
        
        for i, l := range(g.other_generals) {
            if &l != g && &l != commander {
                l.om_algorithm(g, m-1, g.next_order(g.is_traitor, order, i))
            }
        }
        
    }
    
}
func count_orders(orders []string) map[string]int {
    
    attack := 0
    retreat := 0
    for _, order := range(orders) {
        if order == "ATTACK" {
            attack ++
        }
        else if order == "RETREAT" {
            retreat ++
        }
    }
    var mp map[string]int
    mp["ATTACK"] = attack
    mp["RETREAT"] = retreat
    return mp
    
}
func init_generals(gen_spec string) []General {
    var generals []General
    general_spec := []rune(gen_spec)
    for idx, spec := range(general_spec) {
        var temp General
        temp.id = idx
        if spec == 'l' {
            
        }
        else if spec == 't' {
            temp.is_traitor = true
            general_spec = append(general_spec, temp)
        }
        else {
            fmt.Errorf("invalid in generals input %v", gen_spec)
            os.Exit(1)
            
        }
    }
    for _, spec := range(generals) {
        spec.other_generals = generals
    }
    return generals
    
}
func main() {
    var general_spec := "l,t,l,l,l"
    generals := init_generals(general_spec)
    
    fmt.Println("Hello World")
}
