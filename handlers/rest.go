package rest

// func TodoListHandler(w http.ResponseWriter, r *http.Request) {
//    var
// }

// func RestMiddleWare(h interface{}) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		handler := reflect.ValueOf(h)
// 		// 类型判断
// 		if handler.Kind() != reflect.Func {
// 			log.Println("function is not function")
// 			return
// 		}
// 		// 入参判断
// 		parameters := handler.Type().NumIn()
// 		argValues := make([]reflect.Value, 0, parameters)
// 		for i := 0; i < parameters; i++ {
// 			arg := handler.Type().In(i)
// 			switch arg.Kind() {
// 			case reflect.Struct:
// 				model := reflect.New(arg).Elem()

// 				var j map[string]interface{}
// 				resp, _ := ioutil.ReadAll(r.Body)
// 				json.Unmarshal(resp, &j)

// 				for i := 0; i < model.NumField(); i++ {
// 					v := model.Field(i)
// 					t := arg.Field(i)
// 					switch t.Type.Kind() {
// 					case reflect.String:
// 						v.SetString(j[t.Tag.Get("json")].(string))
// 					case reflect.Bool:
// 						v.SetBool(j[t.Tag.Get("json")].(bool))
// 					case reflect.Float32, reflect.Float64:
// 						v.SetFloat((j[t.Tag.Get("json")].(float64)))
// 					case reflect.Uint,
// 						reflect.Uint8,
// 						reflect.Uint16,
// 						reflect.Uint32,
// 						reflect.Uint64,
// 						reflect.Int,
// 						reflect.Int8,
// 						reflect.Int16,
// 						reflect.Int32,
// 						reflect.Int64:
// 						v.SetInt((j[t.Tag.Get("json")].(int64)))
// 					}
// 				}
// 				argValues = append(argValues, model)
// 			}
// 		}
// 		ret, _ := handler.Call(argValues)
// 	}
// }
