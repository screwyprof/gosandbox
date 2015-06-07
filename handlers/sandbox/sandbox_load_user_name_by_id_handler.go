package sandbox

import (
	"fmt"
	SandboxServicePkg "github.com/screwyprof/gosandbox/services/sandbox"
	"log"
	"net/http"
	"time"
)

type loadUserNameByIdHandler struct {
	sandboxService SandboxServicePkg.ISandboxService
}

func NewSandboxLoadUserNameByIdHandler(sandboxService SandboxServicePkg.ISandboxService) *loadUserNameByIdHandler {
	return &loadUserNameByIdHandler{sandboxService}
}

func (handler *loadUserNameByIdHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	t1 := time.Now()
	fmt.Fprintf(w, handler.sandboxService.LoadUserNameById(1))
	t2 := time.Now()
	log.Printf("[%s] %q %v\n", r.Method, r.URL.String(), t2.Sub(t1))
}
