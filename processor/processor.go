package processor

import (
	"github.com/timob/jnigi"
	"log"
)

// CallMethod : normalize, tokenize, extractPhrase
func CallMethod(methodName string) string {
	_, env, err := jnigi.CreateJVM(jnigi.NewJVMInitArgs(false, true, jnigi.DEFAULT_VERSION, []string{"-Djava.class.path=processor-go.jar"}))
	if err != nil {
		log.Fatal(err)
	}

	str, err := env.NewObject("java/lang/String", []byte("한국어를 처리하는 예시입니닼ㅋㅋㅋ"))
	if err != nil {
		log.Fatal(err)
	}

  obj, err := env.NewObject("org/openkoreantext/go/ProcessorGo")
  if err != nil {
  	log.Fatal(err)
  }
  v, err := obj.CallMethod(env, methodName, "java/lang/String", str)
  if err != nil {
  	log.Fatal(err)
  }
	result, err := v.(*jnigi.ObjectRef).CallMethod(env, "getBytes", jnigi.Byte|jnigi.Array, env.GetUTF8String())
	if err != nil {
  	log.Fatal(err)
  }

	return string(result.([]byte))
}
