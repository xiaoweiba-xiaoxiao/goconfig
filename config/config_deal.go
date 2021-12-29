pcackage config
import (
	"fmt"
	"sync"
)
//deal is a singleton
var (
	lock     sync.Mutex
	instance *configDeal
)

type configDeal struct {
}

func NewSingleton() *configDeal {
	if instance == nil {
		lock.Lock()
                // 双重判断
		if instance == nil {
			instance = &configDeal{
			}
		}
		lock.Unlock()
	}
	return instance
}

func (conf * ConfigDeal)Load_perperties(file string)(filepath string,strategyConfg Config)
{

}